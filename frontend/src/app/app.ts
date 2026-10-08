import { Component, OnInit, signal } from '@angular/core';
import { RouterLink, RouterModule, Router, NavigationEnd } from '@angular/router';
import { filter } from 'rxjs';

@Component({
  imports: [RouterModule, RouterLink],
  selector: 'app-root',
  styleUrl: './app.css',
  templateUrl: './app.html',
})
export class App implements OnInit {
  protected readonly title = signal('sb-web');

  currentPage = signal<string>("/inicio");

  navItems = [
    { route: "/inicio", label: "Inicio" as const },
    { route: "/grupos", label: "Grupos" as const },
    { route: "/fechas", label: "Fechas" as const },
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
