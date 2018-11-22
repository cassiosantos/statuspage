import { Component, OnInit } from '@angular/core';
import { IncidentsService } from '../incident/incidents.service'
import { IncidentDetail } from '../incident/incident';
import { load } from '@angular/core/src/render3/instructions';

@Component({
  selector: 'app-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.css']
})
export class HistoryComponent implements OnInit {
  incidents = [];
  incidentsDate = [];
  incidentsDetails = [];
  constructor(private incidentsService: IncidentsService) { }

  
  currentDate = new Date();
  monthView = this.currentDate.getMonth();
  weeks = Array(6).fill(Array(7).fill(""));
  weekDays = ['S','M','T','W','T','F','S'];

  ngOnInit() {
    this.load();
    this.showIncidentsDetail(this.currentDate.getDate())
  }

  load(){
    this.generateCalendar(this.currentDate)
    this.getIncidents(this.currentDate.getMonth());
  }

  getIncidents(month: number) {
    this.incidentsDate = [];
    this.incidentsService.getIncidents(month).subscribe(
      (incidents: IncidentDetail[]) => {
        var incidentsDate = [];
        if (incidents){
          incidents.forEach(function (i){
            var dt = new Date(i.incident.occurrence_date).getDate()
            incidentsDate.push(dt);
          });
          this.incidents = incidents; 
          this.incidentsDate = incidentsDate;
        }
      });
  }

  getIcon(status): string {
    switch(status){
      case 1:
        return "warning"
      case 2:
        return "error"        
    default:
      return "check_circle"
    }
  }

  nextMonth(): void {
    this.currentDate.setMonth(this.currentDate.getMonth()+1);
    this.load();
  }

  prevMonth(): void {
    this.currentDate.setMonth(this.currentDate.getMonth()-1);
    this.load();
  }

  generateCalendar(date: Date): void {
    var weekStart = new Date(date.getFullYear(),date.getMonth(),1).getDay();
    var lastDay = new Date(date.getFullYear(),date.getMonth() +1 ,0).getDate();
    var month = [], weekLenght = 7;
    var preset = Array(weekStart).fill("")
    var days = Array.from({length: lastDay}, (x,i) => (i+1));
    days = preset.concat(days)
    
    while (days.length > 0){
      month.push(days.splice(0, weekLenght));
    }
      this.weeks = month;
  }

  hasIncident(day: number) {
    return this.incidentsDate.indexOf(day) > -1
  }

  isToday(day){
    var today = new Date()
    today.setHours(0,0,0,0)
    var current = new Date(this.currentDate.getFullYear(),this.currentDate.getMonth(),day)
    return today.getTime() == current.getTime();
  }

  showIncidentsDetail(date: number){
    this.incidentsDetails = this.incidents.filter(function(d) {
      var od = new Date(d.incident.occurrence_date).getDate();
      return date == od;
    });
  }
}
