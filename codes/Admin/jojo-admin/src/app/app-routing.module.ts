import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DashboardComponent } from './dashboard/dashboard.component';
import { UserDetailComponent } from './userdetail/userdetail.component';
import { UserComponent } from './user/user.component';
import { LoginComponent } from './login/login.component';
import { InfoComponent } from './info/info.component';
import { InfoDetailComponent } from './infodetail/infodetail.component';
import { UserStatisticComponent } from './user-statistic/user-statistic.component';
import { WebsiteComponent } from './website/website.component';
import { ActivityComponent } from './activity/activity.component';
import { PasswordComponent } from './password/password.component';
import { InfoStatisticComponent } from './info-statistic/info-statistic.component';

const routes: Routes = [
  { path: 'dashboard', component: DashboardComponent },
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'user/:id', component: UserDetailComponent },
  { path: 'user', component: UserComponent},
  { path: 'login', component: LoginComponent},
  { path: 'info', component:  InfoComponent},
  { path: 'info/:id', component:  InfoDetailComponent},
  { path: 'statistic/user', component:  UserStatisticComponent},
  { path: 'statistic/info', component:  InfoStatisticComponent},
  { path: 'website', component:  WebsiteComponent},
  { path: 'activity', component:  ActivityComponent},
  { path: 'password', component:  PasswordComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
