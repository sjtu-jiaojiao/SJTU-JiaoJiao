import { Component, OnInit } from '@angular/core';
import { WebsiteService } from './../website.service';
import { UserService } from './../user.service';
import { InfoService } from '../info.service';
import { sellInfo } from './../entity/info';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  constructor(private ifService: InfoService, private wbService: WebsiteService,private usService: UserService) { }

  site : any;
  userNum: number;
  infoNum: number;
  acInfoNum: number;
  ngOnInit() {
    this.wbService.getSite().subscribe(s => this.site= s);
    this.usService.getUsers().subscribe( e =>this.userNum = e.user.length);
    this.ifService.getSellInfos().subscribe( e => {
      this.infoNum = e.sellInfo.length;
      this.acInfoNum = e.sellInfo.filter( e => e.status < 3).length;
    }
      );
  }

}
