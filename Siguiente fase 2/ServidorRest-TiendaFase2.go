package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"math"

	"github.com/kr/fernet"
	//"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

var cargartienda []Tienda

type Tienda struct {
	DATOSINFO []DATOS `json:"Datos,omitempty"`
}

type DATOS struct {
	INDICE     string          `json:"Indice,omitempty"`
	DEPARTINFO []DEPARTAMENTOS `json:"Departamentos,omitempty"`
}

type DEPARTAMENTOS struct {
	TIPOTIENDA string `json:"Nombre,omitempty"`
	NOTIENDA   []info `json:"Tiendas,omitempty"`
}

type info struct {
	ID           string `json:"Id,omitempty"`
	NOMBRE       string `json:"Nombre,omitempty"`
	DESCRIPCION  string `json:"Descripcion,omitempty"`
	CONTACTO     string `json:"Contacto,omitempty"`
	CALIFICACION int    `json:"Calificacion,omitempty"`
	LOGO         string `json:"Logo,omitempty"`
}

var tiendaespecifica []Especifica

type Especifica struct {
	DepartamentoEspecifico string `json:"Departamento,omitempty"`
	NombreEspecifico       string `json:"Nombre,omitempty"`
	CalificacionEspecifico int    `json:"Calificacion,omitempty"`
}

var eliminar []Eliminacion

type Eliminacion struct {
	NombreEliminar       string `json:"Nombre,omitempty"`
	CategoriaEliminar    string `json:"Categoria,omitempty"`
	CalificacionEliminar int    `json:"Calificacion,omitempty"`
}

var invent []Inventario

type Inventario struct{
	DatosInventario  []InfoInventario  `json:"Inventarios,omitempty"`
}

type InfoInventario struct{
	InvenTienda          string      `json:"Tienda,omitempty"`
	InvenDepartamento    string      `json:"Departamento,omitempty"`
	InvenCalificacion    int         `json:"Calificacion,omitempty"`
	ProductosInventario  []Productos `json:"Productos,omitempty"`
}

type Productos struct{
	NombreProducto      string  `json:"Nombre,omitempty"`
	CodigoProducto      int     `json:"Codigo,omitempty"`
	DescripcionProducto string  `json:"Descripcion,omitempty"`
	PrecioProducto      int `json:"Precio,omitempty"`
	CantidadProducto    int     `json:"Cantidad,omitempty"`
	ImagenProducto      string  `json:"Imagen,omitempty"`
	AlmacenProducto     string  `json:"Almacenamiento,omitempty"`
}


var pedir []Pedidos

type Pedidos struct{
	PEDIDO  []InfoPedido  `json:"Pedidos,omitempty"`
}

type InfoPedido struct{
	FechaPedido         string            `json:"Fecha,omitempty"`
	TiendaPedido        string            `json:"Tienda,omitempty"`
	DepartamentoPedido  string            `json:"Departamento,omitempty"`
	CalificacionPedido  int               `json:"Calificacion,omitempty"`
	ClientePedido       int               `json:"Cliente,omitempty"` 
	ProductosPedido     []PedirProductos  `json:"Productos,omitempty"`
}

type PedirProductos struct{
	ProductoCodigo  int  `json:"Codigo,omitempty"`     
}


var miembro []Cuentas

type Cuentas struct{
	USUARIOS  []Usuarios  `json:"Usuarios,omitempty"`
}

type Usuarios struct{
	Dpi        int       `json:"Dpi,omitempty"`
	Nombre     string    `json:"Nombre,omitempty"`
	Correo     string    `json:"Correo,omitempty"`
	Password   string    `json:"Password,omitempty"`
	Cuenta     string    `json:"Cuenta,omitempty"`

}

var suprimir []EliminarUsuario

type EliminarUsuario struct {
	NombreUsuario       string `json:"Nombre,omitempty"`
	PasswordUsuario     string `json:"Password,omitempty"`
}


var ingresar []Sesion

type Sesion struct {
	NombreUsuario       string `json:"Nombre,omitempty"`
	PasswordUsuario     string `json:"Password,omitempty"`
}


var caminografo []Grafo

type Grafo struct {
	Nodos          []Nodo  `json:"Nodos,omitempty"`
	PosicionRobot  string  `json:"PosicionInicialRobot,omitempty"`
	Entrega        string  `json:"Entrega,omitempty"`
}

type Nodo struct {
	Nombre    string     `json:"Nombre,omitempty"`
	Enlace    []Enlaces  `json:"Enlaces,omitempty"`
}

type Enlaces struct {
	Nombre     string   `json:"Nombre,omitempty"`
	Distancia  int      `json:"Distancia,omitempty"`
}


var Tcomentario []ComentarioTienda

type ComentarioTienda struct{
	IDCliente int `json:"Dpi,omitempty"`
	Tienda   string  `json:"Tienda,omitempty"`
	ComentTienda  []Comentario  `json:"Comentarios,omitempty"`
}

type CTienda struct{
	IDCliente int `json:"Dpi,omitempty"`
	Tienda   string  `json:"Tienda,omitempty"`
}

var Icomentario []ComentarioProducto

type ComentarioProducto struct{
	IDCliente int `json:"Dpi,omitempty"`
	Producto  string  `json:"Producto,omitempty"`
	ComentProducto  []Comentario  `json:"Comentarios,omitempty"`

}

type CProducto struct{
	IDCliente int `json:"Dpi,omitempty"`
	Producto  string  `json:"Producto,omitempty"`
}

type Comentario struct{
	Coment string `json:"Comentario,omitempty"`
	SubComent  []SubComentario `json:"SubComentarios,omitempty"`
}

type SubComentario struct{
	Sub string `json:"SubComentario,omitempty"`
	SubComent  []SubComentario `json:"SubComentarios,omitempty"`
}




var encript []Encriptado

type Encriptado struct{
	Dpi        string       `json:"Dpi,omitempty"`
	Nombre     string    `json:"Nombre,omitempty"`
	Correo     string    `json:"Correo,omitempty"`
	Password   string    `json:"Password,omitempty"`
	Cuenta     string    `json:"Cuenta,omitempty"`

}




var IdPedidos [1000]string

var IdUsuarios [1000]string

var IdTiendas [1000]string

var IdProductos [1000]string


var TM []TiendaMerkle

type TiendaMerkle struct {
	ID           string `json:"Id,omitempty"`
	NOMBRE       string `json:"Nombre,omitempty"`
	DESCRIPCION  string `json:"Descripcion,omitempty"`
	CONTACTO     string `json:"Contacto,omitempty"`
	CALIFICACION int    `json:"Calificacion,omitempty"`
	LOGO         string `json:"Logo,omitempty"`
}

var PM []ProductosMerkle

type ProductosMerkle struct{
	NombreProducto      string  `json:"Nombre,omitempty"`
	CodigoProducto      int     `json:"Codigo,omitempty"`
	DescripcionProducto string  `json:"Descripcion,omitempty"`
	PrecioProducto      int `json:"Precio,omitempty"`
	CantidadProducto    int     `json:"Cantidad,omitempty"`
	ImagenProducto      string  `json:"Imagen,omitempty"`
	AlmacenProducto     string  `json:"Almacenamiento,omitempty"`
}

var UM []UsuariosMerkle

type UsuariosMerkle struct{
	Dpi        int       `json:"Dpi,omitempty"`
	Nombre     string    `json:"Nombre,omitempty"`
	Correo     string    `json:"Correo,omitempty"`
	Password   string    `json:"Password,omitempty"`
	Cuenta     string    `json:"Cuenta,omitempty"`

}

var PeM []PedidosMerkle

type PedidosMerkle struct{
	FechaPedido         string            `json:"Fecha,omitempty"`
	TiendaPedido        string            `json:"Tienda,omitempty"`
	DepartamentoPedido  string            `json:"Departamento,omitempty"`
	CalificacionPedido  int               `json:"Calificacion,omitempty"`
	ClientePedido       int               `json:"Cliente,omitempty"` 
	ProductosPedido     []PedirProductos  `json:"Productos,omitempty"`
}



// no usar esta seccion encacillada{

var lista []infolista

type infolista struct {
	ID           string `json:"Id,omitempty"`
	NOMBRE       string `json:"Nombre,omitempty"`
	DESCRIPCION  string `json:"Descripcion,omitempty"`
	CONTACTO     string `json:"Contacto,omitempty"`
	CALIFICACION int    `json:"Calificacion,omitempty"`
	Next         *infolista
}

func Ordenamiento(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(cargartienda)

	var linea Tienda

	//var multi Linealizacion

	//z := 0

	p1 := len(linea.DATOSINFO)
	p2 := len(linea.DATOSINFO[0].DEPARTINFO)

	// k = Numero de Calificaciones

	k := 5

	multi := p1 * p2 * k

	///lista := list.New()

	vector := make([]infolista, multi)

	z := 0
	for z < len(linea.DATOSINFO) {

		j := 0

		for j < len(linea.DATOSINFO[z].DEPARTINFO) {

			i := 0

			for i < len(linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {
				p3 := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CALIFICACION

				//nom := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE

				n := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE
				d := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].DESCRIPCION
				c := linea.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CONTACTO

				vectorlinea := (((z * p2) + j) * k) + p3

				//vec = append(vec, Linealizacion{NOMBRE: n, DESCRIPCION: d, CONTACTO: c, CALIFICACION: p3})
				//x := lista.PushBack(nom)
				vector[vectorlinea].Next.NOMBRE = n
				vector[vectorlinea].Next.DESCRIPCION = d
				vector[vectorlinea].Next.CONTACTO = c
				vector[vectorlinea].Next.CALIFICACION = p3

				//vector[vectorlinea] = vec

				//json.NewEncoder(w).Encode(vec)

				//vector[vectorlinea] = item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA[i].NOMBRE[vectorlinea]
				i = i + 1

			}

			j = j + 1

		}

		z = z + 1

	}

	json.NewEncoder(w).Encode(lista)

}


// } no usar esta seccion encillada


func subirComentarioTiendaEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var publicacion ComentarioTienda
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &publicacion)

	for z := 0; z < len(publicacion.ComentTienda); z++ {
		for j := 0; j < len(publicacion.ComentTienda[z].SubComent); j++ {
			_ = json.NewDecoder(req.Body).Decode(&publicacion)

		}
	}
	Tcomentario = append(Tcomentario, publicacion)
	json.NewEncoder(w).Encode(publicacion)

}

func ComentarioTiendaEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tcomentario)

}

func subirComentarioProductoEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var publicacion ComentarioProducto
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &publicacion)

	for z := 0; z < len(publicacion.ComentProducto); z++ {
		for j := 0; j < len(publicacion.ComentProducto[z].SubComent); j++ {
			for r := 0; r < len(publicacion.ComentProducto[z].SubComent[j].SubComent); r++{
				_ = json.NewDecoder(req.Body).Decode(&publicacion)

			}

		}
	}
	Icomentario = append(Icomentario, publicacion)
	json.NewEncoder(w).Encode(publicacion)

}

func ComentarioProductoEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Icomentario)

}




func ListadoDeTiendasEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cargartienda)

}

func ListadoDeUsuariosEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(miembro)

}

func ObtenerTipoCuentaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range miembro {
		for z := 0; z < len(item.USUARIOS); z++ {
			if item.USUARIOS[z].Cuenta == params["Cuenta"] {
				json.NewEncoder(w).Encode(item.USUARIOS[z])
				return
			}

		}
	}
	json.NewEncoder(w).Encode(&Cuentas{})
}

func CrearUsuarioEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var avatar Cuentas
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &avatar)

	for z := 0; z < len(avatar.USUARIOS); z++ {
		_ = json.NewDecoder(req.Body).Decode(&avatar)
		
	}
	miembro = append(miembro, avatar)
	json.NewEncoder(w).Encode(avatar)

}

func EliminarUsuarioEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var elimus EliminarUsuario
	json.NewDecoder(req.Body).Decode(&elimus)

	//tienda.ID = params["Id nuevo"]
	suprimir = append(suprimir, elimus)
	//json.NewEncoder(w).Encode(tiendaespecifica)

	for _, item := range miembro {
		z := 0

		for  z < len(item.USUARIOS) {

			if item.USUARIOS[z].Nombre == elimus.NombreUsuario{

				if item.USUARIOS[z].Password == elimus.PasswordUsuario{

					item.USUARIOS = append(item.USUARIOS[:z], item.USUARIOS[z+1:]...)

					break
				}


				break
			}

			z = z + 1


		}

	}

	json.NewEncoder(w).Encode(miembro)

}


func SubirGrafoEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var robot Grafo
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &robot)

	for z := 0; z < len(robot.Nodos); z++ {
		for j := 0; j < len(robot.Nodos[z].Enlace); j++ {
				_ = json.NewDecoder(req.Body).Decode(&robot)
				//tienda.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].ID = strconv.Itoa(len(cargartienda) + 1)

		}
	}
	caminografo = append(caminografo, robot)
	json.NewEncoder(w).Encode(robot)

}


func DatosGrafoEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(caminografo)

}





func ListadoDeInventarioEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invent)

}

func ListadoDePedidoEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedir)

}

func ObtenerIndiceTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {
			if item.DATOSINFO[z].INDICE == params["Indice"] {
				json.NewEncoder(w).Encode(item.DATOSINFO[z])
				return
			}

		}
	}
	json.NewEncoder(w).Encode(&Tienda{})
}

func ObtenerNombreTipoTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {
				if item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA == params["Nombre"] {
					json.NewEncoder(w).Encode(item.DATOSINFO[z].DEPARTINFO[j])
					return
				}
				j = j + 1
			}

		}
	}
	json.NewEncoder(w).Encode(&Tienda{})
}

func ObtenerNombreTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				i := 0

				for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {
					if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE == params["Nombre"] {
						json.NewEncoder(w).Encode(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i])

						return
					}

					i = i + 1

				}

				j = j + 1

			}

		}
	}
	json.NewEncoder(w).Encode(&Tienda{})
}

func ObtenerContactoTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				i := 0

				for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {
					if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CONTACTO == params["Contacto"] {
						json.NewEncoder(w).Encode(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i])

						return
					}

					i = i + 1

				}

				j = j + 1

			}

		}
	}
	json.NewEncoder(w).Encode(&Tienda{})
}

func BorrarTiendaEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range cargartienda {

		z := 0
		for z < len(item.DATOSINFO) {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				i := 0

				for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {
					if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE == params["Nombre"] {
						item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA = append(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[:i], item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i+1:]...)

						//_ = json.NewDecoder(req.Body).Decode(&index)
						break
					}

					i = i + 1

				}

				j = j + 1

			}

			z = z + 1

		}

	}
	json.NewEncoder(w).Encode(cargartienda)

}

func CrearEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tienda Tienda
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &tienda)

	for z := 0; z < len(tienda.DATOSINFO); z++ {
		for j := 0; j < len(tienda.DATOSINFO[z].DEPARTINFO); j++ {
			for i := 0; i < len(tienda.DATOSINFO[z].DEPARTINFO[j].NOTIENDA); i++ {
				_ = json.NewDecoder(req.Body).Decode(&tienda)
				//tienda.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].ID = strconv.Itoa(len(cargartienda) + 1)
			}

		}
	}
	cargartienda = append(cargartienda, tienda)
	json.NewEncoder(w).Encode(tienda)

}

func TiendaEspecificaEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var especifico Especifica
	json.NewDecoder(req.Body).Decode(&especifico)

	//tienda.ID = params["Id nuevo"]
	tiendaespecifica = append(tiendaespecifica, especifico)
	//json.NewEncoder(w).Encode(tiendaespecifica)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				if item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA == especifico.DepartamentoEspecifico {

					i := 0

					for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {

						if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE == especifico.NombreEspecifico {

							if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CALIFICACION == especifico.CalificacionEspecifico {

								json.NewEncoder(w).Encode(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i])
								//z = len(item.DATOSINFO)
								//j = len(item.DATOSINFO[z].DEPARTINFO)
								//i = len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA)

								return
							}

							return
						}

						i = i + 1

					}

					return
				}

				j = j + 1

			}

		}

	}

	json.NewEncoder(w).Encode(&Tienda{})

}

func EliminarRegistroEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var elim Eliminacion
	json.NewDecoder(req.Body).Decode(&elim)

	//tienda.ID = params["Id nuevo"]
	eliminar = append(eliminar, elim)
	//json.NewEncoder(w).Encode(tiendaespecifica)

	for _, item := range cargartienda {
		for z := 0; z < len(item.DATOSINFO); z++ {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				if item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA == elim.CategoriaEliminar {

					i := 0

					for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {

						if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE == elim.NombreEliminar {

							if item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CALIFICACION == elim.CalificacionEliminar {

								item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA = append(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[:i], item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i+1:]...)
								//z = len(item.DATOSINFO)
								//j = len(item.DATOSINFO[z].DEPARTINFO)
								//i = len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA)

								break
							}

							break
						}

						i = i + 1

					}

					break
				}

				j = j + 1

			}

		}

	}

	json.NewEncoder(w).Encode(cargartienda)

}

func InventarioEndPoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var produc Inventario
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &produc)

	for z := 0; z < len(produc.DatosInventario); z++ {
		for j := 0; j < len(produc.DatosInventario[z].ProductosInventario); j++ {
			_ = json.NewDecoder(req.Body).Decode(&produc)

		}
	}
	invent = append(invent, produc)
	json.NewEncoder(w).Encode(produc)


}

func PedidosEndPoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var peticion Pedidos
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar dato valido")
	}

	json.Unmarshal(reqBody, &peticion)

	for z := 0; z < len(peticion.PEDIDO); z++ {
		for j := 0; j < len(peticion.PEDIDO[z].ProductosPedido); j++ {
			_ = json.NewDecoder(req.Body).Decode(&peticion)

		}
	}
	pedir = append(pedir, peticion)
	json.NewEncoder(w).Encode(peticion)

	//for _, item := range invent {

	//	for _, otro := range pedir{

	//		for z := 0; z < len(item.DatosInventario); z++ {

	//			j := 0
	
	//			for j < len(item.DatosInventario[z].TiendaPedido) {
	
	//				q := 0


	
					//if item.DATOSINFO[z].DEPARTINFO[j].TIPOTIENDA == peticion.Pedidos {
	
					//	i := 0
	
						
	
					//	return
					//}
	
	//				j = j + 1
	
	//			}
	
	//		}


	//	}


	//}

	//json.NewEncoder(w).Encode(&Inventario{})
	
}


func InicioSesionEndPoint(w http.ResponseWriter, req *http.Request){

	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var usuario Sesion
	json.NewDecoder(req.Body).Decode(&usuario)

	//tienda.ID = params["Id nuevo"]
	ingresar = append(ingresar, usuario)
	//json.NewEncoder(w).Encode(tiendaespecifica)

	for _, item := range miembro {
		for z := 0; z < len(item.USUARIOS); z++ {

			if item.USUARIOS[z].Nombre == usuario.NombreUsuario {

				if item.USUARIOS[z].Password == usuario.PasswordUsuario {

					json.NewEncoder(w).Encode(item.USUARIOS[z])
					//z = len(item.DATOSINFO)
					//j = len(item.DATOSINFO[z].DEPARTINFO)
					//i = len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA)

					return
				}

				return
			}

		}

	}

	json.NewEncoder(w).Encode(&Cuentas{})
	
	
}



// no usar esta seccion encacillada{


func DescargarArchivoEndPoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename=prueba1.0")
	w.Header().Set("Content-Type", "application/json")

}

func Grafica(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cargartienda)

	g := newGraph()
	// https://graphviz.gitlab.io/_pages/Gallery/directed/fsm.html

	//w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	for _, item := range cargartienda {

		z := 0
		for z < len(item.DATOSINFO) {

			j := 0

			for j < len(item.DATOSINFO[z].DEPARTINFO) {

				i := 0

				for i < len(item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA) {

					p3 := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CALIFICACION

					n := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].NOMBRE
					//d := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].DESCRIPCION
					c := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i].CONTACTO

					n1 := item.DATOSINFO[z].DEPARTINFO[j].NOTIENDA[i+1].NOMBRE

					s := strconv.Itoa(p3)

					g.addEdge(n, n1, "sig")
					g.addEdge(n, c, "contacto")
					g.addEdge(n, s, "calificacion")
					g.addEdge(n1, n, "ante")

					i = i + 1
				}

				j = j + 1

			}

			z = z + 1

		}

	}
	fmt.Println(g)

	path2, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path2, "-Tpng", "grafo.dot").Output()
	mode := int(0777)
	ioutil.WriteFile("grafo.png", cmd, os.FileMode(mode))

}


// }no usar esta seccion encacillada

func ArbolSinCifrar() string{

	//var cuenta Cuenta


	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	for _, nodo := range miembro{

		z := 0
		for z < len(nodo.USUARIOS){
			grafo+=`Nodo`+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+`[shape=record label=<`
	        grafo+=`<table cellspacing="0" border="0" cellborder="1">`
	        grafo+="<tr><td colspan=\"2\">Dpi: "+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "</td><td colspan=\"2\">Nombre: "+nodo.USUARIOS[z].Nombre+ " </td></tr>"
	        grafo+="<tr><td>Correo: "+nodo.USUARIOS[z].Correo+"</td><td>Password: "+nodo.USUARIOS[z].Password+"</td></tr>"
	        grafo+="<tr><td colspan=\"2\">Cuenta: "+nodo.USUARIOS[z].Cuenta+"</td></tr></table>"
	        grafo+=`
	        >];
	        `	

			if z <= 3{

				if z == 3{
					grafo += "Xi->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

				}else{
					grafo += "Xi->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "Xi->X"+strconv.FormatInt(int64(z), 10)+";"

				}
				

			}else if (z>= 4 && z <= 23){

				if z == 23{
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

				}else if (z >= 4 && z <= 7){
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[0].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[0].Dpi), 10)+"->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 8 && z <= 11){
					grafo += "X0->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X0->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 12 && z <= 15){
					grafo += "X1->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X1->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 16 && z <= 19){
					grafo += "X2->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X2->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 20 && z <= 22){
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->X"+strconv.FormatInt(int64(z), 10)+";"

				}else{
					fmt.Println("Funciona Bien")
				}

				//grafo += "Fin->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

			}else{
				grafo += "Inicio->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "->Xi;"

			}

			
			z = z + 1
		}


	}

	//contador=1
	//grafo=recorrerArbol("Nodo0", nodo, grafo)
	//grafo += "{rank=same Nodo0 X Nodo1}"
	//grafo += "Nodo1->Nodo0;\n"
	grafo+="}"
	return grafo

}


func getArbolSinCifrar(w http.ResponseWriter, req *http.Request){

	data := []byte(ArbolSinCifrar())
    err := ioutil.WriteFile("arbolSinCifrar.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearArbolSinCifrar.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./arbolSinCifrar.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}



func ArbolCifrarSuave() string{

	//var cuenta Cuenta


	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	for _, nodo := range miembro{

		z := 0
		for z < len(nodo.USUARIOS){

			encript := [] byte (strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10))
			hash := sha256.Sum256(encript)
			rad := hex.EncodeToString(hash[:])

			k := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok, err := fernet.EncryptAndSign([]byte(nodo.USUARIOS[z].Nombre), k[0])
			zet := hex.EncodeToString(tok[:])

			k1 := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok1, err1 := fernet.EncryptAndSign([]byte(nodo.USUARIOS[z].Correo), k1[0])
			zet1 := hex.EncodeToString(tok1[:])

			if err != nil {
		        panic(err)
	        }

			if err1 != nil {
		        panic(err1)
	        }



			grafo+=`Nodo`+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+`[shape=record label=<`
	        grafo+=`<table cellspacing="0" border="0" cellborder="1">`
	        grafo+="<tr><td colspan=\"2\">Dpi: "+rad+ "</td><td colspan=\"2\">Nombre: "+nodo.USUARIOS[z].Nombre+ " </td></tr>"
	        grafo+="<tr><td>Correo: "+zet+"</td><td>Password: "+zet1+"</td></tr>"
	        grafo+="<tr><td colspan=\"2\">Cuenta: "+nodo.USUARIOS[z].Cuenta+"</td></tr></table>"
	        grafo+=`
	        >];
	        `	

			if z <= 3{

				if z == 3{
					grafo += "Xi->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

				}else{
					grafo += "Xi->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "Xi->X"+strconv.FormatInt(int64(z), 10)+";"

				}
				

			}else if (z>= 4 && z <= 23){

				if z == 23{
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

				}else if (z >= 4 && z <= 7){
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[0].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[0].Dpi), 10)+"->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 8 && z <= 11){
					grafo += "X0->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X0->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 12 && z <= 15){
					grafo += "X1->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X1->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 16 && z <= 19){
					grafo += "X2->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X2->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 20 && z <= 22){
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->X"+strconv.FormatInt(int64(z), 10)+";"

				}else{
					fmt.Println("Funciona Bien")
				}

				//grafo += "Fin->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

			}else{
				grafo += "Inicio->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "->Xi;"

			}

			
			z = z + 1
		}


	}

	//contador=1
	//grafo=recorrerArbol("Nodo0", nodo, grafo)
	//grafo += "{rank=same Nodo0 X Nodo1}"
	//grafo += "Nodo1->Nodo0;\n"
	grafo+="}"
	return grafo

}


func getArbolCifrarSuave(w http.ResponseWriter, req *http.Request){

	data := []byte(ArbolCifrarSuave())
    err := ioutil.WriteFile("arbolCifrarSuave.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	loc := "crearArbolCifrarSuave.bat"
	_, err2 := exec.Command(loc).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./arbolCifrarSuave.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}



func ArbolCifrar() string{

	//var cuenta Cuenta


	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	for _, nodo := range miembro{

		z := 0
		for z < len(nodo.USUARIOS){

			encript := [] byte (strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10))
			hash := sha256.Sum256(encript)
			rad := hex.EncodeToString(hash[:])

			k := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok, err := fernet.EncryptAndSign([]byte(nodo.USUARIOS[z].Nombre), k[0])
			zet := hex.EncodeToString(tok[:])

			k1 := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok1, err1 := fernet.EncryptAndSign([]byte(nodo.USUARIOS[z].Correo), k1[0])
			zet1 := hex.EncodeToString(tok1[:])

			k2 := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok2, err2 := fernet.EncryptAndSign([]byte(nodo.USUARIOS[z].Password), k2[0])
			zet2 := hex.EncodeToString(tok2[:])

			k3 := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok3, err3 := fernet.EncryptAndSign([]byte(nodo.USUARIOS[z].Cuenta), k3[0])
			zet3 := hex.EncodeToString(tok3[:])


	        if err != nil {
		        panic(err)
	        }

			if err1 != nil {
		        panic(err1)
	        }

			if err2 != nil {
		        panic(err2)
	        }

			if err3 != nil {
		        panic(err3)
	        }


			grafo+=`Nodo`+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+`[shape=record label=<`
	        grafo+=`<table cellspacing="0" border="0" cellborder="1">`
	        grafo+="<tr><td colspan=\"2\">Dpi: "+rad+ "</td><td colspan=\"2\">Nombre: "+zet+ " </td></tr>"
	        grafo+="<tr><td>Correo: "+zet1+"</td><td>Password: "+zet2+"</td></tr>"
	        grafo+="<tr><td colspan=\"2\">Cuenta: "+zet3+"</td></tr></table>"
	        grafo+=`
	        >];
	        `	

			if z <= 3{

				if z == 3{
					grafo += "Xi->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

				}else{
					grafo += "Xi->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "Xi->X"+strconv.FormatInt(int64(z), 10)+";"

				}
				

			}else if (z>= 4 && z <= 23){

				if z == 23{
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

				}else if (z >= 4 && z <= 7){
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[0].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[0].Dpi), 10)+"->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 8 && z <= 11){
					grafo += "X0->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X0->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 12 && z <= 15){
					grafo += "X1->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X1->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 16 && z <= 19){
					grafo += "X2->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				    grafo += "X2->X"+strconv.FormatInt(int64(z), 10)+";"

				}else if (z >= 20 && z <= 22){
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
					grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[3].Dpi), 10)+"->X"+strconv.FormatInt(int64(z), 10)+";"

				}else{
					fmt.Println("Funciona Bien")
				}

				//grafo += "Fin->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"

			}else{
				grafo += "Inicio->Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "[color=grey arrowhead=none];"
				grafo += "Nodo"+strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)+ "->Xi;"

			}

			
			z = z + 1
		}


	}

	//contador=1
	//grafo=recorrerArbol("Nodo0", nodo, grafo)
	//grafo += "{rank=same Nodo0 X Nodo1}"
	//grafo += "Nodo1->Nodo0;\n"
	grafo+="}"
	return grafo

}


func getArbolCifrar(w http.ResponseWriter, req *http.Request){

	data := []byte(ArbolCifrar())
    err := ioutil.WriteFile("arbolCifrado.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	tec := "crearArbolCifrar.bat"
	_, err2 := exec.Command(tec).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./arbolCifrado.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}


func GrafoRobot() string{
	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	for _, nodo := range caminografo{

		z := 0
		for z < len(nodo.Nodos){

			if nodo.Nodos[z].Nombre == nodo.PosicionRobot{

				grafo += nodo.Nodos[z].Nombre+"[shape=circle, color=green]"

			}else if nodo.Nodos[z].Nombre == nodo.Entrega{

				grafo += nodo.Nodos[z].Nombre+"[shape=circle, color=yellow]"

			}else{

				grafo += nodo.Nodos[z].Nombre+"[shape=circle]"

			}

			z = z + 1
 
		}

		grafo+=""

		y := 0

		for y < len(nodo.Nodos){

			e := 0

			for e < len(nodo.Nodos[y].Enlace){

				grafo += nodo.Nodos[y].Nombre+"->"+nodo.Nodos[y].Enlace[e].Nombre+"[label="+strconv.FormatInt(int64(nodo.Nodos[y].Enlace[e].Distancia), 10)+"];"

				e = e + 1
			}

			y = y + 1

		}

		
	}

	grafo+=""
	grafo+="}"
	return grafo

}

func getGrafoRobot(w http.ResponseWriter, req *http.Request){

	data := []byte(GrafoRobot())
    err := ioutil.WriteFile("grafoRobot.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearGrafoRobot.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./grafoRobot.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}

func getEncriptado(w http.ResponseWriter, req *http.Request){

	data := []byte(EncriptarInfo())
    err := ioutil.WriteFile("EncriptadoUsuarios.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
}


func EncriptarInfo() string{
	//w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	//var x string
	var encp string
	encp="{"
	encp+=""

	for _, item := range miembro {
		for z := 0; z < len(item.USUARIOS); z++ {

			encript := [] byte (strconv.FormatInt(int64(item.USUARIOS[z].Dpi), 10))
			hash := sha256.Sum256(encript)
			rad := hex.EncodeToString(hash[:])

			k := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok, err := fernet.EncryptAndSign([]byte(item.USUARIOS[z].Nombre), k[0])
			zet := hex.EncodeToString(tok[:])

			k1 := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok1, err1 := fernet.EncryptAndSign([]byte(item.USUARIOS[z].Correo), k1[0])
			zet1 := hex.EncodeToString(tok1[:])

			k2 := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok2, err2 := fernet.EncryptAndSign([]byte(item.USUARIOS[z].Password), k2[0])
			zet2 := hex.EncodeToString(tok2[:])

			k3 := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	        tok3, err3 := fernet.EncryptAndSign([]byte(item.USUARIOS[z].Cuenta), k3[0])
			zet3 := hex.EncodeToString(tok3[:])


	        if err != nil {
		        panic(err)
	        }

			if err1 != nil {
		        panic(err1)
	        }

			if err2 != nil {
		        panic(err2)
	        }

			if err3 != nil {
		        panic(err3)
	        }
			encp+="{"

			encp+="Dpi: "+rad+","
			encp+="Nombre: "+zet+","
			encp+="Correo: "+zet1+","
			encp+="Password: "+zet2+","
			encp+="Cuenta: "+zet3+","
			encp+="}"

			


		}
		
	}
	encp+="}"
	return encp

}

//var vec []Vector

//type Vector struct {
//	id string
// }


func MerkleUsuarioCrear() string{

	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range miembro{

		z := 0
		for z < len(nodo.USUARIOS){

			x := strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)
			y := nodo.USUARIOS[z].Nombre
			r := nodo.USUARIOS[z].Correo
			w := nodo.USUARIOS[z].Password
			v := nodo.USUARIOS[z].Cuenta

			encript := [] byte (x+","+y+","+r+","+w+","+v)
			hash := sha256.Sum256(encript)
			id[dot] = hex.EncodeToString(hash[:])

			IdUsuarios[dot] = hex.EncodeToString(hash[:])

			fmt.Println(id[dot])

			je := strconv.FormatInt(int64(dot), 10)

			grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`"];`	

			z = z + 1
			dot = dot + 1
		}
		fmt.Println("Termino Bien")

	}

	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdUsuarios[dot] = hex.EncodeToString(hash[:])

		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdUsuarios[jas] = hex.EncodeToString(hash[:])


			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdUsuarios[jas] = hex.EncodeToString(hash[:])

						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}

func MerkleUsuarioConfig(w http.ResponseWriter, req *http.Request) string{

	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var datos UsuariosMerkle
	json.NewDecoder(req.Body).Decode(&datos)

	//tienda.ID = params["Id nuevo"]
	UM = append(UM, datos)

	fer := strconv.FormatInt(int64(datos.Dpi), 10)
	fi := datos.Nombre
	fa := datos.Correo
	fo := datos.Password
	fu := datos.Cuenta
	

	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range miembro{

		z := 0
		for z < len(nodo.USUARIOS){

			x := strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)
			y := nodo.USUARIOS[z].Nombre
			r := nodo.USUARIOS[z].Correo
			w := nodo.USUARIOS[z].Password
			v := nodo.USUARIOS[z].Cuenta


			encript := [] byte (x+","+y+","+r+","+w+","+v)
			hash := sha256.Sum256(encript)
			id[dot] = hex.EncodeToString(hash[:])

			IdUsuarios[dot] = hex.EncodeToString(hash[:])

			fmt.Println(id[dot])

			je := strconv.FormatInt(int64(dot), 10)

			if x == fer{

				nodo.USUARIOS[z].Nombre = fi
				nodo.USUARIOS[z].Correo = fa
				nodo.USUARIOS[z].Password = fo
				nodo.USUARIOS[z].Cuenta = fu


				grafo+=`Nodo`+ je +`[shape=record,fillcolor=red,style=filled,label="`+id[dot]+`&#92;n&#92;n`+fer+`&#92;n`+fi+`&#92;n`+fa+`&#92;n`+fo+`&#92;n`+fu+`"];`	

			}else{

				grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`"];`	

			}


			z = z + 1
			dot = dot + 1
		}
		fmt.Println("Termino Bien")

	}

	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdUsuarios[dot] = hex.EncodeToString(hash[:])

		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdUsuarios[jas] = hex.EncodeToString(hash[:])


			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdUsuarios[jas] = hex.EncodeToString(hash[:])

						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}

func MerkleUsuarioArreg() string{

	var datos UsuariosMerkle

	fer := strconv.FormatInt(int64(datos.Dpi), 10)
	fi := datos.Nombre
	fa := datos.Correo
	fo := datos.Password
	fu := datos.Cuenta
	

	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range miembro{

		z := 0
		for z < len(nodo.USUARIOS){

			x := strconv.FormatInt(int64(nodo.USUARIOS[z].Dpi), 10)
			y := nodo.USUARIOS[z].Nombre
			r := nodo.USUARIOS[z].Correo
			w := nodo.USUARIOS[z].Password
			v := nodo.USUARIOS[z].Cuenta


			if x == fer{

				encript := [] byte (fer+","+fi+","+fa+","+fo+","+fu)
			    hash := sha256.Sum256(encript)
			    id[dot] = hex.EncodeToString(hash[:])

			    IdUsuarios[dot] = hex.EncodeToString(hash[:])

			    fmt.Println(id[dot])
 
			    je := strconv.FormatInt(int64(dot), 10)

				grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+fer+`&#92;n`+fi+`&#92;n`+fa+`&#92;n`+fo+`&#92;n`+fu+`"];`	

			}else{

				encript := [] byte (x+","+y+","+r+","+w+","+v)
			    hash := sha256.Sum256(encript)
			    id[dot] = hex.EncodeToString(hash[:])

			    IdUsuarios[dot] = hex.EncodeToString(hash[:])

			    fmt.Println(id[dot])

			    je := strconv.FormatInt(int64(dot), 10)

				grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`"];`	

			}
			

			z = z + 1
			dot = dot + 1
		}
		fmt.Println("Termino Bien")

	}

	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdUsuarios[dot] = hex.EncodeToString(hash[:])

		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdUsuarios[jas] = hex.EncodeToString(hash[:])


			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdUsuarios[jas] = hex.EncodeToString(hash[:])

						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}

func arbolMerkleUsuarioEndPoint(w http.ResponseWriter, req *http.Request){

	data := []byte(MerkleUsuarioCrear())
    err := ioutil.WriteFile("MerkleUsuario.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearMerkleUsuario.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./MerkleUsuario.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}

func ConfigMerkleUsuarioEndPoint(w http.ResponseWriter, req *http.Request){

	data := []byte(MerkleUsuarioConfig(w, req))
    err := ioutil.WriteFile("MerkleUsuario.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearMerkleUsuario.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./MerkleUsuario.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}




func MerkleTiendasCrear() string{

	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range cargartienda{

		for ty := 0; ty < len(nodo.DATOSINFO); ty++{

			yum := 0

			for yum < len(nodo.DATOSINFO[ty].DEPARTINFO){

				z := 0
				for z < len(nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA){
		
					x := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].ID
					y := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].NOMBRE
					r := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].DESCRIPCION
					w := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].CONTACTO
					v := strconv.FormatInt(int64(nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].CALIFICACION), 10)
					q := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].LOGO
		
					encript := [] byte (x+","+y+","+r+","+w+","+v+","+q)
					hash := sha256.Sum256(encript)
					id[dot] = hex.EncodeToString(hash[:])

					IdTiendas[dot] = hex.EncodeToString(hash[:])
		
					fmt.Println(id[dot])
		
					je := strconv.FormatInt(int64(dot), 10)
		
					grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`&#92;n`+q+`"];`	
		
					z = z + 1
					dot = dot + 1
				}

				yum = yum + 1
			}

		}
		fmt.Println("Termino Bien")

	}
	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdTiendas[dot] = hex.EncodeToString(hash[:])


		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		//z = z + 1
		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdTiendas[jas] = hex.EncodeToString(hash[:])


			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdTiendas[jas] = hex.EncodeToString(hash[:])


						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}

func MerkleTiendasConfig(w http.ResponseWriter, req *http.Request) string{

	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var datos TiendaMerkle
	json.NewDecoder(req.Body).Decode(&datos)

	//tienda.ID = params["Id nuevo"]
	TM = append(TM, datos)

	fer := datos.ID
	fi := datos.NOMBRE
	fa := datos.DESCRIPCION
	fo := datos.CONTACTO
	fe := datos.CALIFICACION
	fu := datos.LOGO

	fei := strconv.FormatInt(int64(datos.CALIFICACION), 10)

	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range cargartienda{

		for ty := 0; ty < len(nodo.DATOSINFO); ty++{

			yum := 0

			for yum < len(nodo.DATOSINFO[ty].DEPARTINFO){

				z := 0
				for z < len(nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA){
		
					x := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].ID
					y := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].NOMBRE
					r := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].DESCRIPCION
					w := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].CONTACTO
					v := strconv.FormatInt(int64(nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].CALIFICACION), 10)
					q := nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].LOGO
		
					encript := [] byte (x+","+y+","+r+","+w+","+v+","+q)
					hash := sha256.Sum256(encript)
					id[dot] = hex.EncodeToString(hash[:])

					IdTiendas[dot] = hex.EncodeToString(hash[:])
		
					fmt.Println(id[dot])
		
					je := strconv.FormatInt(int64(dot), 10)

					if w == fo{

						//nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].NOMBRE = fi
						nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].DESCRIPCION = fa
						nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].NOMBRE = fi
						nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].LOGO = fu
						nodo.DATOSINFO[ty].DEPARTINFO[yum].NOTIENDA[z].CALIFICACION = fe
		
		
						grafo+=`Nodo`+ je +`[shape=record,fillcolor=red,style=filled,label="`+id[dot]+`&#92;n&#92;n`+fer+`&#92;n`+fi+`&#92;n`+fa+`&#92;n`+fo+`&#92;n`+fei+`&#92;n`+fu+`"];`	
		
					}else{
		
						grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`&#92;n`+q+`"];`	
		
					}
		
					
					z = z + 1
					dot = dot + 1
				}

				yum = yum + 1
			}

		}
		fmt.Println("Termino Bien")

	}
	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdTiendas[dot] = hex.EncodeToString(hash[:])


		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		//z = z + 1
		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdTiendas[jas] = hex.EncodeToString(hash[:])


			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdTiendas[jas] = hex.EncodeToString(hash[:])


						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}

func arbolMerkleTiendasEndPoint(w http.ResponseWriter, req *http.Request){

	data := []byte(MerkleTiendasCrear())
    err := ioutil.WriteFile("MerkleTiendas.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearMerkleTiendas.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./MerkleTiendas.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}

func ConfigMerkleTiendasEndPoint(w http.ResponseWriter, req *http.Request){

	data := []byte(MerkleTiendasConfig(w, req))
    err := ioutil.WriteFile("MerkleTiendas.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearMerkleTiendas.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./MerkleTiendas.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}




func MerkleProductosCrear() string{

	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range invent{

		for ty := 0; ty < len(nodo.DatosInventario); ty++{

			z := 0
			for z < len(nodo.DatosInventario[ty].ProductosInventario){
	
				x := nodo.DatosInventario[ty].ProductosInventario[z].NombreProducto
				y := strconv.FormatInt(int64(nodo.DatosInventario[ty].ProductosInventario[z].CodigoProducto), 10)
				r := nodo.DatosInventario[ty].ProductosInventario[z].DescripcionProducto
				w := strconv.FormatInt(int64(nodo.DatosInventario[ty].ProductosInventario[z].PrecioProducto), 10)
				v := strconv.FormatInt(int64(nodo.DatosInventario[ty].ProductosInventario[z].CantidadProducto), 10)
				q := nodo.DatosInventario[ty].ProductosInventario[z].ImagenProducto
				e := nodo.DatosInventario[ty].ProductosInventario[z].AlmacenProducto
	
				encript := [] byte (x+","+y+","+r+","+w+","+v+","+q+","+e)
				hash := sha256.Sum256(encript)
				id[dot] = hex.EncodeToString(hash[:])

				IdProductos[dot] = hex.EncodeToString(hash[:])

	
				fmt.Println(id[dot])
	
				je := strconv.FormatInt(int64(dot), 10)
	
				grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`&#92;n`+q+`&#92;n`+e+`"];`	
	
				z = z + 1
				dot = dot + 1
			}
	


		}
		fmt.Println("Termino Bien")

	}
	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdProductos[dot] = hex.EncodeToString(hash[:])

		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdProductos[jas] = hex.EncodeToString(hash[:])



			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdProductos[jas] = hex.EncodeToString(hash[:])

						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}

func MerkleProductosConfig(w http.ResponseWriter, req *http.Request) string{

	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var datos ProductosMerkle
	json.NewDecoder(req.Body).Decode(&datos)

	//tienda.ID = params["Id nuevo"]
	PM = append(PM, datos)

	fer := strconv.FormatInt(int64(datos.CodigoProducto), 10)
	fi := datos.NombreProducto
	fa := datos.DescripcionProducto
	fo := datos.PrecioProducto
	fu := datos.CantidadProducto
	fe := datos.ImagenProducto
	fot := datos.AlmacenProducto

	foi := strconv.FormatInt(int64(datos.PrecioProducto), 10)
	fui := strconv.FormatInt(int64(datos.CantidadProducto), 10)

	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range invent{

		for ty := 0; ty < len(nodo.DatosInventario); ty++{

			z := 0
			for z < len(nodo.DatosInventario[ty].ProductosInventario){
	
				x := nodo.DatosInventario[ty].ProductosInventario[z].NombreProducto
				y := strconv.FormatInt(int64(nodo.DatosInventario[ty].ProductosInventario[z].CodigoProducto), 10)
				r := nodo.DatosInventario[ty].ProductosInventario[z].DescripcionProducto
				w := strconv.FormatInt(int64(nodo.DatosInventario[ty].ProductosInventario[z].PrecioProducto), 10)
				v := strconv.FormatInt(int64(nodo.DatosInventario[ty].ProductosInventario[z].CantidadProducto), 10)
				q := nodo.DatosInventario[ty].ProductosInventario[z].ImagenProducto
				e := nodo.DatosInventario[ty].ProductosInventario[z].AlmacenProducto
	
				encript := [] byte (x+","+y+","+r+","+w+","+v+","+q+","+e)
				hash := sha256.Sum256(encript)
				id[dot] = hex.EncodeToString(hash[:])

				IdProductos[dot] = hex.EncodeToString(hash[:])

	
				fmt.Println(id[dot])
	
				je := strconv.FormatInt(int64(dot), 10)

				if y == fer{

					nodo.DatosInventario[ty].ProductosInventario[z].NombreProducto = fi
					nodo.DatosInventario[ty].ProductosInventario[z].DescripcionProducto = fa
					nodo.DatosInventario[ty].ProductosInventario[z].PrecioProducto = fo
					nodo.DatosInventario[ty].ProductosInventario[z].CantidadProducto = fu
					nodo.DatosInventario[ty].ProductosInventario[z].ImagenProducto = fe
					nodo.DatosInventario[ty].ProductosInventario[z].AlmacenProducto = fot
	
	
					grafo+=`Nodo`+ je +`[shape=record,fillcolor=red,style=filled,label="`+id[dot]+`&#92;n&#92;n`+fi+`&#92;n`+fer+`&#92;n`+fa+`&#92;n`+foi+`&#92;n`+fui+`&#92;n`+fe+`&#92;n`+fot+`"];`	
	
				}else{
	
					grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`&#92;n`+q+`&#92;n`+e+`"];`	
	
				}
	
			
	
				z = z + 1
				dot = dot + 1
			}
	


		}
		fmt.Println("Termino Bien")

	}
	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdProductos[dot] = hex.EncodeToString(hash[:])

		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdProductos[jas] = hex.EncodeToString(hash[:])



			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdProductos[jas] = hex.EncodeToString(hash[:])

						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}

func arbolMerkleProductosEndPoint(w http.ResponseWriter, req *http.Request){

	data := []byte(MerkleProductosCrear())
    err := ioutil.WriteFile("MerkleProductos.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearMerkleProductos.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./MerkleProductos.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}

func ConfigMerkleProductosEndPoint(w http.ResponseWriter, req *http.Request){

	data := []byte(MerkleProductosConfig(w, req))
    err := ioutil.WriteFile("MerkleProductos.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearMerkleProductos.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./MerkleProductos.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}



func MerklePedidosCrear() string{

	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range pedir{

		z := 0
		for z < len(nodo.PEDIDO){

			x := nodo.PEDIDO[z].FechaPedido
			y := nodo.PEDIDO[z].TiendaPedido
			r := nodo.PEDIDO[z].DepartamentoPedido
			w := strconv.FormatInt(int64(nodo.PEDIDO[z].CalificacionPedido), 10)
			v := strconv.FormatInt(int64(nodo.PEDIDO[z].ClientePedido), 10)

			//var cod [100]int

			//tas := 0

			//for tas < len(nodo.PEDIDO[z].ProductosPedido){

			//	cod[tas] = strconv.FormatInt(int64(nodo.PEDIDO[z].ProductosPedido[tas].ProductoCodigo), 10)

			//	tas = tas + 1
			//}

			encript := [] byte (x+","+y+","+r+","+w+","+v)
			hash := sha256.Sum256(encript)
			id[dot] = hex.EncodeToString(hash[:])

			IdPedidos[dot] = hex.EncodeToString(hash[:])

			fmt.Println(id[dot])

			je := strconv.FormatInt(int64(dot), 10)

			grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`"];`	

			z = z + 1
			dot = dot + 1
		}

		fmt.Println("Termino Bien")

	}

	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdPedidos[dot] = hex.EncodeToString(hash[:])

		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdPedidos[jas] = hex.EncodeToString(hash[:])


			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdPedidos[jas] = hex.EncodeToString(hash[:])


						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}

func MerklePedidosConfig(w http.ResponseWriter, req *http.Request) string{

	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(req)
	var datos PedidosMerkle
	json.NewDecoder(req.Body).Decode(&datos)

	//tienda.ID = params["Id nuevo"]
	PeM = append(PeM, datos)

	fer := strconv.FormatInt(int64(datos.ClientePedido), 10)
	fi := datos.FechaPedido
	fa := datos.TiendaPedido
	fo := datos.DepartamentoPedido
	fu := datos.CalificacionPedido

	fui := strconv.FormatInt(int64(datos.CalificacionPedido), 10)


	var grafo string
	grafo="digraph G{\n"
	grafo+="graph [compound=true, labelloc=\"b\"];\n"
	grafo+=""

	var id [1000]string

	dot := 0

	for _, nodo := range pedir{

		z := 0
		for z < len(nodo.PEDIDO){

			x := nodo.PEDIDO[z].FechaPedido
			y := nodo.PEDIDO[z].TiendaPedido
			r := nodo.PEDIDO[z].DepartamentoPedido
			w := strconv.FormatInt(int64(nodo.PEDIDO[z].CalificacionPedido), 10)
			v := strconv.FormatInt(int64(nodo.PEDIDO[z].ClientePedido), 10)

			//var cod [100]int

			//tas := 0

			//for tas < len(nodo.PEDIDO[z].ProductosPedido){

			//	cod[tas] = strconv.FormatInt(int64(nodo.PEDIDO[z].ProductosPedido[tas].ProductoCodigo), 10)

			//	tas = tas + 1
			//}

			encript := [] byte (x+","+y+","+r+","+w+","+v)
			hash := sha256.Sum256(encript)
			id[dot] = hex.EncodeToString(hash[:])

			IdPedidos[dot] = hex.EncodeToString(hash[:])

			fmt.Println(id[dot])

			je := strconv.FormatInt(int64(dot), 10)

			if v == fer{

				nodo.PEDIDO[z].FechaPedido = fi
				nodo.PEDIDO[z].TiendaPedido = fa
				nodo.PEDIDO[z].DepartamentoPedido = fo
				nodo.PEDIDO[z].CalificacionPedido = fu


				grafo+=`Nodo`+ je +`[shape=record,fillcolor=red,style=filled,label="`+id[dot]+`&#92;n&#92;n`+fi+`&#92;n`+fa+`&#92;n`+fo+`&#92;n`+fui+`&#92;n`+fer+`"];`	

			}else{

				grafo+=`Nodo`+ je +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n`+x+`&#92;n`+y+`&#92;n`+r+`&#92;n`+w+`&#92;n`+v+`"];`	

			}

			

			z = z + 1
			dot = dot + 1
		}

		fmt.Println("Termino Bien")

	}

	if (dot % 2 == 0){
		fmt.Println("Nodos Pares")

	}else{
		fmt.Println("Nodos Impares")

		j := strconv.FormatInt(int64(dot), 10)

		encript := [] byte ("-"+j)
		hash := sha256.Sum256(encript)
		id[dot] = hex.EncodeToString(hash[:])

		IdPedidos[dot] = hex.EncodeToString(hash[:])

		grafo+=`Nodo`+ j +`[shape=record,fillcolor=green,style=filled,label="`+id[dot]+`&#92;n&#92;n-`+j+`"];`	

		dot = dot + 1
	}

	a := dot
	d := 0
	//r := 30
	jas := dot
	rar := dot
	cont := 0
	rot := 0
	zo := dot

	for rot < 1000{


		zo = a
		fmt.Println(zo)
		a = 0

		for a < zo {
			encript := [] byte (id[d]+","+id[d+1])
			hash := sha256.Sum256(encript)
			id[jas] = hex.EncodeToString(hash[:])

			IdPedidos[jas] = hex.EncodeToString(hash[:])


			fmt.Println(a)

			jo := strconv.FormatInt(int64(jas), 10)

			pr := strconv.FormatInt(int64(d), 10)

			ps := strconv.FormatInt(int64(d+1), 10)

			grafo+=`Nodo`+ jo +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n`+id[d]+`&#92;n`+id[d+1]+`"];`	

			grafo += "Nodo"+pr+" -> Nodo"+jo+";"
			grafo += "Nodo"+ps+" -> Nodo"+jo+";"

			a = a + 2
			d = d + 2
			cont = cont + 1
			jas = jas + 1

			if a == zo{
				fmt.Println(zo, a)
				if a == 2{
					fmt.Println("Paso final")
					rot = 1000 + 1

				}else{

					if (cont % 2 == 0){
						fmt.Println("Nodos Pares")
						fmt.Println("Otro Nivel")
						//rot = 1000 + 1
			
					}else{
						fmt.Println("Nodos Impares")
						fmt.Println("Otro Nivel")

						tot := strconv.FormatInt(int64(jas), 10)
			
						encript := [] byte ("-"+tot)
						hash := sha256.Sum256(encript)
						id[jas] = hex.EncodeToString(hash[:])

						IdPedidos[jas] = hex.EncodeToString(hash[:])


						grafo+=`Nodo`+ tot +`[shape=record,fillcolor=green,style=filled,label="`+id[jas]+`&#92;n&#92;n-`+tot+`"];`	

						jas = jas + 1
						//rot = 1000 + 1
					}

					d = rar
					rar = jas
					//r = r + 1
					cont = 0
					
				}
				

			}else{
				fmt.Println("Sigue la secuencia")
			}

		}


		a = a/2

		if (a % 2 == 0){
			fmt.Println(a)
			

		}else{
			a = a + 1
			fmt.Println(a)

		}

	}

	grafo+=""
	grafo+="}"
	return grafo

}


func arbolMerklePedidosEndPoint(w http.ResponseWriter, req *http.Request){

	data := []byte(MerklePedidosCrear())
    err := ioutil.WriteFile("MerklePedidos.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearMerklePedidos.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./MerklePedidos.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}

func ConfigMerklePedidosEndPoint(w http.ResponseWriter, req *http.Request){

	data := []byte(MerklePedidosConfig(w, req))
    err := ioutil.WriteFile("MerklePedidos.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearMerklePedidos.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./MerklePedidos.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}


func HashTiendasCrear() string{

	//var idCounter = make(map[CTienda]int)
	//var idCounter = make(map[int]int)
	//var keys = []int{}
	A := 0.0050

	//rot := 0

	//for _, id := range Tcomentario{
	//	hId := id.IDCliente
	//	hTienda := id.Tienda
	//	idCounter[hId]++
	//}

	//for ka := range idCounter {
	//	keys = append(keys, ka)
	//}

    //for rot < 1000{
	
		n := 7
		fot := 0
		var HKey [1000]int


		for _, k := range Tcomentario{

			x := k.IDCliente

			HKey[fot] = x

			dot := ((n+1)/2)-1

			fmt.Println(n)
			fmt.Println(dot)

			if dot == fot{
				n = n + 2
				jet := 0

				fmt.Println("Paso el igual")

				for jet < 1000{
					
					d := 2
					hum := n 
					for d < hum{
						if n % d == 0{
							fmt.Println(n)
							fmt.Println("No es primo")

							n = n + 2
							d = hum + 2

						}else{
							if d == (hum - 1){
								fmt.Println(n)
							    fmt.Println("Si es primo")

								jet = 1000 + 1
								d = hum + 2

							}else{
								d = d + 1
							}
							
						}

					}
					


				}
				//n = 11
				fot = fot + 1

			}else{
				fot = fot + 1

			}

		}
		m := n

		var HK [1000]int

		for _, k := range Tcomentario{

			x := k.IDCliente
			//mis := float64(x)

			trot := float64(x)*A

			red := float64(m)*math.Mod(trot, float64(1))
			j := int(red)

			if HK[j] == 0{
				HK[j] = x

			}else{
				i := 1
				jot := 0

				for jot < 1000{
					trot := float64(x)*A
					ris := math.Mod((float64(m)*math.Mod(trot, float64(1)) + float64(i*i)), float64(m))
					a := int(ris)

					if HK[a] == 0{
						HK[a] = x
						jot = 1000 + 1

					}else{
						i = i + 1
					}


				}

			}

		}

		var grafo string
	    grafo="digraph G{\n"
	    grafo+="graph [compound=true, labelloc=\"b\"];\n"
	    grafo+=""


		deta := 0

		for deta < m{

			nodet := HK[deta]

			if nodet == 0{
				bay := strconv.FormatInt(int64(deta), 10)

				grafo+=`N`+ bay +`[shape=record, label="`+bay+`&#92;n&#92;n----&#92;n----"];`
				deta = deta + 1

			}else{
				for _, nodo := range Tcomentario{

					cli := nodo.IDCliente
					tie := nodo.Tienda

					if nodet == cli{

						bay := strconv.FormatInt(int64(deta), 10)
						jay := strconv.FormatInt(int64(nodet), 10)
	
						grafo+=`N`+ bay +`[shape=record, label="`+bay+`&#92;n&#92;n`+jay+`&#92;n`+tie+`"];`
						
						deta = deta + 1
	
	
					}else{
						fmt.Println("Se Sigue Buscando")
					}
		
		
		
				}


			}

		}

		jeta := 1

		for jeta < m{
			bay := strconv.FormatInt(int64(jeta), 10)
			biy := strconv.FormatInt(int64(jeta - 1), 10)

			grafo += "N"+biy+" -> N"+bay+";"
			jeta = jeta + 1

		}

		for _, nodo := range Tcomentario{
			cli := nodo.IDCliente

			z := 0
		    for z < len(nodo.ComentTienda){

				co := nodo.ComentTienda[z].Coment

				y := 0
				for y < len(nodo.ComentTienda[z].SubComent){

					su := nodo.ComentTienda[z].SubComent[y].Sub

					f := 0
					for f < len(nodo.ComentTienda[z].SubComent[y].SubComent){

						su2 := nodo.ComentTienda[z].SubComent[y].SubComent[f].Sub

						heta := 0
						for heta < m{
							jodet := HK[heta]
			
							if jodet == cli{
								bay := strconv.FormatInt(int64(heta), 10)
								day := strconv.FormatInt(int64(heta+z+y+f), 10)

								grafo+=`C`+ day +`[shape=record, label="Comentario: `+co+`&#92;n&#92;n SubComentarios: `+su+`&#92;n`+su2+`"];`

								grafo += "N"+bay+" -> C"+day+";"

								heta = m + 1
			
							}else{
								heta = heta + 1
							}
						}


						f = f + 1

					}
					y = y + 1
				}
				z = z + 1
			}

		}


	grafo+=""
	grafo+="}"
	return grafo

	
	//}
		

}

func TablaHashTiendas(w http.ResponseWriter, req *http.Request){

	data := []byte(HashTiendasCrear())
    err := ioutil.WriteFile("HashTiendas.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearHashTiendas.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./HashTiendas.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}


func HashProductosCrear() string{

	//var idCounter = make(map[CTienda]int)
	//var idCounter = make(map[int]int)
	//var keys = []int{}
	A := 0.0050

	//rot := 0

	//for _, id := range Tcomentario{
	//	hId := id.IDCliente
	//	hTienda := id.Tienda
	//	idCounter[hId]++
	//}

	//for ka := range idCounter {
	//	keys = append(keys, ka)
	//}

    //for rot < 1000{
	
		n := 7
		fot := 0
		var HKey [1000]int


		for _, k := range Icomentario{

			x := k.IDCliente

			HKey[fot] = x

			dot := ((n+1)/2)-1

			fmt.Println(n)
			fmt.Println(dot)

			if fot == dot{
				n = n + 2
				jet := 0

				fmt.Println("Paso el igual")

				for jet < 1000{
					
					d := 2
					hum := n 
					for d < hum{
						if n % d == 0{
							fmt.Println(n)
							fmt.Println("No es primo")

							n = n + 2
							d = hum + 2

						}else{
							if d == (hum - 1){
								fmt.Println(n)
							    fmt.Println("Si es primo")

								jet = 1000 + 1
								d = hum + 2

							}else{
								d = d + 1
							}
							
						}

					}
					


				}
				//n = 11
				fot = fot + 1

			}else{
				fot = fot + 1

			}

		}
		m := n

		var HK [1000]int

		for _, k := range Icomentario{

			x := k.IDCliente
			//mis := float64(x)

			trot := float64(x)*A

			red := float64(m)*math.Mod(trot, float64(1))
			j := int(red)

			if HK[j] == 0{
				HK[j] = x

			}else{
				i := 1
				jot := 0

				for jot < 1000{
					trot := float64(x)*A
					ris := math.Mod((float64(m)*math.Mod(trot, float64(1)) + float64(i*i)), float64(m))
					a := int(ris)

					if HK[a] == 0{
						HK[a] = x
						jot = 1000 + 1

					}else{
						i = i + 1
					}


				}

			}

		}

		var grafo string
	    grafo="digraph G{\n"
	    grafo+="graph [compound=true, labelloc=\"b\"];\n"
	    grafo+=""


		deta := 0

		for deta < m{

			nodet := HK[deta]

			if nodet == 0{
				bay := strconv.FormatInt(int64(deta), 10)

				grafo+=`N`+ bay +`[shape=record, label="`+bay+`&#92;n&#92;n----&#92;n----"];`
				deta = deta + 1

			}else{
				for _, nodo := range Icomentario{

					cli := nodo.IDCliente
					tie := nodo.Producto

					if nodet == cli{

						bay := strconv.FormatInt(int64(deta), 10)
						jay := strconv.FormatInt(int64(nodet), 10)
	
						grafo+=`N`+ bay +`[shape=record, label="`+bay+`&#92;n&#92;n`+jay+`&#92;n`+tie+`"];`
						
						deta = deta + 1
	
	
					}else{
						fmt.Println("Se Sigue Buscando")
					}
		
		
		
				}


			}

		}

		jeta := 1

		for jeta < m{
			bay := strconv.FormatInt(int64(jeta), 10)
			biy := strconv.FormatInt(int64(jeta - 1), 10)

			grafo += "N"+biy+" -> N"+bay+";"
			jeta = jeta + 1

		}

		for _, nodo := range Icomentario{
			cli := nodo.IDCliente

			z := 0
		    for z < len(nodo.ComentProducto){

				co := nodo.ComentProducto[z].Coment

				y := 0
				for y < len(nodo.ComentProducto[z].SubComent){

					su := nodo.ComentProducto[z].SubComent[y].Sub

					f := 0
					for f < len(nodo.ComentProducto[z].SubComent[y].SubComent){

						su2 := nodo.ComentProducto[z].SubComent[y].SubComent[f].Sub

						heta := 0
						for heta < m{
							jodet := HK[heta]
			
							if jodet == cli{
								bay := strconv.FormatInt(int64(heta), 10)
								day := strconv.FormatInt(int64(heta+z+y+f), 10)

								grafo+=`C`+ day +`[shape=record, label="Comentario: `+co+`&#92;n&#92;n SubComentarios: `+su+`&#92;n`+su2+`"];`

								grafo += "N"+bay+" -> C"+day+";"

								heta = m + 1
			
							}else{
								heta = heta + 1
							}
						}


						f = f + 1

					}
					y = y + 1
				}
				z = z + 1
			}

		}


	grafo+=""
	grafo+="}"
	return grafo

	
	//}
		

}


func TablaHashProductos(w http.ResponseWriter, req *http.Request){

	data := []byte(HashProductosCrear())
    err := ioutil.WriteFile("HashProductos.dot", data, 0644)
    if err != nil {
        log.Fatal(err)
    }
	//Generamos la imagen
	app := "crearHashProductos.bat"
	_, err2 := exec.Command(app).Output()
	if err2 != nil {
		fmt.Println("errrooor :(")
		fmt.Println(err2)
	} else {
		fmt.Println("Todo bien")
	}
	//abrimos la imagen
	img, err3 := os.Open("./HashProductos.png")
    if err3 != nil {
        log.Fatal(err3) // perhaps handle this nicer
    }
    defer img.Close()
	//devolvemos como respuesta la imagen
    w.Header().Set("Content-Type", "image/png")
    io.Copy(w, img)


}






func main() {
	router := mux.NewRouter().StrictSlash(true)

	//endpoints
	router.HandleFunc("/cargartienda", ListadoDeTiendasEndpoint).Methods("GET")

	router.HandleFunc("/cargarinventario", ListadoDeInventarioEndpoint).Methods("GET")

	router.HandleFunc("/cargarpedido", ListadoDePedidoEndpoint).Methods("GET")



	router.HandleFunc("/cargarusuarios", ListadoDeUsuariosEndpoint).Methods("GET")

	router.HandleFunc("/cargarusuarios/{Cuenta}", ObtenerTipoCuentaEndpoint).Methods("GET")

	router.HandleFunc("/nuevoUsuarios", CrearUsuarioEndPoint).Methods("POST")

	router.HandleFunc("/EliminarUsuario", EliminarUsuarioEndPoint).Methods("POST")

	router.HandleFunc("/InicioSesion", InicioSesionEndPoint).Methods("POST")


	router.HandleFunc("/subirGrafo", SubirGrafoEndPoint).Methods("POST")

	router.HandleFunc("/verDatosGrafo", DatosGrafoEndPoint).Methods("GET")



	router.HandleFunc("/arbolMerkleUsuario", arbolMerkleUsuarioEndPoint).Methods("GET")
	router.HandleFunc("/arbolMerkleTiendas", arbolMerkleTiendasEndPoint).Methods("GET")
	router.HandleFunc("/arbolMerkleProductos", arbolMerkleProductosEndPoint).Methods("GET")
	router.HandleFunc("/arbolMerklePedidos", arbolMerklePedidosEndPoint).Methods("GET")


	router.HandleFunc("/ConfigMerkleUsuario", ConfigMerkleUsuarioEndPoint).Methods("POST")
	router.HandleFunc("/ConfigMerkleTiendas", ConfigMerkleTiendasEndPoint).Methods("POST")
	router.HandleFunc("/ConfigMerkleProductos", ConfigMerkleProductosEndPoint).Methods("POST")
	router.HandleFunc("/ConfigMerklePedidos", ConfigMerklePedidosEndPoint).Methods("POST")

	//router.HandleFunc("/ArregMerkleUsuario", ArregMerkleUsuarioEndPoint).Methods("GET")
	//router.HandleFunc("/ArregMerkleTiendas", ArregMerkleTiendasEndPoint).Methods("GET")
	//router.HandleFunc("/ArregMerkleProductos", ArregMerkleProductosEndPoint).Methods("GET")
	//router.HandleFunc("/ArregMerklePedidos", ArregMerklePedidosEndPoint).Methods("GET")

	router.HandleFunc("/tablaHashTiendas", TablaHashTiendas).Methods("GET")
	router.HandleFunc("/tablaHashProductos", TablaHashProductos).Methods("GET")




	router.HandleFunc("/subirComentarioTienda", subirComentarioTiendaEndPoint).Methods("POST")

	router.HandleFunc("/verComentarioTienda", ComentarioTiendaEndPoint).Methods("GET")


	router.HandleFunc("/subirComentarioProducto", subirComentarioProductoEndPoint).Methods("POST")

	router.HandleFunc("/verComentarioProducto",ComentarioProductoEndPoint).Methods("GET")






	router.HandleFunc("/cargartienda/indice/{Indice}", ObtenerIndiceTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/tipotienda/{Nombre}", ObtenerNombreTipoTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/nombre/{Nombre}", ObtenerNombreTiendaEndpoint).Methods("GET")
	router.HandleFunc("/cargartienda/calificacion/{Contacto}", ObtenerContactoTiendaEndpoint).Methods("GET")

	//---router.HandleFunc("/cargartienda/{Id nuevo}", CrearTiendaEndpoint).Methods("POST")

	//prueba post
	router.HandleFunc("/cargartienda/subir", CrearEndPoint).Methods("POST")

	router.HandleFunc("/cargartienda/{Nombre}", BorrarTiendaEndpoint).Methods("DELETE")

	router.HandleFunc("/TiendaEspecifica", TiendaEspecificaEndPoint).Methods("POST")

	router.HandleFunc("/Eliminar", EliminarRegistroEndPoint).Methods("POST")



	router.HandleFunc("/Inventario", InventarioEndPoint).Methods("POST")
	router.HandleFunc("/Pedidos", PedidosEndPoint).Methods("POST")


	router.HandleFunc("/ArbolSinCifrar", getArbolSinCifrar).Methods("GET")
	router.HandleFunc("/ArbolCifrarSuave", getArbolCifrarSuave).Methods("GET")
	router.HandleFunc("/ArbolCifrar", getArbolCifrar).Methods("GET")

	router.HandleFunc("/Encriptacion", getEncriptado).Methods("GET")

	router.HandleFunc("/imgGrafo", getGrafoRobot).Methods("GET")


	







    // no usar esta seccion encacillada{

	router.HandleFunc("/Descargar", DescargarArchivoEndPoint)

	router.HandleFunc("/orden", Ordenamiento).Methods("GET")

	router.HandleFunc("/getArreglo", Grafica).Methods("GET")

	// }no usar esta seccion encacillada





	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}






// no usar esta seccion encacillada{

type edge struct {
	node  string
	label string
}
type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(from, to, label string) {
	g.nodes[from] = append(g.nodes[from], edge{node: to, label: label})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (e *edge) String() string {
	return fmt.Sprintf("%v", e.node)
}

func (g *graph) String() string {
	out := `digraph finite_state_machine {
		rankdir=LR;
		size="8,5"
		node [shape = circle];`
	for k := range g.nodes {
		for _, v := range g.getEdges(k) {
			out += fmt.Sprintf("\t%s -> %s\t[ label = \"%s\" ];\n", k, v.node, v.label)
		}
	}
	out += "}"
	return out
}
// }no usar esta seccion encacillada