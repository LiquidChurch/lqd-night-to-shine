package gqlSchema

import (
)

type Resolver struct{}

var Schema = `
  schema {
    query: Query
  }
  
  # Query requests. Get fields.
  type Query {
    health(): HealthDetail
  }

  type HealthDetail {
    status: String!
  } 
 `
