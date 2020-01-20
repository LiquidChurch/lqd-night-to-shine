import { Injectable } from '@angular/core';
import { Query, Mutation } from 'apollo-angular';
import gql from 'graphql-tag';
import { SessionDetail } from '../graphqls';

export interface ItemDetail {
  id: string;
  parentId: string;
  type: string;
  refId: string;
  name: string;
  description: string;
  color: string;
  webURL: string;
}

export interface ItemFilter {
  id: string;
  idType: string;
  type: string;
  parentId: string;
}

export interface GetGuest {
  getItem: ItemDetail;
  sessionDetail: SessionDetail;
}
@Injectable({
  providedIn: 'root',
})
export class GetGuestGQL extends Query<GetGuest> {
  document = gql`
    query GuestLookup($lookup: ItemFilter!) {
      getItem(lookup: $lookup) {
        id
        parentId
        type
        name
        extId
        description
        webURL
        color
      }
      sessionDetail{
        sessionToken
        status
        userID
        expiration
      }
    }`;
}
