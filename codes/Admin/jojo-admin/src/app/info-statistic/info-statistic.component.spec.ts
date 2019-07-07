import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InfoStatisticComponent } from './info-statistic.component';

describe('InfoStatisticComponent', () => {
  let component: InfoStatisticComponent;
  let fixture: ComponentFixture<InfoStatisticComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InfoStatisticComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InfoStatisticComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
