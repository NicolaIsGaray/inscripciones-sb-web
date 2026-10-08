import { Routes } from '@angular/router';
import { GroupsComponent } from './features/groups/groups.component';
import { HomeComponent } from './features/home/home.component';
import { DatesComponent } from './features/dates/dates.component';

export const routes: Routes = [
    {path: "inicio", component: HomeComponent},
    {path: "grupos", component: GroupsComponent},
    {path: "fechas", component: DatesComponent}
];
