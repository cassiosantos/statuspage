import { Component, OnInit } from '@angular/core';
import { Resource } from '../resource/resource';
import { ResourceDetailService } from '../resource/resource-detail.service';
import { Router, ActivatedRoute } from '@angular/router';


@Component({
  selector: 'app-resource-incidents',
  templateUrl: './resource-incidents.component.html',
  styleUrls: ['./resource-incidents.component.css']
})
export class ResourceIncidentsComponent implements OnInit{
  resource: Resource = new Resource();
  maxResources: number = 10;
  constructor(private resourceDetailService: ResourceDetailService, private activeRoute: ActivatedRoute, private router : Router) { }

  ngOnInit() {
    this.getResource()
  }

  hasMoreIncidents(): boolean {
    if (this.resource.incidents_history === undefined) return false;
    return this.resource.incidents_history.length > this.maxResources
  }

  getResource(): void {
    this.activeRoute.params.subscribe(routeParams => {

      this.resourceDetailService.getResource(routeParams.id).subscribe(

        (resource: Resource) => {
          this.resource = resource;
        });

    });
  }


  showMoreIncidents(): void {
    this.maxResources += 10;
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
}
