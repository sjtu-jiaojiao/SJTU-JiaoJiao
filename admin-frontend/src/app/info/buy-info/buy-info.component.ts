
import { Component, OnInit, ɵConsole } from '@angular/core';
import { buyInfo } from '../../entity/info';
import { InfoService } from '../../info.service';
import { filter } from 'rxjs/operators';
import { Format } from 'src/app/Formatter/format';
@Component({
  selector: 'app-buy-info',
  templateUrl: './buy-info.component.html',
  styleUrls: ['./buy-info.component.css']
})
export class BuyInfoComponent implements OnInit {
  tags = ['测试','数据'];
  searchTag: string[]=[];
  buyinfos: buyInfo[];
  current : number = 1;
  size :number = 4;
  count : number;
  selectedType: number=0;
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
    this.infoService.searchBuyInfos(this.searchUser,this.size, this.current*this.size-this.size)
    .subscribe(infos => {
      if(!infos) return;
      this.buyinfos = infos.buyInfo;
      this.checkcount();
    });
  }

  stringToDate(params) {
      const date = new Date(params);
      return Format(date,'yyyy-MM-dd HH:mm:ss');
      }

  checkcount(){
    if(this.buyinfos && this.buyinfos.length === this.size)
    this.count = (this.current + 1) * this.size;
    else 
    this.count = this.current * this.size;
  }

  getinfos(): void {
    this.infoService.getPageBuyInfos(this.size, this.current*this.size-this.size)
    .subscribe(infos => {
      if(!infos) return;
      this.buyinfos = infos.buyInfo;
      this.checkcount();
    });
  }

  onChange(){
    this.searchByUser();
  }


}
