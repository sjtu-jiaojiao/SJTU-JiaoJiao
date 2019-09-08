
import { Component, OnInit, ɵConsole } from '@angular/core';
import { sellInfo,buyInfo } from '../../entity/info';
import { InfoService } from '../../info.service';
import { filter } from 'rxjs/operators';
import { Format } from 'src/app/Formatter/format';
import { InfoComService } from '../infocom.service';
import { FileService } from './../../file.service';

@Component({
  selector: 'app-sell-info',
  templateUrl: './sell-info.component.html',
  styleUrls: ['./sell-info.component.css']
})
export class SellInfoComponent implements OnInit {
  tags = ['测试','数据'];
  sellinfos: sellInfo[];
  gridspan: number;
  current : number = 1;
  size : number = 4;
  count : number;
  searchTag: string[]=[];
  searchUserID: string;
  searchStatus: number;
  searchGoodName: string;
  constructor(private gs: InfoComService, private infoService: InfoService, private fileService: FileService) { }

  ngOnInit() {
    this.getinfos();
    this.gridspan = this.gs.get();
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
        case 5:
          return '关闭';
      }
  }
  getinfos(): void {    
    const st = this.gs.unstorage();
    this.searchUserID=st.u;
    this.searchGoodName =st.g;
    this.searchStatus =st.s;
    this.infoService.getSellInfos(this.searchUserID,this.searchStatus, this.searchGoodName, this.size, this.current*this.size-this.size)
    .subscribe(infos => {
      if(!infos) return;
      this.sellinfos = infos.sellInfo;
      this.checkcount();
      this.getcontent();
    });
  }

  stringToDate(params) {
      const date = new Date(params*1000);
      return Format(date,'yyyy-MM-dd HH:mm:ss');
  }

  checkcount(){
    if(this.sellinfos && this.sellinfos.length === this.size)
    this.count = (this.current + 1) * this.size;
    else 
    this.count = this.current * this.size;
  }


  getcontent(): void{
      this.sellinfos.forEach(
        info => {
          this.fileService.getContent(info.contentID).subscribe(
            e => {
                if(e){
                info.tags = e.tags;
                }
            }
        )
        }
      )
  }
}
