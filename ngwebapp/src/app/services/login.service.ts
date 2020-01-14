import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { BsModalService, BsModalRef } from 'ngx-foundation/modal';

import { LoginModalComponent } from '../modals';

@Injectable()
export class LoginModalService {
  /**
   * ignore
   */
  bsModalRef: BsModalRef;

  /**
   * Component constructor
   */
  constructor(
    private bsModalService: BsModalService,
  ) { }

  /**
   * Open Login Modal Component and subscribe to modal action.
   */
  open(): Observable<string> {
//  open(): void {
    const initialState = {
      class: 'small'
    };
    this.bsModalRef = this.bsModalService.show(LoginModalComponent, { initialState });

    return new Observable<string>(this.getConfirmSubscriber());
  }

  private getConfirmSubscriber() {
    return (observer) => {
      const subscription = this.bsModalService.onHidden.subscribe((reason: string) => {
        if (reason !== null) {
          observer.next(reason);
        } else {
          observer.next(this.bsModalRef.content.action);
        }
        observer.complete();
      });

      return {
        unsubscribe() {
          subscription.unsubscribe();
        }
      };
    };
  }
}
