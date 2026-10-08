import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';
import { Observable } from 'rxjs';
import { Group } from '../model/group';

@Service()
export class GroupService {
    private apiUrl = 'http://localhost:8080/api';
    private httpClient = inject(HttpClient);

    // Obtener todos los grupos
    getGroups(): Observable<Group[]> {
        return this.httpClient.get<Group[]>(`${this.apiUrl}/groups`);
    }

    // Obtener un grupo por ID
    getGroup(id: string): Observable<Group> {
        return this.httpClient.get<Group>(`${this.apiUrl}/groups/${id}`);
    }

    // Crear un nuevo grupo
    createGroup(group: Group): Observable<Group> {
        return this.httpClient.post<Group>(`${this.apiUrl}/groups`, group);
    }

    // Actualizar un grupo existente
    updateGroup(id: string, group: Group): Observable<Group> {
        return this.httpClient.put<Group>(`${this.apiUrl}/groups/${id}`, group);
    }

    // Eliminar un grupo
    deleteGroup(id: string): Observable<void> {
        return this.httpClient.delete<void>(`${this.apiUrl}/groups/${id}`);
    }

    // Quitar un integrante de un grupo
    removeMember(groupId: string, memberId: string): Observable<Group> {
        return this.httpClient.delete<Group>(`${this.apiUrl}/admin/groups/${groupId}/members/${memberId}`);
    }
}
