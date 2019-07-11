import { TestBed } from '@angular/core/testing';

import { AuthService } from './auth.service';
import { HttpClient } from '@angular/common/http';

describe('AuthService', () => {

  beforeEach(() => {
    const httpClientSpy = jasmine.createSpyObj('HttpClient', ['get', 'put', 'post', 'delete']);
    TestBed.configureTestingModule({
      providers: [{ provide: HttpClient, useValue: httpClientSpy }]
    });

  });

  it('should be created', () => {
    const service: AuthService = TestBed.get(AuthService);
    expect(service).toBeTruthy();
  });


  it('should work normal', () => {
    const service: AuthService = TestBed.get(AuthService);
    expect(service.isLoggedIn).toEqual(false);
    service.login({ token: '123' });
    service.logout();
  });

});
