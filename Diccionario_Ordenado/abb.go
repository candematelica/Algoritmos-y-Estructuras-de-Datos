package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(clave1, clave2 K) int
}

type iteradorABB[K comparable, V any] struct {
	abb                *abb[K, V]
	desde              *K
	hasta              *K
	elementosOrdenados TDAPila.Pila[nodoAbb[K, V]]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func (nodo *nodoAbb[K, V]) buscarNodoABB(clave K, nodopadre *nodoAbb[K, V], funcCmp func(K, K) int) (bool, *nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return false, nodopadre, nodo
	}
	if funcCmp(nodo.clave, clave) == 0 {
		return true, nodopadre, nodo
	}
	if funcCmp(nodo.clave, clave) > 0 {
		return nodo.izquierdo.buscarNodoABB(clave, nodo, funcCmp)
	}

	return nodo.derecho.buscarNodoABB(clave, nodo, funcCmp)
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodoEsta, nodoPadre, nodo := abb.raiz.buscarNodoABB(clave, abb.raiz, abb.cmp) //nodoPadre y nodo son punteros
	if !nodoEsta {
		nodoNuevo := &nodoAbb[K, V]{clave: clave, dato: dato}
		if nodoPadre == nil {
			abb.raiz = nodoNuevo
		} else if abb.cmp(nodoPadre.clave, clave) > 0 {
			nodoPadre.izquierdo = nodoNuevo
		} else {
			nodoPadre.derecho = nodoNuevo
		}
		abb.cantidad++
	} else {
		nodo.dato = dato
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodoEsta, _, _ := abb.raiz.buscarNodoABB(clave, abb.raiz, abb.cmp)
	return nodoEsta
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodoEsta, _, nodo := abb.raiz.buscarNodoABB(clave, abb.raiz, abb.cmp)
	if !nodoEsta {
		panic("La clave no pertenece al diccionario")
	}

	return nodo.dato
}

func (abb *abb[K, V]) buscarRemplazo(nodo, padreNodo *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo.izquierdo == nil {
		return nodo, padreNodo
	}

	return abb.buscarRemplazo(nodo.izquierdo, nodo)
}

func (nodo *nodoAbb[K, V]) borrarNodo(clave K, nodoPadre *nodoAbb[K, V], abb *abb[K, V]) V {
	dato := nodo.dato
	//CASOS:
	if nodo.izquierdo == nil && nodo.derecho == nil { //Borrar un elemento sin hijos
		if abb.cmp(nodoPadre.clave, clave) > 0 { // Es hijo izquierdo del padre
			nodoPadre.izquierdo = nil
		} else if abb.cmp(nodoPadre.clave, clave) < 0 { // Es hijo derecho del padre
			nodoPadre.derecho = nil
		} else { // El nodo a borrar es la raiz
			abb.raiz = nil
		}
		(abb.cantidad)--
	} else if nodo.izquierdo == nil && nodo.derecho != nil { // El nodo tiene 1 solo hijo (y es el derecho)
		if abb.cmp(nodoPadre.clave, clave) > 0 { // Es hijo izquierdo del padre
			nodoPadre.izquierdo = nodo.derecho
		} else if abb.cmp(nodoPadre.clave, clave) < 0 { // Es hijo derecho del padre
			nodoPadre.derecho = nodo.derecho
		} else { // El nodo a borrar es la raiz
			abb.raiz = nodo.derecho
		}
		(abb.cantidad)--
	} else if nodo.derecho == nil && nodo.izquierdo != nil { // El nodo tiene 1 solo hijo (y es el izquierdo)
		if abb.cmp(nodoPadre.clave, clave) > 0 { // Es hijo izquierdo del padre
			nodoPadre.izquierdo = nodo.izquierdo
		} else if abb.cmp(nodoPadre.clave, clave) < 0 { // Es hijo derecho del padre
			nodoPadre.derecho = nodo.izquierdo
		} else { // El nodo a borrar es la raiz
			abb.raiz = nodo.izquierdo
		}
		(abb.cantidad)--
	} else if nodo.izquierdo != nil && nodo.derecho != nil { //Borrar elementos con dos hijos (buscamos los posibles remplazantes)
		nodoRemplazo, nodoPadreRemplazo := abb.buscarRemplazo(nodo.derecho, nodo)
		nodoRemplazo.borrarNodo(nodoRemplazo.clave, nodoPadreRemplazo, abb)
		nodoRemplazo.izquierdo = nodo.izquierdo
		nodoRemplazo.derecho = nodo.derecho
		if abb.cmp(nodoPadre.clave, clave) > 0 { // Es hijo izquierdo del padre
			nodoPadre.izquierdo = nodoRemplazo
		} else if abb.cmp(nodoPadre.clave, clave) < 0 { // Es hijo derecho del padre
			nodoPadre.derecho = nodoRemplazo
		} else { // El nodo a borrar es la raiz
			abb.raiz = nodoRemplazo
		}
	}
	return dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodoEsta, nodoPadre, nodo := abb.raiz.buscarNodoABB(clave, abb.raiz, abb.cmp)
	if !nodoEsta {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).borrarNodo(clave, nodoPadre, abb)
}

func (nodo *nodoAbb[K, V]) iteradorInterno(visitar func(clave K, dato V) bool, desde, hasta *K, funcCmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}

	if desde != nil && funcCmp(nodo.clave, *desde) < 0 {
		// Si la clave del nodo es menor que el límite inferior del rango, continuar con el subárbol derecho
		return nodo.derecho.iteradorInterno(visitar, desde, hasta, funcCmp)
	} else if hasta != nil && funcCmp(nodo.clave, *hasta) > 0 {
		// Si la clave del nodo es mayor que el límite superior del rango, continuar con el subárbol izquierdo
		return nodo.izquierdo.iteradorInterno(visitar, desde, hasta, funcCmp)
	} else {
		// La clave del nodo está dentro del rango, continuar con ambos subárboles
		if !nodo.izquierdo.iteradorInterno(visitar, desde, hasta, funcCmp) {
			return false
		}
		if !visitar(nodo.clave, nodo.dato) {
			// Si la función de visita retorna falso, terminar la iteración
			return false
		}
	}

	return nodo.derecho.iteradorInterno(visitar, desde, hasta, funcCmp)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iteradorInterno(visitar, desde, hasta, abb.cmp)
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.raiz.iteradorInterno(visitar, nil, nil, abb.cmp)
}

func (nodo *nodoAbb[K, V]) apilarHijosIzquierdos(pila TDAPila.Pila[nodoAbb[K, V]], funcCmp func(K, K) int, desde, hasta *K) TDAPila.Pila[nodoAbb[K, V]] {
	if nodo == nil {
		return pila
	}
	if desde == nil && hasta == nil {
		pila.Apilar(*nodo)
	} else if desde == nil {
		if funcCmp(nodo.clave, *hasta) > 0 {
			return nodo.izquierdo.apilarHijosIzquierdos(pila, funcCmp, desde, hasta)
		}
		pila.Apilar(*nodo)
	} else if hasta == nil {
		if funcCmp(nodo.clave, *desde) < 0 {
			return nodo.derecho.apilarHijosIzquierdos(pila, funcCmp, desde, hasta)
		}
		pila.Apilar(*nodo)
	} else {
		if funcCmp(nodo.clave, *desde) < 0 {
			return nodo.derecho.apilarHijosIzquierdos(pila, funcCmp, desde, hasta)
		}
		if funcCmp(nodo.clave, *hasta) > 0 {
			return nodo.izquierdo.apilarHijosIzquierdos(pila, funcCmp, desde, hasta)
		}
		pila.Apilar(*nodo)
	}

	return nodo.izquierdo.apilarHijosIzquierdos(pila, funcCmp, desde, hasta)
}

func (abb *abb[K, V]) crearPilaElementosOrdenados(desde, hasta *K) TDAPila.Pila[nodoAbb[K, V]] {
	pilaElementosOrdenados := TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	pilaElementosOrdenados = abb.raiz.apilarHijosIzquierdos(pilaElementosOrdenados, abb.cmp, desde, hasta)
	return pilaElementosOrdenados
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return &iteradorABB[K, V]{abb, desde, hasta, abb.crearPilaElementosOrdenados(desde, hasta)}
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (iter *iteradorABB[K, V]) HaySiguiente() bool {
	return !iter.elementosOrdenados.EstaVacia()
}

func (iter *iteradorABB[K, V]) chequearIterador() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {
	iter.chequearIterador()
	return iter.elementosOrdenados.VerTope().clave, iter.elementosOrdenados.VerTope().dato
}

func (iter *iteradorABB[K, V]) Siguiente() {
	iter.chequearIterador()
	desapilado := iter.elementosOrdenados.Desapilar()
	if desapilado.derecho != nil {
		iter.elementosOrdenados = desapilado.derecho.apilarHijosIzquierdos(iter.elementosOrdenados, iter.abb.cmp, iter.desde, iter.hasta)
	}
}
