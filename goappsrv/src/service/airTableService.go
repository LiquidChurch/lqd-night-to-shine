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

func LoadAirTable(c helper.ContextDetail, airTableDetail model.ItemDetail) (*model.AirTableList, error) {
  
  var airTableList = new(model.AirTableList)
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

    var respJson = new(model.AirTableList)

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
  qrImage := model.QrImageStruct {
    Url: "https://api.qrserver.com/v1/create-qr-code/?size=250x250&data=https://nts.lqd.ch/" + itemId,
    FileName: "qrcode",
  }
  
  qrImageArray := []model.QrImageStruct{qrImage}

  bodyItem := model.AirTableRecord {
    Id: itemId,
  }
  
  helper.Log(c,"info","airtable record data", "airtablerecord", bodyItem.Id, "", "")
  
  bodyItem.Fields.QRImage = qrImageArray
  bodyItem.Fields.QRValue = itemId
  var bodyJson = new(model.AirTableList)
  
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