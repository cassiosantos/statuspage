import { Component, Input } from '@angular/core';
import { Resource } from '../resource/resource';
import { ResourceDetailService } from '../resource/resource-detail.service';
import { Router, ActivatedRoute } from '@angular/router';


@Component({
  selector: 'app-resource-details',
  templateUrl: './resource-details.component.html',
  styleUrls: ['./resource-details.component.css']
})
export class ResourceDetailsComponent {
  @Input() resource: Resource;
  constructor(private resourceDetailService: ResourceDetailService, private activeRoute: ActivatedRoute, private router : Router) { }

  showIncidents(resourceID: string): void {
      this.router.navigateByUrl("/resource/"+resourceID);
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
