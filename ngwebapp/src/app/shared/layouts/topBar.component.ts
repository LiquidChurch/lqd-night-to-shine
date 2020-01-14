import { Component, Input, OnInit } from '@angular/core';
import { Observable, Subscription } from 'rxjs';

import { UserDetail } from '../../graphqls';
import { LoginModalComponent } from '../../modals';
import { CurrentUserController } from '../../controllers';
import { LoginModalService } from '../../services';
/**
 * Top Bar Component
 */
@Component({
  selector: 'topbar',
  templateUrl: 'topBar.component.html',
  styleUrls: ['layouts.component.css'],
})

export class TopBarComponent implements OnInit {
  /**
   * Topbar Title
   */
  @Input() topbarTitle: string;

  /**
   * CurrentUser Object
   */
  currentUser: Observable<UserDetail>;

  /**
   * @ignore
   */
  private subscription: Subscription;
  
  constructor(private loginModalService: LoginModalService,
              private currentUserController: CurrentUserController) {
  }
  
  /**
   * Initialize Top Bar Component and calls checkUser()
   */
  ngOnInit() {
    this.currentUserController.updateCurrentUser();
    this.currentUser = this.currentUserController.getCurrentUser();
  }

  /**
   * Opens [LoginModalComponent]{@link LoginModalComponent}
   */
  openModal(): void {
    this.loginModalService.open().subscribe((action: string) => {
    });
  }

}
