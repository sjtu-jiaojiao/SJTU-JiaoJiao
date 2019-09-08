import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { sellInfo } from 'src/app/entity/info';
import { InfoService } from 'src/app/info.service';
import { Format } from 'src/app/Formatter/format';
import { ActivitydetailComponent } from '../../activitydetail/activitydetail.component';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from '../../app-routing.module';
import { AppComponent } from '../../app.component';
import { UserDetailComponent } from '../../userdetail/userdetail.component';
import { UserComponent } from '../../user/user.component';
import { InfoComponent } from '../../info/info.component';
import { LoginComponent } from '../../login/login.component';
import { InfoDetailComponent } from '../../infodetail/infodetail.component';
import { WebsiteComponent } from '../../website/website.component';
import { ActivityComponent } from '../../activity/activity.component';
import { InfoStatisticComponent } from '../../info-statistic/info-statistic.component';
import { NgxEchartsModule } from 'ngx-echarts';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InMemoryDataService } from '../../inmemory-data.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DashboardComponent } from '../../dashboard/dashboard.component';
import { AuthService } from '../../auth.service';
import { SellInfoComponent } from '../../info/sell-info/sell-info.component';
import { BuyInfoComponent } from '../../info/buy-info/buy-info.component';

describe('SellInfoComponent', () => {
  let c: SellInfoComponent;
  let fixture: ComponentFixture<SellInfoComponent>;

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
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SellInfoComponent);
    c = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(c).toBeTruthy();
    const service: AuthService = TestBed.get(AuthService);
    service.login({ token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjMzMzk3MDYsImlkIjozLCJyb2xlIjoyfQ.woB67gYA8hTMljeg6lqwG_3fSJm4Q7SD6Ln8w2Ol4xk' });
    expect(c.size).toEqual(4);
    expect(c.current).toEqual(1);
    expect(c.getstate(1)).toEqual('待预约');
    expect(c.getstate(2)).toEqual('预约');
    expect(c.getstate(3)).toEqual('完成');
    expect(c.getstate(4)).toEqual('失效');
    expect(c.getstate(5)).toEqual('关闭');
    expect(c.stringToDate(1563134054)).toEqual('2019-07-15 03:54:14');
    c.checkcount();
    c.searchUserID = '';
    c.sellinfos= [new sellInfo];
    c.size = 1;
    c.getcontent();
    c.checkcount();
  });
});
