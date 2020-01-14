package model

import (
  "context"
  "google.golang.org/appengine/datastore"
)

type JWTPayload struct {
  Sub         string          `json:"sub"`      //subject (user id)
  Aud         string          `json:"aud"`      //audience (client id)
  Jti         string          `json:"jti"`      //jwt id
  Iat         int64           `json:"iat"`      //issue at time
  Exp         int64           `json:"exp"`      //expiration time
  Auth        AuthPayload     `json:"auths"` //metadata         
}

type AuthPayload struct {
  SiteID        string          `json:"siteId"`
  Subdomain     string          `json:"subdomain"`        
  Scopes        []ScopePayload  `json:"scopes"`
  UID           string          `json:"uId"`
  TranID        string          `json:"tranId"`
  SessionID     string          `json:"sId"`  
}

type ScopePayload struct {
  Scope   string  `json:"scope"`
  Access  string  `json:"access"`
}

var NullAuthPayload = &AuthPayload {
  SiteID: "",
  Scopes: nil,
}

var NullJWTPayload = &JWTPayload {
  Sub: "",
  Aud: "",
  Jti: "",
  Iat: int64(0),
  Exp: int64(0),
  Auth: *NullAuthPayload,
}

func SaveJWTPayload(ctx context.Context, saveRecord *JWTPayload) error {
  //modify for datatype
  dbKind := "JWTPayload"
  indexKey := saveRecord.Aud

  //generic save func, no parent key
  key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
  if _, err := datastore.Put(ctx, key, saveRecord); err != nil {
      return err
  }
  return nil
}

func LoadJWTPayloadByAud(ctx context.Context, aud string) (*JWTPayload, error) {
  //modify for datatype
  dbKind := "JWTPayload"
  var foundRecord JWTPayload
  indexKey := aud

  //generic load func, no parentkey
  key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
  if err := datastore.Get(ctx, key, &foundRecord); err != nil {
    return NullJWTPayload, err             
  }
  return &foundRecord, nil
}