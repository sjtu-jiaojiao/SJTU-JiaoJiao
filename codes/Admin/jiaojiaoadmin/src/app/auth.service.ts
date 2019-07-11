import { Injectable, Inject } from '@angular/core';

import { Observable, of } from 'rxjs';
import { tap, delay } from 'rxjs/operators';
import { DA_SERVICE_TOKEN, ITokenService, JWTTokenModel, ITokenModel } from '@delon/auth';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})

export class AuthService {
  isLoggedIn = false;
  private authUrl = 'auth';  // URL to web api
  constructor(private http: HttpClient, @Inject(DA_SERVICE_TOKEN) private tokenService: ITokenService) {
  }
  // store the URL so we can redirect after logging in
  redirectUrl: string = '/dashboard';

  login(res): void {
    //JWTTokenModely
    this.isLoggedIn = true;
    this.tokenService.set(res);
  }

  logout(): void {
    this.isLoggedIn = false;
  }
}