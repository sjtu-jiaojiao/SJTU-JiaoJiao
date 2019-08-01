import { TestBed } from '@angular/core/testing';

import { TransactionService } from './transaction.service';
import { HttpClient } from '@angular/common/http';
import { Transaction } from './entity/transaction';

describe('TransactionService', () => {
  beforeEach(() => {
    const httpClientSpy = jasmine.createSpyObj('HttpClient', ['get','put','post','delete']);
    TestBed.configureTestingModule({
    providers: [{provide: HttpClient, useValue: httpClientSpy} ]
  });
});

  it('should be created', () => {
    const service: TransactionService = TestBed.get(TransactionService);
    expect(service).toBeTruthy();
  });
});
