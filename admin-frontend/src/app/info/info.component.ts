import { Component, OnInit } from '@angular/core';
import { Info } from '../entity/info';
import { InfoService } from '../info.service';
import { filter } from 'rxjs/operators';

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: ['./info.component.css'],
})
export class InfoComponent implements OnInit {
  searchTag: string[]=[];
  infos: Info[];
  current : number = 1;
  curinfos: Info[];
  size :number = 4;
  count : number;
  Tthreshold: number;
  Ythreshold: number;
  selectedType: number=-1;
  threshold: number;
  constructor(private infoService: InfoService) { }

  ngOnInit() {
    this.getinfos();
  }


  selectType(type: number) :void {
    this.infoService.getInfos()
    .subscribe(infos => {
      this.infos = infos;        
        this.infos = this.infos.filter( ele => ele.type !== (1-type) );
        this.searchTag.forEach( e => {
          this.infos = this.infos.filter( arr => arr.tags.indexOf(e) >= 0 );
        });
        this.count = this.infos.length;
        this.current=1;
        this.switchPage(this.current, this.size);
    }
        );
  }

  selectTag(tags: string[]) :void {
    this.infoService.getInfos()
    .subscribe(infos => {
      this.infos = infos;
      this.infos = this.infos.filter( ele => ele.type !== (1-this.selectedType) );
      tags.forEach( e => {
        this.infos = this.infos.filter( arr => arr.tags.indexOf(e) >= 0 );
      });
      this.count = this.infos.length; 
      this.current=1;
      this.switchPage(this.current, this.size);
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
}
