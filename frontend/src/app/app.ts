import { Component, OnInit, computed, signal } from '@angular/core';
import { RouterLink, RouterModule, Router, NavigationEnd } from '@angular/router';
import { filter } from 'rxjs';
import { IconName } from './core/models/school.model';
import { IconComponent } from './shared/components/icon/icon.component';

@Component({
  imports: [RouterModule, RouterLink, IconComponent],
  selector: 'app-root',
  styleUrl: './app.css',
  templateUrl: './app.html',
})
export class App implements OnInit {
  protected readonly title = signal('sb-web');

  currentPage = signal<string>("/inicio");

  // El panel de administración ocupa toda la pantalla: sin header, barra
  // inferior ni footer del sitio (ver .app--admin en styles.css)
  isAdmin = computed(() => this.currentPage().startsWith('/admin'));

  navItems: { route: string; label: string; icon: IconName }[] = [
    { route: "/inicio", label: "Inicio", icon: "home" },
    { route: "/grupos", label: "Grupos", icon: "users" },
    { route: "/fechas", label: "Fechas", icon: "calendar" },
  ];

  constructor(private router: Router) {}

  ngOnInit() {
    // Escuchar cambios de navegación
    this.router.events
      .pipe(filter(event => event instanceof NavigationEnd))
      .subscribe((event: NavigationEnd) => {
        this.currentPage.set(event.urlAfterRedirects);
      });
  }
}
