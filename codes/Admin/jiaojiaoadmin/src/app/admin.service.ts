import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
}; 

export class Admin{
  public password: string;
  public id: string;
}

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  private adminsUrl = 'api/admins';  // URL to web api
 
  constructor(private http: HttpClient){}
 
  login (username: string, password: string): Observable<boolean> {   
    const url = `${this.adminsUrl}/${username}`;
    return this.http.get<Admin>(url).pipe(
      map(ad => ad.password ===password)
    ); 
  }

  editpwd( username: string, password: string, newpassword: string) {
    this.login(username, password).subscribe(    
      e => {
      if(e) {
      const admin: Admin = new Admin();
      admin.password = newpassword;
      admin.id = username;
      this.http.put<Admin>(this.adminsUrl, admin, httpOptions).subscribe();
      }}
      );
  }
}
