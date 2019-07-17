import { TestBed } from '@angular/core/testing';

import { InMemoryDataService } from './inmemory-data.service';

describe('InmemoryDataService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: InMemoryDataService = TestBed.get(InMemoryDataService);
    expect(service).toBeTruthy();
  });

  it('should create datebase', () => {
    const service: InMemoryDataService = TestBed.get(InMemoryDataService);
    service.createDb();
  });
});
