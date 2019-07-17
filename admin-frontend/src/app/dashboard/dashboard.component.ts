import { Component, OnInit } from '@angular/core';
import { WebsiteService } from './../website.service';
import { UserService } from './../user.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  constructor(private wbService: WebsiteService,private usService: UserService) { }

  site : any;
  userNum: any;
  ngOnInit() {
    this.wbService.getSite().subscribe(s => this.site= s);
    this.usService.getUsers().subscribe( e =>this.userNum = e.user.length);
  }

}
