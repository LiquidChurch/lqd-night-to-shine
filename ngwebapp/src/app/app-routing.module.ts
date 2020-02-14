import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomePageComponent, NotFoundComponent, GuestLookupComponent, BarcodeLookupComponent, SiteMapComponent, EventScheduleComponent, RoleDocsComponent } from './pages';

const routes: Routes = [
  {path: '', component: HomePageComponent, data: {title: 'NTS'}},
  {path: 'guest-lookup', component: GuestLookupComponent, data : {title: 'QR Lookup - NTS'}},
  {path: 'site-plan', component: SiteMapComponent, data : {title: 'Site Plan - NTS'}},
  {path: 'schedule', component: EventScheduleComponent, data : {title: 'Schedule - NTS'}},
  {path: 'role-docs', component: RoleDocsComponent, data : {title: 'Role Docs - NTS'}},
  {path: '404', component: NotFoundComponent, data: {title: '404'}},
  {path: ':barcode', component: BarcodeLookupComponent, data: {title: 'Info - NTS'}},
  {path: '**', redirectTo: '/404'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
