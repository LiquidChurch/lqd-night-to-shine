package service

import (
  "time"
  
  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

func PostGrant(c helper.ContextDetail, actorID string,  itemID string, grantType string, grant string) (*model.GrantDetail, error) {
  helper.Log(c, "info", "PostGrant Entered", "itemID", itemID, "actorID", actorID)

  grantID, err := helper.GenerateULID()  
  if err != nil {
    helper.Log(c, "error", "generate ULID", "error", err.Error())
    return model.NullGrantDetail, err
  }
  
  newGrant := &model.GrantDetail {
    ID: grantID,
    Type: grantType,
    ItemID: itemID,
    ActorID: actorID,
    Grant: grant,
    Status: 1,
    CreatedTime: time.Now(),
    UpdatedTime: time.Now(),
  }
  
  err = model.SaveGrantDetail(c.Ctx, grantID, newGrant)
  if err != nil {
    helper.Log(c, "error", "Save Grant Error", "error", err.Error())
    return model.NullGrantDetail, err
  }
  
  return newGrant, nil
}

func GetGrantDetailsByItem(c helper.ContextDetail, itemID string) (*[]model.GrantDetail, error) {
  helper.Log(c, "info", "GetGrantDetailByItem Entered", "itemID", itemID)
  
  
  grants, err := model.LoadGrantDetailsByItem(c.Ctx, itemID)

  if err != nil {
    helper.Log(c, "error", "LoadGrantDetailsByItem", "error", err.Error())
    return nil, err
  }
  
  return grants, nil
}

func GetGrantDetailsByActor(c helper.ContextDetail, actorID string, grantType string) (*[]model.GrantDetail, error) {
  helper.Log(c, "info", "GetAuthDetailByActor Entered", "actorID", actorID)
  
  grants, err := model.LoadGrantDetailsByActor(c.Ctx, actorID, grantType)
  if err != nil {
    helper.Log(c, "error", "LoadAuthDetailByActor", "error", err.Error())
    return nil, err
  }
  
  return grants, nil
}