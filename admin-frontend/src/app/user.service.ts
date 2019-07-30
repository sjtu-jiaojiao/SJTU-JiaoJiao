import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { User, UserResponse } from './entity/user'
const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({ providedIn: 'root' })
export class UserService {

  private usersUrl = 'api/user';  // URL to web api

  constructor(private http: HttpClient) { }

  /** GET users from the server */
  getPageUsers(limit:number, offset: number): Observable<UserResponse> {
    const url = `${this.usersUrl}?limit=${limit}&offset=${offset}`;
    return this.http.get<UserResponse>(url,httpOptions)
      .pipe(
        catchError(this.handleError<UserResponse>('getUsers'))
      );
  }
  /** GET users from the server */
  getUsers(): Observable<UserResponse> {
    const url = `${this.usersUrl}?limit=100000`;
    return this.http.get<UserResponse>(url,httpOptions)
      .pipe(
        catchError(this.handleError<UserResponse>('getUsers'))
      );
  }
  /** GET user by id. Will 404 if id not found */
  getUser(id: string): Observable<User> {
    const url = `${this.usersUrl}/${id}`;
    return this.http.get<User>(url).pipe(
      catchError(this.handleError<User>(`getUser id=${id}`))
    );
  }

  /* GET users whose name contains search term */
  searchUsers(term: string,limit:number, offset: number): Observable<UserResponse> {
    const url =`${this.usersUrl}?userName=${term}&limit=${limit}&offset=${offset}`;
    return this.http.get<UserResponse>(url).pipe(
      catchError(this.handleError<UserResponse>('searchUsers'))
    );
  }


  /** PUT: update the user on the server */
  updateUser(user): Observable<any> {
    return this.http.put(this.usersUrl, user, httpOptions).pipe(
      catchError(this.handleError<any>('updateUser'))
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