import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule, FormsModule } from "@angular/forms";
import { HttpClientModule } from "@angular/common/http";

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { InicioComponent } from './componentes/inicio/inicio.component';
import { CargarTiendasComponent } from './componentes/cargar-tiendas/cargar-tiendas.component';
import { InventarioComponent } from './componentes/inventario/inventario.component';
import { CarritoDeComprasComponent } from './componentes/carrito-de-compras/carrito-de-compras.component';
import { CalendarioPedidosComponent } from './componentes/calendario-pedidos/calendario-pedidos.component';
import { CargarPedidosComponent } from './componentes/cargar-pedidos/cargar-pedidos.component';
import { CargarInventarioComponent } from './componentes/cargar-inventario/cargar-inventario.component';
import { PruebaComponent } from './componentes/prueba/prueba.component';
import { TiendasComponent } from './componentes/tiendas/tiendas.component';
import { ObjToArrayPipe } from './objToArray.pipe';
import { BusquedaespecificaComponent } from './componentes/busquedaespecifica/busquedaespecifica/busquedaespecifica.component';
import { EliminarespecificaComponent } from './componentes/eliminarespecifica/eliminarespecifica/eliminarespecifica.component';

@NgModule({
  declarations: [
    AppComponent,
    InicioComponent,
    CargarTiendasComponent,
    InventarioComponent,
    CarritoDeComprasComponent,
    CalendarioPedidosComponent,
    CargarPedidosComponent,
    CargarInventarioComponent,
    PruebaComponent,
    TiendasComponent,
    ObjToArrayPipe,
    BusquedaespecificaComponent,
    EliminarespecificaComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }