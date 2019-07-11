import { TestBed, async, inject } from '@angular/core/testing';

import { AuthGuard } from './auth.guard';
import { AuthService } from '../auth.service';
import { Router, RouterStateSnapshot } from '@angular/router';
import { HttpClient } from '@angular/common/http';

describe('AuthGuard', () => {
  const routerSpy = jasmine.createSpyObj('Router', ['navigate']);
  const httpClientSpy = jasmine.createSpyObj('HttpClient', ['get','put','post','delete']);
  beforeEach(() => {
    TestBed.configureTestingModule({
    providers: [AuthGuard , AuthService, {provide: Router, useValue: routerSpy},
      {provide: HttpClient, useValue: httpClientSpy}
  ]
  });  

});
  it('should ...', inject([AuthGuard], (guard: AuthGuard) => {
    expect(guard).toBeTruthy();
  }));

  it('should guard', inject([AuthGuard], (guard: AuthGuard) => {
    expect(guard.checkLogin('/users')).toEqual(false);
    const service: AuthService = TestBed.get(AuthService);
    service.isLoggedIn = true;
    expect(guard.checkLogin('/users')).toEqual(true);
    guard.canActivate(null,null);
  }));
});
