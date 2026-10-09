import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';
import { Observable } from 'rxjs';
import { Group } from '../model/group';
import { API_URL } from '../core/config/api-url.token';

@Service()
export class GroupService {
    private apiUrl = inject(API_URL);
    private httpClient = inject(HttpClient);

    // Obtener todos los grupos (público)
    getGroups(): Observable<Group[]> {
        return this.httpClient.get<Group[]>(`${this.apiUrl}/groups`);
    }

    // Obtener un grupo por ID (público)
    getGroup(id: string): Observable<Group> {
        return this.httpClient.get<Group>(`${this.apiUrl}/groups/${id}`);
    }

    // Crear un nuevo grupo (requiere token de administrador)
    createGroup(group: Group): Observable<Group> {
        return this.httpClient.post<Group>(`${this.apiUrl}/admin/groups`, group);
    }

    // Actualizar un grupo existente (requiere token de administrador)
    updateGroup(id: string, group: Group): Observable<Group> {
        return this.httpClient.put<Group>(`${this.apiUrl}/admin/groups/${id}`, group);
    }

    // Eliminar un grupo (requiere token de administrador)
    deleteGroup(id: string): Observable<void> {
        return this.httpClient.delete<void>(`${this.apiUrl}/admin/groups/${id}`);
    }

    // Quitar un integrante de un grupo
    removeMember(groupId: string, memberId: string): Observable<Group> {
        return this.httpClient.delete<Group>(`${this.apiUrl}/admin/groups/${groupId}/members/${memberId}`);
    }
}
