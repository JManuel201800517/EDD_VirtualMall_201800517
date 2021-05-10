import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Info } from 'src/app/models/info/info';
import { Infopedido } from 'src/app/models/infopedido/infopedido';
import { Productos } from 'src/app/models/Productos/productos';
import { Usuario } from 'src/app/models/usuario/usuario';
import { InventarioService } from 'src/app/services/inventario/inventario.service';
import { PedidosService } from 'src/app/services/pedidos/pedidos.service';
import { SubirtiendaService } from 'src/app/services/tienda/subirtienda.service';
import { UsuarioService } from 'src/app/services/usuario/usuario.service';

@Component({
  selector: 'app-arbol-merkle',
  templateUrl: './arbol-merkle.component.html',
  styleUrls: ['./arbol-merkle.component.css']
})
export class ArbolMerkleComponent implements OnInit {

  mostrarMensaje = false;
  mostrarMensajeError = false;
  mensajeError = ''

  Fecha = new FormControl('')
  Tienda = new FormControl('')
  Departamento = new FormControl('')
  Calificacion = new FormControl('')
  Dpi = new FormControl("")

  CuentaUsuario!: string;
  DpiUsuario = new FormControl('')
  NombreUsuario = new FormControl('')
  CorreoUsuario = new FormControl('')
  PasswordUsuario = new FormControl('')

  CodigoProducto = new FormControl('')
  NombreProducto = new FormControl('')
  DescripcionProducto = new FormControl('')
  PrecioProducto = new FormControl('')
  CantidadProducto = new FormControl("")
  ImagenProducto = new FormControl('')
  AlmacenProducto = new FormControl("")

  NombreTienda = new FormControl('')
  DescripcionTienda = new FormControl('')
  ContactoTienda = new FormControl('')
  CalificacionTienda = new FormControl('')
  LogoTienda = new FormControl("")
  //SesionTienda = new FormControl("")


  ArPedidos = false;
  ArTiendas = false;
  ArProductos = false;
  ArUsuarios = false;

  NodoPedido = true;
  NodoUsuario = true;
  NodoTienda = true;
  NodoProducto = true;

  ArreglarPedido = true;
  ArreglarUsuario = true;
  ArreglarTienda = true;
  ArreglarProducto = true;


  constructor(private pedidosservice: PedidosService, private inventarioservice: InventarioService, private subirtiendaservice: SubirtiendaService,
    private usuarioService: UsuarioService) { }

  ngOnInit(): void {
  }

  ArbolUsuario(){
    this.ArPedidos = false;
    this.ArTiendas = false;
    this.ArProductos = false;
    this.ArUsuarios = true;

  }

  ArbolTiendas(){
    this.ArPedidos = false;
    this.ArTiendas = true;
    this.ArProductos = false;
    this.ArUsuarios = false;

  }

  ArbolProductos(){
    this.ArPedidos = false;
    this.ArTiendas = false;
    this.ArProductos = true;
    this.ArUsuarios = false;

  }

  ArbolPedidos(){
    this.ArPedidos = true;
    this.ArTiendas = false;
    this.ArProductos = false;
    this.ArUsuarios = false;

  }

  CambiarNodoPedidos(){

    const pedido: Infopedido={
      Fecha: this.Fecha.value,
      Tienda: this.Tienda.value,
      Departamento: this.Departamento.value,
      Calificacion: Number(this.Calificacion.value),
      Cliente: Number(this.Dpi.value),
      Productos:[]
    }


    this.pedidosservice.ConfigMerklePedidos(pedido).subscribe((res:any)=>{
      this.Fecha.setValue("")
      this.Tienda.setValue("")
      this.Departamento.setValue("")
      this.Calificacion.setValue("")
      this.Dpi.setValue("")
      console.log("Arbol Merkle Pedidos Configurado")
    },(err)=>{
      this.mostrarMensajeError=true
    })


    this.NodoPedido = false;
    this.ArreglarPedido = true;
  }

  CambiarNodoUsuarios(){

    const usuario: Usuario={
      Dpi: Number(this.DpiUsuario.value),
      Nombre: this.NombreUsuario.value,
      Correo: this.CorreoUsuario.value,
      Password: this.PasswordUsuario.value,
      Cuenta: this.CuentaUsuario
    }

    this.usuarioService.ConfigMerkleUsuarios(usuario).subscribe((res:any)=>{
      this.DpiUsuario.setValue("")
      this.NombreUsuario.setValue("")
      this.CorreoUsuario.setValue("")
      this.PasswordUsuario.setValue("")
      console.log("Arbol Merkle Usuario Configurado")
    },(err)=>{
      this.mostrarMensajeError=true
    })


    this.NodoUsuario = false;
    this.ArreglarUsuario = true;  



  }

  CambiarNodoProductos(){

    const producto: Productos={
      Nombre: this.NombreProducto.value,
      Codigo: Number(this.CodigoProducto.value),
      Descripcion: this.DescripcionProducto.value,
      Precio: Number(this.PrecioProducto.value),
      Cantidad: Number(this.CantidadProducto.value),
      Imagen: this.ImagenProducto.value,
      Almacenamiento: this.AlmacenProducto.value
    }

    this.inventarioservice.ConfigMerkleProductos(producto).subscribe((res:any)=>{
      this.NombreProducto.setValue("")
      this.CodigoProducto.setValue("")
      this.DescripcionProducto.setValue("")
      this.PrecioProducto.setValue("")
      this.CantidadProducto.setValue("")
      this.ImagenProducto.setValue("")
      this.AlmacenProducto.setValue("")
      console.log("Arbol Merkle Productos Configurado")
    },(err)=>{
      this.mostrarMensajeError=true
    })


    this.NodoProducto = false;
    this.ArreglarProducto = true;  

  }

  CambiarNodoTiendas(){

    const tienda: Info={
      Id:Number(null),
      Nombre:this.NombreTienda.value,
      Descripcion:this.DescripcionTienda.value,
      Contacto:this.ContactoTienda.value,
      Calificacion:Number(this.CalificacionTienda.value),
      Logo:this.LogoTienda.value
    }

    this.subirtiendaservice.ConfigMerkleTiendas(tienda).subscribe((res:any)=>{
      this.NombreTienda.setValue("")
      this.DescripcionTienda.setValue("")
      this.ContactoTienda.setValue("")
      this.CalificacionTienda.setValue("")
      this.LogoTienda.setValue("")  
      console.log("Arbol Merkle Tiendas Configurado")
    },(err)=>{
      this.mostrarMensajeError=true
    })


    this.NodoTienda = false;
    this.ArreglarTienda = true; 
    
  }

  ArreglarUsuarios(){
    this.usuarioService.getMerkleUsuarios().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Usuarios'
    })


    this.NodoUsuario = true;
    //this.ArreglarUsuario = false;  

  }

  ArreglarTiendas(){
    this.subirtiendaservice.getMerkleTiendas().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Tiendas'
    })


    this.NodoTienda = true;
   // this.ArreglarTienda = false;  

  }

  ArreglarPedidos(){

    this.pedidosservice.getMerklePedidos().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Pedidos'
    })

    this.NodoPedido = true;
    //this.ArreglarPedido = false;  


  }

  ArreglarProductos(){
    this.inventarioservice.getMerkleProductos().subscribe((dataList:any)=>{
      console.log(dataList)
      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError='No se pudo cargar la lista de Productos'
    })


    this.NodoProducto = true;
    //this.ArreglarProducto = false;  
    
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }


}
