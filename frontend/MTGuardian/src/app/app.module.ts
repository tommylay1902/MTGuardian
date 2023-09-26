import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { PrescriptionComponent } from './components/prescription/prescription.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { MenuBarComponent } from './components/menu-bar/menu-bar.component';
import { MenuItemComponent } from './components/menu-item/menu-item.component';
import { TreeTableModule } from 'primeng/treetable';

@NgModule({
  declarations: [
    AppComponent,
    PrescriptionComponent,
    MenuBarComponent,
    MenuItemComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    TreeTableModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
