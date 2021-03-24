export class Productos {
    NombreProducto: string
    CodigoProducto: number
    DescripcionProducto: string
    PrecioProducto: number
    CantidadProducto: number
    ImagenProducto: string

    constructor(_NombreProducto: string, 
        _CodigoProducto: number, _DescripcionProducto: string, 
        _PrecioProducto: number, _CantidadProducto: number, _ImagenProducto: string){
        this.NombreProducto = _NombreProducto
        this.CodigoProducto = _CodigoProducto
        this.DescripcionProducto = _DescripcionProducto
        this.PrecioProducto = _PrecioProducto
        this.CantidadProducto = _CantidadProducto
        this.ImagenProducto = _ImagenProducto
    }
}
