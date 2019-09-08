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
import { LoginComponent } from '../login/login.component';
import { InfoDetailComponent } from '../infodetail/infodetail.component';
import { ActivityComponent } from '../activity/activity.component';
import { InfoStatisticComponent } from '../info-statistic/info-statistic.component';
import { NgxEchartsModule } from 'ngx-echarts';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InfoComponent } from '../info/info.component';
import { InMemoryDataService } from '../inmemory-data.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { WebsiteComponent } from './website.component';
import { AuthService } from '../auth.service';
import { BuyInfoComponent } from '../info/buy-info/buy-info.component';
import { SellInfoComponent } from '../info/sell-info/sell-info.component';

describe('WebsiteComponent', () => {
  let component: WebsiteComponent;
  let fixture: ComponentFixture<WebsiteComponent>;

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
    fixture = TestBed.createComponent(WebsiteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    const service: AuthService = TestBed.get(AuthService);
    service.login({ token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjMzMzk3MDYsImlkIjozLCJyb2xlIjoyfQ.woB67gYA8hTMljeg6lqwG_3fSJm4Q7SD6Ln8w2Ol4xk' });

    expect(component).toBeTruthy();
    component.submitForm();
    component.update();
  });
});
