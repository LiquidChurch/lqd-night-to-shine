package service

import (
  "os"
  "time"
  "context"
  "strconv"
  "net/http"
  "google.golang.org/appengine"

  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

var jwtRefreshMin, _ = strconv.ParseInt(os.Getenv("JWT_REFRESH_MIN"),10, 64)
var RefreshTime = int64(60 * jwtRefreshMin)

func AuthCheck(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)

  c := helper.ContextDetail {
    Ctx: appengine.NewContext(r),
    FunctionName: "AuthorizationCheck",
    TranID: helper.GenerateTranID(),
    UID: "00000000000000000000000000",
    ProductID: "",
  }    

  helper.Log(c, "info", "Started")
  
	authType := ""
  authTypeCookie, err1 := r.Cookie("authType")
	
  if err1 != nil {
    helper.Log(c, "error", "AuthType Check", "error", err1.Error())
	} else {
	  authType = authTypeCookie.Value
	  helper.Log(c, "info", "AuthType Check", "authType", authType)
	}
	
	authToken := ""
	authTokenCookie, err2 := r.Cookie("authToken")
	if err2 != nil {
    helper.Log(c, "error", "AuthToken Check", "error", err2.Error())
	} else {
	  authToken = authTokenCookie.Value
	}

  // sessionToken := ""
  aud := os.Getenv("GCLOUD_AUD")
	
	// Provider Authorization
	var tokenInfo *helper.TokenInfo
  
  var authPayload model.AuthPayload
  authPayload.SessionID = ""
  authPayload.TranID = c.TranID
  
  // Google Authorization
	if authType == "Google" && authToken != "" {
    certInfo, err3 :=  helper.GetGoogleCerts(ctx)
    if err3 != nil {
      helper.Log(c, "error", "fetch google cert", "error", err3.Error())
    }

    if certInfo == nil {
      helper.Log(c, "error", "fetch google cert", "error", "Google Cert Not Fetched")
    } else {
      var err error

      tokenInfo, err = helper.VerifyGoogleIDToken(authToken, certInfo, aud)	

      if err != nil{
        helper.Log(c, "error", "google id token validation error", "error", err.Error())  
      }
    }

    if tokenInfo != nil {
      helper.Log(c, "info", "UserDetailIdbyGoogleToken called", "")
      var err error        
      c.UID, err = PostAccountDetailbyGoogleToken(c, *tokenInfo)

      authPayload.UID = c.UID
      
      if err != nil {
        helper.Log(c, "error", "provider user lookup error", "error", err.Error()) 
      }
    }
	}
	
  // Session Authorization
	if authType == "Session" && authToken != "" {
	  helper.Log(c, "info", "Session Token Check")
    // sessionToken = authToken
    jwtPayload, errPayload := ValidateJWTSession(c, authToken)

    // helper.Log(c, "info", "JWT Token", "jwtPayload", jwtPayload)
    authPayload = jwtPayload.Auth
    authPayload.SessionID = authToken
    
    if errPayload != nil {
      helper.Log(c, "error", "Session Token Check", "error", errPayload.Error()) 
		} else {
		  helper.Log(c, "debug", "Session Token Check", "jwtPayload", jwtPayload.Sub) 
		  c.UID = jwtPayload.Sub
      authPayload.UID = c.UID
		}
		
		currentTimeInUnix := time.Now().Unix()
		
		tokenAge := currentTimeInUnix - jwtPayload.Iat
    helper.Log(c, "debug", "Check Token Age", "tokenAge", strconv.FormatInt(tokenAge, 10)) 

    if tokenAge > RefreshTime {
      helper.Log(c, "debug", "Refresh Token Triggered") 
      // sessionToken = RefreshJWTSession(c, jwtPayload.Aud, jwtPayload.Jti)
      authPayload.SessionID = RefreshJWTSession(c, jwtPayload.Aud, jwtPayload.Jti)
      if authPayload.SessionID == "not_refreshed" {
        // sessionToken = authToken
        authPayload.SessionID = authToken
      }
    }
	}

  if authType != "Session" && c.UID != "00000000000000000000000000" {
    // sessionToken = CreateJWTSession(c)
    authPayload, authPayload.SessionID = CreateJWTSession(c)
    authPayload.UID = c.UID
  }

  if authPayload.SessionID == "not_valid" {
    c.UID = "00000000000000000000000000"
    authPayload.UID = c.UID
  }

  helper.Log(c, "info", "Completed")
    
  next.ServeHTTP(w, requestWithAppengineContext(r, c.Ctx, authPayload))
  })
}

func requestWithAppengineContext(r *http.Request, c context.Context, auth model.AuthPayload ) *http.Request {
	ctx := context.WithValue(r.Context(), "GAE", c)
	return r.WithContext(context.WithValue(ctx, "AUTH", auth))
}
