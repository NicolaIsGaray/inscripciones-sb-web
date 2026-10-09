import { InjectionToken } from '@angular/core';

/**
 * URL base de la API.
 * Se provee en app.config.ts desde `environment.apiUrl`, de modo que
 * development y production apunten a destinos distintos sin tocar los servicios.
 */
export const API_URL = new InjectionToken<string>('API_URL');