package gqlSchema

import (
  "time"
  "errors"
  "context"
  "encoding/json"

  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
  "X/goappsrv/src/service"
)

type updateInput struct {
  AirTableId    string
  RecordType    string
  ForceUpdate   bool
}

func (r *Resolver) UpdateAirTable(ctx context.Context, args *struct{Params updateInput}) (*importStatusResolver, error) {
  auth := ctx.Value("AUTH").(model.AuthPayload)
  c := helper.ContextDetail {
      Ctx: ctx.Value("GAE").(context.Context),
      FunctionName: "UpdateAirTableMutation",
      TranID: auth.TranID,
      UID: auth.UID,
      ProductID: "",
  }   
  
  helper.Log(c, "info", "Started")
  if (c.UID == "00000000000000000000000000") {
    helper.Log(c, "warning", "No User Logged in", "uid", c.UID)
    err := errors.New("No Logged in user`")
    return nil, err
  }

  
  isAuthorized := helper.AuthCheck(auth, "site", "write")
  
  if (c.UID == "00000000000000000000000000" || auth.SiteID != "00000000000000000000000000" || isAuthorized == false) {
    helper.Log(c, "warning", "Update Airtable not authorized", "uid", c.UID)
    err := errors.New("Not Authorized")
    return nil, err
  }
 
  airTableDetail, loadAirTableDetailErr := model.LoadItemDetail(c.Ctx, args.Params.AirTableId)
  
  if (loadAirTableDetailErr != nil || airTableDetail.ID != args.Params.AirTableId ){
    helper.Log(c, "error", "invalid airtable Id", "uid", c.UID, "airTableId", args.Params.AirTableId)
    err := errors.New("AirTableId not found")
    return nil, err    
  }
  
  airTableRecords, err := service.LoadAirTable(c, *airTableDetail)
  
  if err != nil {
    helper.Log(c, "warning", "Error loading air table", "uid", c.UID, "error", err.Error())
    err := errors.New("Error accessing airtable API")
    return nil, err
  }

  // var guestItemsResolver []*itemDetailResolver
  
  importStatus := &ImportStatus {
    Created: 0,
    Modified: 0,
    Skipped: 0,
    Total: 0,
  }
  
  for i, _ := range airTableRecords.Records {
    itemID := "" 
    time.Sleep(50 * time.Millisecond)
    foundItem, loadItemErr := model.LoadItemDetailByExtID(c.Ctx, airTableRecords.Records[i].Id)
    if loadItemErr != nil {
      helper.Log(c, "warning", "Error loading item detail by ref ID", "uid", c.UID, "refId", airTableRecords.Records[i].Id)
    }
    if (*foundItem).ID != "" {
      itemID = (*foundItem).ID
    } 

    bytesField, _ := json.Marshal(airTableRecords.Records[i].Fields)
    guestItem := &model.ItemDetail {
      ID: itemID,
      ParentID: args.Params.AirTableId,
      Type: args.Params.RecordType,
      ExtID: airTableRecords.Records[i].Id,
      Name: airTableRecords.Records[i].Fields.Name,
      Description: string(bytesField),
      ExtSync: airTableRecords.Records[i].Fields.LastModified,
    }
    
    if (args.Params.ForceUpdate) {
      foundItem.ExtSync = ""
    }
    
    if foundItem.ExtSync != guestItem.ExtSync {
      if guestItem.ID == "" {
        helper.Log(c, "info", "Item Created", "itemID", guestItem.ID, "extID", guestItem.ExtID)
        importStatus.Created = importStatus.Created + 1
      } else {
        helper.Log(c, "info", "Item Modified", "itemID", guestItem.ID, "extID", guestItem.ExtID)
        importStatus.Modified = importStatus.Modified + 1
      }
      guestItem, err = service.PostItemDetail(c, guestItem)
      
    } else {
      importStatus.Skipped = importStatus.Skipped + 1   
    }
  }
  
  importStatus.Total = importStatus.Created + importStatus.Modified + importStatus.Skipped

  helper.Log(c, "info", "Completed")
  return &importStatusResolver{c, importStatus}, nil
}
