import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { Info, InfoResponse } from './entity/info';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class InfoService {

  private infosUrl = 'api/sellInfo';  // URL to web api

  constructor(private http: HttpClient) { }

  /** GET infos from the server */
  getInfos(): Observable<InfoResponse> {
    return this.http.get<InfoResponse>(this.infosUrl)
      .pipe(
        catchError(this.handleError<InfoResponse>('getInfos'))
      );
  }

  /** GET infos from the server */
  getPageInfos(limit:number, offset: number): Observable<InfoResponse> {    
    const url = `${this.infosUrl}?limit=${limit}&offset=${offset}`;
    return this.http.get<InfoResponse>(url)
      .pipe(
        catchError(this.handleError<InfoResponse>('getInfos'))
      );
  }

  /** GET info by id. Will 404 if id not found */
  getInfo(id: string): Observable<Info> {
    const url = `${this.infosUrl}/${id}`;
    return this.http.get<Info>(url).pipe(
      catchError(this.handleError<Info>(`getInfo id=${id}`))
    );
  }

  /** GET infos from the server */
  searchInfos(term: string,limit:number, offset: number): Observable<InfoResponse> {    
    const url = `${this.infosUrl}?userId=${term}&limit=${limit}&offset=${offset}`;
    return this.http.get<InfoResponse>(url)
      .pipe(
        catchError(this.handleError<InfoResponse>('getInfos'))
      );
  }

  /** PUT: update the info on the server */
  updateInfo(info: any): Observable<any> {
    return this.http.post(this.infosUrl, info, httpOptions).pipe(
      catchError(this.handleError<any>('updateInfo'))
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
