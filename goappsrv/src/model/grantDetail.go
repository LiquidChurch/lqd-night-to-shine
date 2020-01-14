package model

import (
  "time"
  "context"
  "google.golang.org/appengine/datastore"
)

type GrantDetail struct {
  ID              string
  Type            string
  ItemID          string
  ActorID         string
  Grant           string
  Status          int
  CreatedTime     time.Time
  UpdatedTime	  time.Time    
}

var NullGrantDetail = &GrantDetail {
  ID: "",
  Type: "",
  ItemID: "",
  ActorID: "",
  Grant: "",
  Status: 0,
  CreatedTime: time.Now(),
  UpdatedTime: time.Now(),
}

func SaveGrantDetail(ctx context.Context, indexKey string, saveRecord *GrantDetail) error {
  dbKind := "GrantDetail"
    
  key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
  if _, err := datastore.Put(ctx, key, saveRecord); err != nil {
    return err
  }
  return nil
}

func LoadGrantDetail(ctx context.Context, indexKey string) (*GrantDetail, error) {
  dbKind := "GrantDetail"
  var foundRecord GrantDetail
    
  key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
  if err := datastore.Get(ctx, key, &foundRecord); err != nil {
    return NullGrantDetail, err             
  }
  return &foundRecord, nil
}

func LoadGrantDetailsByItem(ctx context.Context, itemID string) (*[]GrantDetail, error) {
  dbKind := "GrantDetail"
  
  var foundRecords []GrantDetail
  
  query := datastore.NewQuery(dbKind).
           Filter("ItemID =", itemID)

  if _, err := query.GetAll(ctx, &foundRecords); err != nil {
    return nil, err
  }
  
  if len(foundRecords) == 0 {
    foundRecords = append(foundRecords, *NullGrantDetail)
  }
  
  return &foundRecords, nil
}

func LoadGrantDetailsByActor(ctx context.Context, actorID string, grantType string) (*[]GrantDetail, error) {
  dbKind := "GrantDetail"
  
  var foundRecords []GrantDetail
  
  query := datastore.NewQuery(dbKind).
           Filter("ActorID =", actorID).
           Filter("Type =", grantType)

  if _, err := query.GetAll(ctx, &foundRecords); err != nil {
    return nil, err
  }
  
  if len(foundRecords) == 0 {
    foundRecords = append(foundRecords, *NullGrantDetail)
  }
  
  return &foundRecords, nil
}
