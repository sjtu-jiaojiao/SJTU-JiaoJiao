import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-activity',
  templateUrl: './activity.component.html',
  styleUrls: ['./activity.component.css']
})
export class ActivityComponent implements OnInit {
  activitys: any;
  constructor() { }

  ngOnInit() {
    this.activitys = [ {id: 4396, intro: '高星信用评价获得尊贵图标', state: '正在进行'},
    {id: 396, intro: '毕业季', state: '正在进行'},
    {id: 96, intro: '垃圾分类新时尚', state: '已结束'},
    {id: 6, intro: '我们需要你输入的tag!',  state: '未开始'}];
  }

}
