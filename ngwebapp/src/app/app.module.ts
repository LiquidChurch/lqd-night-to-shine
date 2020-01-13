import { BrowserModule, Title } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { FontAwesomeModule, FaIconLibrary } from '@fortawesome/angular-fontawesome';
import { faBars, faExclamationTriangle } from '@fortawesome/pro-regular-svg-icons';

import { AlertModule, BsDropdownModule, ModalModule } from 'ngx-foundation';

import { HomePageComponent, NotFoundComponent } from './pages';
import { TopBarComponent, BottomBarComponent } from './shared/layouts';

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    NotFoundComponent,
    TopBarComponent,
    BottomBarComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FontAwesomeModule,
    AlertModule.forRoot(),
    BsDropdownModule.forRoot(),
    ModalModule.forRoot(),
  ],
  providers: [
    Title
  ],
  bootstrap: [AppComponent]
})
export class AppModule { 
  constructor(library: FaIconLibrary) {
    library.addIcons(faBars);
    library.addIcons(faExclamationTriangle);
  }
}
