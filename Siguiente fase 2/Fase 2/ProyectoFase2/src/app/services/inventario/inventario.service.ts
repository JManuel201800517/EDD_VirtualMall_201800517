import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from 'rxjs';
import { Inventario } from "../../models/inventario/inventario";

@Injectable({
  providedIn: 'root'
})
export class InventarioService {

  constructor(private http: HttpClient) { }


  postInventario(invent: string):Observable<Inventario>{
    const httpOptions = {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      }),
    };
    return this.http.post<Inventario>(baseURL+"Inventario", invent, httpOptions)
  }

  getListaInventarios():Observable<Inventario[]>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<Inventario[]>(baseURL + 'cargarinventario', httpOptions);
  }
}
