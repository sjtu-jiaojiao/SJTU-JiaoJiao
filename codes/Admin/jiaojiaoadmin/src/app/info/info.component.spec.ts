
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
import { InfoComponent } from './info.component';
import { ComponentFixture, TestBed, async } from '@angular/core/testing';
import { InfoService } from '../info.service';
import { Info } from '../entity/info';
import { CallbackComponent } from '../callback/callback.component';

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
    const c = new InfoComponent(TestBed.get(InfoService));
    expect(c.size).toEqual(4);
    expect(c.current).toEqual(1);
    expect(c.getstate(0)).toEqual('可预约');
    expect(c.getstate(1)).toEqual('预约中');
    expect(c.getstate(2)).toEqual('已完成');
    expect(c.getstate(3)).toEqual('待评价');
    expect(c.getstate(4)).toEqual('强制结束');
    c.infos =[new Info, new Info, new Info];
    c.count =3;
    c.size = 2;
    c.current = 1;
    c.pageChange(2);
    expect(c.curinfos.length).toEqual(1);
    c.sizeChange(2);
    expect(c.curinfos.length).toEqual(2);
    c.selectTag(['黑']);
    c.selectType(1);
    c.end();
    c.delete(new Info());
  });
});
