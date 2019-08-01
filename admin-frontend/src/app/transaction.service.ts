import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Transaction } from './entity/transaction';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class TransactionService {

  private trUrl = 'api/transaction';  // URL to web api
  tr: Transaction[]=[];
  constructor(private http: HttpClient) { }
  
  getAllTR(status,lowCreateTime,highCreateTime): Transaction[]{
    if(this.tr.length==0)
    {
      this.getTransactions(null,status,null,lowCreateTime,highCreateTime).subscribe(e=>
        {
        this.tr = e.transactions;
        this.getMoreTransaction(100,false,status,lowCreateTime,highCreateTime);
        })
    }
    return this.tr;
  }

  getMoreTransaction(offset, dynamic,status,lowCreateTime,highCreateTime){
    if(!(this.tr.length%100) && !dynamic)
      this.getTransactions(null,null,null,lowCreateTime,highCreateTime, null,offset).subscribe(
          e => {
              if(e&&Object.keys(e).length!=0){
              this.tr=this.tr.concat(e.transactions);
              if(e.buyInfo.length!=100)
              dynamic =! dynamic;
            }
              this.getMoreTransaction(offset+100, dynamic,status,lowCreateTime,highCreateTime);
            }
      );
      else{
          setTimeout(() => {
              this.getTransactions(null,null,null,lowCreateTime,highCreateTime, null,offset).subscribe(
                  e => {
                      if(e&&Object.keys(e).length!=0)
                      this.tr=this.tr.concat(e.buyInfo);
                      this.getMoreTransaction(this.tr.length-1, true,status,lowCreateTime,highCreateTime);
          }
      );}, 5000 );
      }
}

  /** GET infos from the server */
  getTransactions(userID: string = null, status: number = null, infoID: number = null, 
    lowCreateTime :number = null, highCreateTime: number =null, limit:number=null, offset: number=null): 
    Observable<any> 
    {    
    let url = `${this.trUrl}?`;
    if(userID && userID.trim()) url += `userID=${userID}&`;
    if(status) url+= `status=${status}&`;
    if(infoID) url+=`goodName=${infoID}&`;
    if(limit) url+= `limit=${limit}&`;
    if(lowCreateTime) url+=`lowCreateTime=${lowCreateTime}&`;
    if(highCreateTime) url+=`highCreateTime=${highCreateTime}&`;
    if(offset) url+=`offset=${offset}&`;
    return this.http.get<any>(url).pipe(
        catchError(this.handleError<any>('getTransactions'))
      );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   *
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
