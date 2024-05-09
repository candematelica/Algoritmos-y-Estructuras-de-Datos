package lista

type Lista[T any] interface {
	//EstaVacia devuelve verdadero si la lista no contiene ningun elemento, falso en caso contario.
	EstaVacia() bool
	//InsertarPrimero inserta el elemento pasado por parametro al inicio de la lista
	InsertarPrimero(T)
	//InsertarUltimo inserta el elemento pasado por parametro al final de la lista
	InsertarUltimo(T)
	//BorrarPrimero borra el primer elemento de la lista, devuelve un panic si esta vacia
	BorrarPrimero() T
	//VerPrimero muestra el primer elemento de la lista, devuelve un panic si esta vacia
	VerPrimero() T
	//VerUltimo muestra el primer elemento de la lista, devuelve un panic si esta vacia
	VerUltimo() T
	//Largo devuelve la cantidad de elementos que tiene la lista
	Largo() int
	//Iterar es el iterador interno del TDA
	Iterar(visitar func(T) bool)
	//Iterador es el iterador externo del TDA
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	//VerActual devuelve el elemento sobre el cual estamos parados en la lista, en caso de que
	//se llame a esta funcion cuando el iterador ya haya terminado de iterar, devolvera un panic
	VerActual() T
	//HaySiguiente devuelve true si se puede seguir iterando (es decir si hay elementos despues
	//del cual estemos parados). En caso contrario devuelve false
	HaySiguiente() bool
	//Siguiente devuelve el siguiente elemento a la posicion en la que estamos parados en la lista, en caso
	//de que se llame a esta funcion cuando el iterador ya haya terminado de iterar, devolvera un panic
	Siguiente()
	//Insertar inserta un elemento pasado por parametro en la posicion en la que estamos parados en la lista
	Insertar(T)
	//Borrar borra y devuelve el elemento sobre el cual estemos parados en la lista, en caso de que
	//se llame a esta funcion cuando el iterador ya haya terminado de iterar, devolvera un panic
	Borrar() T
}
