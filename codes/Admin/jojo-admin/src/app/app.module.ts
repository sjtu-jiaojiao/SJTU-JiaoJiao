import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgZorroAntdModule, NZ_I18N, zh_CN } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
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
import { PasswordComponent } from './password/password.component';
import { InfoStatisticComponent } from './info-statistic/info-statistic.component';
import { ActivitydetailComponent } from './activitydetail/activitydetail.component';

registerLocaleData(zh);

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
    PasswordComponent,
    InfoStatisticComponent,
    ActivitydetailComponent
  ],
  imports: [
    ReactiveFormsModule,
    BrowserModule,
    AppRoutingModule,
    NgZorroAntdModule,
    FormsModule,
    NgxEchartsModule,
    HttpClientModule,
    BrowserAnimationsModule
  ],
  providers: [{ provide: NZ_I18N, useValue: zh_CN }],
  bootstrap: [AppComponent]
})
export class AppModule { }
