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
    # Get Item By RefId
    getItem(lookup: ItemFilter!): ItemDetail
  }

  # Mutation requsts. Upsert fields
  type Mutation {
    # Get Session Details
    sessionDetail(): SessionDetail
    # Refresh Current User Details
    getUser(): UserDetail
    # Create Item Detail endpoint
    postItemDetail(postItem: ItemInput!): ItemDetail
    # Get Guest update
    updateGuests(airTableId: String!): [ItemDetail]
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

  # Item Detail
  type ItemDetail{
    # Item ID
    id: String!
    # Item Parent ID
    parentId: String!
    # Item Type
    type: String!
    # Item Name
    name: String!
    # Item External ID
    extId: String!
    # Item Description
    description: String!
    # Item PictureURL
    webURL: String!
    # Item Background Color
    color: String!
  }

  # Get Item Filter 
  input ItemFilter {
    id: String!
    type: String!
    parentId: String!
    idType: String!
  }

  # Post Item Input
  input ItemInput {
    id: String!
    type: String!
    name: String!
    description: String!
    webURL: String!
    color: String!
    extId: String!
  }
 `
