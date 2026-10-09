import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { scheduleExtras, scheduleStages } from '../../core/data/schedule-data';
import { IconComponent } from '../../shared/components/icon/icon.component';

@Component({
  imports: [CommonModule, IconComponent],
  selector: 'app-dates',
  templateUrl: './dates.component.html',
})
export class DatesComponent {
  readonly stages = scheduleStages;
  readonly extras = scheduleExtras;
}