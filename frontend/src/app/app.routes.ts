import { Routes } from '@angular/router';
import { GroupsComponent } from './features/groups/groups.component';
import { HomeComponent } from './features/home/home.component';
import { DatesComponent } from './features/dates/dates.component';
import { LoginComponent } from './features/auth/login.component';
import { AdminComponent } from './features/admin/admin.component';
import { AuthGuard } from './guards/auth.guard';

export const routes: Routes = [
    {path: "", redirectTo: "inicio", pathMatch: "full"},
    {path: "inicio", component: HomeComponent},
    {path: "grupos", component: GroupsComponent},
    {path: "fechas", component: DatesComponent},
    {path: "login", component: LoginComponent},
    {path: "admin", component: AdminComponent, canActivate: [AuthGuard]}
];
