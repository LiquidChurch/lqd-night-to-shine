import { BrowserModule, Title } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { FontAwesomeModule, FaIconLibrary } from '@fortawesome/angular-fontawesome';
import { faBars, faExclamationTriangle, faSignIn, faTimesSquare, faPencilAlt, faCheckSquare } from '@fortawesome/pro-regular-svg-icons';
import { CookieService } from 'ngx-cookie-service';
import { SocialLoginModule, AuthServiceConfig } from "angularx-social-login";
import { AlertModule, BsDropdownModule, ModalModule } from 'ngx-foundation';

import { AuthConfig } from './configs';
import { GraphQLModule } from './graphql.module';

import { HomePageComponent, NotFoundComponent } from './pages';
import { TopBarComponent, BottomBarComponent } from './shared/layouts';
import { LoginModalComponent } from './modals';

import { LoginModalService } from './services';

import { CurrentUserController } from './controllers';

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    NotFoundComponent,
    TopBarComponent,
    BottomBarComponent,
    LoginModalComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    ModalModule.forRoot(),
    AlertModule.forRoot(),
    BsDropdownModule.forRoot(),
    FontAwesomeModule,
    SocialLoginModule,
    GraphQLModule,
    HttpClientModule,
  ],
  providers: [
    Title,
    CookieService,
    LoginModalService,
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
    library.addIcons(faBars);
    library.addIcons(faExclamationTriangle);
    library.addIcons(faSignIn);
    library.addIcons(faTimesSquare);
    library.addIcons(faPencilAlt);
    library.addIcons(faCheckSquare);
  }
}
