import { Component, OnInit, Input } from '@angular/core';

/**
 * Generic Display component
 */
@Component({
  selector: 'guestInfoTable',
  templateUrl: 'guestInfoTable.component.html',
  styleUrls: ['guestInfo.component.css'],
})

export class GuestInfoTable implements OnInit {
  @Input() guestInfo: any
  
  name: string
  obj: any
  
  constructor() {}
  
  ngOnInit() {
    console.log("guest info table");
    console.log(this.guestInfo);
    
    this.obj = JSON.parse(this.guestInfo)
    console.log(this.obj['Guest Name']);
  
  }
}