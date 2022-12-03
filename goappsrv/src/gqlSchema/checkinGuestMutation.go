package gqlSchema

import (
  //"time"
  "context"
  "encoding/json"
  
  "X/goappsrv/src/helper"
  "X/goappsrv/src/service"
  "X/goappsrv/src/model"
)

type checkinDetail struct {
  Id          string
  Description string
}

func (r *Resolver) PostCheckinDetail(ctx context.Context, args *struct{CheckinInput checkinDetail}) (*itemDetailResolver, error) {
  auth := ctx.Value("AUTH").(model.AuthPayload)
  c := helper.ContextDetail {
    Ctx: ctx.Value("GAE").(context.Context),
    FunctionName: "PostCheckinDetail",
    TranID: auth.TranID,
    UID: auth.UID,
    ProductID: "",
  } 
  
  helper.Log(c, "info", "Started")
  helper.Log(c, "info", "Checkin ID", "checkinId", args.CheckinInput.Id)
  
  var foundItem *model.ItemDetail
  var parentItem *model.ItemDetail
  var loadItemErr error
  
  foundItem, loadItemErr = model.LoadItemDetailByExtID(c.Ctx, args.CheckinInput.Id)
  
  if loadItemErr != nil {
    helper.Log(c, "error", "Error loading Checkin Id", "checkinId", args.CheckinInput.Id, "error", loadItemErr.Error())
    foundItem = model.NullItemDetail
  }
   
  if foundItem.ID == "" {
    helper.Log(c, "warning", "Checkin ID not found", "checkinId", args.CheckinInput.Id)
    return &itemDetailResolver{c, foundItem}, nil
  }

  helper.Log(c, "info", "Found Checkin Item", "checkinItem", foundItem.ID)

  parentItem, loadItemErr = model.LoadItemDetail(c.Ctx, foundItem.ParentID)

  if loadItemErr != nil {
    helper.Log(c, "error", "Error loading Parent ID", "parentId", foundItem.ParentID, "error", loadItemErr.Error())
    foundItem = model.NullItemDetail
  }
   
  if parentItem.ID == "" {
    helper.Log(c, "warning", "Parent ID not found", "parentId", foundItem.ParentID)
    return &itemDetailResolver{c, foundItem}, nil
  }

  helper.Log(c, "info", "Found Parent Item", "parentItem", foundItem.ID)
  
  guestFields := model.GuestFields{}
  json.Unmarshal([]byte(args.CheckinInput.Description), &guestFields)
  
  guestFields.CheckedIn = true
  guestFields.LastModified = ""
  guestFields.QRValue = ""
  guestFields.QRImage = nil
  
  patchRecord := model.AirTableRecord {
    Id: foundItem.ExtID,
    Fields: guestFields,
  }
  
  airTableRecords, err := service.PatchAirtable(c, *parentItem, patchRecord)
  if err != nil {
    helper.Log(c, "warning", "Error checkin guest", "error", err.Error())
  }
  
    bytesField, _ := json.Marshal(airTableRecords.Records[0].Fields)
    guestItem := &model.ItemDetail {
      ID: foundItem.ID,
      ParentID: parentItem.ID,
      Type: "guest",
      ExtID: airTableRecords.Records[0].Id,
      Name: airTableRecords.Records[0].Fields.Name,
      Description: string(bytesField),
      ExtSync: airTableRecords.Records[0].Fields.LastModified,
    }
      
    guestItem, _ = service.PostItemDetail(c, guestItem)
      
  helper.Log(c, "info", "Update GuestInfo", "guestItem", guestItem.Description)
  
  helper.Log(c, "info", "Completed")
  return &itemDetailResolver{c, guestItem}, nil
}