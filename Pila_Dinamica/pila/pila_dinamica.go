package pila

const TAM_INICIAL = 20
const VALOR_DE_REDIMENSION = 2

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func redimensionarPila[T any](pila *pilaDinamica[T]) {
	if cap(pila.datos) == pila.cantidad {

		vectorAuxiliar(pila, cap(pila.datos)*VALOR_DE_REDIMENSION)

	} else if pila.cantidad*4 <= cap(pila.datos) && (cap(pila.datos)/VALOR_DE_REDIMENSION > TAM_INICIAL) {

		vectorAuxiliar(pila, cap(pila.datos)/VALOR_DE_REDIMENSION)
	}

}

func vectorAuxiliar[T any](pila *pilaDinamica[T], tam int) {

	aux := make([]T, tam)
	copy(aux, pila.datos)
	pila.datos = aux

}

func CrearPilaDinamica[T any]() Pila[T] {
	// definir una constante para la creacion
	pila := new(pilaDinamica[T])

	pila.datos = make([]T, TAM_INICIAL)

	pila.cantidad = 0

	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {

	return pila.cantidad == 0
}

func chequearPilaVacia[T any](pila *pilaDinamica[T]) {
	if pila.EstaVacia() {

		panic("La pila esta vacia")
	}

}

func (pila *pilaDinamica[T]) VerTope() T {

	chequearPilaVacia(pila)

	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elem T) {

	redimensionarPila(pila)

	pila.datos[pila.cantidad] = elem

	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {

	chequearPilaVacia(pila)

	pila.cantidad--

	redimensionarPila(pila)

	return pila.datos[pila.cantidad]
}
