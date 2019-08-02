import { Component, OnInit, ɵConsole, ViewChild, AfterViewInit } from '@angular/core';
import { sellInfo,buyInfo } from '../entity/info';
import { InfoService } from '../info.service';
import { filter } from 'rxjs/operators';
import { SellInfoComponent } from './sell-info/sell-info.component';
import { BuyInfoComponent } from './buy-info/buy-info.component';
import { InfoComService } from './infocom.service';

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: ['./info.component.css'],
})
export class InfoComponent implements OnInit, AfterViewInit {
  searchTag: string[]=[];
  searchType: number=-1;
  searchUserID: string;
  searchStatus: number;
  searchGoodName: string;
  @ViewChild(SellInfoComponent, {static: false})
  schild: SellInfoComponent;
  @ViewChild(BuyInfoComponent, {static: false})
  bchild: BuyInfoComponent;
  constructor(private gs: InfoComService) { }

  ngOnInit() {      
    this.gs.set(12);   
  }

  ngAfterViewInit() {
    setTimeout(()=> this.selectType(-1), 0);
  }
  search(){
    this.gs.storage(this.searchUserID,this.searchStatus,this.searchGoodName);
    if(this.searchType!==0){
    this.bchild.getinfos();
    }
    if(this.searchType!==1){
    this.schild.getinfos(); 
    }
  }
  selectType(type: number){
    if(type===-1)
    {
      this.gs.set(12);   
  }
    else{
      this.gs.set(6);
    }
    this.search();
  }
  selectTag(tag: string[]){
    if(this.searchType!==0)
    this.bchild.searchTag =tag;
    if(this.searchType!==1)
    this.schild.searchTag =tag;
  }
}
