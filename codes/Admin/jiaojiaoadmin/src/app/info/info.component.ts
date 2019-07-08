import { Component, OnInit } from '@angular/core';
import { Info } from '../entity/info';
import { InfoService } from '../info.service';
import { filter } from 'rxjs/operators';

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: ['./info.component.css']
})
export class InfoComponent implements OnInit {
  searchTag: string[]=[];
  infos: Info[];
  current : number;
  curinfos: Info[];
  size :number;
  count : number;
  Tthreshold: number;
  Ythreshold: number;
  threshold: number;
  constructor(private infoService: InfoService) { }

  ngOnInit() {
    this.getinfos();
  }

  selectTag(tags: string[]) :void {
    this.infoService.getInfos()
    .subscribe(infos => {
      this.infos = infos;
      tags.forEach( e => {
        this.infos = this.infos.filter( arr => arr.tags.indexOf(e) >= 0 );
        this.count = this.infos.length; 
        this.current=1;
        this.switchPage(this.current, this.size);
      });
    });
  }
  getstate(statecode: number): string {
      switch (statecode){
        case 0:
          return '可预约';
        case 1: 
          return '预约中';
        case 2:
          return '已完成';
        case 3:
          return '待评价';
        case 4: 
          return '强制结束';
      }
  }
  end():void {
    this.infos.filter(h => new Date().getTime() - new Date(h.time).getTime() /1000/60/60/24 > this.Tthreshold && h.count < this.Ythreshold)
    .map(h => { h.state=4 ; return h;}).forEach(element => 
      this.infoService.updateInfo(element).subscribe());
  }	
  getinfos(): void {
    this.infoService.getInfos()
    .subscribe(infos => {
      this.infos = infos; 
      this.count = this.infos.length; 
      this.current=1;
      this.size=4;
      this.switchPage(this.current, this.size);
    });
  }
  switchPage(page, size) {
    if(page* size< this.count)
    this.curinfos = this.infos.slice( (page-1)*size, page* size)
    else 
    this.curinfos = this.infos.slice( (page-1)*size );
  }
  pageChange(page){
    this.switchPage(page,this.size);
  }
  sizeChange(size){
    this.switchPage(this.current,size);
  }
  delete(info: Info): void {
    this.infos = this.infos.filter(h => h !== info);
    this.infoService.deleteInfo(info).subscribe(_ =>{
      this.count = this.infos.length; 
      this.switchPage(this.current, this.size);
    });
  }

  add(id: string): void {
    id = id.trim();
    if (!id) { return; }
    this.infoService.addInfo({ id } as Info)
      .subscribe(info => {
        this.infos.push(info);
      });
  }

}
