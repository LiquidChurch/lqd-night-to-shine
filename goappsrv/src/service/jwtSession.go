package service

import (
  "os"
  "strings"
  "strconv"
  "errors"
  "encoding/json"
  "math/rand"
  "time"
  
  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

var jwtDurationDay, _ = strconv.ParseInt(os.Getenv("JWT_DURATION_DAY"),10, 64)
var expireDuration = int64(86400 * jwtDurationDay) //in seconds
// var refreshPeriod = int64(2)

func CreateJWTSession(c helper.ContextDetail) (model.AuthPayload, string) {
  //functionName := "CreateJWTSession"
  helper.Log(c, "info", "entered")    

  rand.Seed(time.Now().UnixNano())

  a := make([]byte, 32)
  _, _ = rand.Read(a)

  b := make([]byte, 32)
  _, _ = rand.Read(b)

  audToken := helper.Base64Encode(a)
  jtiToken := helper.Base64Encode(b)

  currentTimeInUnix := time.Now().Unix()
  expireTimeInUnix := currentTimeInUnix + expireDuration
  
  helper.Log(c, "info", "Loading site grant by UID", "uid", c.UID)
  userSites, loadGrantErr := model.LoadGrantDetailsByActor(c.Ctx, c.UID, "site")
  if loadGrantErr != nil {
      helper.Log(c, "error", "Error loading user site", "uid", c.UID, "error", loadGrantErr.Error())
  }
  
  if len(*userSites) > 1 {
      helper.Log(c, "error", "Multiple User Sites Found", "uid", c.UID)
  }
    
  var userSiteID string
  var userSiteGrant string
  
  if (*userSites)[0].ItemID == "" {
    helper.Log(c, "info", "No site grants found for user", "uid", c.UID)
    userSiteID = "00000000000000000000000000"
    userSiteGrant = "none"
  } else {
    helper.Log(c, "info", "Site grant set for user", "site", (*userSites)[0].ItemID, "grant", (*userSites)[0].Grant)
    userSiteID = (*userSites)[0].ItemID
    userSiteGrant = (*userSites)[0].Grant
  }
  
  scopePayload := &[]model.ScopePayload {
      {
        Scope: "site",
        Access: userSiteGrant,
      },
      {
        Scope:"checkin",
        Access:"read",
      },
      {
        Scope:"group",
        Access:"write",
      },
  }
  
  authPayload := &model.AuthPayload {
    SiteID: userSiteID,
    Scopes: *scopePayload,
  }
  
  jwtPayload := &model.JWTPayload {
    Sub: c.UID,
    Aud: audToken,
    Jti: jtiToken,
    Iat: currentTimeInUnix,
    Exp: expireTimeInUnix,
    Auth: *authPayload,
  }

  saveError := model.SaveJWTPayload(c.Ctx, jwtPayload)

  if saveError != nil {
      helper.Log(c, "error", "Saving JWTSession", "aud", jwtPayload.Aud, "error", saveError.Error())
  }

  sessionToken := helper.SignJWTToken(jwtPayload)

  helper.Log(c, "info", "exited")    

  return *authPayload, sessionToken
}

func RefreshJWTSession(c helper.ContextDetail, aud string, jti string) string {
  c.FunctionName = "RefreshJWTSession"
  helper.Log(c, "info", "Stated") 

  jwtPayload, loadError := model.LoadJWTPayloadByAud(c.Ctx, aud)
  currentTimeInUnix := time.Now().Unix()

  if loadError != nil {
    helper.Log(c, "error", "Loading JWTSession", "error", loadError.Error())
  }

  if jwtPayload.Jti != jti {        
    if jwtPayload.Exp >= currentTimeInUnix + (expireDuration - 2)  {
      helper.Log(c, "warn", "JTI just refreshed", "aud", jwtPayload.Aud)
      return "not_refreshed"
    } else {
      helper.Log(c, "error", "JTI check fails", "aud", jwtPayload.Aud, "error", "JTI doesn't match")
      return "not_valid"               
    }
  }

  b := make([]byte, 32)
  _, _ = rand.Read(b)

  jwtPayload.Jti = helper.Base64Encode(b)

  expireTimeInUnix := currentTimeInUnix + expireDuration

  jwtPayload.Iat = currentTimeInUnix
  jwtPayload.Exp = expireTimeInUnix

  saveError := model.SaveJWTPayload(c.Ctx, jwtPayload)

  if saveError != nil {
    helper.Log(c, "error", "Saving JWTSession", "aud", jwtPayload.Aud, "error", saveError.Error())
  }

  sessionToken := helper.SignJWTToken(jwtPayload)

  helper.Log(c, "info", "Completed") 
  return sessionToken
}

func ValidateJWTSession(c helper.ContextDetail, sessionToken string) (*model.JWTPayload, error) {    
  token := strings.Split(sessionToken, ".")

  header := helper.JWTHeader{}

  if len(token) != 3 {
    splitErr := errors.New("Invalid JWT token")
    return model.NullJWTPayload, splitErr
  }

  encodedHeader, errHeader := helper.Base64Decode(token[0])
  errHeader = json.Unmarshal([]byte(encodedHeader), &header)

  if errHeader != nil {
    errHeader := errors.New("Invalid Header in JWT token")
    return model.NullJWTPayload, errHeader
  }

  if header.Typ != "JWT" {return model.NullJWTPayload, errors.New("Invalid token type, only JWT allowed")}
  if header.Alg != "HS256" {return model.NullJWTPayload, errors.New("Invalid algorithm type, only HS256 allowed")}    

  payload, errDecode := helper.DecodeJWTToken(token[1])

  if errDecode != nil {
    return  model.NullJWTPayload, errDecode
  }

  currentTimeInUnix := time.Now().Unix()

  if currentTimeInUnix > payload.Exp {return model.NullJWTPayload, errors.New("Token is expired")}

  isValidated := helper.CompareHmac(token[2], token[0] + "." + token[1])            

  if isValidated {
    return payload, nil
  } else {
    err := errors.New("Invalid JWT token")
    return model.NullJWTPayload, err
  }
}