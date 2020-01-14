package gqlSchema

import (
  "strconv"
  "context"
  
  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

func (r *Resolver) SessionDetail(ctx context.Context) *sessionDetailResolver {
  auth := ctx.Value("AUTH").(model.AuthPayload)
  c := helper.ContextDetail {
      Ctx: ctx.Value("GAE").(context.Context),
      FunctionName: "GetItemDetail",
      TranID: auth.TranID,
      UID: auth.UID,
      ProductID: "",
  }   
  
  helper.Log(c, "info", "Started")
  
  //sID := value[2]
  sID := auth.SessionID
  eXP := int32(0)

  if sID == "" {
    helper.Log(c, "info", "Session Detail", "error", "No Session")
    sID = "not_valid"
  }

  if sID != "not_valid"  {
    eXP = helper.CalculateEXP(sID)
    helper.Log(c, "debug", "Session Detail", "exp", strconv.FormatInt(int64(eXP), 10))
  }

  helper.Log(c, "info", "Completed")
  return &sessionDetailResolver{c.UID, sID, eXP}
}

type sessionDetailResolver struct {
  uID string
  sID string
  eXP int32
}

func (r *sessionDetailResolver) SessionToken() string {
  if r.uID == "00000000000000000000000000" {
    return ""
  } else {
    return r.sID
  }
}

func (r *sessionDetailResolver) UserID() string {
  if r.uID == "00000000000000000000000000" {
    return ""
  } else {
    return r.uID
  }
}

func (r *sessionDetailResolver) Status() string {
  if (r.uID == "00000000000000000000000000" || r.uID == "") {
    return "Unauthorized"
  } else {
    if r.eXP == (86400 * 30) {
      return "Refresh"
    } else {
      return "Authorized"
    }
  }
}

func (r *sessionDetailResolver) Expiration() int32 {
  if r.uID == "00000000000000000000000000" {
    return int32(0)
  } else {
    return r.eXP
  }
}