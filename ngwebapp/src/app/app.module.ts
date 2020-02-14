import { BrowserModule, Title } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { FontAwesomeModule, FaIconLibrary } from '@fortawesome/angular-fontawesome';
import { faSignIn, faTimesSquare, faPencilAlt, faCheckSquare, faPlusCircle, faCopyright, faBars, faUserAlt, faExclamationTriangle, faFlashlight, faArrowAltSquareLeft, faQrcode, faHome } from '@fortawesome/pro-regular-svg-icons';
import { CookieService } from 'ngx-cookie-service';
import { SocialLoginModule, AuthServiceConfig } from "angularx-social-login";
import { AlertModule, BsDropdownModule, ModalModule, OffcanvasModule } from 'ngx-foundation';
import { ZXingScannerModule } from '@zxing/ngx-scanner';

import { AuthConfig } from './configs';
import { GraphQLModule } from './graphql.module';

import { HomePageComponent, NotFoundComponent, GuestLookupComponent, BarcodeLookupComponent } from './pages';
import { SiteMapComponent, EventScheduleComponent, RoleDocsComponent } from './pages';

import { TopBarComponent, BottomBarComponent } from './shared/layouts';
import { BarcodeScannerComponent } from './shared/tools';
import { GuestInfoComponent, GuestInfoTable, TeamLeadTable } from './shared/displays';

import { LoginModalComponent } from './modals';

import { LoginModalService, BarcodeService } from './services';

import { CurrentUserController } from './controllers';

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    NotFoundComponent,
    GuestLookupComponent,
    BarcodeLookupComponent,
    TopBarComponent,
    BottomBarComponent,
    GuestInfoComponent,
    GuestInfoTable,
    TeamLeadTable,
    LoginModalComponent,
    BarcodeScannerComponent,
    SiteMapComponent,
    EventScheduleComponent,
    RoleDocsComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    ModalModule.forRoot(),
    AlertModule.forRoot(),
    BsDropdownModule.forRoot(),
    OffcanvasModule.forRoot(),
    FontAwesomeModule,
    ZXingScannerModule,
    SocialLoginModule,
    GraphQLModule,
    HttpClientModule,
  ],
  providers: [
    Title,
    CookieService,
    LoginModalService,
    BarcodeService,
    CurrentUserController,
    { provide: AuthServiceConfig, useFactory: AuthConfig }
  ],
  bootstrap: [AppComponent],
  entryComponents: [
    LoginModalComponent,
  ]
})
export class AppModule { 
  constructor(library: FaIconLibrary) {
    library.addIcons(faCopyright);
    library.addIcons(faSignIn);
    library.addIcons(faTimesSquare);
    library.addIcons(faPencilAlt);
    library.addIcons(faCheckSquare);
    library.addIcons(faPlusCircle);
    library.addIcons(faBars);
    library.addIcons(faUserAlt);
    library.addIcons(faExclamationTriangle);
    library.addIcons(faFlashlight);
    library.addIcons(faArrowAltSquareLeft);
    library.addIcons(faQrcode);
    library.addIcons(faHome);
  }
}
