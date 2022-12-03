import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomePageComponent, NotFoundComponent, GuestLookupComponent, BarcodeLookupComponent, SiteMapComponent, EventScheduleComponent, RoleDocsComponent, CheckinGuestComponent } from './pages';

const routes: Routes = [
  {path: '', component: HomePageComponent, data: {title: 'Kings & Queens Prom'}},
  {path: 'guest-lookup', component: GuestLookupComponent, data : {title: 'Guest Lookup'}},
  {path: 'site-plan', component: SiteMapComponent, data : {title: 'Event Map'}},
  {path: 'schedule', component: EventScheduleComponent, data : {title: 'Schedule'}},
  {path: 'role-docs', component: RoleDocsComponent, data : {title: 'Role Guides'}},
  {path: '404', component: NotFoundComponent, data: {title: '404'}},
  {path: ':barcode', component: BarcodeLookupComponent, data: {title: 'Guest Info'}},
  {path: ':barcode/checkin', component: CheckinGuestComponent, data: {title: 'Guest Checkin'}},
  {path: '**', redirectTo: '/404'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
