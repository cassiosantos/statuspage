import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
 name: 'orderByStatus'
})
export class OrderByStatus implements PipeTransform{

 transform(array: Array<string>, args: string): Array<string> {

  if(!array || array === undefined || array.length === 0) return null;

    array.sort((a: any, b: any) => {
      var lastA = a.incidents_history[a.incidents_history.length -1]
      var lastB = b.incidents_history[b.incidents_history.length -1]
      if (lastA.status < lastB.status) {
        return -1;
      } else if (lastA.status > lastB.status) {
        return 1;
      } else {
        return 0;
      }
    });
    return array;
  }

}