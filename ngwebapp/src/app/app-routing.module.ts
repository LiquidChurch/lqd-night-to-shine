import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomePageComponent, NotFoundComponent } from './pages';

const routes: Routes = [
  {path: '', component: HomePageComponent, data: {title: 'Home'}},
  {path: '404', component: NotFoundComponent, data: {title: '404'}},
  {path: '**', redirectTo: '/404'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
