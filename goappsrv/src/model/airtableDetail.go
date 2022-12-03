package model

type AirTableErrorResp struct {
  Error AirTableError `json:"error"`
}

type AirTableError struct {
  Type string `json:"type"`
  Message string `json:"message"`
}

type AirTableList struct {
  Records []AirTableRecord  `json:"records"`
  Offset  string            `json:"offset,omitempty"` 
}

type AirTableRecord struct {
  Id      string            `json:"id"`
  Fields  GuestFields       `json:"fields"`
}

type GuestFields struct {
  Name            string          `json:"Name,omitempty"`
  FirstName       string          `json:"Guest First Name,omitempty"`
  LastName        string          `json:"Guest Last Name,omitempty"`
  PromDay         string          `json:"Prom Day,omitempty"`
  Gender          string          `json:"Gender,omitempty"`
  LOSupervision   string          `json:"Level of Supervision,omitempty"`
  SNDescription   string          `json:"SN Description,omitempty"`
  RespiteRoom     string          `json:"Respite Room"`
  SpecificBuddy   string          `json:"Specific Buddy,omitempty"`
  LOBathroom      string          `json:"Bathroom Assistance,omitempty"`
  Medication      string          `json:"Medication During Prom,omitempty"`
  DRestriction    []string        `json:"Dietary Restrictions,omitempty"`
  Sensory         []string        `json:"Sensory,omitempty"`
  Limo            string          `json:"Limo,omitempty"`
  ContactName     string          `json:"Contact Name,omitempty"`
  ContactNumber   string          `json:"Contact #,omitempty"`
  ContactEmail    string          `json:"Email,omitempty"`
  Notes           string          `json:"NOTES,omitempty"`
  ArrivalTime     string          `json:"Arrival Time,omitempty"`
  PagerNumber     string          `json:"Pager Number"`
  TimeOfMed       string          `json:"Time of Medication"`
  LastModified    string          `json:"Last Modified,omitempty"`
  QRValue         string          `json:"QR Value,omitempty"`
  QRImage         []QrImageStruct `json:"QR Image,omitempty"`
  CheckedIn       bool            `json:"Checked In,omitempty"`
  MobilityAsst    string          `json:"Mobility Assistance,omitempty"`
  CommuAsst       string          `json:"Communication Assistance,omitempty"`
  Teams           string          `json:"TEAMS,omitempty"`
  Role            string          `json:"ROLE,omitempty"`
  ROverview       []RoleOverview  `json:"ROLE OVERVIEW,omitempty"`
  TeamRoster      string          `json:"Team Roster List,omitempty"`
}

type QrImageStruct struct {
  Id        string        `json:"id,omitempty"`
  Url       string        `json:"url"`
  FileName  string        `json:"filename"`
}

type RoleOverview struct {
  FileName  string        `json:"filename"`
  Url       string        `json:"url"`
}
