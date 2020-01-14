package helper

import (
  "X/goappsrv/src/model"
)

func AuthCheck(auth model.AuthPayload, scopeCheck string, accessCheck string) bool {
  accessMatch := false
  
  for _, scope := range auth.Scopes {
    if scopeCheck == scope.Scope {
      if accessCheck == scope.Access {
        accessMatch = true
      }
    }
  }  
  return accessMatch
}