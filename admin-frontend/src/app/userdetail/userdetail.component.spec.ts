import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivitydetailComponent } from '../activitydetail/activitydetail.component';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from '../app-routing.module';
import { DashboardComponent } from '../dashboard/dashboard.component';
import { AppComponent } from '../app.component';
import { UserComponent } from '../user/user.component';
import { LoginComponent } from '../login/login.component';
import { InfoDetailComponent } from '../infodetail/infodetail.component';
import { WebsiteComponent } from '../website/website.component';
import { ActivityComponent } from '../activity/activity.component';
import { InfoStatisticComponent } from '../info-statistic/info-statistic.component';
import { NgxEchartsModule } from 'ngx-echarts';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InfoComponent } from '../info/info.component';
import { InMemoryDataService } from '../inmemory-data.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { UserDetailComponent, fFormatter } from './userdetail.component';
import { CallbackComponent } from '../callback/callback.component';
import { AuthService } from '../auth.service';
import { User } from '../entity/user';
import { SellInfoComponent } from '../info/sell-info/sell-info.component';
import { BuyInfoComponent } from '../info/buy-info/buy-info.component';
import { DelonAuthModule } from '@delon/auth';

describe('UserdetailComponent', () => {
  let component: UserDetailComponent;
  let fixture: ComponentFixture<UserDetailComponent>;

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
        CallbackComponent,
        WebsiteComponent,
        ActivityComponent,
        InfoStatisticComponent,
        ActivitydetailComponent,
        SellInfoComponent,
        BuyInfoComponent
      ],
      imports: [   
    //import { BuyInfoComponent } from './../info/buy-info/buy-info.component';
    DelonAuthModule,
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
    fixture = TestBed.createComponent(UserDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
    const service: AuthService = TestBed.get(AuthService);
    service.login({ token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjMzMzk3MDYsImlkIjozLCJyb2xlIjoyfQ.woB67gYA8hTMljeg6lqwG_3fSJm4Q7SD6Ln8w2Ol4xk' });

    component.user= new User();
    component.forbid = true;
    component.save();
    
    component.forbid = false;
    component.save();
    
    component.goBack();
    expect(component.getstate(1)).toEqual('待预约');
    expect(component.getstate(2)).toEqual('预约');
    expect(component.getstate(3)).toEqual('完成');
    expect(component.getstate(4)).toEqual('失效');
    expect(fFormatter({value: 5, name: 'Jan'})).toEqual('5 activities in Jan');
    expect(component.stringToDate(new Date(1563134054))).toEqual('1970-01-19 10:12:14');
    expect(component.typeof(1)).toEqual('number');
    component.infos = [];
    component.graph();
  });
});
