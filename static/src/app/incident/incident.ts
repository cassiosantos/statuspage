export class Incident {
  status: number
  description: string
  occurrence_date: string
}

export class IncidentDetail {
  component: string
  incident: Incident
}