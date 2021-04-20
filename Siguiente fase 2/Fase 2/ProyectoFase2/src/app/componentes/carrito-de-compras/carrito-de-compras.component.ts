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

  pedir: Pedidos[]=[]

  Fecha = new FormControl('')
  Tienda = new FormControl('')
  Departamento = new FormControl('')
  Calificacion = new FormControl('')
  Codigo = new FormControl('')
  CodigoElim = new FormControl('')
  Dpi = new FormControl("")

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
    console.log("Funciona muy bien")
    //console.log(this.inventario.value)

    for(var x = 0; x < this.pedir.length; x++) {

      this.inventarioService.postCarritoDeCompras(this.pedir[x]).subscribe((res:any)=>{
        this.mostrarMensaje=true
        this.Codigo.setValue("")
        console.log("Pedido Subido")
        console.log(res)
  
      }, (err)=>{
        this.mostrarMensajeError=true
      })

    }

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
      Cliente: Number(this.Dpi.value),
      Productos:[productos]
    }

    const pedido: Pedidos={
      Pedidos: [infopedido]
    }

    console.log(this.numeroCodigos)
    //console.log(infopedido)
    //console.log(pedido)

    this.pedir[this.numero] = pedido

    this.numero = this.numero + 1

    console.log(this.pedir)

    this.Codigo.setValue("")


  }
  

  eliminarProducto(){
    const productos: Pedirproductos={
      Codigo: Number(this.CodigoElim.value),
    }


    const infopedido: Infopedido={
      Fecha: this.Fecha.value,
      Tienda: this.Tienda.value,
      Departamento: this.Departamento.value,
      Calificacion: Number(this.Calificacion.value),
      Cliente: Number(this.Dpi.value),
      Productos:[productos]
    }

    const pedido: Pedidos={
      Pedidos: [infopedido]
    }

    for(var x = 0; x < this.pedir.length; x++){

      console.log(this.pedir[x])
      console.log(pedido)

      if(this.pedir[x].Pedidos[0].Productos[0].Codigo == pedido.Pedidos[0].Productos[0].Codigo){
        this.pedir[x].Pedidos[0].Productos[0].Codigo == null
        this.pedir[x].Pedidos[0].Calificacion == null
        this.pedir[x].Pedidos[0].Departamento == null
        this.pedir[x].Pedidos[0].Cliente == null
        this.pedir[x].Pedidos[0].Fecha == null
        this.pedir[x].Pedidos[0].Tienda == null

      }else{
        console.log("Continuar...")
      }
    }

    console.log(this.pedir)




  }

}
