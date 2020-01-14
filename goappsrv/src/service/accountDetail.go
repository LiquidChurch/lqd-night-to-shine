package service

import (
  "os"
  "encoding/base64"
  "strings"
	"time"
	"errors"

  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

var adminSID = os.Getenv("ADMIN_SID")

func PostAccountDetailbyGoogleToken(c helper.ContextDetail, tokenInfo helper.TokenInfo) (string, error) {
  //functionName := "ProviderUserLookup"
  helper.Log(c, "info", "entered")

  if tokenInfo.Sub == "" {
    helper.Log(c, "error", "token check", "error", "ProviderID Missing")
    err := errors.New("ProviderID Missing")
    return "", err
  }

	accountDetail := &model.AccountDetail {
    Provider: "Google",
    ProviderID: tokenInfo.Sub,
    Email: tokenInfo.Email,
    Name: tokenInfo.Name,
    PictureUrl: "",
    CreatedTime: time.Now(),
    LastLoginTime: time.Now(),
    UserDetailID: "",
	}

  accountDetailKey := accountDetail.Provider + "=" + accountDetail.ProviderID
  encAccountDetailKey := strings.TrimRight(base64.StdEncoding.EncodeToString([]byte(accountDetailKey)), "=")
	
	if tokenInfo.Picture != "" {
    accountDetail.PictureUrl = tokenInfo.Picture
	}
	
  helper.Log(c, "info", "AccountDetailByProvider called", "")

	foundAccountDetail, loadErr := model.LoadAccountDetail(c.Ctx, encAccountDetailKey)
	if loadErr != nil {
    helper.Log(c, "error", "load account detail", "key", encAccountDetailKey, "error", loadErr.Error())
	}
	
	if foundAccountDetail.ProviderID != accountDetail.ProviderID {   
    helper.Log(c, "info", "create account detail", "key", encAccountDetailKey)  
		
		accountDetail.UserDetailID = PostUserDetail(c, accountDetail)
    helper.Log(c, "info", "provider sub", "sub", tokenInfo.Sub) 
    if tokenInfo.Sub == adminSID {
      _, grantErr := PostGrant(c, accountDetail.UserDetailID, "00000000000000000000000000", "site", "write")
      
      if grantErr != nil {
        helper.Log(c, "error", "save admin grant error", "UID", accountDetail.UserDetailID, "error", grantErr.Error())
	    } 
      time.Sleep(1 * time.Second)
    }
    
	} else {
    helper.Log(c, "info", "update account detail", "key", encAccountDetailKey)

    accountDetail.CreatedTime = foundAccountDetail.CreatedTime
    accountDetail.UserDetailID = foundAccountDetail.UserDetailID
	}
	
	saveErr := model.SaveAccountDetail(c.Ctx, encAccountDetailKey, accountDetail)
	if saveErr != nil {
    helper.Log(c, "error", "save account detail", "key", encAccountDetailKey, "error", saveErr.Error())
	}
	
  helper.Log(c, "trace", "foundAccountDetail", "userDetailID", accountDetail.UserDetailID)
  helper.Log(c, "info", "exited")

  return accountDetail.UserDetailID, nil
}