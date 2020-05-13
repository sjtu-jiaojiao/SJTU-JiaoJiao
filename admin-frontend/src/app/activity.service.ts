import { Injectable } from "@angular/core";
import { HttpHeaders, HttpClient } from "@angular/common/http";
import { Observable, of } from "rxjs";
import { catchError, max } from "rxjs/operators";

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class ActivityService {

  private actUrl = 'api/activity';
  constructor(private http: HttpClient) { }

  /** GET info by id. Will 404 if id not found */
  getAct(id): Observable<any> {
    const url = `${this.actUrl}/${id}`;
    return this.http.get<any>(url).pipe(
      catchError(this.handleError<any>(`getWebsiteHistory`))
    );
  }

  /** GET info by id. Will 404 if id not found */
  getActs(): Observable<any[]> {
    const url = `${this.actUrl}`;
    return this.http.get<any>(url).pipe(
      catchError(this.handleError<any>(`getWebsiteHistory`))
    );
  }

  deleteAct(item: any){
    const url = `${this.actUrl}/${item.id}`;
    return this.http.delete<any>(url).pipe(
      catchError(this.handleError<any>(`deleteWebsiteHistory`))
    );
  }

  updateAct(item: any){
  const url = `${this.actUrl}`;
  return this.http.put<any>(url,item,httpOptions).pipe(
    catchError(this.handleError<any>(`updateWebsiteHistory`))
  );
  }

  addAct(item: any){
    const url = `${this.actUrl}`;
    return this.http.post<any>(url,item,httpOptions).pipe(
      catchError(this.handleError<any>(`addWebsiteHistory`))
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
