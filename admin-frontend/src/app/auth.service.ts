import { Injectable, Inject } from '@angular/core';

import { Observable, of } from 'rxjs';
import { tap, delay } from 'rxjs/operators';
import { DA_SERVICE_TOKEN, ITokenService, JWTTokenModel, ITokenModel } from '@delon/auth';
import { HttpClient } from '@angular/common/http';


@Injectable({
  providedIn: 'root',
})

export class AuthService {
  constructor(private http: HttpClient, @Inject(DA_SERVICE_TOKEN) private tokenService: ITokenService) {
  }

  login(res): void {
    //JWTTokenModely
    //这里会自动补全payload,然而测试不认账，nmdwsm

    this.tokenService.set(res);
    const jwt :JWTTokenModel =  this.tokenService.get<JWTTokenModel>(JWTTokenModel);
    if(jwt.payload.role !== 1) {
      this.logout();
    console.log( 'not admin');
  }
  }

  logout(): void {
    this.tokenService.clear();
  }

}