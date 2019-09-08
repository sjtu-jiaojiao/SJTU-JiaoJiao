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
import { WebsiteComponent } from '../website/website.component';
import { ActivityComponent } from '../activity/activity.component';
import { InfoStatisticComponent, fdgFormatter } from './info-statistic.component';
import { NgxEchartsModule } from 'ngx-echarts';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InfoComponent } from '../info/info.component';
import { InMemoryDataService } from '../inmemory-data.service';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import 'echarts-wordcloud';
import { SellInfoComponent } from '../info/sell-info/sell-info.component';
import { BuyInfoComponent } from '../info/buy-info/buy-info.component';
import { buyInfo, sellInfo } from '../entity/info';
import { Transaction } from '../entity/transaction';
describe('InfoStatisticComponent', () => {
  let component: InfoStatisticComponent;
  let fixture: ComponentFixture<InfoStatisticComponent>;

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
    fixture = TestBed.createComponent(InfoStatisticComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
    component.bi = [ new buyInfo()];
    component.si = [ new sellInfo()];
    expect(component.fmt(new Date(1563134054000))).toEqual('2019/7/15');
    component.pl=false;
    component.pauseLine();
    expect(component.pl).toEqual(true);
    component.clickForce({dataType: 'node', data:{name: '1'}});
    component.tr= [new Transaction()];
    component.bi = [new buyInfo()];
    component.si = [new sellInfo()];
    expect(fdgFormatter({dataType: 'node',data:{name: '1', value: '2'}})).toEqual('1 has completed 2 transaction');   
    expect(fdgFormatter({dataType: 'edge',data:{source: '1', target: '2', value: '2'}})).
    toEqual('1 has selled 2 goods to 2');
    component.onBrushSelected({batch:[{selected: [{dataIndex:[]}]}]});
    component.tr= [
       {transactionID: 2, infoID: 1, category: 2, fromUserID: 10003, toUserID: 3,createTime :1 ,status:6},
      {transactionID: 3, infoID: 1, category: 3, fromUserID: 10003, toUserID: 3,createTime :1 ,status:6},
      {transactionID: 4, infoID: 2, category: 2, fromUserID: 10003, toUserID: 3,createTime :1 ,status:6},
       {transactionID: 5, infoID: 6, category: 2, fromUserID: 10003, toUserID: 10002,createTime :1 ,status:6},
       {transactionID: 6, infoID: 6, category: 1, fromUserID: 10003, toUserID: 1366,createTime :1 ,status:6},
       {transactionID: 7, infoID: 200001, category: 1, fromUserID: 3, toUserID: 10003,createTime :1 ,status:6},
       {transactionID: 8, infoID: 200002, category: 1, fromUserID: 10002, toUserID: 10003,createTime :1 ,status:6},
       {transactionID: 9, infoID: 200003, category: 1, fromUserID: 1366, toUserID: 10003,createTime :1 ,status:6},
       {transactionID: 10, infoID: 16473, category: 1, fromUserID: 1366, toUserID: 9368,createTime :1 ,status:6}]
    component.forceGraph();
    component.randamColor();
    expect(component.cldsz([1,1])).toEqual(1);
    expect(component.cldfm({data:[1,1]})).toEqual('1 completed transactions created in 1');
  });
});
