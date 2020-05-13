import { Component, OnInit } from '@angular/core';
import { ActivityService } from '../activity.service';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';


@Component({
  selector: 'app-activitydetail',
  templateUrl: './activitydetail.component.html',
  styleUrls: ['./activitydetail.component.css']
})
export class ActivitydetailComponent implements OnInit {
  act: any;
  constructor(
    private route: ActivatedRoute,
    private actService: ActivityService,
    private location: Location
  ) {}

  ngOnInit(): void {
  this.getAct();
}
    getAct(): void {
    const id = this.route.snapshot.paramMap.get('id');
    this.actService.getAct(id).subscribe( e=> {
      this.act = e;
      this.act.validDate=new Date(e.validTime); 
      this.act.releaseDate=new Date(e.releaseTime); 
    });
  }
  goBack(): void {
  this.location.back();
}
save(): void{ 
  this.act.validTime=this.act.validDate.getTime(); 
  this.act.releaseTime=this.act.releaseDate.getTime(); 
  this.actService.updateAct(this.act).subscribe();
}

}
