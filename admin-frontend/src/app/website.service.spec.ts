import { TestBed } from '@angular/core/testing';

import { WebsiteService } from './website.service';
import { HttpClient } from 'selenium-webdriver/http';

describe('WebsiteService', () => {
  let service: WebsiteService;
  beforeEach(() => {
    const httpClientSpy = jasmine.createSpyObj('HttpClient', ['get','put','post','delete']);
    TestBed.configureTestingModule({
    providers: [{provide: HttpClient, useValue: httpClientSpy} ]
  });  

  it('should be created', () => {
    const service: WebsiteService = TestBed.get(WebsiteService);
    expect(service).toBeTruthy();
  });
});
