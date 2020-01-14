package model

import (
	"time"
  "context"
  "google.golang.org/appengine/datastore"
)

type UserDetail struct {
  ID              string
  Email           string
  Name            string
  Company         string
  Role            string
  PictureURL      string
  CreatedTime     time.Time
	UpdatedTime		time.Time
} 

var NullUserDetail = &UserDetail {
  ID: "",
  Email: "",
  Name: "",
  Company: "",
  Role: "",
  PictureURL: "",
  CreatedTime: time.Now(),
  UpdatedTime: time.Now(),
}

func SaveUserDetail(ctx context.Context, indexKey string, saveRecord *UserDetail) error {
    //modify for datatype
    dbKind := "UserDetail"
    
    //generic save func, no parent key
    key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
    if _, err := datastore.Put(ctx, key, saveRecord); err != nil {
        return err
    }
    return nil
}

func LoadUserDetail(ctx context.Context, indexKey string) (*UserDetail, error) {   
    //modify for datatype
    dbKind := "UserDetail"
    var foundRecord UserDetail
    
    //generic load func, no parentkey
    key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
    if err := datastore.Get(ctx, key, &foundRecord); err != nil {
        return NullUserDetail, err             
	}
    return &foundRecord, nil
}