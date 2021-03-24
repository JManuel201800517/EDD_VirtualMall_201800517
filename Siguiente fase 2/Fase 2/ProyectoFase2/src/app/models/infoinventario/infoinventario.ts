import { Productos } from "../Productos/productos";

export class Infoinventario {
    InvenTienda: string
    InvenDepartamento: string
    InvenCalificacion: number
    ProductosInventario: Productos[]

    constructor(_InvenTienda: string, _InvenDepartamento: string,
        _InvenCalificacion: number, _ProductosInventario: Productos[]){
        this.InvenTienda = _InvenTienda
        this.InvenDepartamento = _InvenDepartamento
        this.InvenCalificacion = _InvenCalificacion
        this.ProductosInventario = _ProductosInventario

    }
}
