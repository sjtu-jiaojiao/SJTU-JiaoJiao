import { Component, OnInit } from '@angular/core';
import { WebsiteService } from './../website.service';
import { UserService } from './../user.service';
import { InfoService } from '../info.service';
import { sellInfo } from './../entity/info';
import { getCurrencySymbol } from '@angular/common';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  constructor(private ifService: InfoService, private wbService: WebsiteService,private usService: UserService) { }

  site : any;
  hst: any[];
  userNum: number;
  infoNum: number;
  acInfoNum: number;
  ngOnInit() {
    this.getHistory();
    this.wbService.getSite().subscribe(s => this.site= s);
    this.usService.getUsers().subscribe( e =>this.userNum = e.user.length);
    this.ifService.getSellInfos().subscribe( e => {
      this.infoNum = e.sellInfo.length;
      this.acInfoNum = e.sellInfo.filter( e => e.status < 3).length;
    }
      );
  }
  getHistory(){
    this.wbService.getSiteHistory().subscribe( s =>this.hst =s.sort((a,b)=> a.time-b.time));
  }
  add() {
    const newitem = {id: '',time:'2019-01-01',description:'init' };
    const tmp = this.hst;
    newitem.id = tmp.sort((a,b)=> b.id-a.id)[0].id +1;
    this.hst.push(newitem);
    this.hst =this.hst.sort((a,b)=> a.time-b.time)
    this.wbService.addSiteHistory(newitem).subscribe(
      _ => this.getHistory());
  }          
  getColor(item) {
    switch(item.type){
      case 1: 
      return 'green';
      case -1:
        return 'red';
      default:
          return 'blue';
    }
  }
  save() {
    this.hst.forEach( e => {
      if(e.description.trim())
        this.wbService.updateSiteHistory(e).subscribe(
          _ => this.getHistory());
      else 
        this.wbService.deleteSiteHistory(e).subscribe(
          _ => this.getHistory()
        );
    });
  }
}
