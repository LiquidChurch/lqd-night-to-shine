import { Component, ViewChild } from '@angular/core';
import { TabsetComponent } from 'ngx-foundation';
/**
 * GuestLookup component
 */
@Component({
  templateUrl: 'roleDocs.component.html',
  styleUrls: ['guestLookup.component.css'],
})

export class RoleDocsComponent {
  constructor() {
  }
  @ViewChild('wteTabs', {static: false}) wteTabs: TabsetComponent;
  
  
  
}
