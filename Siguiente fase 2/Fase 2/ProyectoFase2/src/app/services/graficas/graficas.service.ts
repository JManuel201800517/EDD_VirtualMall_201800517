import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from 'rxjs';
import { Grafo } from "../../models/grafo/grafo";

@Injectable({
  providedIn: 'root'
})
export class GraficasService {

  constructor(private http: HttpClient) { }


  postGrafo(grafica: string):Observable<Grafo>{
    const httpOptions = {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      }),
    };
    return this.http.post<Grafo>(baseURL+"subirGrafo", grafica, httpOptions)
  }
}
