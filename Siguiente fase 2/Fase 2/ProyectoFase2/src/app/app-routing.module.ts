import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { InicioComponent } from './componentes/inicio/inicio.component';
import { CargarTiendasComponent } from './componentes/cargar-tiendas/cargar-tiendas.component';
import { InventarioComponent } from './componentes/inventario/inventario.component';
import { CarritoDeComprasComponent } from './componentes/carrito-de-compras/carrito-de-compras.component';
import { CalendarioPedidosComponent } from './componentes/calendario-pedidos/calendario-pedidos.component';
import { CargarPedidosComponent } from './componentes/cargar-pedidos/cargar-pedidos.component';
import { CargarInventarioComponent } from './componentes/cargar-inventario/cargar-inventario.component';
import { PruebaComponent } from './componentes/prueba/prueba.component';
import { TiendasComponent } from "./componentes/tiendas/tiendas.component";
import { BusquedaespecificaComponent } from './componentes/busquedaespecifica/busquedaespecifica/busquedaespecifica.component';
import { EliminarespecificaComponent } from './componentes/eliminarespecifica/eliminarespecifica/eliminarespecifica.component';
import { CreacionUsuarioComponent } from './componentes/creacion-usuario/creacion-usuario.component';
import { EliminacionUsuarioComponent } from './componentes/eliminacion-usuario/eliminacion-usuario.component';

const routes: Routes = [
  {
    path: '',
    component: InicioComponent,
  },
  {
    path: 'cargarTiendas',
    component: CargarTiendasComponent,
  },
  {
    path: 'inventario',
    component: InventarioComponent,
  },
  {
    path: 'carritoDeCompras',
    component: CarritoDeComprasComponent,
  },
  {
    path: 'vistaCalendario',
    component: CalendarioPedidosComponent,
  },
  {
    path: 'cargarPedidos',
    component: CargarPedidosComponent,
  },
  {
    path: 'cargarInventario',
    component: CargarInventarioComponent,
  },
  {
    path: 'prueba',
    component: PruebaComponent,
  },
  {
    path: 'tiendas',
    component: TiendasComponent,
  },
  {
    path: 'busquedaEspecifica',
    component: BusquedaespecificaComponent,
  },
  {
    path: 'eliminarEspecifica',
    component: EliminarespecificaComponent,
  },
  {
    path: "CreacionUsuario",
    component: CreacionUsuarioComponent,
  },
  {
    path: "EliminarUsuario",
    component: EliminacionUsuarioComponent,
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
