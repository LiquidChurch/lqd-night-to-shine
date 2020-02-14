import { Component } from '@angular/core';
import { Router, ActivatedRoute } from "@angular/router";
import { BarcodeService } from '../services';

/**
 * GuestLookup component
 */
@Component({
  templateUrl: 'barcodeLookup.component.html',
  styleUrls: ['guestLookup.component.css'],
})

export class BarcodeLookupComponent {
  barcode: string
  
  constructor(private route: ActivatedRoute,
              private router: Router,
              private barcodeService: BarcodeService) {
  }

  ngOnInit() {
    this.route.params.subscribe( params => {
      this.barcode = params.barcode;
    })
    console.log(this.barcode.substring(0,3))
    console.log(this.barcode.length)
    
    if (this.barcode.substring(0,3) !== "rec" || this.barcode.length !== 17) {
      this.router.navigate(['/404']);
    }
    this.barcodeService.success(this.barcode);
  }
  
}
