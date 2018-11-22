import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Resource } from './resource'

@Injectable({
  providedIn: 'root'
})
export class ResourceDetailService {

  constructor(private http: HttpClient) { }

  getResource(id) {
    return this.http.get<Resource>(`http://localhost:4200/api/component/` + id);
  }
}
