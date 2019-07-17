import { TestBed } from '@angular/core/testing';

import { WebsiteService } from './website.service';
import { HttpClient } from '@angular/common/http';

let httpClientSpy: { get: jasmine.Spy, put: jasmine.Spy, post: jasmine.Spy, delete: jasmine.Spy };
describe('UserService', () => {
  beforeEach(() => {
    httpClientSpy = jasmine.createSpyObj('HttpClient', ['get', 'put', 'post', 'delete']);
    TestBed.configureTestingModule({
      providers: [{ provide: HttpClient, useValue: httpClientSpy }]
    });

  });

  it('should be created', () => {
    const service: WebsiteService = TestBed.get(WebsiteService);
    expect(service).toBeTruthy();
  });

});