import { Component, OnInit } from '@angular/core';
import { WebsiteService } from './../website.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  constructor(private wbService: WebsiteService) { }

  site : any;
  ngOnInit() {
    this.wbService.getSite().subscribe(s => this.site= s);
  }

}
