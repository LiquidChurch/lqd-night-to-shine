import { Component, Input } from '@angular/core';

/**
 * Top Bar Component
 */
@Component({
  selector: 'topbar',
  templateUrl: 'topBar.component.html',
  styleUrls: ['layouts.component.css'],
})

export class TopBarComponent {
  /**
   * Topbar Title
   */
  @Input() topbarTitle: string;

  constructor() {
  }
}
