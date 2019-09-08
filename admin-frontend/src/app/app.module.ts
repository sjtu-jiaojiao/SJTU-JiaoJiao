import { BrowserModule } from '@angular/platform-browser';
import { NgModule, ModuleWithProviders } from '@angular/core';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InMemoryDataService }  from './inmemory-data.service';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgZorroAntdModule, NZ_I18N, zh_CN, NZ_ICONS } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgxEchartsModule } from 'ngx-echarts';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { UserDetailComponent } from './userdetail/userdetail.component';
import { UserComponent } from './user/user.component';
import { InfoComponent } from './info/info.component';
import { InfoDetailComponent } from './infodetail/infodetail.component';
import { LoginComponent } from './login/login.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { WebsiteComponent } from './website/website.component';
import { ActivityComponent } from './activity/activity.component';
import { InfoStatisticComponent } from './info-statistic/info-statistic.component';
import { ActivitydetailComponent } from './activitydetail/activitydetail.component';
import { IconDefinition } from '@ant-design/icons-angular';
import { DelonAuthModule, JWTInterceptor, DA_STORE_TOKEN, MemoryStore } from '@delon/auth';



import { DashboardOutline, UserOutline, LeftCircleOutline ,SaveOutline,  ProfileOutline, TransactionOutline, ContactsOutline,
BulbOutline, LoginOutline, KeyOutline, DeleteOutline, SearchOutline, ControlOutline, LockOutline, LogoutOutline, CopyrightOutline } from '@ant-design/icons-angular/icons';
import { SellInfoComponent } from './info/sell-info/sell-info.component';
import { BuyInfoComponent } from './info/buy-info/buy-info.component'
const icons: IconDefinition[] = [ControlOutline,LeftCircleOutline ,SaveOutline, DashboardOutline, UserOutline, ProfileOutline, TransactionOutline, ContactsOutline,
  BulbOutline, LoginOutline, LockOutline, KeyOutline, DeleteOutline, SearchOutline ,LogoutOutline, CopyrightOutline];
registerLocaleData(zh);

export class DelonModule {
  static forRoot(): ModuleWithProviders {
    return {
      ngModule: DelonModule,
      providers: [
        { provide: DA_STORE_TOKEN, useClass: MemoryStore }
      ]
    };
  }
}

@NgModule({
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
    DelonAuthModule,
    ReactiveFormsModule,
    BrowserModule,
    AppRoutingModule,    
    NgZorroAntdModule.forRoot(),
    FormsModule,
    NgxEchartsModule,
    HttpClientModule,
    HttpClientInMemoryWebApiModule.forRoot(
    InMemoryDataService, { dataEncapsulation: false,passThruUnknownUrl: true }),
    BrowserAnimationsModule
  ],
  providers: [    
    { provide: NZ_I18N, useValue: zh_CN } , { provide: NZ_ICONS, useValue: icons },{ provide: HTTP_INTERCEPTORS, useClass: JWTInterceptor, multi: true}],
  bootstrap: [AppComponent]
})
export class AppModule { }
