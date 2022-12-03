import { Component, OnInit, Input } from '@angular/core';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

import { CheckinGuestGQL, ItemDetail, CheckinDetail } from '../../graphqls';
/**
 * Generic Display component
 */
@Component({
  selector: 'verifyInfoForm',
  templateUrl: 'verifyInfoForm.component.html',
  styleUrls: ['guestInfo.component.css'],
})

export class VerifyInfoForm implements OnInit {
  @Input() guestInfo: any;
  name: string;
  obj: any;
  contactName: string;
  contactNumber: string;
  guestDetail: Observable<ItemDetail>;
  
  checkinDetail: CheckinDetail;
  
  constructor(private checkGuestGQL: CheckinGuestGQL,
              private router: Router) {}
  
  ngOnInit() {
    this.obj = JSON.parse(this.guestInfo)
    this.contactName = this.obj['Contact Name']
    this.contactNumber = this.obj['Contact #']
    this.checkinDetail = {
      id: this.obj['QR Value'],
      description: ''
    };
  }
  
  confirmCheckin() {
    console.log("checkin confirmed")
    this.checkinDetail.description = JSON.stringify(this.obj)
    console.log("checkinDetail", this.checkinDetail)
    
    this.checkGuestGQL.mutate({
      checkinInput: this.checkinDetail
    },{
      update: (proxy, {data: {postCheckinDetail}}) => {
        console.log(postCheckinDetail);
      }
    })  
    .subscribe(({data: {postCheckinDetail}}) =>  {
      this.router.navigate([this.obj['QR Value']]);
    }, (error) => {
      console.log("error with query", error);
    });
  }
}