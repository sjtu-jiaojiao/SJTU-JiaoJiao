import { Component, OnInit } from '@angular/core';
import { ActivityService } from '../activity.service';
import { Format } from '../Formatter/format';

@Component({
  selector: 'app-activity',
  templateUrl: './activity.component.html',
  styleUrls: ['./activity.component.css']
})
export class ActivityComponent implements OnInit {
  acts: any[];
  ID: string;
  constructor(private actService: ActivityService) { }
  ngOnInit() {
    this.actService.getActs().subscribe(e =>this.refresh(e));
  }
  delete(item){
    this.actService.deleteAct(item).subscribe(
      e =>
      this.acts = this.acts.filter( e => e.id != item.id)
    )
  }
  refresh(e){
    this.acts = e;
    this.acts.forEach(e=> e.validDate=new Date(e.validTime)); 
    this.acts.forEach(e=> e.releaseDate=new Date(e.releaseTime)); 
  }
  add(){
    const item = 
    {id:this.ID, title: '',
    description: '',
    releaseTime: '150000', validTime: '150000', weight: 1,
    pic:0,isNew:false
    };
    this.actService.addAct(item).subscribe(
      a =>
      this.actService.getActs().subscribe(e => this.refresh(e))
    )
  }
  save(item){
    item.validTime=item.validDate.getTime(); 
    item.releaseTime=item.releaseDate.getTime(); 
    this.actService.updateAct(item).subscribe(
      a =>
      this.actService.getActs().subscribe(e => this.refresh(e))
    )
  }
}
