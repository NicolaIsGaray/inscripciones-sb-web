import { Component, OnInit, Input, computed, signal } from '@angular/core';
import { Observable } from 'rxjs';
import { Student } from '../../../model/student';
import { StudentService } from '../../../service/student.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IconComponent } from '../../../shared/components/icon/icon.component';

@Component({
  imports: [CommonModule, FormsModule, IconComponent],
  selector: 'app-student-list',
  templateUrl: './student-list.component.html',
})
export class StudentListComponent implements OnInit {
  @Input() school: 'secundaria' | 'deportiva' = 'secundaria';

  // Data from backend
  studentList$!: Observable<Student[]>;
  students = signal<Student[]>([]);

  // Filters
  selectedInstances = signal<string[]>([]);
  selectedTitles = signal<string[]>([]);

  // Search
  dniQuery = signal<string>('');

  // Confirmation flow
  confirming = signal<boolean>(false);
  justConfirmed = signal<boolean>(false);
  selected = signal<Student | null>(null);
  email = signal<string>('');
  submitting = signal<boolean>(false);
  confirmError = signal<string>('');

  // Pagination
  currentPage = signal<number>(1);
  itemsPerPage: number = 5;

  // Email inválido (solo formato: el email lo carga el propio ingresante)
  invalidEmail = computed(() => {
    const value = this.email().trim();
    return value.length > 0 && !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value);
  });

  constructor(private studentService: StudentService) {}

  ngOnInit(): void {
    this.loadStudents();
  }

  loadStudents(): void {
    this.studentList$ = this.studentService.getStudents();
    this.studentList$.subscribe((data) => {
      this.students.set(data.filter((item) => item.school === this.school));
    });
  }

  // Computed values
  filtered = computed(() => {
    const instances = this.selectedInstances();
    const titles = this.selectedTitles();
    return this.students().filter((item) => {
      const schoolMatch = item.school === this.school;
      const instanceMatch = instances.length === 0 || instances.includes(item.instance);
      const titleMatch = titles.length === 0 || titles.includes(item.title);
      return schoolMatch && instanceMatch && titleMatch;
    });
  });

  // Opciones de filtro disponibles (derivadas de los alumnos de esta escuela)
  availableInstances = computed(() =>
    Array.from(
      new Set(
        this.students()
          .map((s) => s.instance)
          .filter(Boolean),
      ),
    ),
  );

  availableTitles = computed(() =>
    Array.from(
      new Set(
        this.students()
          .map((s) => s.title)
          .filter(Boolean),
      ),
    ),
  );

  // Paginación: derivada de `filtered()`, se repinta sola cuando cambian los datos
  totalPages = computed(() => Math.ceil(this.filtered().length / this.itemsPerPage));

  paginatedStudents = computed(() => {
    const total = this.totalPages();
    const page = Math.min(Math.max(this.currentPage(), 1), Math.max(total, 1));
    const start = (page - 1) * this.itemsPerPage;
    return this.filtered().slice(start, start + this.itemsPerPage);
  });

  normalizedQuery = computed(() => this.dniQuery().replace(/\D/g, ''));

  matches = computed(() => {
    const q = this.normalizedQuery();
    return q.length >= 8
      ? this.students().filter((item) => item.dni === q && item.school === this.school)
      : [];
  });

  // Filter methods
  toggleInstance(instance: string) {
    this.selectedInstances.update((current) => (current.includes(instance) ? [] : [instance]));
    this.currentPage.set(1);
  }

  toggleTitle(title: string) {
    this.selectedTitles.update((current) =>
      current.includes(title) ? current.filter((item) => item !== title) : [...current, title],
    );
    this.currentPage.set(1);
  }

  clearFilters() {
    this.selectedInstances.set([]);
    this.selectedTitles.set([]);
    this.currentPage.set(1);
  }

  // Search methods
  selectStudent(student: Student) {
    // Si el banco ya estaba confirmado, mostramos el estado final y no el formulario
    if (student.confirmed) {
      this.selected.set(student);
      this.confirming.set(false);
      this.justConfirmed.set(true);
      return;
    }

    this.selected.set(student);
    this.dniQuery.set(student.dni);
    this.email.set('');
    this.confirmError.set('');
  }

  // Confirmation methods
  submitConfirmation(event: Event) {
    event.preventDefault();

    const student = this.selected();
    const email = this.email().trim();

    if (!student || this.submitting()) {
      return;
    }

    if (!email || this.invalidEmail()) {
      this.confirmError.set('Ingresá un correo electrónico válido.');
      return;
    }

    this.submitting.set(true);
    this.confirmError.set('');

    this.studentService.confirmStudent(student.id, email).subscribe({
      next: (updated) => {
        // Reemplazamos el alumno en el signal: la fila pasa a ✓ sola
        this.students.update((list) => list.map((item) => (item.id === updated.id ? updated : item)));
        this.selected.set(null);
        this.email.set('');
        this.dniQuery.set('');
        this.submitting.set(false);
        this.confirming.set(false);
        this.justConfirmed.set(true);
      },
      error: () => {
        this.confirmError.set('No pudimos confirmar tu banco. Intentá nuevamente en unos minutos.');
        this.submitting.set(false);
      },
    });
  }

  resetConfirmation() {
    this.confirming.set(false);
    this.justConfirmed.set(false);
    this.selected.set(null);
    this.email.set('');
    this.dniQuery.set('');
    this.submitting.set(false);
    this.confirmError.set('');
  }

  // Pagination methods
  changePage(direction: number): void {
    const total = Math.max(this.totalPages(), 1);
    const newPage = this.currentPage() + direction;
    if (newPage >= 1 && newPage <= total) {
      this.currentPage.set(newPage);
    }
  }
}
