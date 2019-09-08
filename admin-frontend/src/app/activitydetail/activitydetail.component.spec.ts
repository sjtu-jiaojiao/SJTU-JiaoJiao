import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ActivitydetailComponent } from './activitydetail.component';
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
import { HttpClientModule } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InMemoryDataService } from '../inmemory-data.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AuthService } from '../auth.service';
import { SellInfoComponent } from '../info/sell-info/sell-info.component';
import { BuyInfoComponent } from '../info/buy-info/buy-info.component';
import { DelonAuthModule } from '@delon/auth';

describe('ActivitydetailComponent', () => {
  let component: ActivitydetailComponent;
  let fixture: ComponentFixture<ActivitydetailComponent>;

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
    fixture = TestBed.createComponent(ActivitydetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {

    expect(component).toBeTruthy();
    component.act ={id: 1, title: '新人体验',
    description: '用户首次发布交易信息可进行申报,要求多媒体数>=1,获得特别加权推荐',
    releaseTime: 1567957148885, validTime: 1567957148885, 
    releaseDate: new Date(), validDate:new Date(),weight: 3,
    pic:1,isNew:true
    };
    component.save();
    component.goBack();
  });
});
