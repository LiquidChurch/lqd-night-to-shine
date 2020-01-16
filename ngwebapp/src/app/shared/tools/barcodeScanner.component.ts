import { Component, OnInit, OnChanges } from '@angular/core';

import { BehaviorSubject } from 'rxjs';
import { map } from 'rxjs/operators';

import { Router } from '@angular/router';
import { BarcodeFormat } from '@zxing/library';

import { BarcodeService } from '../../services';

/**
 * Barcode Scanner component
 */
@Component({
  selector: 'barcodeScanner',
  templateUrl: 'barcodeScanner.component.html',
  styleUrls: ['barcodeScanner.component.css'],
})

export class BarcodeScannerComponent {
  /**
   * @ignore
   */
  isDebug = true;

  availableDevices: MediaDeviceInfo[];
  currentDevice: MediaDeviceInfo = null;
  hasDevices: boolean;
  hasPermission: boolean;
  formatsEnabled: BarcodeFormat[] = [
    BarcodeFormat.QR_CODE,
  ];
  
  torchEnabled = false;
  torchAvailable$ = new BehaviorSubject<boolean>(false);
  
  qrResultString: string;
  
  constructor(private barcodeService: BarcodeService) { }

  onHasPermission(has: boolean) {
    this.hasPermission = has;
  }
  
  onCamerasFound(devices: MediaDeviceInfo[]): void {
    console.log("Camera(s) Found");
    this.availableDevices = devices;
    this.hasDevices = Boolean(devices && devices.length);
    this.currentDevice = devices[0];
  }

  onCamerasNotFound(): void {
    console.log("Camera Not Found");
    setTimeout(() => this.hasPermission = true, 10);
  }
  
  onTorchCompatible(isCompatible: boolean): void {
    console.log("Torch Compatiable")
    this.torchAvailable$.next(isCompatible || false);
  }
  
  toggleTorch(): void {
    this.torchEnabled = !this.torchEnabled;
  }
  
  onDeviceSelectChange(selected: string) {
    const device = this.availableDevices.find(x => x.deviceId === selected);
    this.currentDevice = device || null;
  }

  onCodeResult(resultString: string) {
    this.barcodeService.success(resultString);
    console.log(resultString)
  }
}
