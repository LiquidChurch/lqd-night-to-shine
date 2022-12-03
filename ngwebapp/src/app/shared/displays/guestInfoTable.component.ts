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
    this.obj = JSON.parse(this.guestInfo)
  }
}