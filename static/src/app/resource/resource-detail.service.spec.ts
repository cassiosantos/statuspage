import { TestBed } from '@angular/core/testing';

import { ResourceDetailService } from './resource-detail.service';

describe('ResourceDetailService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: ResourceDetailService = TestBed.get(ResourceDetailService);
    expect(service).toBeTruthy();
  });
});
