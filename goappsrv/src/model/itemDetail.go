package model

import (
  "time"
  "context"
  "google.golang.org/appengine/datastore"
)

type ItemDetail struct {
  ID              string
  ParentID        string
  Type            string
  Name            string
  Description     string
  WebURL          string
  Color           string
  Role            string
  ExtID           string
  ExtSync         string
  CreatedTime     time.Time
  UpdatedTime	    time.Time
} 

var NullItemDetail = &ItemDetail {
  ID: "",
  ParentID: "",
  Type: "",
  Name: "",
  Description: "",
  WebURL: "",
  Color:"",
  Role: "",
  ExtID: "",
  ExtSync: "",
  CreatedTime: time.Now(),
  UpdatedTime: time.Now(),
}

func SaveItemDetail(ctx context.Context, indexKey string, saveRecord *ItemDetail) error {
  dbKind := "ItemDetail"
    
  //generic save func, no parent key
  key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
  if _, err := datastore.Put(ctx, key, saveRecord); err != nil {
    return err
  }
  return nil
}

func LoadItemDetail(ctx context.Context, indexKey string) (*ItemDetail, error) {
  dbKind := "ItemDetail"
  var foundRecord ItemDetail
    
  //generic load func, no parentkey
  key := datastore.NewKey(ctx, dbKind, indexKey, 0, nil)
  if err := datastore.Get(ctx, key, &foundRecord); err != nil {
    return NullItemDetail, err             
  }
  return &foundRecord, nil
}

func LoadItemsByType(ctx context.Context, itemType string, parentID string) (*[]ItemDetail, error) {
  dbKind := "ItemDetail"
  
  var foundRecords []ItemDetail
  
  query := datastore.NewQuery(dbKind).
           Filter("Type =", itemType).
           Filter("ParentID =", parentID)
  
  if _, err := query.GetAll(ctx, &foundRecords); err != nil {
    return nil, err
  }
  
  if len(foundRecords) == 0 {
    foundRecords = append(foundRecords, *NullItemDetail)
  }
  
  return &foundRecords, nil
}

func LoadItemDetailByID(ctx context.Context, extID string, itemType string, parentID string) (*ItemDetail, error) {
  dbKind := "ItemDetail"
  
  var foundRecords []ItemDetail
  
  query := datastore.NewQuery(dbKind).
           Filter("ExtID =", extID).
           Filter("Type =", itemType).
           Filter("ParentID =", parentID)
  
  if _, err := query.GetAll(ctx, &foundRecords); err != nil {
    return nil, err
  }
  
  if len(foundRecords) == 0 {
    foundRecords = append(foundRecords, *NullItemDetail)
  }
  
  return &foundRecords[0], nil  
}

func LoadItemDetailByExtID(ctx context.Context, extID string, itemType string, parentID string) (*ItemDetail, error) {
  dbKind := "ItemDetail"
  
  var foundRecords []ItemDetail
  
  query := datastore.NewQuery(dbKind).
           Filter("ExtID =", extID).
           Filter("Type =", itemType).
           Filter("ParentID =", parentID)
  
  if _, err := query.GetAll(ctx, &foundRecords); err != nil {
    return nil, err
  }
  
  if len(foundRecords) == 0 {
    foundRecords = append(foundRecords, *NullItemDetail)
  }
  
  return &foundRecords[0], nil  
}
