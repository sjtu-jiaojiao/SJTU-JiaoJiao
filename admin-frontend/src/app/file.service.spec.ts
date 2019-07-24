import { TestBed } from '@angular/core/testing';

import { FileService } from './file.service';
import { HttpClient } from '@angular/common/http';

describe('FileService', () => {
  beforeEach(() => {
    const httpClientSpy = jasmine.createSpyObj('HttpClient', ['get','put','post','delete']);
    TestBed.configureTestingModule({
    providers: [{provide: HttpClient, useValue: httpClientSpy} ]
  });  

});
  it('should be created', () => {
    const service: FileService = TestBed.get(FileService);
    expect(service).toBeTruthy();
  });
});
