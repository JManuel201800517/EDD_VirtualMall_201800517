import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from 'rxjs';
import { Tienda } from "../../models/tienda/tienda";
import { Datos } from 'src/app/models/datos/datos';

@Injectable({
  providedIn: 'root'
})
export class SubirtiendaService {

  constructor(private http: HttpClient) { 
    
  }

  //postTienda(tienda):Observable<any>{
    //const httpOptions = {
      //headers: new HttpHeaders({
      //  "Content-Type": "application/json"
      //}),
    //};
    //return this.http.post<Datos[]>(baseURL+"cargartienda/subir", tienda, httpOptions)
  //}
}
