import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { WebsiteComponent } from './website.component';

describe('WebsiteComponent', () => {
  let component: WebsiteComponent;
  let fixture: ComponentFixture<WebsiteComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ WebsiteComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(WebsiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
