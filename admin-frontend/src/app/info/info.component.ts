import { Component, OnInit, ɵConsole } from '@angular/core';
import { Info } from '../entity/info';
import { InfoService } from '../info.service';
import { filter } from 'rxjs/operators';

export function Format (date, fmt) { //author: meizz 
  var o = {
      "M+": date.getMonth() + 1, //月份 
      "d+": date.getDate(), //日 
      "H+": date.getHours(), //小时 
      "m+": date.getMinutes(), //分 
      "s+": date.getSeconds(), //秒 
      "q+": Math.floor((date.getMonth() + 3) / 3), //季度 
      "S": date.getMilliseconds() //毫秒 
  };
  if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (date.getFullYear() + "").substr(4 - RegExp.$1.length));
  for (var k in o)
  if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
  return fmt;
}

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: ['./info.component.css'],
})
export class InfoComponent implements OnInit {
  searchTag: string[]=[];
  infos: Info[];
  current : number = 1;
  size :number = 4;
  count : number;
  Tthreshold: number;
  Ythreshold: number;
  selectedType: number=-1;
  threshold: number;
  searchUser: string;
  constructor(private infoService: InfoService) { }

  ngOnInit() {
    this.getinfos();
  }

  getstate(statecode: number): string {
      switch (statecode){	
        case 1:
          return '待预约';
        case 2: 
          return '预约';
        case 3:
          return '完成';
        case 4: 
          return '失效';
      }
  }
  searchByUser(): void {    
    if (!this.searchUser || !this.searchUser.trim()) {
    this.getinfos();
    // if not search term, return all info array.
    return;
  }
    this.infoService.searchInfos(this.searchUser,this.size, this.current*this.size-this.size)
    .subscribe(infos => {
      if(!infos) return;
      this.infos = infos.sellInfo;
      this.checkcount();
    });
  }
  stringToDate(params) {
      const date = new Date(params);
      return Format(date,'yyyy-MM-dd HH:mm:ss');
      }
  checkcount(){
    if(this.infos && this.infos.length === this.size)
    this.count = (this.current + 1) * this.size;
    else 
    this.count = this.current * this.size;
  }
  getinfos(): void {
    this.infoService.getPageInfos(this.size, this.current*this.size-this.size)
    .subscribe(infos => {
      if(!infos) return;
      this.infos = infos.sellInfo;
      this.checkcount();
    });
  }

  onChange(){
    this.searchByUser();
  }
}
