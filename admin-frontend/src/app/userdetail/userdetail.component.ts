import { Component, OnInit, Input } from '@angular/core';
import { UserService } from '../user.service';
import { ActivatedRoute } from '@angular/router';
import { User } from '../entity/user';
import { Location } from '@angular/common';
import { InfoService } from '../info.service';
import { Format } from '../Formatter/format';
import { SellInfoComponent } from '../info/sell-info/sell-info.component';
import { sellInfo, buyInfo } from '../entity/info';
const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
export function fFormatter(params) {
     return params.value + ' activities in ' + params.name;
}
@Component({
  selector: 'app-userdetail',
  templateUrl: './userdetail.component.html',
  styleUrls: ['./userdetail.component.css']
})
export class UserDetailComponent implements OnInit {
  option: any;
  @Input() user: User;
  forbid : boolean;
  userName : string;
  infos : any[];
  constructor(
  private route: ActivatedRoute,
  private userService: UserService,
  private infoService: InfoService,
  private location: Location
) {}
stringToDate(params) {
    const date = new Date(parseInt(params+'000'));
    return Format(date,'yyyy-MM-dd HH:mm:ss');
    }

  ngOnInit(): void {
  this.getuser();
}
  typeof(obj): string {
    if( obj['sellInfoID'])
    return 'sellInfo';
    if( obj['buyInfoID'])
    return 'buyInfo';
  }
    goBack(): void {
    this.location.back();
  }
    getuser(): void {
    const id = this.route.snapshot.paramMap.get('id');
    this.userService.getUser(id)
      .subscribe(user => {this.user = user; 
        this.forbid = this.user.status === 2;
        this.userName = this.user.userName;
        this.infoService.getSellInfos(id).subscribe(e => {
          if(e.sellInfo)
          {
        this.infos = e.sellInfo;
          }
          this.infoService.getBuyInfos(id).subscribe(e => {
            if(e.buyInfo)
          this.infos = this.infos.concat(e.buyInfo);
          this.infos= this.infos.sort((a,b )=>parseInt(a.releaseTime)-
      parseInt(b.releaseTime));
          this.graph();
            }
          );


        });
      });
   
  }
    save(): void {
    if(!this.user) return;
    const status = this.forbid? 2: 1;
    this.userService.updateUser({userID: this.user.userID, status: status, userName: this.userName, role: this.user.role})
      .subscribe(() => this.goBack());
  }

  getstate(statecode: number): string {
    switch (statecode){	
      case 1:
        return '待预约';
      case 2: 
        return '预约';
      case 3:
        return '完成';
      case 4: 
        return '失效';
      case 5: 
        return '关闭';
    }
}





  graph() {

const sellData = [ 0 , 0 , 0 , 0 , 0 ,0 ,0 , 0 ,0 , 0 ,0 ,0];
const buyData = [ 0 , 0 , 0 , 0 , 0 ,0 ,0 , 0 ,0 , 0 ,0 ,0];
this.infos.forEach(element => {
  const y =new Date(element.releaseTime*1000).getFullYear(); 
  const m =new Date(element.releaseTime*1000).getMonth() + 1;
  //if((ynow == y + 1 && m > mnow) || (ynow == y && m <= mnow) )
    if(this.typeof(element)=='sellInfo')
      sellData[m-1] +=1;
    else 
      buyData[m-1] +=1;
});
    this.option = {
title: {
},
legend: {
    data: ['出售', '求购']
},
polar: {},
tooltip: {
  formatter: fFormatter
},
angleAxis: {
    type: 'category',
    data: months,
    boundaryGap: false,
    splitLine: {
        show: true,
        lineStyle: {
            color: '#999',
            type: 'dashed'
        }
    },
    axisLine: {
        show: false
    }
},
radiusAxis: {
    type: 'value',
      minInterval: 1
},
series: [{
  name: 'Sell Record',
  type: 'bar',
  data: sellData,
  coordinateSystem: 'polar'
},
{
  name: 'Buy Record',
  type: 'bar',
  data: buyData,
  coordinateSystem: 'polar'
}]
};
  }
}
