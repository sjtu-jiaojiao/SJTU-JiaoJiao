import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { Info } from './entity/info';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class InfoService {

  private infosUrl = 'api/info';  // URL to web api

  constructor(private http: HttpClient) { }

  /** GET infos from the server */
  getInfos(): Observable<Info[]> {
    return this.http.get<Info[]>(this.infosUrl)
      .pipe(
        catchError(this.handleError<Info[]>('getInfos', []))
      );
  }

  /** GET info by id. Will 404 if id not found */
  getInfo(id: string): Observable<Info> {
    const url = `${this.infosUrl}/${id}`;
    return this.http.get<Info>(url).pipe(
      catchError(this.handleError<Info>(`getInfo id=${id}`))
    );
  }

  /** DELETE: delete the info from the server */
  deleteInfo(id: string): Observable<Info> {
    const url = `${this.infosUrl}/${id}`;
    return this.http.delete<Info>(url, httpOptions).pipe(
      catchError(this.handleError<Info>('deleteInfo'))
    );
  }

  /** PUT: update the info on the server */
  updateInfo(info: Info): Observable<any> {
    return this.http.put(this.infosUrl, info, httpOptions).pipe(
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
