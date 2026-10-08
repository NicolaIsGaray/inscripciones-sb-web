import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';
import { Observable } from 'rxjs';
import { Student } from '../model/student';

// Respuesta del endpoint de importación
export interface ImportResult {
  success: number;
  errors: string[];
}

@Service()
export class StudentService {
    private apiUrl = 'http://localhost:8080/api';
    private httpClient = inject(HttpClient);

    // Obtener todos los alumnos
    getStudents(): Observable<Student[]> {
        return this.httpClient.get<Student[]>(`${this.apiUrl}/students`);
    }

    // Obtener un alumno por ID
    getStudent(id: string): Observable<Student> {
        return this.httpClient.get<Student>(`${this.apiUrl}/students/${id}`);
    }

    // Crear un nuevo alumno
    createStudent(student: Student): Observable<Student> {
        return this.httpClient.post<Student>(`${this.apiUrl}/students`, student);
    }

    // Actualizar un alumno existente
    updateStudent(id: string, student: Student): Observable<Student> {
        return this.httpClient.put<Student>(`${this.apiUrl}/students/${id}`, student);
    }

    // Confirmar el banco de un alumno (flujo público del ingresante)
    confirmStudent(id: string, email: string): Observable<Student> {
        return this.httpClient.post<Student>(`${this.apiUrl}/students/${id}/confirm`, { email });
    }

    // Eliminar un alumno
    deleteStudent(id: string): Observable<void> {
        return this.httpClient.delete<void>(`${this.apiUrl}/students/${id}`);
    }

    // Importar alumnos desde Excel
    // La escuela se toma de la columna E del archivo; si falta, queda secundaria
    importStudents(file: File): Observable<ImportResult> {
        const formData = new FormData();
        formData.append('file', file);
        return this.httpClient.post<ImportResult>(`${this.apiUrl}/admin/students/import`, formData);
    }

    // Guardar alumno (crear o actualizar)
    saveStudent(student: Student): Observable<Student> {
        if (student.id) {
            return this.updateStudent(student.id, student);
        }
        return this.createStudent(student);
    }
}
