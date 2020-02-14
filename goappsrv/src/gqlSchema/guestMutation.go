package gqlSchema

import (
  "log"
  "time"
  "context"
  "encoding/json"

  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
  "X/goappsrv/src/service"
)

func (r *Resolver) UpdateGuests(ctx context.Context, args *struct{AirTableId string}) (*importStatusResolver, error) {
  auth := ctx.Value("AUTH").(model.AuthPayload)
  c := helper.ContextDetail {
      Ctx: ctx.Value("GAE").(context.Context),
      FunctionName: "UpdateGuests",
      TranID: auth.TranID,
      UID: auth.UID,
      ProductID: "",
  }   
  
  helper.Log(c, "info", "Started")
//  isAuthorized := helper.AuthCheck(auth, "site", "write")
  
//  if (auth.SiteID != "00000000000000000000000000" || isAuthorized == false) {
//    helper.Log(c, "warning", "update church not authorized", "uid", c.UID)
//    err := errors.New("Not Authorized")
//    return nil, err
//  }
 
  airTableDetail, loadAirTableDetailErr := model.LoadItemDetail(c.Ctx, args.AirTableId)
  
  if loadAirTableDetailErr != nil {
    helper.Log(c, "error", "invalid airtable Id", "uid", c.UID, "airTableId", args.AirTableId)
  }
  
  airTableRecords, err := service.LoadAirTable(c, *airTableDetail)
  
  if err != nil {
    helper.Log(c, "warning", "Error loading air table", "uid", c.UID, "error", err.Error())
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
      ParentID: args.AirTableId,
      Type: "guest",
      ExtID: airTableRecords.Records[i].Id,
      Name: airTableRecords.Records[i].Fields.Name,
      Description: string(bytesField),
      ExtSync: airTableRecords.Records[i].Fields.LastModified,
    }
    
    if foundItem.ExtSync != guestItem.ExtSync {
      if guestItem.ID == "" {
        helper.Log(c, "info", "Guest Item Created", "id", guestItem.ID, "extID", guestItem.ExtID)
        log.Println(foundItem)
        importStatus.Created = importStatus.Created + 1
      } else {
        importStatus.Modified = importStatus.Modified + 1
        helper.Log(c, "info", "Guest Item Modified", "id", guestItem.ID, "extID", guestItem.ExtID)
      }
      guestItem, err = service.PostItemDetail(c, guestItem)
      
    } else {
      // helper.Log(c, "info", "No change to guest detail", "RefID", guestItem.ExtID)
      importStatus.Skipped = importStatus.Skipped + 1   
    }
    // guestItemsResolver = append(guestItemsResolver, &itemDetailResolver{c, guestItem})
  }
  
  importStatus.Total = importStatus.Created + importStatus.Modified + importStatus.Skipped

  helper.Log(c, "info", "Completed")
  return &importStatusResolver{c, importStatus}, nil
}
