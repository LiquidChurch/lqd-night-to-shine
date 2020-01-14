package main

import (
  "time"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/etherlabsio/healthcheck"
  "google.golang.org/appengine"
  graphql "github.com/graph-gophers/graphql-go"
  "github.com/graph-gophers/graphql-go/relay"

   "X/goappsrv/src/gqlSchema"
)

var webAppDist = "goappsrv/dist/webapp"
var graphiQLDist = "goappsrv/dist/graphiql"
var schema *graphql.Schema

func main() {
  schema = graphql.MustParseSchema(gqlSchema.Schema, &gqlSchema.Resolver{})
  r := mux.NewRouter()
  
  r.HandleFunc("/_ah/warmup", WarmupHandler)
  r.HandleFunc("/graphiql", GraphiQLHandler)
  r.HandleFunc("/query", QueryHandler)  
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

func QueryHandler(w http.ResponseWriter, r *http.Request) {
  http.Handle("/query", &relay.Handler{Schema: schema})
}

func GraphiQLHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, graphiQLDist + "/index.html")
}
