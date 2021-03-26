import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from 'rxjs';
import { Tienda } from "../../models/tienda/tienda";
import { Datos } from 'src/app/models/datos/datos';
import { Eliminarespecifico } from 'src/app/models/eliminarespecifico/eliminarespecifico';
import { Buscarespecifico } from 'src/app/models/buscarespecifico/buscarespecifico';

@Injectable({
  providedIn: 'root'
})
export class SubirtiendaService {

  constructor(private http: HttpClient) { 
    
  }

  postTienda(tienda: any):Observable<Tienda>{
    const httpOptions = {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      }),
    };
    return this.http.post<Tienda>(baseURL+"cargartienda/subir", tienda, httpOptions)
  }

  deleteTienda(eliminar: any):Observable<Eliminarespecifico>{
    const httpOptions = {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      }),
    };
    return this.http.post<Eliminarespecifico>(baseURL+"Eliminar", eliminar, httpOptions)
  }

  buscarTienda(buscar: any):Observable<Tienda>{
    const httpOptions = {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      }),
    };
    return this.http.post<Tienda>(baseURL+"TiendaEspecifica", {buscar: buscar}, httpOptions)
  }

  //obtenerTienda():Observable<Tienda[]>{
    //const httpOptions = {
      //headers: new HttpHeaders({
        //'Content-Type': 'application/json',
      //}),
    //};
    //return this.http.get<Tienda[]>(baseURL + 'TiendaEspecifica', httpOptions);
  //}

  getListaTiendas():Observable<Tienda[]>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<Tienda[]>(baseURL + 'cargartienda', httpOptions);
  }

}
