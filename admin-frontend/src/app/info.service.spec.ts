import { TestBed } from '@angular/core/testing';

import { InfoService } from './info.service';
import { HttpClient } from '@angular/common/http';
import { sellInfo, buyInfo } from './entity/info';

describe('InfoService', () => {
  let service: InfoService;
  beforeEach(() => {
    const httpClientSpy = jasmine.createSpyObj('HttpClient', ['get','put','post','delete']);
    TestBed.configureTestingModule({
    providers: [{provide: HttpClient, useValue: httpClientSpy} ]
  });  

});

  it('should be created', () => {
    service = TestBed.get(InfoService);
    expect(service).toBeTruthy();

  });
  });
