import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { sellInfo, buyInfo, InfoResponse } from './entity/info';
import { query } from '\@angular/animations';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class InfoService {

  private sellinfoUrl = 'api/sellInfo';  // URL to web api
  private buyinfoUrl = 'api/buyInfo';
  constructor(private http: HttpClient) { }
  bi: buyInfo[] = [];
  si: sellInfo[] = [];

  getAcInfo(){
    return this.bi.filter( a => a.status<3).length+ this.si.filter(a =>a.status<3).length;
  }
  getInfoNum(){
    return this.bi.length + this.si.length;
  }
  getReserveInfoNum(){
    return this.bi.filter( a=> a.status ==2).length+ this.si.filter(a=> a.status ==2).length;
  }
  getAllSellInfo(): sellInfo[]{
    if(this.si.length==0)
    {
      this.getSellInfos().subscribe(e=>
        {
        this.si = e.sellInfo;
        this.getMoreSell(100,false);
        })
    }
    return this.si;
  }
  
  getAllBuyInfo(): buyInfo[]{
    if(this.bi.length==0)
    {
      this.getBuyInfos().subscribe(e=>
        {
        this.bi = e.buyInfo;
      this.getMoreBuy(100,false);
        })
    }
    return this.bi;
  }
  
  getMoreBuy(offset, dynamic){
    if(!(this.bi.length%100) && !dynamic)
      this.getBuyInfos(null,null,null,null,offset).subscribe(
          e => {
              if(e&&Object.keys(e).length!=0){
              this.bi=this.bi.concat(e.buyInfo);
              if(e.buyInfo.length!=100)
              dynamic =! dynamic;
            }
              this.getMoreBuy(offset+100, dynamic);
            }
      );
      else{
          setTimeout(() => {
              this.getBuyInfos(null,null,null,null,offset).subscribe(
                  e => {
                      if(e&&Object.keys(e).length!=0)
                      this.bi=this.bi.concat(e.buyInfo);
                      this.getMoreBuy(this.bi.length-1, true);
          }
      );}, 5000 );
      }
}

getMoreSell(offset, dynamic){
    if(!(this.si.length%100) && !dynamic)
      this.getSellInfos(null,null,null,null,offset).subscribe(
          e => {
              if(e&&Object.keys(e).length!=0){
              this.si= this.si.concat(e.sellInfo);   
              if(e.sellInfo.length!=100)
                dynamic=!dynamic;
              }
              this.getMoreSell(offset+100,dynamic)
            }
      );
  else{
      setTimeout(() => {
          this.getSellInfos(null,null,null,null,offset).subscribe(
              e => {
                  if(e&&Object.keys(e).length!=0)
                  this.si=this.si.concat(e.sellInfo);
                  this.getMoreSell(this.si.length-1,true);
      }
  );}, 5000 );
  }
}
  /** GET info by id. Will 404 if id not found */
  getSellInfo(id: number): Observable<sellInfo> {
    const url = `${this.sellinfoUrl}/${id}`;
    return this.http.get<sellInfo>(url).pipe(
      catchError(this.handleError<sellInfo>(`getInfo id=${id}`))
    );
  }

  /** GET infos from the server */
  getSellInfos(userID: string = null, status: number = null, goodName: string = null,limit:number=null, offset: number=null): Observable<InfoResponse> {    
    let url = `${this.sellinfoUrl}?`;
    if(userID && userID.trim()) url += `userID=${userID}&`;

    if(status) url+= `status=${status}&`;
    if(goodName && goodName.trim()) url+=`goodName=${goodName}&`;
    if(limit) url+= `limit=${limit}&`;
    if(offset) url+=`offset=${offset}&`;

    return this.http.get<InfoResponse>(url)
      .pipe(
        catchError(this.handleError<InfoResponse>('getSellInfos'))
      );

    }

  /** PUT: update the info on the server */
  updateSellInfo(info: any): Observable<any> {
    return this.http.put(this.sellinfoUrl, info, httpOptions).pipe(
      catchError(this.handleError<any>('updateSellInfo'))
    );
  }

  /** GET infos from the server */
  getBuyInfos(userID: string = null, status: number = null, goodName: string = null,limit:number=null, offset: number=null):
   Observable<InfoResponse> {    
    let url = `${this.buyinfoUrl}?`;
    if(userID && userID.trim()) url += `userID=${userID}&`;
    if(status) url+= `status=${status}&`;
    if(goodName && goodName.trim()) url+=`goodName=${goodName}&`;
    if(limit) url+= `limit=${limit}&`;
    if(offset) url+=`offset=${offset}&`;
    return this.http.get<InfoResponse>(url)
      .pipe(
        catchError(this.handleError<InfoResponse>('getBuyInfos'))
      );
  }

  /** GET info by id. Will 404 if id not found */
  getBuyInfo(id: number): Observable<buyInfo> {
    const url = `${this.buyinfoUrl}/${id}`;
    return this.http.get<buyInfo>(url).pipe(
      catchError(this.handleError<buyInfo>(`getBuy id=${id}`))
    );
  }

  /** PUT: update the info on the server */
  updateBuyInfo(info: any): Observable<any> {
    return this.http.put(this.buyinfoUrl, info, httpOptions).pipe(
      catchError(this.handleError<any>('updateBuyInfo'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation, result?: T) {
    return (error: any): Observable<T> => {
      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead
      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

}
