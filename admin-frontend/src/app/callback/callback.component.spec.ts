import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ActivitydetailComponent } from '../activitydetail/activitydetail.component';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from '../app-routing.module';
import { DashboardComponent } from '../dashboard/dashboard.component';
import { AppComponent } from '../app.component';
import { UserDetailComponent } from '../userdetail/userdetail.component';
import { UserComponent } from '../user/user.component';
import { InfoComponent } from '../info/info.component';
import { LoginComponent } from '../login/login.component';
import { InfoDetailComponent } from '../infodetail/infodetail.component';
import { WebsiteComponent } from '../website/website.component';
import { ActivityComponent } from '../activity/activity.component';
import { InfoStatisticComponent } from '../info-statistic/info-statistic.component';
import { NgxEchartsModule } from 'ngx-echarts';
import { HttpClientModule, HttpClient } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InMemoryDataService } from '../inmemory-data.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { CallbackComponent } from './callback.component';
import { Router, ActivatedRoute } from '@angular/router';
import { SocialService } from '@delon/auth';
import { SellInfoComponent } from '../info/sell-info/sell-info.component';
import { BuyInfoComponent } from '../info/buy-info/buy-info.component';

describe('Callback', () => {
  let component: CallbackComponent;
  let fixture: ComponentFixture<CallbackComponent>;

  beforeEach(async(() => {
    const httpClientSpy = jasmine.createSpyObj('HttpClient', ['get','put','post','delete']);
    const socialSpy = jasmine.createSpyObj('SocialService',['callback']);
    TestBed.configureTestingModule({
    providers: [
          {provide: HttpClient, useValue: httpClientSpy},
        {provide: SocialService, useValue: socialSpy}],
      declarations: [
        AppComponent,
        UserDetailComponent,
        UserComponent,
        InfoComponent,
        InfoDetailComponent,
        LoginComponent,
        DashboardComponent,
        CallbackComponent,
        WebsiteComponent,
        ActivityComponent,
        InfoStatisticComponent,
        ActivitydetailComponent,
        SellInfoComponent,
        BuyInfoComponent 
      ],
      imports: [   
        ReactiveFormsModule,
        BrowserModule,
        AppRoutingModule,    
        NgZorroAntdModule,
        FormsModule,
        NgxEchartsModule,
        HttpClientModule,
        HttpClientInMemoryWebApiModule.forRoot(
          InMemoryDataService, { dataEncapsulation: false }),
        BrowserAnimationsModule
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CallbackComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
    component.mockModel('eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjMzMzk3MDYsImlkIjozLCJyb2xlIjoyfQ.woB67gYA8hTMljeg6lqwG_3fSJm4Q7SD6Ln8w2Ol4xk');;
  });
});
