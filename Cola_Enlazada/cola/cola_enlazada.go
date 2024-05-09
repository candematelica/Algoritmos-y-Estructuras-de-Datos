package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {

	return &colaEnlazada[T]{primero: nil, ultimo: nil}
}
func crearNodo[T any](dato T) *nodoCola[T] {

	return &nodoCola[T]{dato: dato, prox: nil}
}
func (cola *colaEnlazada[T]) EstaVacia() bool {

	return cola.primero == nil

}

func chequearColaVacia[T any](cola *colaEnlazada[T]) {
	if cola.EstaVacia() {

		panic("La cola esta vacia")
	}

}

func (cola *colaEnlazada[T]) VerPrimero() T {
	chequearColaVacia(cola)

	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elem T) {

	nuevoNodo := crearNodo(elem)

	if cola.EstaVacia() {
		cola.primero = nuevoNodo

	} else {
		cola.ultimo.prox = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	chequearColaVacia(cola)
	dato := cola.primero.dato

	cola.primero = cola.primero.prox

	if cola.primero == nil {
		cola.ultimo = nil
	}

	return dato

}
