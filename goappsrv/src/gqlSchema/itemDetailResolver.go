package gqlSchema

import (
  "context"

  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

type itemFilter struct {
  Id        string
  IdType    string
  Type      string
  ParentId  string
}

func (r *Resolver) GetItem(ctx context.Context, args *struct{Lookup itemFilter}) (*itemDetailResolver, error) {
  auth := ctx.Value("AUTH").(model.AuthPayload)
  c := helper.ContextDetail {
      Ctx: ctx.Value("GAE").(context.Context),
      FunctionName: "GetItemDetail",
      TranID: auth.TranID,
      UID: auth.UID,
      ProductID: "",
  }   
  
  helper.Log(c, "info", "Started")
  helper.Log(c, "info", "Item Lookup", "uid", c.UID, )

  
  var foundItem *model.ItemDetail
  var loadItemErr error
  
  switch args.Lookup.IdType {
    case "Int":
      helper.Log(c, "info", "Item Lookup by Item ID", "uid", c.UID, "itemId", args.Lookup.Id)    
      foundItem, loadItemErr = model.LoadItemDetail(c.Ctx, args.Lookup.Id)
    case "Ext":
      helper.Log(c, "info", "Item Lookup by External ID", "uid", c.UID, "extId", args.Lookup.Id)    
      foundItem, loadItemErr = model.LoadItemDetailByExtID(c.Ctx, args.Lookup.Id)
    case "ETP":
      helper.Log(c, "info", "Item Lookup by ETP", "uid", c.UID, "extId", args.Lookup.Id, "type", args.Lookup.Type, "parentId", args.Lookup.ParentId)   
      foundItem, loadItemErr = model.LoadItemDetailByETP(c.Ctx, args.Lookup.Id, args.Lookup.Type, args.Lookup.ParentId)
    default:
      helper.Log(c, "warning", "IdType not valid", "uid", c.UID, "id", args.Lookup.Id)
      foundItem = model.NullItemDetail
  }

  if loadItemErr != nil {
    helper.Log(c, "warning", "Error loading item detail", "uid", c.UID, "id", args.Lookup.Id)
    foundItem = model.NullItemDetail
  }
  
  if foundItem.ID == "" {
    helper.Log(c, "warning", "No item found item", "uid", c.UID, "id", args.Lookup.Id)
    foundItem = model.NullItemDetail
  }
  
  // var item *model.ItemDetail
   
  helper.Log(c, "info", "Completed")

  return &itemDetailResolver{c, foundItem}, nil
}

type itemDetailResolver struct {
  c helper.ContextDetail
  u *model.ItemDetail
}

func (r *itemDetailResolver) Id() string {
  return r.u.ID
}

func (r *itemDetailResolver) ParentId() string {
  return r.u.ParentID
}

func (r *itemDetailResolver) Type() string {
  return r.u.Type
}

func (r *itemDetailResolver) Name() string {
  return r.u.Name
}

func (r *itemDetailResolver) Description() string {
  return r.u.Description
}

func (r *itemDetailResolver) WebURL() string {
  return r.u.WebURL
}

func (r *itemDetailResolver) Color() string {
  return r.u.Color
}

func (r *itemDetailResolver) ExtId() string {
  return r.u.ExtID
}
