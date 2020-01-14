package service

import (
  "math/rand"
  "encoding/base64"
  "strings"
	"time"

  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

func PostUserDetail(c helper.ContextDetail, newAccountDetail *model.AccountDetail) string {
  userDetailID, err := helper.GenerateULID()

  if err != nil {
      helper.Log(c, "error", "generate ULID", "error", err.Error())
  }

  helper.Log(c, "info", "entered", "userDetailID", userDetailID)

  userDetail := &model.UserDetail {
    ID: userDetailID,
    Email: newAccountDetail.Email,
    Name: newAccountDetail.Name,
    Company: "",
    Role: "",
    PictureURL: newAccountDetail.PictureUrl,
    CreatedTime: time.Now(),
    UpdatedTime: time.Now(),
	}

  saveErr := model.SaveUserDetail(c.Ctx, userDetailID, userDetail)    
	if saveErr != nil {
    helper.Log(c, "error", "save login detail", "userDetailID", userDetailID, "error", saveErr.Error())
	}
  
  helper.Log(c, "info", "exited", "userDetailID", userDetailID)
  return userDetailID
}

func GetUserDetail(c helper.ContextDetail) (*model.UserDetail, error) {
  foundUserDetail, loadErr := model.LoadUserDetail(c.Ctx, c.UID)
  if loadErr != nil {
    return nil, loadErr
  }

  if foundUserDetail.ID == "" {
    helper.Log(c, "error", "user not found", "userDetailID", c.UID)
  }
  
  helper.Log(c, "info", "exited", "userDetailID", c.UID)
  return foundUserDetail, nil
}

func generateUserID(email string) (string, error) {
    rand.Seed(time.Now().UnixNano())
    
    base64Email := base64.URLEncoding.EncodeToString([]byte(email))
    
    if len(email) > 15 {
        email = email[0:15]
    }
    size := 30-len(email)
     
    b := make([]byte, size)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    
    base64Prefix := base64.URLEncoding.EncodeToString(b)
    
    userID := strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(base64Prefix + base64Email)), "=")
    
    return userID, nil
}