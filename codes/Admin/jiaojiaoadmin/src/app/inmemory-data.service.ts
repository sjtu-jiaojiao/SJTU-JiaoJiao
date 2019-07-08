import { InMemoryDbService } from 'angular-in-memory-web-api';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class InMemoryDataService implements InMemoryDbService {
  createDb() {
    const users =  [ {id: '4396', name: '4396', forbid: '是', score: 5, active: 100,intro: '我是练习市场两年半的偶像练习生,我是练习市场两年半的偶像练习生，我喜欢唱跳rap和篮球'},
    {id: '216', name: '王牌测试员', forbid: '否', score: 4.2, active: 20,intro: '我是lboss的小粉丝,一生唯爱lboss，一见lboss误终生'}, 
    {id: '426', name: '艾克', forbid: '是', score: 4.396, active: 200,intro: 'cnmooc天下第二！jboss！我爱学习强国！'},
    {id: '123', name: '垃圾分类小助手', forbid: '是', score: 3, active: 10,intro: '你是什么垃圾？我是本群的垃圾分类小助手，帮助你回收你的垃圾'}];
    
    const infos = [ 
    {id: '4396', source:'4396', type: 0, price: 100, tags: ['黑色','塑料','垃圾桶'], time: '2019-01-01', state: 0, count: 2, intro: '黑色塑料分类垃圾桶'},
    {id: '196', source: '123', type: 0, price: 35, tags:['C++','精装','中文','书籍'],time: '2019-07-01', state: 1, count: 10, intro: '深度探索C++对象模型中文精装本'},
    {id: '42396', source: '1234', type: 1,price: 52,  time: '2019-02-11',tags:['不锈钢','保温杯'],state: 2, count: 2, intro: '虎牌不锈钢保温杯'},
    {id: '43396', source: '4396', type: 1,price: 200,  time: '2019-05-10', tags:['凤凰牌','自行车'], state: 3, count: 0, intro: '凤凰牌使用1年的自行车'},
    {id: '2396', source: '213', type: 0, price: 23, tags: ['塑料','垃圾袋'], time: '2019-01-01', state: 0, count: 2, intro: '100抽塑料垃圾袋'},
    {id: '195', source: '123', type: 0, price: 15, tags:['C++','中文','书籍'],time: '2019-6-01', state: 1, count: 10, intro: 'Effective中文版'},
    {id: '42296', source: '14', type: 1,price: 52,  time: '2019-02-11',tags:['黑色','保温杯'],state: 2, count: 2, intro: '黑色保温杯'},
    {id: '4196', source: '4', type: 1,price: 20,  time: '2019-05-10', tags:['小排量','电动','自行车'], state: 3, count: 0, intro: '小排量电动自行车'}];

    return {users , infos};
  }
}