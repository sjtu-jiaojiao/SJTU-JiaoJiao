import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivitydetailComponent } from '../activitydetail/activitydetail.component';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from '../app-routing.module';
import { DashboardComponent } from '../dashboard/dashboard.component';
import { AppComponent } from '../app.component';
import { UserDetailComponent } from '../userdetail/userdetail.component';
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
import { UserComponent } from './user.component';
import { UserService } from '../user.service';
import { User } from '../entity/user';
import { CallbackComponent } from '../callback/callback.component';

describe('UserComponent', () => {
  let component: UserComponent;
  let fixture: ComponentFixture<UserComponent>;

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
        CallbackComponent,
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
    fixture = TestBed.createComponent(UserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should deal with page', () => {
    const c = new UserComponent(TestBed.get(UserService));
    expect(c.size).toEqual(4);
    expect(c.current).toEqual(1);
    c.users =[new User, new User, new User];
    c.count =3;
    c.size = 2;
    c.current = 1;
    c.threshold = 1;
    c.curusers = [new User()];
    c.pageChange(2);
    expect(c.curusers.length).toEqual(1);
    c.sizeChange(2);
    expect(c.curusers.length).toEqual(2);
    c.forbid();
    c.delete(new User());
    c.searchName="";
    c.searchByName();
    c.searchName="4396";
    c.searchByName();
  });
});
