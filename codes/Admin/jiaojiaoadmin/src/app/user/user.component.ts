import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  users: any;
  threshold: number;
  constructor() { }

  ngOnInit() {
    this.users = [ {id: 4396, name: 4396, forbid: '是', score: 5, count: 100},
    {id: 216, name: '王牌测试员', forbid: '否', score: 4.2, count: 20}, {id: 426, name: '艾克', forbid: '是', score: 4.396, count: 200},
    {id: 123, name: '垃圾分类小助手', forbid: '是', score: 3, count: 10}];
  }

}
