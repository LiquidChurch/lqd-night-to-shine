package gqlSchema

import (
  "context"
  
  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
  "X/goappsrv/src/service"
)

type favContextKey string

func (r *Resolver) GetUser(ctx context.Context) (*userDetailResolver, error) {

  auth := ctx.Value("AUTH").(model.AuthPayload)
  c := helper.ContextDetail {
      Ctx: ctx.Value("GAE").(context.Context),
      FunctionName: "GetItemDetail",
      TranID: auth.TranID,
      UID: auth.UID,
      ProductID: "",
  }   
    
  helper.Log(c, "info", "Started")

  var user *model.UserDetail
  var err error

  if c.UID != "00000000000000000000000000" {
   user, err = service.GetUserDetail(c)
   if err != nil {
       helper.Log(c, "error", "User Detail Load", "error", err.Error())
        user = model.NullUserDetail
     }
  } else {
     helper.Log(c, "info", "User Detail Load", "error", "No UID")
     user = model.NullUserDetail
  }

  helper.Log(c, "info", "AuthPayload SiteID", "siteID", auth.SiteID)
  if user.Name != "" {
      helper.Log(c, "info", "User Detail Load")    
  }

  helper.Log(c, "info", "Completed")
  return &userDetailResolver{c, user}, nil
}

type userDetailResolver struct {
  c helper.ContextDetail
  u *model.UserDetail
}

func (r *userDetailResolver) Id() string {
  return "1"
}

func (r *userDetailResolver) Uid() string {
  return r.u.ID
}

func (r *userDetailResolver) Email() string {
  return r.u.Email
}

func (r *userDetailResolver) Name() string {
  return r.u.Name
}

func (r *userDetailResolver) Company() string {
  return r.u.Company
}

func (r *userDetailResolver) Role() string {
  return r.u.Role
}

func (r *userDetailResolver) PictureURL() string {
  return r.u.PictureURL
}
