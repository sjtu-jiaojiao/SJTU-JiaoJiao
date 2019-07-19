import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { sellInfo, buyInfo, InfoResponse } from './entity/info';

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

  /** GET infos from the server */
  getSellInfos(): Observable<InfoResponse> {
    return this.http.get<InfoResponse>(this.sellinfoUrl)
      .pipe(
        catchError(this.handleError<InfoResponse>('getSellInfos'))
      );
  }

  /** GET infos from the server */
  getPageSellInfos(limit:number, offset: number): Observable<InfoResponse> {    
    const url = `${this.sellinfoUrl}?limit=${limit}&offset=${offset}`;
    return this.http.get<InfoResponse>(url)
      .pipe(
        catchError(this.handleError<InfoResponse>('getPageSellInfos'))
      );
  }

  /** GET info by id. Will 404 if id not found */
  getSellInfo(id: string): Observable<sellInfo> {
    const url = `${this.sellinfoUrl}/${id}`;
    return this.http.get<sellInfo>(url).pipe(
      catchError(this.handleError<sellInfo>(`getInfo id=${id}`))
    );
  }

  /** GET infos from the server */
  searchSellInfos(term: string,limit:number, offset: number): Observable<InfoResponse> {    
    const url = `${this.sellinfoUrl}?userId=${term}&limit=${limit}&offset=${offset}`;
    return this.http.get<InfoResponse>(url)
      .pipe(
        catchError(this.handleError<InfoResponse>('getInfos'))
      );
  }

  /** PUT: update the info on the server */
  updateSellInfo(info: any): Observable<any> {
    return this.http.post(this.sellinfoUrl, info, httpOptions).pipe(
      catchError(this.handleError<any>('updateSellInfo'))
    );
  }

  /** GET infos from the server */
  getBuyInfos(): Observable<InfoResponse> {
    return this.http.get<InfoResponse>(this.buyinfoUrl)
      .pipe(
        catchError(this.handleError<InfoResponse>('getBuyInfos'))
      );
  }

  /** GET infos from the server */
  getPageBuyInfos(limit:number, offset: number): Observable<InfoResponse> {    
    const url = `${this.buyinfoUrl}?limit=${limit}&offset=${offset}`;
    return this.http.get<InfoResponse>(url)
      .pipe(
        catchError(this.handleError<InfoResponse>('getPageBuyInfos'))
      );
  }

  /** GET info by id. Will 404 if id not found */
  getBuyInfo(id: string): Observable<buyInfo> {
    const url = `${this.buyinfoUrl}/${id}`;
    return this.http.get<buyInfo>(url).pipe(
      catchError(this.handleError<buyInfo>(`getBuy id=${id}`))
    );
  }

  /** GET infos from the server */
  searchBuyInfos(term: string,limit:number, offset: number): Observable<InfoResponse> {    
    const url = `${this.buyinfoUrl}?userId=${term}&limit=${limit}&offset=${offset}`;
    return this.http.get<InfoResponse>(url)
      .pipe(
        catchError(this.handleError<InfoResponse>('getBuyInfos'))
      );
  }

  /** PUT: update the info on the server */
  updateBuyInfo(info: any): Observable<any> {
    return this.http.post(this.buyinfoUrl, info, httpOptions).pipe(
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
