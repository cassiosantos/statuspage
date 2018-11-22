import { Incident } from '../incident/incident'
export class Resource {
    name: string
    groups: string[]
    address: string
    incidents_history: Incident[]
  }