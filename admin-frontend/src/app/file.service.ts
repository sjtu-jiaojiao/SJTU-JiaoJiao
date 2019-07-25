import { Injectable } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { ContentResponse } from './entity/content';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};
@Injectable({
  providedIn: 'root'
})

export class FileService {

  private contentUrl = 'api/content';  // URL to web api
  constructor(private http: HttpClient) { }
  /** GET info by id. Will 404 if id not found */
  getContent(id: string): Observable<ContentResponse> {
    const url = `${this.contentUrl}/${id}`;
    return this.http.get<ContentResponse>(url).pipe(
      catchError(this.handleError<ContentResponse>(`getContent id=${id}`))
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
