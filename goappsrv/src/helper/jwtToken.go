package helper

import (
  "os"
  "time"
  "strings"
  "errors"
  "net/url"
  "encoding/base64"
  "encoding/json"
  "crypto/hmac"
  "crypto/sha256"

  "X/goappsrv/src/model"
)

var secret = os.Getenv("JWT_SECRET")

type JWTHeader struct {
  Alg         string  `json:"alg"`
  Typ         string  `json:"typ"`
}

func DecodeJWTToken(encodedJWTPayload string) (*model.JWTPayload, error) {
  payload := model.JWTPayload{}

  encodedPayload, errPayload := Base64Decode(encodedJWTPayload)
  errPayload = json.Unmarshal([]byte(encodedPayload), &payload)

  if errPayload != nil {
    errPayload := errors.New("Invalid Payload in JWT token")
    return model.NullJWTPayload, errPayload
  }

  return &payload, nil
}

func SignJWTToken(jwtPayload *model.JWTPayload) string {    
  jwtHeader := JWTHeader {
    Alg: "HS256",
    Typ: "JWT",
  }

  header, _ := json.Marshal(jwtHeader)
  payload, _ := json.Marshal(jwtPayload)

  encodedHeader := Base64Encode(header)
  encodedPayload := Base64Encode(payload)
  encodedSignature := strings.Replace(url.QueryEscape(hmac256(encodedHeader + "." + encodedPayload, secret)), "%", "", -1)

  sessionToken := encodedHeader + "." + encodedPayload + "." + encodedSignature

  return sessionToken
}

func Base64Encode(src []byte) string {
  str := strings.TrimRight(base64.StdEncoding.EncodeToString(src), "=")
  return str
}

func Base64Decode(src string) (string, error) {
  remainder := len(src)%4
  if remainder > 0 {
    for i := remainder; i <= 3 ; i++ { src = src + "=" }
  }

  decoded, err := base64.StdEncoding.DecodeString(src)

  if err != nil {
    return "", err
  }

  return string(decoded), nil
}

func hmac256(src string, secret string) string {
  key := []byte(secret)
  h := hmac.New(sha256.New, key)
  h.Write([]byte(src))
  return Base64Encode(h.Sum(nil))
}

func CompareHmac(messageHmac string, message string) bool {
  signature := strings.Replace(url.QueryEscape(hmac256(message, secret)), "%", "", -1)
  return messageHmac == signature
}

func CalculateEXP(sID string) int32 {
  jwtPayload := strings.Split(sID, ".")
  payload, _ := DecodeJWTToken(jwtPayload[1])

  currentTimeInUnix := time.Now().Unix()

  return int32(payload.Exp - currentTimeInUnix)
}