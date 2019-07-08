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
import { PasswordComponent } from './password/password.component';
import { InfoStatisticComponent } from './info-statistic/info-statistic.component';
import { ActivitydetailComponent } from './activitydetail/activitydetail.component';
import { AuthGuard } from './auth/auth.guard';

const routes: Routes = [
  { path: 'dashboard', component: DashboardComponent,canActivate: [AuthGuard] },
  { path: '', redirectTo: '/dashboard', pathMatch: 'full',canActivate: [AuthGuard] },
  { path: 'user/:id', component: UserDetailComponent ,canActivate: [AuthGuard]},
  { path: 'user', component: UserComponent,canActivate: [AuthGuard]},
  { path: 'login', component: LoginComponent},
  { path: 'info', component:  InfoComponent,canActivate: [AuthGuard]},
  { path: 'info/:id', component:  InfoDetailComponent,canActivate: [AuthGuard]},
  { path: 'activity/:id', component:  ActivitydetailComponent,canActivate: [AuthGuard]},
  { path: 'statistic', component:  InfoStatisticComponent,canActivate: [AuthGuard]},
  { path: 'website', component:  WebsiteComponent,canActivate: [AuthGuard]},
  { path: 'activity', component:  ActivityComponent,canActivate: [AuthGuard]},
  { path: 'password', component:  PasswordComponent,canActivate: [AuthGuard]},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
