package gqlSchema

import (
)

type Resolver struct{}

var Schema = `
  schema {
    query: Query
    mutation: Mutation
  }
  
  # Query requests. Get fields.
  type Query {
    # Simple Health Check
    health(): HealthDetail
    # Get Session Details
    sessionDetail(): SessionDetail
    # Get Current User Details
    getUser(): UserDetail
  }

  # Mutation requsts. Upsert fields
  type Mutation {
    # Get Session Details
    sessionDetail(): SessionDetail
    # Refresh Current User Details
    getUser(): UserDetail
  }

  type HealthDetail {
    status: String!
  }

  # Session information
  type SessionDetail {
    # Session JWT token.
    sessionToken: String!
    # Application User ID
    userID: String!
    # User Status
    status: String!
    # Session Expiration timer in sec
    expiration: Int!
  }

  # User information
  type UserDetail {
    # Application User ID
    id: String!
    # User ID
    uid: String!
    # User display name
    name: String!
    # User email address
    email: String!
    # User picture link
    pictureURL: String!
    # User company
    company: String!
    # User Role
    role: String!
  }
 `
