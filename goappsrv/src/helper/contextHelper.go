package helper

import (
  "context"
  cRand "crypto/rand"
  mRand "math/rand"
  "encoding/base64"
  "github.com/oklog/ulid"
  "google.golang.org/appengine/log"    
)

type ContextDetail struct {
  Ctx context.Context
  FunctionName string
  TranID string
  UID string
  ProductID string
}

func Log(c ContextDetail, level string, action string, args ...string) {
  funcString := "\"Function\":\"" + c.FunctionName + "\""
  txnIDString := "\"TranID\":\"" + c.TranID + "\""
  uIDString := "\"UID\":\"" + c.UID + "\""
  levelString := "\"Level\":\"" + level + "\""
  actionString := "\"Action\":\"" + action + "\""
  argsString := ""

  if len(args) >=2 {
    argsString = ", \"" + args[0] + "\":\"" + args[1] + "\"" 
  } 

  if len(args) >= 4 {
    argsString = argsString + "\", \"" + args[2] + "\":\"" + args[3] + "\""    
  }

  logString := "{" + funcString + ", " + txnIDString +  ", " + uIDString + ", " + levelString + ", " + actionString + argsString + "}"

  switch level {
    case "warning": log.Warningf(c.Ctx, logString)
    case "info": log.Infof(c.Ctx, logString)
    case "error": log.Errorf(c.Ctx, logString)  
    case "critical": log.Criticalf(c.Ctx, logString)
    default: log.Debugf(c.Ctx, logString)
  }
}

func GenerateTranID() string {   
  b := make([]byte, 12)
  _, err := mRand.Read(b)
  if err != nil {
    return ""
  }

  return base64.URLEncoding.EncodeToString(b)
}

func GenerateULID() (string, error) {
	id, err := ulid.New(ulid.Now(), cRand.Reader)
	if err != nil {
	  return "", err
  }
	return id.String(), nil
}