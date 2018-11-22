import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Resource } from '../resource/resource';
import { ResourceService } from '../resource/resource.service';
import { Incident } from '../incident/incident';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: [ './dashboard.component.css' ]
})
export class DashboardComponent implements OnInit {
  resources: Resource[] = [];
  interval: any;
  searchResource: string = '';
  constructor(private resourceService: ResourceService, private router:Router) { }

  ngOnInit() {
    this.getResources();
  }

  getLastIncident(incidents): Incident {
    return incidents[incidents.length -1]
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
  
  getResources(): void {
    this.resourceService.getResources()
      .subscribe(resources => this.resources = resources);
  }
}