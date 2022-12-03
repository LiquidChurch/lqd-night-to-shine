import { Injectable } from '@angular/core';
import { Query, Mutation } from 'apollo-angular';
import gql from 'graphql-tag';
import { ItemDetail, SessionDetail } from '../graphqls';

export interface CheckinDetail {
  id: string;
  description: string;
}

export interface CheckinGuest {
  postCheckinDetail: ItemDetail
  sessionDetail: SessionDetail;
}
@Injectable({
  providedIn: 'root',
})
export class CheckinGuestGQL extends Mutation<CheckinGuest> {
  document = gql`
    mutation CheckinGuest($checkinInput: CheckinDetail!) {
      postCheckinDetail(checkinInput: $checkinInput) {
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