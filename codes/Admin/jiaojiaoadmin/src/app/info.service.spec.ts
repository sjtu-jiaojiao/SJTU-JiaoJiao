import { TestBed } from '@angular/core/testing';

import { InfoService } from './info.service';
import { HttpClient } from '@angular/common/http';
import { Info } from './entity/info';

describe('InfoService', () => {
  let service: InfoService;
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    service = TestBed.get(InfoService);
    expect(service).toBeTruthy();
  });

  it('#getObservableValue should return value from observable',
    (done: DoneFn) => {
    service.getInfo('4396').subscribe(value => {
      expect(value).toBe(
        (<Info>{id: '4396', source: '4396', type: 0, price: 100, 
          tags: ['黑色','塑料','垃圾桶'], time: '2019-01-01', state: 0, count: 2, intro: '黑色塑料分类垃圾桶'}));
      done();
    });
  });
 
  });
