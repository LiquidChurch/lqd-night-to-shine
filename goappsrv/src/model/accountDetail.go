package model

import (
	"time"
  "context"
  "google.golang.org/appengine/datastore"
)

type AccountDetail struct {
  Provider        string
  ProviderID      string
  Email           string
  Name            string
  PictureUrl      string
  CreatedTime     time.Time
	LastLoginTime	time.Time
	UserDetailID    string
} 

var nullAccountDetail = &AccountDetail {
  Provider: "",
  ProviderID: "",
  Email: "",
  Name: "",
  PictureUrl: "",
  CreatedTime: time.Now(),
  LastLoginTime: time.Now(),
  UserDetailID: "",
}

func SaveAccountDetail(ctx context.Context, indexKey string, saveRecord *AccountDetail) error {
  //modify for datatype
  dbKind := "AccountDetail"

  //generic save func, no parent key
  key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
  if _, err := datastore.Put(ctx, key, saveRecord); err != nil {
    return err
  }
  return nil
}

func LoadAccountDetail(ctx context.Context, indexKey string) (*AccountDetail, error) {
  //modify for datatype
  dbKind := "AccountDetail"
  var foundRecord AccountDetail

  //generic load func, no parentkey
  key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
  if err := datastore.Get(ctx, key, &foundRecord); err != nil {
    return nullAccountDetail, err             
  }
  return &foundRecord, nil
}