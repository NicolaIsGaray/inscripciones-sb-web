import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';
import { API_URL } from '../core/config/api-url.token';
import { AppSettings, UpdateAppSettings } from '../model/setting';

@Injectable({
  providedIn: 'root'
})
export class SettingsService {
  private apiUrl = inject(API_URL);
  private httpClient = inject(HttpClient);

  // Estado de la formación de grupos (público)
  getSettings(): Observable<AppSettings> {
    return this.httpClient.get<AppSettings>(`${this.apiUrl}/settings`);
  }

  // Habilitar o deshabilitar la formación de grupos (requiere token de administrador)
  updateSettings(changes: UpdateAppSettings): Observable<AppSettings> {
    return this.httpClient.patch<AppSettings>(`${this.apiUrl}/admin/settings`, changes);
  }
}