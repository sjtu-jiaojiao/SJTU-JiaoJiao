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
  ]
    return { site , history};
  }
}