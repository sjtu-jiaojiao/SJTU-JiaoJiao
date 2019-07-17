
import { ActivitydetailComponent } from '../activitydetail/activitydetail.component';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from '../app-routing.module';
import { DashboardComponent } from '../dashboard/dashboard.component';
import { AppComponent } from '../app.component';
import { UserDetailComponent } from '../userdetail/userdetail.component';
import { UserComponent } from '../user/user.component';
import { LoginComponent } from '../login/login.component';
import { InfoDetailComponent } from '../infodetail/infodetail.component';
import { WebsiteComponent } from '../website/website.component';
import { ActivityComponent } from '../activity/activity.component';
import { InfoStatisticComponent } from '../info-statistic/info-statistic.component';
import { NgxEchartsModule } from 'ngx-echarts';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InMemoryDataService } from '../inmemory-data.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { InfoComponent, Format } from './info.component';
import { ComponentFixture, TestBed, async } from '@angular/core/testing';
import { InfoService } from '../info.service';
import { Info } from '../entity/info';
import { CallbackComponent } from '../callback/callback.component';
import { AuthService } from '../auth.service';

describe('InfoComponent', () => {
  let component: InfoComponent;
  let fixture: ComponentFixture<InfoComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        AppComponent,
        UserDetailComponent,
        UserComponent,
        CallbackComponent,
        InfoComponent,
        InfoDetailComponent,
        LoginComponent,
        DashboardComponent,
        WebsiteComponent,
        ActivityComponent,
        InfoStatisticComponent,
        ActivitydetailComponent
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
    fixture = TestBed.createComponent(InfoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should deal with page', () => {
    const service: AuthService = TestBed.get(AuthService);
    service.login({ token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjMzMzk3MDYsImlkIjozLCJyb2xlIjoyfQ.woB67gYA8hTMljeg6lqwG_3fSJm4Q7SD6Ln8w2Ol4xk' });

    const c = new InfoComponent(TestBed.get(InfoService));
    expect(c.size).toEqual(4);
    expect(c.current).toEqual(1);
    expect(c.getstate(1)).toEqual('待预约');
    expect(c.getstate(2)).toEqual('预约');
    expect(c.getstate(3)).toEqual('完成');
    expect(c.getstate(4)).toEqual('失效');
    expect(Format(new Date(1563134054),'yyyy')).toEqual('1970');
    expect(c.stringToDate(new Date(1563134054))).toEqual('1970-01-19 10:12:14')
    c.checkcount();
    c.searchUser = '';
    c.searchByUser();
    c.searchUser = '1';
    c.onChange();
    c.infos= [new Info];
    c.size = 1;
    c.checkcount();
  });
});
