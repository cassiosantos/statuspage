import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
 name: 'filterResourceName'
})
export class FilterResourceName implements PipeTransform{

  transform(resources: any, searchText: any): any {
    if(searchText == null) return resources;
    if(resources == null) return resources;

    return resources.filter(function(resource){
      return resource.name.toLowerCase().indexOf(searchText.toLowerCase()) > -1;
    })
  }

}