package gqlSchema

import (
  "context"
  "net/http"
	"io/ioutil"
  "encoding/json"
  "errors"
  "log"

  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
  "X/goappsrv/src/service"
)

type airTableRecord struct {
  Id      string        `json:"id"`
  Fields  guestField    `json:"fields"`
}

type airTableList struct {
  Records []airTableRecord  `json:"records"`
}

type guestField struct {
  GuestName       string    `json:"Guest Name"`
  LOSupervision   string    `json:"Level of Supervision"`
  SNDescription   string    `json:"SN Description"`
  LOBathroom      string    `json:"Level of Bathroom Assistance"`
  Medication      string    `json:"Medication During Prom"`
  DRestriction    []string    `json:"Dietary Restrictions"`
  Notes           string    `json:"Notes"`
  Limo            string    `json:"Limo"`
  ContactName     string    `json:"Contact Name"`
  ContactNumber   string    `json:"Contact #"`
  ContactEmail    string    `json:"Email"`
  WantsCall       string    `json:"Wants a Call"`
  LastModified    string    `json:"Last Modified"`
}

func (r *Resolver) UpdateGuests(ctx context.Context, args *struct{AirTableId string}) (*[]*itemDetailResolver, error) {
  auth := ctx.Value("AUTH").(model.AuthPayload)
  c := helper.ContextDetail {
      Ctx: ctx.Value("GAE").(context.Context),
      FunctionName: "UpdateGuests",
      TranID: auth.TranID,
      UID: auth.UID,
      ProductID: "",
  }   

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
  
  url := "https://api.airtable.com/v0/" + airTableDetail.WebURL
  client := &http.Client{
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    },
  }
  
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Println(err)
  }
    
  req.Header.Set("Authorization", "Bearer " + airTableDetail.ExtID)
  
  resp, err := client.Do(req)
  if err != nil {
    log.Println(err)
  }
  
  defer resp.Body.Close()
  
  if resp.StatusCode != 200 {
    helper.Log(c, "warning", "http not successfu", "uid", c.UID, "response code", resp.Status)
    err := errors.New("Subdomain not found")
    return nil, err
  }
  
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Println(err)
  }
  
  var respJson = new(airTableList)
  var guestItemsResolver []*itemDetailResolver
  
  err = json.Unmarshal(body, &respJson)
  if err != nil {
    log.Println(err)
  }
  
  for i, _ := range respJson.Records {
    log.Println(respJson.Records[i].Fields)
    itemID := "" 

    foundItem, loadItemErr := model.LoadItemDetailByExtID(c.Ctx, respJson.Records[i].Id)
    if loadItemErr != nil {
      helper.Log(c, "warning", "Error loading item detail by ref ID", "uid", c.UID, "refId", respJson.Records[i].Id)
    }
    if (*foundItem).ID != "" {
      itemID = (*foundItem).ID
    }

    bytesField, _ := json.Marshal(respJson.Records[i].Fields)
    guestItem := &model.ItemDetail {
      ID: itemID,
      ParentID: args.AirTableId,
      Type: "guest",
      ExtID: respJson.Records[i].Id,
      Name: respJson.Records[i].Fields.GuestName,
      Description: string(bytesField),
      WebURL: respJson.Records[i].Fields.LastModified,
    }
    
    if foundItem.WebURL != guestItem.WebURL {
      guestItem, err = service.PostItemDetail(c, guestItem)
      helper.Log(c, "info", "No change to guest detail", "RefID", guestItem.ExtID)
    }
    guestItemsResolver = append(guestItemsResolver, &itemDetailResolver{c, guestItem})
  }
  
  return &guestItemsResolver, nil
}
