import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IconName } from '../../../core/models/school.model';

@Component({
  selector: 'app-icon',
  standalone: true,
  imports: [CommonModule],
  template: `
    <svg [attr.width]="size" [attr.height]="size" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true" [ngSwitch]="name">
      <ng-container *ngSwitchCase="'home'"><path d="m3 11 9-8 9 8" /><path d="M5 10v11h14V10M9 21v-7h6v7" /></ng-container>
      <ng-container *ngSwitchCase="'users'"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" /><circle cx="9" cy="7" r="4" /><path d="M22 21v-2a4 4 0 0 0-3-3.87M16 3.13a4 4 0 0 1 0 7.75" /></ng-container>
      <ng-container *ngSwitchCase="'calendar'"><rect x="3" y="5" width="18" height="16" rx="2" /><path d="M16 3v4M8 3v4M3 11h18" /></ng-container>
      <ng-container *ngSwitchCase="'school'"><path d="m3 10 9-6 9 6-9 6-9-6Z" /><path d="M7 14v4c3 2 7 2 10 0v-4M21 10v6" /></ng-container>
      <ng-container *ngSwitchCase="'search'"><circle cx="11" cy="11" r="7" /><path d="m20 20-4-4" /></ng-container>
      <ng-container *ngSwitchCase="'check'"><path d="m5 12 4 4L19 6" /></ng-container>
      <ng-container *ngSwitchCase="'arrow'"><path d="M5 12h14M13 6l6 6-6 6" /></ng-container>
      <ng-container *ngSwitchCase="'trash'"><path d="M3 6h18M8 6V4h8v2M6 6l1 15h10l1-15M10 11v5M14 11v5" /></ng-container>
      <ng-container *ngSwitchCase="'book'"><path d="M4 19.5A2.5 2.5 0 0 1 6.5 17H20V3H6.5A2.5 2.5 0 0 0 4 5.5v14Z" /><path d="M8 7h8M8 11h6" /></ng-container>
      <ng-container *ngSwitchCase="'mail'"><rect x="3" y="5" width="18" height="14" rx="2" /><path d="m3 7 9 6 9-6" /></ng-container>
      <ng-container *ngSwitchCase="'clock'"><circle cx="12" cy="12" r="9" /><path d="M12 7v6l4 2" /></ng-container>
      <ng-container *ngSwitchCase="'plus'"><path d="M12 5v14M5 12h14" /></ng-container>
      <ng-container *ngSwitchCase="'close'"><path d="m6 6 12 12M18 6 6 18" /></ng-container>
      <ng-container *ngSwitchCase="'shield'"><path d="M12 22s8-3 8-10V5l-8-3-8 3v7c0 7 8 10 8 10Z" /><path d="m9 12 2 2 4-5" /></ng-container>
    </svg>
  `
})
export class IconComponent {
  @Input({ required: true }) name!: IconName;
  @Input() size: number = 22;
}