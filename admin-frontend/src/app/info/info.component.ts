import { Component, OnInit, ɵConsole, ViewChild, AfterViewInit } from '@angular/core';
import { sellInfo,buyInfo } from '../entity/info';
import { InfoService } from '../info.service';
import { filter } from 'rxjs/operators';
import { SellInfoComponent } from './sell-info/sell-info.component';
import { BuyInfoComponent } from './buy-info/buy-info.component';

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: ['./info.component.css'],
})
export class InfoComponent implements OnInit {
  searchTag: string[]=[];
  selectedType: number=0;
  searchUser: string;
  @ViewChild(SellInfoComponent, {static: false})
  schild: SellInfoComponent;
  @ViewChild(BuyInfoComponent, {static: false})
  bchild: BuyInfoComponent;
  constructor() { }

  ngOnInit() {
  }
  searchByUser(){
    if(this.selectedType!==0){
    this.bchild.searchUser=this.searchUser;
    this.bchild.searchByUser();
    }
    
    if(this.selectedType!==1){
    this.schild.searchUser= this.searchUser;
    this.schild.searchByUser(); 
    }
  }
  
  selectTag(tag: string[]){
    if(this.selectedType!==0)
    this.bchild.searchTag =tag;
    if(this.selectedType!==1)
    this.schild.searchTag =tag;
  }
}
