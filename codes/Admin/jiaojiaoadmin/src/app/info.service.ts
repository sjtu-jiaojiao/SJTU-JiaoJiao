import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
 
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { Info} from './entity/info';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};
 
@Injectable({
  providedIn: 'root'
})
export class InfoService {

  private infosUrl = 'api/infos';  // URL to web api
 
  constructor(private http: HttpClient){}
 
  /** GET infos from the server */
  getInfos (): Observable<Info[]> {
    return this.http.get<Info[]>(this.infosUrl)
      .pipe(
        catchError(this.handleError<Info[]>('getInfos', []))
      );
  }
 
  /** GET info by id. Return `undefined` when id not found */
  getInfoNo404<Data>(id: string): Observable<Info> {
    const url = `${this.infosUrl}/?id=${id}`;
    return this.http.get<Info[]>(url)
      .pipe(
        map(infos => infos[0]), // returns a {0|1} element array
        catchError(this.handleError<Info>(`getInfo id=${id}`))
      );
  }
 
  /** GET info by id. Will 404 if id not found */
  getInfo(id: string): Observable<Info> {
    const url = `${this.infosUrl}/${id}`;
    return this.http.get<Info>(url).pipe(
      catchError(this.handleError<Info>(`getInfo id=${id}`))
    );
  }
 
  /* GET infos whose name contains search term */
  searchInfos(term: string): Observable<Info[]> {
    if (!term.trim()) {
      // if not search term, return empty info array.
      return of([]);
    }
    return this.http.get<Info[]>(`${this.infosUrl}/?intro=${term}`).pipe(
      catchError(this.handleError<Info[]>('searchInfos', []))
    );
  }
 
  //////// Save methods //////////
 
  /** POST: add a new info to the server */
  addInfo (info: Info): Observable<Info> {
    return this.http.post<Info>(this.infosUrl, info, httpOptions).pipe(
      catchError(this.handleError<Info>('addInfo'))
    );
  }
 
  /** DELETE: delete the info from the server */
  deleteInfo (info: Info | string): Observable<Info> {
    const id = typeof info === 'string' ? info : info.id;
    const url = `${this.infosUrl}/${id}`;
 
    return this.http.delete<Info>(url, httpOptions).pipe(
      catchError(this.handleError<Info>('deleteInfo'))
    );
  }
 
  /** PUT: update the info on the server */
  updateInfo (info: Info): Observable<any> {
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
  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
 
      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
 
}
