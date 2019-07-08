import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree, Router, CanActivate } from '@angular/router';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class AuthGuard implements CanActivate {
  
  constructor(private router: Router) {}

canActivate(
  next: ActivatedRouteSnapshot,
  state: RouterStateSnapshot): boolean {
    return true;
}
}
