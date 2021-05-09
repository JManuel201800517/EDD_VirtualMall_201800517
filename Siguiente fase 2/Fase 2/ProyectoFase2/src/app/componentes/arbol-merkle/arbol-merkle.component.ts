import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

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


  ArPedidos = false;
  ArTiendas = false;
  ArProductos = false;
  ArUsuarios = false;

  NodoPedido = true;
  NodoUsuario = true;
  NodoTienda = true;
  NodoProducto = true;

  ArreglarPedido = false;
  ArreglarUsuario = false;
  ArreglarTienda = false;
  ArreglarProducto = false;


  constructor() { }

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
    this.NodoPedido = false;
    this.ArreglarPedido = true;  

  }

  CambiarNodoUsuarios(){
    this.NodoUsuario = false;
    this.ArreglarUsuario = true;  

  }

  CambiarNodoProductos(){
    this.NodoProducto = false;
    this.ArreglarProducto = true;  

  }

  CambiarNodoTiendas(){
    this.NodoTienda = false;
    this.ArreglarTienda = true;  

  }

  ArreglarUsuarios(){
    this.NodoUsuario = true;
    this.ArreglarUsuario = false;  

  }

  ArreglarTiendas(){
    this.NodoTienda = true;
    this.ArreglarTienda = false;  

  }

  ArreglarPedidos(){
    this.NodoPedido = true;
    this.ArreglarPedido = false;  


  }

  ArreglarProductos(){
    this.NodoProducto = true;
    this.ArreglarProducto = false;  
    
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }


}
