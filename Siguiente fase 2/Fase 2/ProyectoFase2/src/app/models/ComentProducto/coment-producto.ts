import { Comentarios } from "../Comentarios/comentarios";

export class ComentProducto {
    Producto: string
    Comentarios: Comentarios[]

    constructor(_Producto: string, _Comentarios: Comentarios[]){
        this.Producto = _Producto
        this.Comentarios = _Comentarios

    }
}
