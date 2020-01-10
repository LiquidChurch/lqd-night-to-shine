package main

import (
  "time"
  "net/http"
  "google.golang.org/appengine"
  "github.com/gorilla/mux"
  "github.com/etherlabsio/healthcheck"
)

var webAppDist = "dist/webapp"

func main() {
  r := mux.NewRouter()
  
  r.HandleFunc("/_ah/warmup", WarmupHandler)
  
  r.PathPrefix("/{_:.*}").HandlerFunc(WebAppHandler)
  
  http.Handle("/", r)
  
  appengine.Main()
}


func WarmupHandler(w http.ResponseWriter, r *http.Request) {
  http.Handle("/_ah/warmup", healthcheck.Handler(
    healthcheck.WithTimeout(5*time.Second),
  )) 
}

func WebAppHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, webAppDist)
}
