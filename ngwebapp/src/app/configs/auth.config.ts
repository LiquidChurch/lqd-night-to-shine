import { AuthServiceConfig, GoogleLoginProvider } from 'angularx-social-login';

const config = new AuthServiceConfig([
  {
    id: GoogleLoginProvider.PROVIDER_ID,
    provider: new GoogleLoginProvider('418258098418-24oc4f3fv7rintio2pd0uc0i0v22m8nq.apps.googleusercontent.com')
  }
]);

export function AuthConfig() {
  return config;
}
