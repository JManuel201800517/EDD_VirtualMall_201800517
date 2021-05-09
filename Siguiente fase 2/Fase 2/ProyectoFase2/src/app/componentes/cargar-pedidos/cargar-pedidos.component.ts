import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms';
import { PedidosService } from "../../services/pedidos/pedidos.service";

@Component({
  selector: 'app-cargar-pedidos',
  templateUrl: './cargar-pedidos.component.html',
  styleUrls: ['./cargar-pedidos.component.css']
})
export class CargarPedidosComponent implements OnInit {

  peticion = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;
  mensajeError = ''

  constructor(private pedidosservice: PedidosService) { }

  ngOnInit(): void {
  }

  subirPedidos(){
    console.log("Funciona muy bien")
    //console.log(this.inventario.value)


    this.pedidosservice.postPedido(this.peticion.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.peticion.setValue("")
      console.log("Pedidos Cargados")
      console.log(res)

    }, (err)=>{
      this.mostrarMensajeError=true
    })


    this.pedidosservice.getMerklePedidos().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

}
