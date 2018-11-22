import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { IncidentDetail } from '../incident/incident'

@Injectable({
  providedIn: 'root'
})
export class IncidentsService {

  constructor(private http: HttpClient) { }

  getIncidents(month: number) {
    return this.http.get<IncidentDetail[]>(`http://localhost:4200/api/incidents?month=` + month);
  }
}
