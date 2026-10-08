import { Component } from '@angular/core';
import { StudentListComponent } from './student-list/student-list.component';
import { IconComponent } from '../../shared/components/icon/icon.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [StudentListComponent, IconComponent],
  templateUrl: './home.component.html'
})
export class HomeComponent {
}
