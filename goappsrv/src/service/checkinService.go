package service

import (
  "net/http"
  "bytes"
  "errors"
  "io/ioutil"  
  "encoding/json"
  
  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

func PatchAirtable(c helper.ContextDetail, airTableDetail model.ItemDetail, patchRecord model.AirTableRecord) (*model.AirTableList, error) {
  
  // Declare variables
  var airTableList = new(model.AirTableList)
  var payloadRecords =new (model.AirTableList)

  // Set URL  
  url := "https://api.airtable.com/v0/" + airTableDetail.WebURL
  helper.Log(c, "info", "Patch Airtable Info", "url", url)
  
  //. Initiate Client
  client := &http.Client{
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      return http.ErrUseLastResponse
    },
  }
  
  // Set Patch Payload Json
  payloadRecords.Records = append(payloadRecords.Records, patchRecord)
  payload, _ := json.Marshal(payloadRecords)
  helper.Log(c, "info", "Patch Airtable Info", "payload", string(payload[:]))
  
  // Initate HTTP Request
  req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(payload))
  if err != nil {
    helper.Log(c, "error", "Error Loading air table", "airtableId", airTableDetail.ID, "error", err.Error())
    err := errors.New("Air Table API Error")
    return nil, err
  }
  
  // Set HTTP Request Headers
  req.Header.Set("Authorization", "Bearer " + airTableDetail.ExtID)
  req.Header.Set("Content-Type", "application/json; charset=utf-8")

  // Execute HTTP Request
  resp, err := client.Do(req)
  if err != nil {
    helper.Log(c, "error", "Error Loading air table", "airtableId", airTableDetail.ID, "error", err.Error())
    err := errors.New("Air Table API Error")
    return nil, err
  }
  defer resp.Body.Close()

  // Parse HTTP Response Body
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    helper.Log(c, "error", "Error parsing response body", "airtableId", airTableDetail.ID, "error", err.Error())
    err := errors.New("Error Parsing Air Table")
    return nil, err
  }
  
  // Log is response status is not 200
  if resp.StatusCode != 200 {
    helper.Log(c, "error", "Http call not successful", "airtableId", airTableDetail.ID, "response code", resp.Status)
    helper.Log(c, "error", "Http call not successful", "airtableId", airTableDetail.ID, "error", string(body[:]))
    err := errors.New("Air Table API Error")
    return nil, err
  }
  
  // Convert Response from Json to Object
  err = json.Unmarshal(body, &airTableList)
  if err != nil {
    helper.Log(c, "error", "Error parsing response body", "airtableId", airTableDetail.ID, "error", err.Error())
    err := errors.New("Error Parsing Air Table")
    return nil, err
  }
  
  return airTableList, nil
}