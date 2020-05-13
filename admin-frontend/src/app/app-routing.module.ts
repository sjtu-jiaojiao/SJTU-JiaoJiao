import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { UserDetailComponent } from './userdetail/userdetail.component';
import { UserComponent } from './user/user.component';
import { LoginComponent } from './login/login.component';
import { InfoComponent } from './info/info.component';
import { InfoDetailComponent } from './infodetail/infodetail.component';
import { WebsiteComponent } from './website/website.component';
import { ActivityComponent } from './activity/activity.component';
import { InfoStatisticComponent } from './info-statistic/info-statistic.component';
import { ActivitydetailComponent } from './activitydetail/activitydetail.component';
import { JWTGuard } from '@delon/auth';

const routes: Routes = [
  { path: 'dashboard', component: DashboardComponent,canActivate: [JWTGuard] },
  { path: '', redirectTo: '/dashboard', pathMatch: 'full',canActivate: [JWTGuard] },
  { path: 'user/:id', component: UserDetailComponent ,canActivate: [JWTGuard]},
  { path: 'user', component: UserComponent,canActivate: [JWTGuard]},
  { path: 'login', component: LoginComponent},
  { path: 'info', component:  InfoComponent,canActivate: [JWTGuard]},
  { path: 'info/:type/:id', component:  InfoDetailComponent,canActivate: [JWTGuard]},
  { path: 'activity/:id', component:  ActivitydetailComponent,canActivate: [JWTGuard]},
  { path: 'statistic', component:  InfoStatisticComponent,canActivate: [JWTGuard]},
  { path: 'website', component:  WebsiteComponent,canActivate: [JWTGuard]},
  { path: 'activity', component:  ActivityComponent,canActivate: [JWTGuard]},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
