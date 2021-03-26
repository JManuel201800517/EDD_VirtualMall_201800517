import { Component, OnInit } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import listadotiendas from "src/json/Tiendas.json";
import { FormControl } from '@angular/forms';
//import { InventarioService } from "../../services/inventario/inventario.service";
import { SubirtiendaService } from 'src/app/services/tienda/subirtienda.service';


@Component({
  selector: 'app-eliminarespecifica',
  templateUrl: './eliminarespecifica.component.html',
  styleUrls: ['./eliminarespecifica.component.css']
})
export class EliminarespecificaComponent implements OnInit {

  eliminacion = new FormControl('');
  mostrarMensaje = false;
  mostrarMensajeError = false;


  constructor(private eliminarService: SubirtiendaService) { }

  ngOnInit(): void {
  }

  subirEliminacion(){
    console.log("Funciona muy bien")
    //console.log(this.inventario.value)


    this.eliminarService.deleteTienda(this.eliminacion.value).subscribe((res:any)=>{
      this.mostrarMensaje=true
      this.eliminacion.setValue("")
      console.log("Inventario Cargado")
      console.log(res)

    }, (err)=>{
      this.mostrarMensajeError=true
    })
  }

  desactivarMensaje(){
    this.mostrarMensaje=false
    this.mostrarMensajeError=false
  }

}
