import { Injectable } from '@angular/core';
import { Query, Mutation } from 'apollo-angular';
import gql from 'graphql-tag';

export interface UserDetail {
  id: string;
  uid: string;
  name: string;
  email: string;
  pictureURL: string;
//  roles: ItemDetail;
}

export interface SessionDetail {
  sessionID: string;
  status: string;
  userID: string;
  expiration: number;
}

export interface UserResponse {
  getUser: UserDetail;
  sessionDetail: SessionDetail;
}

@Injectable({
  providedIn: 'root',
})
export class GetUserGQL extends Query<UserResponse> {
  document = gql`
    query GetUser {
      getUser {
        id
        uid
        name
        email
        pictureURL
      }
      sessionDetail {
        sessionToken
        status
        userID
        expiration
      }
    }`;
}

@Injectable({
  providedIn: 'root',
})
export class RefreshUserGQL extends Mutation<UserResponse> {
  document = gql`
    mutation refreshUser {
      getUser {
        id
        uid
        name
        email
        pictureURL
      }
      sessionDetail{
        sessionToken
        status
        userID
        expiration
      }
    }`;
}

@Injectable({
  providedIn: 'root',
})
export class PutUserNameGQL extends Mutation {
  document = gql`
    mutation putUserName($updateName: String!) {
      putUserName(name: $updateName) {
        id
        uid
        name
        email
        pictureURL
      }
      sessionDetail{
        sessionToken
        status
        userID
        expiration
      }
    }`;
}
