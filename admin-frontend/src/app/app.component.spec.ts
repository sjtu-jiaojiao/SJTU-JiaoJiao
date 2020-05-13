import { TestBed, async } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';

import { ActivitydetailComponent } from './activitydetail/activitydetail.component';
import { NgZorroAntdModule, NzNotificationService } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { DashboardComponent } from './dashboard/dashboard.component';
import { UserComponent } from './user/user.component';
import { LoginComponent } from './login/login.component';
import { InfoDetailComponent } from './infodetail/infodetail.component';
import { WebsiteComponent } from './website/website.component';
import { ActivityComponent } from './activity/activity.component';
import { InfoStatisticComponent } from './info-statistic/info-statistic.component';
import { NgxEchartsModule } from 'ngx-echarts';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InfoComponent } from './info/info.component';
import { InMemoryDataService } from './inmemory-data.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { UserDetailComponent } from './userdetail/userdetail.component';
import { AuthService } from './auth.service';
import { Router } from '@angular/router';
import { SellInfoComponent } from './info/sell-info/sell-info.component';
import { BuyInfoComponent } from './info/buy-info/buy-info.component';
describe('AppComponent', () => {  
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        AppComponent,
        UserDetailComponent,
        UserComponent,
        InfoComponent,
        InfoDetailComponent,
        LoginComponent,
        DashboardComponent,
        WebsiteComponent,
        ActivityComponent,
        InfoStatisticComponent,
        ActivitydetailComponent,
        SellInfoComponent,
        BuyInfoComponent
      ],
      imports: [
        //    DelonAuthModule,
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
    }).compileComponents();
  }));

  it('should deal with log out ', () => {
    const eg = new AppComponent(TestBed.get(Router), TestBed.get(AuthService),TestBed.get(NzNotificationService));
    eg.logout();
  })
});
