package diccionario

import (
	"fmt"
)

const (
	COEFICIENTE_REDIMENSION float64 = 0.7
	FACTOR_REDIMENSION      int     = 2
	VACIO                   int     = 0
	OCUPADO                 int     = 1
	BORRADO                 int     = 2

	TAMANO int = 6
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int // Solo hace referencia a ocupados
	tam      int
	borrados int
}

type iterDiccionario[K comparable, V any] struct {
	hash hashCerrado[K, V]
	pos  int
}

func crearCeldaHash[K comparable, V any]() celdaHash[K, V] {
	nuevaCelda := new(celdaHash[K, V])
	nuevaCelda.estado = 0
	return *nuevaCelda
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

	nuevoHash := new(hashCerrado[K, V])
	nuevoHash.tam = TAMANO
	nuevoHash.cantidad = 0
	nuevoHash.borrados = 0

	nuevoHash.tabla = crearTabla[K, V](nuevoHash.tam)

	return nuevoHash
}

func crearTabla[K comparable, V any](capacidad int) []celdaHash[K, V] {
	nuevaTabla := make([]celdaHash[K, V], capacidad)
	for i := 0; i < capacidad; i++ {
		nuevaCelda := crearCeldaHash[K, V]()
		nuevaTabla = append(nuevaTabla, nuevaCelda)
	}
	return nuevaTabla
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

	factorCarga := float64((hash.cantidad + hash.borrados)) / float64(hash.tam)
	if factorCarga >= COEFICIENTE_REDIMENSION {
		hash.redimensionar(FACTOR_REDIMENSION * hash.tam)
	}

	posicion, err := hash.buscar(clave)

	if err != nil { // el elemento no está
		hash.tabla[posicion].dato = dato
	} 

	hash.tabla[posicion].clave = clave
	hash.tabla[posicion].dato = dato
	hash.tabla[posicion].estado = OCUPADO
	hash.cantidad++

}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, err := hash.buscar(clave)
	return err != nil
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos, err := hash.buscar(clave)
	if err != nil {
		panic(err)
	}
	return hash.tabla[pos].dato
	
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	posicion, err := hash.buscar(clave)

	if err != nil{
		panic(err)
	}

	elemento := hash.tabla[posicion]
	elemento.estado = BORRADO
	hash.cantidad --
	hash.borrados ++
	return elemento.dato
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, valor V) bool) {

	for _,elem := range(hash.tabla){

		if elem.estado == VACIO || elem.estado == BORRADO{
			continue
		}
		if !visitar(elem.clave,elem.dato){
			break
		}

	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

	return &iterDiccionario[K, V]{}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return true
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	var clave K
	var valor V
	return clave, valor
}

func (iter *iterDiccionario[K, V]) Siguiente() {

}

func (hash *hashCerrado[K, V]) redimensionar(nuevaCapacidad int) {

	tablaAnterior := hash.tabla
	hash.tabla = crearTabla[K, V](nuevaCapacidad)
	hash.tam = nuevaCapacidad
	hash.borrados = 0

	for _, elem := range tablaAnterior {
		if elem.estado == OCUPADO {
			K, V := elem.clave, elem.dato
			hash.Guardar(K, V)
		}
	}
}

func convertirABytes[K comparable](clave K) []byte {

	return []byte(fmt.Sprintf("%v", clave))
}

func sdbmHash(data []byte) uint64 {
	var hash uint64

	for _, b := range data {
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

func (hash *hashCerrado[K, V]) hashear(clave K) int {
	claveByte := convertirABytes(clave)
	hashing := sdbmHash(claveByte)
	return int(hashing) % hash.tam
}

func (hash *hashCerrado[K, V]) buscar(clave K) (int, error) {
	posicion := hash.hashear(clave)
	primeraPorcion := hash.tabla[posicion:hash.tam]
	porcionAuxiliar := hash.tabla[:posicion]

	for posicion, celdaActual := range(primeraPorcion) {

		if celdaActual.estado == OCUPADO && celdaActual.clave == clave {
			return  posicion, nil
		} else if celdaActual.estado == VACIO {
			return -1, fmt.Errorf("La clave no pertenece al diccionario")
		} 

		continue
	}

	for _, celdaActual := range porcionAuxiliar {
		if celdaActual.estado == OCUPADO && celdaActual.clave == clave {
			return posicion, nil
		} else if celdaActual.estado == VACIO {
			return -1, fmt.Errorf("La clave no pertenece al diccionario")
		}
		continue
	}

	return posicion, nil
}
