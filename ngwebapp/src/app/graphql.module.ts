import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { Router } from '@angular/router';

import { ApolloModule, Apollo } from 'apollo-angular';
import { ApolloLink, Observable } from 'apollo-link';
import { HttpLinkModule, HttpLink } from 'apollo-angular-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { onError } from 'apollo-link-error';
import { CookieService } from 'ngx-cookie-service';

const uri = '../query';

@NgModule({
  exports: [
    HttpClientModule,
    ApolloModule,
    HttpLinkModule
  ]
})

export class GraphQLModule {
  constructor(
    apollo: Apollo,
    httpLink: HttpLink,
    private cookie: CookieService,
    private router: Router
  ) {

  const http = httpLink.create({ uri });

  const afterwareLink = new ApolloLink((operation, forward) => {
    return forward(operation).map((response) => {
      const { response: {body} } = operation.getContext();

      if (body.data.sessionDetail.status === 'Authorized') {
        console.log('Authorized User');
        // this.router.navigate(['overview']);
      } else if (body.data.sessionDetail.status === 'Refresh') {
        console.log('Refreshing Session Token')
        const expDate = new Date();
        expDate.setSeconds(expDate.getSeconds() + body.data.sessionDetail.expiration);
        setTimeout(() => this.cookie.set('authToken', body.data.sessionDetail.sessionToken, null, '/', '', true, 'Lax'), 0);
        setTimeout(() => this.cookie.set('authType', 'Session', null, '/', '', true, 'Lax'),0);
        // this.router.navigate(['overview']);
      } else {
        console.log('Unauthorized User');
        setTimeout(() => this.cookie.delete('authToken', '/'), 0);
        setTimeout(() => this.cookie.delete('authType', '/'), 0);
        // this.router.navigate(['']);
      }
      return response;      
    });
  });

  const testLink =
    apollo.create({
      link: afterwareLink.concat(http),
      cache: new InMemoryCache()
    });
  }  
  
}
