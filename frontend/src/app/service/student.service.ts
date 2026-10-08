import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';
import { Observable } from 'rxjs';
import { Student } from '../model/student';

@Service()
export class StudentService {
    private httpClient = inject(HttpClient);

    getStudents(): Observable<Student[]> {
        return this.httpClient.get<Student[]>('/students');
    }
}
