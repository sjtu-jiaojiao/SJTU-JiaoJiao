import { Injectable } from '@angular/core';
@Injectable({
  providedIn: 'root'
})
export class InfoComService {
  gridspan:number;
  searchUserID: string;
  searchStatus: number;
  searchGoodName: string;
  constructor() { }
  set(gs){
    this.gridspan = gs;
  }
  get(){{
    return this.gridspan;
  }}
  storage(u,s,g){
    this.searchUserID=u;
    this.searchGoodName=g;
    this.searchStatus=s;
  }
  unstorage(){
    return{
      u: this.searchUserID,
      s: this.searchStatus,
      g: this.searchGoodName
    }
  }
}
