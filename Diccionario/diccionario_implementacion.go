package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	LARGO_MAX_LISTA      = 3
	FACTOR_DE_CARGA_MAX  = 3
	FACTOR_DE_CARGA_MIN  = 0.5
	VALOR_DE_REDIMENSION = 2
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashImplementacion[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	cantidad int
}

type iteradorHash[K comparable, V any] struct {
	diccionario   *hashImplementacion[K, V]
	pos           int
	iteradorLista TDALista.IteradorLista[parClaveValor[K, V]]
}

func crearTabla[K comparable, V any](tam int) []TDALista.Lista[parClaveValor[K, V]] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)
	return tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashImplementacion[K, V]{crearTabla[K, V](0), 0}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// Funcion de hash de Bob Jenkins
func hash(clave []byte, largo int) uint32 {
	var hash uint32

	for _, c := range clave {
		hash += uint32(c)
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}
	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)

	return hash % uint32(largo)
}

func (diccionario *hashImplementacion[K, V]) buscarParClaveValor(clave K) (bool, parClaveValor[K, V], TDALista.IteradorLista[parClaveValor[K, V]]) {
	claveABytes := convertirABytes[K](clave)
	pos := int(hash(claveABytes, len(diccionario.tabla)))
	if diccionario.tabla[pos] == nil {
		diccionario.tabla[pos] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
		iter := diccionario.tabla[pos].Iterador()
		return false, parClaveValor[K, V]{}, iter
	}

	iter := diccionario.tabla[pos].Iterador()
	for iter.HaySiguiente() {
		if iter.VerActual().clave == clave {
			return true, iter.VerActual(), iter
		}
		iter.Siguiente()
	}

	return false, parClaveValor[K, V]{}, iter
}

func (diccionario *hashImplementacion[K, V]) redimensionar(tamNuevo int) {
	nuevaTabla := crearTabla[K, V](tamNuevo)

	for _, lista := range diccionario.tabla {
		if lista != nil {
			lista.Iterar(func(nodo parClaveValor[K, V]) bool {
				claveABytes := convertirABytes[K](nodo.clave)
				pos := int(hash(claveABytes, tamNuevo))
				if nuevaTabla[pos] == nil {
					nuevaTabla[pos] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
				}
				nuevaTabla[pos].InsertarUltimo(nodo)
				return true
			})
		}
	}

	diccionario.tabla = nuevaTabla
}

func (diccionario *hashImplementacion[K, V]) Guardar(clave K, dato V) {
	var nodoNuevo parClaveValor[K, V]
	nodoNuevo.clave = clave
	nodoNuevo.dato = dato

	if len(diccionario.tabla) == 0 {
		nuevaTabla := crearTabla[K, V](2)
		diccionario.tabla = nuevaTabla
	}

	claveEsta, _, iterLista := diccionario.buscarParClaveValor(clave)
	if claveEsta {
		iterLista.Borrar()
	} else {
		(diccionario.cantidad)++
	}
	iterLista.Insertar(nodoNuevo)

	if float32(diccionario.cantidad/len(diccionario.tabla)) > FACTOR_DE_CARGA_MAX {
		tamNuevo := len(diccionario.tabla) * VALOR_DE_REDIMENSION
		diccionario.redimensionar(tamNuevo)
	}
}

func (diccionario *hashImplementacion[K, V]) Pertenece(clave K) bool {
	claveEsta, _, _ := diccionario.buscarParClaveValor(clave)
	return claveEsta
}

func (diccionario *hashImplementacion[K, V]) Obtener(clave K) V {

	claveEsta, claveValor, _ := diccionario.buscarParClaveValor(clave)
	if claveEsta {
		return claveValor.dato
	}

	panic("La clave no pertenece al diccionario")
}

func (diccionario *hashImplementacion[K, V]) Borrar(clave K) V {

	claveEsta, claveValor, iterLista := diccionario.buscarParClaveValor(clave)
	if claveEsta {
		iterLista.Borrar()
		(diccionario.cantidad)--
		if float32(diccionario.cantidad/len(diccionario.tabla)) < FACTOR_DE_CARGA_MIN {
			tamNuevo := len(diccionario.tabla) / 2
			diccionario.redimensionar(tamNuevo)
		}
		return claveValor.dato
	}

	panic("La clave no pertenece al diccionario")
}

func (diccionario *hashImplementacion[K, V]) Cantidad() int {
	return diccionario.cantidad
}

func (diccionario *hashImplementacion[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	var corte bool
	for _, lista := range diccionario.tabla {
		if lista != nil {
			lista.Iterar(func(nodo parClaveValor[K, V]) bool {
				corte = visitar(nodo.clave, nodo.dato)
				return corte
			})
			if !corte {
				return
			}
		}
	}
}

func (diccionario *hashImplementacion[K, V]) Iterador() IterDiccionario[K, V] {

	var iterLista TDALista.IteradorLista[parClaveValor[K, V]]
	pos, _ := diccionario.buscarLista(0)
	lista := diccionario.tabla[pos]

	iterLista = lista.Iterador()

	return &iteradorHash[K, V]{diccionario: diccionario, pos: pos, iteradorLista: iterLista}
}

func (dicc hashImplementacion[K, V]) buscarLista(pos int) (int, bool) {
	for i := pos; i < dicc.Cantidad(); i++ {
		if dicc.tabla[i] != nil && !dicc.tabla[i].EstaVacia() {
			return i, true
		}
	}

	return -1, false
}

func (iterador *iteradorHash[K, V]) HaySiguiente() bool {
	if iterador.diccionario.cantidad == 0 {
		return false
	}
	if iterador.diccionario.tabla[iterador.pos] != nil {
		if iterador.iteradorLista.HaySiguiente() {
			return true
		} else {
			_, hayOtraLista := iterador.diccionario.buscarLista(iterador.pos + 1)
			return hayOtraLista
		}
	}

	return false
}

func (iterador *iteradorHash[K, V]) chequearIterador() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iterador *iteradorHash[K, V]) VerActual() (K, V) {
	iterador.chequearIterador()
	return iterador.iteradorLista.VerActual().clave, iterador.iteradorLista.VerActual().dato
}

func (iterador *iteradorHash[K, V]) Siguiente() {
	iterador.chequearIterador()

	if iterador.iteradorLista.HaySiguiente() {
		iterador.iteradorLista.Siguiente()
	}
	if !iterador.iteradorLista.HaySiguiente() {
		posLista, hayOtraLista := iterador.diccionario.buscarLista(iterador.pos + 1)
		if hayOtraLista {
			iterador.iteradorLista = iterador.diccionario.tabla[posLista].Iterador()
			iterador.pos = posLista
		}
	}
}
