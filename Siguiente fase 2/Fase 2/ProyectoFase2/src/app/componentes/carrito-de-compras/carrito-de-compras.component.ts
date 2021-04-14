import { Component, OnInit } from '@angular/core';
import { Inventario } from 'src/app/models/inventario/inventario';
import { InventarioService } from "../../services/inventario/inventario.service";
import { HttpClient } from "@angular/common/http";
//import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms';
import { Pedidos } from "../../models/pedidos/pedidos";
import { Infopedido } from "../../models/infopedido/infopedido";
import { Pedirproductos } from "../../models/pedirproductos/pedirproductos";

@Component({
  selector: 'app-carrito-de-compras',
  templateUrl: './carrito-de-compras.component.html',
  styleUrls: ['./carrito-de-compras.component.css']
})
export class CarritoDeComprasComponent implements OnInit {

  lista_estudiantes: number[]=[]

  cargarinventario: any
  //tienda: Tienda

  mostrarMensajeError=false
  mostrarMensaje=false
  mensajeError = ''

  opcion!: string;

  numeroCodigos: number[]=[]

  Fecha = new FormControl('')
  Tienda = new FormControl('')
  Departamento = new FormControl('')
  Calificacion = new FormControl('')
  Codigo = new FormControl('')
  CodigoElim = new FormControl('')

  numero = 0;

  constructor(private inventarioService: InventarioService) { }

  ngOnInit(): void {

    this.inventarioService.getListaInventarios().subscribe((dataList:Inventario[])=>{
      console.log(dataList)


      for(var contador=0; contador<dataList.length; contador++ ){
        console.log(dataList[contador])
        this.cargarinventario = dataList[contador]


        var data=dataList[contador].Inventarios
        console.log(data)

        for(var contar=0; contar<data.length; contar++ ){

          var inf = data[contar].Productos
          console.log(inf)

          for(var num=0; num<inf.length; num++){

            console.log(this.cargarinventario)

          }


        }

      }
      
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Inventarios'
    })
    
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

  crearPedido(){

  }

  agregarProducto(){

    this.numeroCodigos[this.numero] = Number(this.Codigo.value)

    //var produc ={
      //Codigo: Number(this.numeroCodigos)
    //}


    const productos: Pedirproductos={
      Codigo: Number(this.numeroCodigos[this.numero]),
    }


    const infopedido: Infopedido={
      Fecha: this.Fecha.value,
      Tienda: this.Tienda.value,
      Departamento: this.Departamento.value,
      Calificacion: Number(this.Calificacion.value),
      Productos:[productos]
    }

    const pedido: Pedidos={
      Pedidos: [infopedido]
    }

    console.log(this.numeroCodigos)
    console.log(infopedido)
    console.log(pedido)

    this.numero = this.numero + 1

  }

  eliminarProducto(){

  }

}
