import { CookieService } from 'ngx-cookie-service';
import { Component, TemplateRef, Input, OnInit } from '@angular/core';
import { BsModalRef, BsModalService } from 'ngx-foundation/modal';

import { trigger, state, style, animate, transition} from '@angular/animations';
import { Router } from '@angular/router';
import { AuthService, GoogleLoginProvider } from 'angularx-social-login';

import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

import { CurrentUserController } from '../controllers';
import { UserDetail, PutUserNameGQL } from '../graphqls';


/**
 * Login Modal component
 */
@Component({
    selector: 'loginModal',
    templateUrl: 'login.component.html',
    styleUrls: ['modals.component.css']
})

export class LoginModalComponent implements OnInit {
  /**
   * Current logged in user
   */
  // currentUser: Observable<UserDetail>;
  currentUser: any
  
  /**
   * Flag to edit name
   */
  isEditingName: boolean;

  /**
   * New user display name
   */
  newName: string;

  /**
   * Component constructor
   */
  constructor(
              /**
               * Imports BsModalRef
               */
              public bsModalRef: BsModalRef,
              private router: Router,
              private socialAuthService: AuthService,
              private cookie: CookieService,
              private currentUserController: CurrentUserController,
              private putUserNameGQL: PutUserNameGQL) {}

  /**
   * @ignore
   */
  isDebug = true;

  /**
   * Initialize modal component by loading current user.
   */
  ngOnInit() {
    if (this.isDebug) { console.log('loginModal ngOnInit'); }

    this.currentUser = this.currentUserController.getCurrentUser();
    //this.currentUser = null;
    this.isEditingName = false;
  }

  /**
   * Close the modal.
   */
  closeModal() {
    this.bsModalRef.content.action = 'canceled';
    this.bsModalRef.hide();
  }

  /**
   * Triggers sign in using Google.
   * Sets Google Auth Token.
   * Triggers current user refresh to update name.
   */
  signinWithGoogle () {
    const socialPlatformProvider = GoogleLoginProvider.PROVIDER_ID;

    this.socialAuthService.signIn(socialPlatformProvider).then(
      (userData) => {
          console.log('socialAuthService userData', userData);
          setTimeout(() => this.cookie.set('authToken', userData.idToken, null, '/', '', true, 'Lax'), 0);
          this.cookie.set('authType', 'Google', null, '/', '', true, 'Lax');
          setTimeout(() => this.currentUserController.currentUserRefresh(), 300);
          setTimeout(() => {
              this.bsModalRef.content.action = 'logged in';
              this.bsModalRef.hide();
              }, 300);
    });
  }

  /**
   * Triggers sign out.
   */
  signOut(): void {
    this.bsModalRef.content.action = 'logged out';
    this.cookie.delete('authToken', '/');
    this.cookie.delete('authType', '/');
    this.socialAuthService.signOut();
    this.isEditingName = false;
    setTimeout(() => this.currentUserController.currentUserRefresh(), 200);
  }

  /**
   * Triggers display name edit.
   */
  editNameOn(): void {
    this.currentUser.subscribe((val) => {
      this.newName = val.name;
    });
    this.isEditingName = true;
  }

  /**
   * Cancels display name edit.
   */
  editNameCancel(): void {
    setTimeout(() => this.isEditingName = false, 40);
    this.currentUser.subscribe((val) => {
      this.newName = val.name;
    });
  }

  /**
   * Saves new display name if not the same as current display name.
   */
  editNameSubmit(): void {
    /**
    this.currentUser.subscribe((val) => {
      if (this.newName.trim() !== val.name) {
        this.putUserNameGQL
            .mutate({
                updateName: this.newName.trim(),
            })
            .subscribe(result => {
            });

          setTimeout(() => this.isEditingName = false, 40);
      } else {
        this.editNameCancel();
      }
    });
    */
  }
}
