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
  templateUrl: 'guestInfo.component.html',
  styleUrls: ['guestInfo.component.css'],
})

export class GuestInfoComponent implements OnInit {
  private subscription: Subscription
  barcode: any;
  refId: string;
  itemFilter: ItemFilter;
  itemType: string;
  guestDetail: Observable<ItemDetail>;
  isOffcanvas = true;
  itemFound = true;
  
  constructor(private barcodeService: BarcodeService,
              private getGuestGQL: GetGuestGQL) { }
  
  ngOnInit() {
    this.itemFilter = {
      id: '',
      idType: 'Ext',
      type: '',
      parentId: ''
    };
    
    this.subscription = this.barcodeService.getBarcode().subscribe((barcode) => {
      this.barcode = barcode; 
      this.barcodeLookup();
    })
    
    
  }
  
  barcodeLookup() {
    console.log("GuestInfo DoCheck entered");
    
    if (this.barcode != null) {
      this.itemFilter.id = this.barcode.text;
      
      this.guestDetail = this.getGuestGQL.watch({
        lookup: this.itemFilter
        })
        .valueChanges
        .pipe(
          map(result => result.data.getItem)
        );

      this.guestDetail.subscribe(val => {
        console.log(val);
          if (val.type == "") {
            console.log("no object found");
            this.itemFound = false;
          } else if (val.type == "guest") {
            console.log("Item is Guest");
            this.itemType = "guest"
          } else if (val.type == "teamlead") {
            console.log("Item is Team Lead");
            this.itemType = "teamlead"
          };
        }
      )
      
      this.isOffcanvas = false;
    }
  }
  
  
  closePanel() {
    this.isOffcanvas = true;
  }
}