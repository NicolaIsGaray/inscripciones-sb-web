import { Component, Input, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-button',
  standalone: true,
  imports: [CommonModule],
  template: `
    <button 
      [className]="'button button--' + variant + ' ' + className" 
      [type]="type" 
      [disabled]="disabled" 
      (click)="onClick.emit()">
      <ng-content></ng-content>
    </button>
  `
})
export class ButtonComponent {
  @Input() variant: "primary" | "secondary" | "ghost" | "danger" = "primary";
  @Input() type: "button" | "submit" = "button";
  @Input() disabled: boolean = false;
  @Input() className: string = "";
  @Output() onClick = new EventEmitter<void>();
}