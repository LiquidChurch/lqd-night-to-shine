package service

import (
  "net/http"
  "bytes"
  "time"
  "errors"
  "io/ioutil"
  "encoding/json"
  
  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

type airTableRecord struct {
  Id      string        `json:"id"`
  Fields  guestField    `json:"fields"`
}

type AirTableList struct {
  Records []airTableRecord  `json:"records"`
  Offset  string            `json:"offset,omitempty"` 
}

type qrImageStruct struct {
  Id        string        `json:"id,omitempty"`
  Url       string        `json:"url"`
  FileName  string        `json:"filename"`
}

type roleOverview struct {
  FileName  string        `json:"filename"`
  Url       string        `json:"url"`
}

type guestField struct {
  Name            string          `json:"Name,omitempty"`
  FirstName       string          `json:"Guest First Name,omitempty"`
  LastName        string          `json:"Guest Last Name,omitempty"`
  PromDay         string          `json:"Prom Day,omitempty"`
  Gender          string          `json:"Gender,omitempty"`
  LOSupervision   string          `json:"Level of Supervision,omitempty"`
  SNDescription   string          `json:"SN Description,omitempty"`
  RespiteRoom     []string        `json:"Respite Room,omitempty"`
  SpecificBuddy   string          `json:"Specific Buddy,omitempty"`
  LOBathroom      string          `json:"Level of Bathroom Assistance,omitempty"`
  Medication      string          `json:"Medication During Prom,omitempty"`
  DRestriction    []string        `json:"Dietary Restrictions,omitempty"`
  Sensory         []string        `json:"Sensory,omitempty"`
  CherryOnTop     string          `json:"Cherry On Top,omitempty"`
  Limo            string          `json:"Limo,omitempty"`
  ContactName     string          `json:"Contact Name,omitempty"`
  ContactNumber   string          `json:"Contact #,omitempty"`
  ContactEmail    string          `json:"Email,omitempty"`
  MailingAddress  string          `json:"Mailing Address,omitempty"`
  Notes           string          `json:"NOTES,omitempty"`
  ArrivalTime     string          `json:"Arrival Time,omitempty"`
  PagerNumber     string          `json:"Pager Number,omitempty"`
  TimeOfMed       string          `json:"Time of Medication,omitempty"`
  LastModified    string          `json:"Last Modified,omitempty"`
  QRValue         string          `json:"QR Value,omitempty"`
  QRImage         []qrImageStruct `json:"QR Image,omitempty"`
  Teams           string          `json:"TEAMS,omitempty"`
  Role            string          `json:"ROLE,omitempty"`
  ROverview       []roleOverview  `json:"ROLE OVERVIEW,omitempty"`
  TeamRoster      string          `json:"Team Roster List,omitempty"`
}

func LoadAirTable(c helper.ContextDetail, airTableDetail model.ItemDetail) (*AirTableList, error) {
  
  var airTableList = new(AirTableList)
  var offset = ""
  var isEnd = false;
  
  for ok := true; ok; ok = (!isEnd) {
  
    url := "https://api.airtable.com/v0/" + airTableDetail.WebURL + "?view=QRAppView&offset=" + offset
    //url := "https://api.airtable.com/v0/" + airTableDetail.WebURL + "?view=QRAppView&maxRecords=15&offset=" + offset
    
    helper.Log(c, "info", "Loading air table", "uid", c.UID, "url", url)
    
    client := &http.Client{
      CheckRedirect: func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
      },
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
      helper.Log(c, "error", "Error Loading air table", "airtableId", airTableDetail.ID, "error", err.Error())
      err := errors.New("Air Table API Error")
      return nil, err
    }

    req.Header.Set("Authorization", "Bearer " + airTableDetail.ExtID)

    resp, err := client.Do(req)
    if err != nil {
      helper.Log(c, "error", "Error Loading air table", "airtableId", airTableDetail.ID, "error", err.Error())
      err := errors.New("Air Table API Error")
      return nil, err
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
      helper.Log(c, "error", "Http call not successful", "airtableId", airTableDetail.ID, "response code", resp.Status)
      err := errors.New("Air Table API Error")
      return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      helper.Log(c, "error", "Error parsing response body", "airtableId", airTableDetail.ID, "error", err.Error())
      err := errors.New("Error Parsing Air Table")
      return nil, err
    }

    var respJson = new(AirTableList)

    err = json.Unmarshal(body, &respJson)
    if err != nil {
      helper.Log(c, "error", "Error parsing response body", "airtableId", airTableDetail.ID, "error", err.Error())
      err := errors.New("Error Parsing Air Table")
      return nil, err
    }

    airTableList.Records = append(airTableList.Records, respJson.Records...)
    
    for i, _ := range respJson.Records {
      if (len(respJson.Records[i].Fields.QRImage) == 0) {
        helper.Log(c, "info", "Generating QR Code", "airtableId", airTableDetail.ID, "extId", respJson.Records[i].Id)
        err := LoadQRCode(c, airTableDetail, respJson.Records[i].Id )
        if err != nil {
          helper.Log(c, "error", "Error Generating QR Code", "extId", respJson.Records[i].Id, "error", err.Error())
        }
      }
    }
    
    if respJson.Offset != "" {
      offset = respJson.Offset
      isEnd = false
      time.Sleep(200 * time.Millisecond)
    } else {
      offset = ""
      isEnd = true
    }
  
  }
    
  return airTableList, nil
}

func LoadQRCode(c helper.ContextDetail, airTableDetail model.ItemDetail, itemId string) error {
  url := "https://api.airtable.com/v0/" + airTableDetail.WebURL
  client := &http.Client{
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    },
  }
  qrImage := qrImageStruct {
    Url: "https://api.qrserver.com/v1/create-qr-code/?size=250x250&data=https://nts.lqd.ch/" + itemId,
    FileName: "qrcode",
  }
  
  qrImageArray := []qrImageStruct{qrImage}

  bodyItem := airTableRecord {
    Id: itemId,
  }

  bodyItem.Fields.QRImage = qrImageArray
  bodyItem.Fields.QRValue = itemId
  var bodyJson = new(AirTableList)
  
  bodyJson.Records = append(bodyJson.Records, bodyItem)
  
  bytesField, _ := json.Marshal(bodyJson)
  
  req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bytesField))
  
  req.Header.Set("Authorization", "Bearer " + airTableDetail.ExtID)
  req.Header.Set("Content-Type", "application/json")
  
  resp, err := client.Do(req)
  if err != nil {
    helper.Log(c, "warning", "AirTable QR Update Error", "extId", itemId, "error", err.Error())
  }

  defer resp.Body.Close()
  if resp.StatusCode != 200 {
      helper.Log(c, "warning", "AirTable QR Update Error", "extId", itemId, "response code", resp.Status)
      err := errors.New("QR Code Generation error")
      return err
   }
  return nil
}