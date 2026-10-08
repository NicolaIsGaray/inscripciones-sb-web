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
  students: Student[] = [];
  filteredStudents: Student[] = [];
  paginatedStudents: Student[] = [];

  // Filters
  selectedInstances = signal<string[]>([]);
  selectedTitles = signal<string[]>([]);

  // Search
  dniQuery = signal<string>("");

  // Confirmation flow
  confirming = signal<boolean>(false);
  justConfirmed = signal<boolean>(false);
  selected = signal<Student | null>(null);
  email = signal<string>("");
  confirmedDnIs = signal<string[]>([]);

  // Pagination
  currentPage: number = 1;
  itemsPerPage: number = 5;
  totalPages: number = 1;

  constructor(
    private studentService: StudentService
  ) {}

  ngOnInit(): void {
    this.loadStudents();
  }

  loadStudents(): void {
    this.studentList$ = this.studentService.getStudents();
    this.studentList$.subscribe(data => {
      this.students = data;
      this.applyFilters();
    });
  }

  // Computed values
  filtered = computed(() => {
    const instances = this.selectedInstances();
    const titles = this.selectedTitles();
    return this.students.filter(item => {
      const instanceMatch = instances.length === 0 || instances.includes(item.instance);
      const titleMatch = titles.length === 0 || titles.includes(item.title);
      return instanceMatch && titleMatch;
    });
  });

  normalizedQuery = computed(() => this.dniQuery().replace(/\D/g, ""));

  matches = computed(() => {
    const q = this.normalizedQuery();
    return q.length >= 8 ? this.students.filter(item => item.dni === q) : [];
  });

  // Filter methods
  toggleInstance(instance: string) {
    this.selectedInstances.update(current => current.includes(instance) ? [] : [instance]);
  }

  toggleTitle(title: string) {
    this.selectedTitles.update(current =>
      current.includes(title) ? current.filter(item => item !== title) : [...current, title]
    );
  }

  clearFilters() {
    this.selectedInstances.set([]);
    this.selectedTitles.set([]);
  }

  // Search methods
  selectStudent(student: Student) {
    this.selected.set(student);
    this.dniQuery.set(student.dni);
    this.email.set("");
  }

  // Confirmation methods
  submitConfirmation(event: Event) {
    event.preventDefault();
    if (this.email() && this.selected()) {
      this.confirmedDnIs.update(current => [...current, this.selected()!.dni]);
      this.justConfirmed.set(true);
      this.confirming.set(false);
    }
  }

  resetConfirmation() {
    this.confirming.set(false);
    this.justConfirmed.set(false);
    this.selected.set(null);
    this.email.set("");
    this.dniQuery.set("");
  }

  // Pagination methods
  applyFilters(): void {
    this.totalPages = Math.ceil(this.filtered().length / this.itemsPerPage);
    this.currentPage = 1;
    this.updatePagination();
  }

  updatePagination(): void {
    const start = (this.currentPage - 1) * this.itemsPerPage;
    const end = start + this.itemsPerPage;
    this.paginatedStudents = this.filtered().slice(start, end);
  }

  changePage(direction: number): void {
    const newPage = this.currentPage + direction;
    if (newPage >= 1 && newPage <= this.totalPages) {
      this.currentPage = newPage;
      this.updatePagination();
    }
  }
}
