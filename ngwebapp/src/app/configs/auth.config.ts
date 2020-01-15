import { AuthServiceConfig, GoogleLoginProvider } from 'angularx-social-login';

const config = new AuthServiceConfig([
  {
    id: GoogleLoginProvider.PROVIDER_ID,
    provider: new GoogleLoginProvider('1039066952699-og870ghai66l563k6orgdas5hii0qdap.apps.googleusercontent.com')
  }
]);

export function AuthConfig() {
  return config;
}
