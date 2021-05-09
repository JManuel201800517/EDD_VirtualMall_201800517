import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from 'rxjs';
import { Pedidos } from "../../models/pedidos/pedidos";

@Injectable({
  providedIn: 'root'
})
export class PedidosService {

  constructor(private http: HttpClient) { }

  getMerklePedidos():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<any>(baseURL + 'arbolMerklePedidos', httpOptions);
  }

  postPedido(pedir: string):Observable<Pedidos>{
    const httpOptions = {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      }),
    };
    return this.http.post<Pedidos>(baseURL+"Pedidos", pedir, httpOptions)
  }

  getListaPedidos():Observable<Pedidos[]>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.get<Pedidos[]>(baseURL + 'cargarpedido', httpOptions);
  }
}
