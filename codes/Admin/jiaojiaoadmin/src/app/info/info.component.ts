import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: ['./info.component.css']
})
export class InfoComponent implements OnInit {

  infos: any;
  Tthreshold: number;
  Ythreshold: number;
  constructor() { }

  ngOnInit() {
    this.infos = [ {id: 4396, source: 4396, type: '买', time: '2019-01-01', state: '预约中', count: 2, intro: '想要买分类垃圾桶'},
    {id: 196, source: 123, type: '买', time: '2019-10-01', state: '交易完成', count: 10, intro: '求女朋友'},
    {id: 42396, source: 1234, type: '卖', time: '2019-02-11', state: '待评价', count: 2, intro: '卖前男友尸体'},
    {id: 43396, source: 4396, type: '卖', time: '2019-05-10', state: '可预约', count: 0, intro: '爱情不是你想买'}];
  }

}
