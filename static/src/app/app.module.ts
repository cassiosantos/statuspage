import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http'; 
import { FormsModule } from '@angular/forms';
import { InfiniteScrollModule } from 'ngx-infinite-scroll';
import { MatCardModule,MatListModule,MatIconModule,MatGridListModule,MatExpansionModule,MatButtonModule,MatInputModule,MatStepperModule,MatToolbarModule } from '@angular/material';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { DashboardComponent } from './dashboard/dashboard.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ResourceDetailsComponent } from './resource-details/resource-details.component';
import { ResourceIncidentsComponent } from './resource-incidents/resource-incidents.component';

import { HistoryComponent } from './history/history.component';
import { ReversePipe } from './reverse.pipe';
import { OrderByStatus } from './pipes/order-by-status.pipe';
import { FilterResourceName } from './pipes/filter-resource-name.pipe';

@NgModule({
  declarations: [
    AppComponent,
    DashboardComponent,
    ResourceDetailsComponent,
    ResourceIncidentsComponent,
    ReversePipe,
    HistoryComponent,
    OrderByStatus,
    FilterResourceName
  ],
  imports: [
    InfiniteScrollModule,
    MatToolbarModule,
    MatStepperModule,
    FormsModule,
    MatInputModule,
    MatButtonModule,
    MatExpansionModule,
    MatCardModule,
    MatListModule,
    MatGridListModule,
    MatIconModule,
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
