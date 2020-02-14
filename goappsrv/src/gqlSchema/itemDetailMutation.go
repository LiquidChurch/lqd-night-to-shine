package gqlSchema

import (
  "errors"
  "context"
  "strings"
  "time"
  
  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
  "X/goappsrv/src/service"
)

type itemInput struct {
  Id            string
  Type          string
  Name          string
  Description   string
  WebURL        string
  Color         string
  ExtId         string
}

func (r *Resolver) PostItemDetail(ctx context.Context, args *struct{PostItem itemInput}) (*itemDetailResolver, error) {
  auth := ctx.Value("AUTH").(model.AuthPayload)
  c := helper.ContextDetail {
      Ctx: ctx.Value("GAE").(context.Context),
      FunctionName: "PostItemDetail",
      TranID: auth.TranID,
      UID: auth.UID,
      ProductID: "",
  }   
  
  helper.Log(c, "info", "Started")
 
  var err error
  createdTime := time.Now()
  if c.UID == "00000000000000000000000000" {
    err = errors.New("User Not Found") 
  }
  
  if helper.AuthCheck(auth, "site", "write") != true {
    err = errors.New("User Not Authorized")
  }

  if args.PostItem.Id != "" {
    foundItem, loadItemErr := model.LoadItemDetail(c.Ctx, args.PostItem.Id) 
  
    if loadItemErr != nil {
      helper.Log(c, "error", "Loading Item By ID Error", "uid", c.UID, "error", loadItemErr.Error())
    }
    
    if foundItem.ParentID != auth.SiteID {
      err = errors.New("User Grant Mismatch")
    } else {
      createdTime = foundItem.CreatedTime
    }
  }
  
  switch strings.ToLower(args.PostItem.Type) {
    case "airtable":
      args.PostItem.Type = "airtable"
    default:
      err = errors.New("Invalid Item Type") 
      helper.Log(c, "error", "Invalid Item Type", "uid", c.UID, "type", args.PostItem.Type)
      return &itemDetailResolver{c, model.NullItemDetail}, err
  }

  if err != nil {
    helper.Log(c, "error", "Authorization error", "uid", c.UID, "error", err.Error())
    return &itemDetailResolver{c, model.NullItemDetail}, err
  }  
  
  item := &model.ItemDetail {
    ID: args.PostItem.Id,
    ParentID: auth.SiteID,
    Type: args.PostItem.Type,
    Name: args.PostItem.Name,
    Description: args.PostItem.Description,
    WebURL: args.PostItem.WebURL,
    Color: args.PostItem.Color,
    ExtID: args.PostItem.ExtId,
    CreatedTime: createdTime,
  }
  
  item, err = service.PostItemDetail(c, item)

  if err != nil {
    helper.Log(c, "error", "Saving item error", "uid", c.UID, "error", err.Error())
    return &itemDetailResolver{c, model.NullItemDetail}, err
  }  

  helper.Log(c, "info", "Completed") 
  return &itemDetailResolver{c, item}, nil
}

