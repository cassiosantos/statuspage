import { TestBed } from '@angular/core/testing';

import { IncidentsService } from './incidents.service';

describe('IncidentsService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: IncidentsService = TestBed.get(IncidentsService);
    expect(service).toBeTruthy();
  });
});
