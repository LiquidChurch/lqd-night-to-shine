import { Component, OnInit, Input } from '@angular/core';

/**
 * Generic Display component
 */
@Component({
  selector: 'teamLeadTable',
  templateUrl: 'teamLeadInfo.component.html',
  styleUrls: ['guestInfo.component.css'],
})

export class TeamLeadTable implements OnInit {
  @Input() teamLead: any
  
  name: string
  obj: any
  
  constructor() {}
  
  ngOnInit() {
    console.log("team lead info table");
    console.log(this.teamLead);
    
    this.obj = JSON.parse(this.teamLead)  
  }
}