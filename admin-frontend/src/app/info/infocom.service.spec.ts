import { TestBed } from '@angular/core/testing';

import { InfoComService } from './infocom.service';

describe('InfoComService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: InfoComService = TestBed.get(InfoComService);
    expect(service).toBeTruthy();
  });
});
