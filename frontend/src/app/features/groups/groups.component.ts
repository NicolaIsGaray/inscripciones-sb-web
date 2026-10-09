import { Component, OnInit, inject, signal } from '@angular/core';
import { IconComponent } from '../../shared/components/icon/icon.component';
import { SettingsService } from '../../service/settings.service';

@Component({
  imports: [IconComponent],
  selector: 'app-groups',
  templateUrl: './groups.component.html'
})
export class GroupsComponent implements OnInit {
  private settingsService = inject(SettingsService);

  // Refleja la disponibilidad real que habilita el secretario desde el panel.
  // Arranca en false (estado seguro) y se actualiza al entrar a la sección.
  protected readonly groupsEnabled = signal(false);

  ngOnInit(): void {
    this.settingsService.getSettings().subscribe({
      next: (settings) => this.groupsEnabled.set(settings.groupsEnabled === true),
      // Si la API no responde se mantiene el estado cerrado
      error: () => this.groupsEnabled.set(false)
    });
  }
}