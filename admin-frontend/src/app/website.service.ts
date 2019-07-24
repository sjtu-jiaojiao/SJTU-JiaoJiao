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
export class WebsiteService {

  private siteUrl = 'api/site';  // URL to web api
  private hstUrl = 'api/history';
  constructor(private http: HttpClient) { }

  /** GET info by id. Will 404 if id not found */
  getSite(): Observable<any> {
    const url = `${this.siteUrl}/0`;
    return this.http.get<any>(url).pipe(
      catchError(this.handleError<any>(`getWebsite`))
    );
  }

  /** PUT: update the info on the server */
  updateSite(site): Observable<any> {
    const url = `${this.siteUrl}`;
    return this.http.put(this.siteUrl, site, httpOptions).pipe(
      catchError(this.handleError<any>('updateInfo'))
    );
  }

  /** GET info by id. Will 404 if id not found */
  getSiteHistory(): Observable<any[]> {
    const url = `${this.hstUrl}`;
    return this.http.get<any>(url).pipe(
      catchError(this.handleError<any>(`getWebsiteHistory`))
    );
  }

  deleteSiteHistory(item: any){
    const url = `${this.hstUrl}/${item.id}`;
    return this.http.delete<any>(url).pipe(
      catchError(this.handleError<any>(`deleteWebsiteHistory`))
    );
  }

  updateSiteHistory(item: any){
  const url = `${this.hstUrl}`;
  return this.http.put<any>(url,item,httpOptions).pipe(
    catchError(this.handleError<any>(`updateWebsiteHistory`))
  );
  }

  addSiteHistory(item: any){
    const url = `${this.hstUrl}`;
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
