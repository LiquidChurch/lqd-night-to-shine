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
  Name            string `datastore:",noindex"`
  Description     string `datastore:"-"`
  ByteDescription []byte `datastore:",noindex"`
  WebURL          string `datastore:",noindex"`
  Color           string `datastore:",noindex"`
  Role            string `datastore:",noindex"`
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
  
  saveRecord.ByteDescription = []byte(saveRecord.Description)
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
  
  foundRecord.Description = string(foundRecord.ByteDescription)
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

func LoadItemDetailByExtID(ctx context.Context, extID string) (*ItemDetail, error) {
  dbKind := "ItemDetail"
  
  var foundRecords []ItemDetail
  
  query := datastore.NewQuery(dbKind).
           Filter("ExtID =", extID)
  
  if _, err := query.GetAll(ctx, &foundRecords); err != nil {
    return nil, err
  }
  
  if len(foundRecords) == 0 {
    foundRecords = append(foundRecords, *NullItemDetail)
  }
  
  foundRecords[0].Description = string(foundRecords[0].ByteDescription)
  return &foundRecords[0], nil  
}

func LoadItemDetailByETP(ctx context.Context, extID string, itemType string, parentID string) (*ItemDetail, error) {
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

  foundRecords[0].Description = string(foundRecords[0].ByteDescription)
  return &foundRecords[0], nil  
}

