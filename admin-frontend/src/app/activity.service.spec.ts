import { TestBed } from '@angular/core/testing';

import { ActivityService } from './activity.service';
import { HttpClient } from '@angular/common/http';

describe('ActivityService', () => {
  beforeEach(() => {
    const httpClientSpy = jasmine.createSpyObj('HttpClient', ['get','put','post','delete']);
    TestBed.configureTestingModule({
    providers: [{provide: HttpClient, useValue: httpClientSpy} ]
  });  

});
  it('should be created', () => {
    const service: ActivityService = TestBed.get(ActivityService);
    expect(service).toBeTruthy();
  });
});
