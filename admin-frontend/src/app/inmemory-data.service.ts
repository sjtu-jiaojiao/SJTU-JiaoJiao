import { InMemoryDbService } from 'angular-in-memory-web-api';
import { Injectable } from '@angular/core';
import { sellInfo } from 'src/app/entity/info';

@Injectable({
  providedIn: 'root',
})
export class InMemoryDataService implements InMemoryDbService {
  createDb() {
    const site = [{id: 0, name: 'jiaojiao', status: true, cp: '2019 jiaojiao'}];
    const history = [ {id: 0, description: '项目启动', time: '2019-07-01', type: 0},
    {id: 1, description: '第一次迭代', time: '2019-07-12', type: 1},
    {id: 2, description: '第二次迭代', time: '2019-07-23', type: 1},
    {id: 3, description: '紧急维护', time: '', type: -1}
  ];
    const activity = [
      {id: 0, title: '开学季大甩卖', 
      description: '出售信息满足开学季需求可进行申报,要求多媒体数>=3,获得优先加权推荐', 
      releaseTime: 1567957148885, validTime: 1567957148885, weight: 2, 
    pic:3,isNew:false
      },
      {id: 1, title: '新人体验',
      description: '用户首次发布交易信息可进行申报,要求多媒体数>=1,获得特别加权推荐',
      releaseTime: 1567957148885, validTime: 1567957148885, weight: 3,
      pic:1,isNew:true
      },
      {id:2, title: '互饮家乡水',
      description: '商品属于土特产类可进行申报,要求多媒体数>=3,获得加权推荐',
      releaseTime: 1567957148885, validTime: 1567957148885, weight: 1,
      pic:3,isNew:false
      }
    ];
    return { site , history, activity};
  }
}