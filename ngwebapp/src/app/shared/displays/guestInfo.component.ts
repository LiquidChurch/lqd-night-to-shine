import { Component, OnInit, OnChanges, DoCheck } from '@angular/core';
import { Observable, Subscription } from 'rxjs';
import { map } from 'rxjs/operators';

import { BarcodeService } from '../../services';
import { ItemDetail, ItemFilter, GetGuestGQL } from '../../graphqls';
/**
 * Guest Info Display component
 */
@Component({
  selector: 'guestInfo',
  templateUrl: 'guestInfo.component.html'
})

export class GuestInfoComponent implements OnInit {
  private subscription: Subscription
  barcode: any;
  refId: string;
  itemFilter: ItemFilter;
  guestDescription: any;
  private guestDetail: Observable<ItemDetail>;
  
  constructor(private barcodeService: BarcodeService,
              private getGuestGQL: GetGuestGQL) { }
  
  ngOnInit() {
    this.itemFilter = {
      id: '',
      idType: 'Ext',
      type: 'guest',
      parentId: ''
    };
    
    this.subscription = this.barcodeService.getBarcode().subscribe((barcode) => {
      console.log('barcode', barcode);
      
      this.barcode = barcode;
      this.guestLookup();
      // setTimeout(() => console.log("guestDescription1", this.guestDetail.Description), 100);
      setTimeout(() => console.log("guestDescription2", this.guestDescription), 100);
    })
  }
  
  guestLookup() {
    console.log("GuestInfo DoCheck entered");
    
    if (this.barcode != null) {
      this.itemFilter.id = this.barcode.text;
      
      this.guestDescription = this.getGuestGQL.watch({
        lookup: this.itemFilter
        })
        .valueChanges
        .pipe(
          map(result => result.data.getItem.description)
        );
    }

  }
  
}