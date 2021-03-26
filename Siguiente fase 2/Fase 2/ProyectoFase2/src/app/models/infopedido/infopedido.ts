import { Pedirproductos } from "../pedirproductos/pedirproductos";

export class Infopedido {
    Fecha: string
    Tienda: string
    Departamento: string
    Calificacion: number
    Productos: Pedirproductos[]

    constructor(_Fecha: string, _Tienda: string, _Departamento: string,
        _Calificacion: number, _Productos: Pedirproductos[]){
        this.Fecha = _Fecha
        this.Tienda = _Tienda
        this.Departamento = _Departamento
        this.Calificacion = _Calificacion
        this.Productos = _Productos

    }
}
