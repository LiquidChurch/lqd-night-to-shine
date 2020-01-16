import { Injectable } from '@angular/core';
import { Router, NavigationStart } from '@angular/router';
import { Observable, Subject } from 'rxjs';

@Injectable()
export class BarcodeService {
  private subject = new Subject<any>();
  private keepAfterNavigationChange = false;

  constructor(private router: Router) {
    // clear alert message on route change
    router.events.subscribe(event => {
      if (event instanceof NavigationStart) {
        if (this.keepAfterNavigationChange) {
          // only keep for a single location change
          this.keepAfterNavigationChange = false;
        } else {
          // clear alert
          this.subject.next();
        }
      }
    });
  }

  success(barcode: string, keepAfterNavigationChange = false) {
    this.keepAfterNavigationChange = keepAfterNavigationChange;
    setTimeout(() => this.subject.next({ type: 'success', text: barcode }), 0);
  }

  error(barcode: string, keepAfterNavigationChange = false) {
    console.log(barcode);
    this.keepAfterNavigationChange = keepAfterNavigationChange;
    setTimeout(() => this.subject.next({ type: 'alert', text: barcode }), 0);
  }

  getBarcode(): Observable<any> {
    return this.subject.asObservable();
  }
}
