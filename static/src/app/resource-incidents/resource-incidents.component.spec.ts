import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ResourceIncidentsComponent } from './resource-incidents.component';

describe('ResourceIncidentsComponent', () => {
  let component: ResourceIncidentsComponent;
  let fixture: ComponentFixture<ResourceIncidentsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ResourceIncidentsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ResourceIncidentsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
