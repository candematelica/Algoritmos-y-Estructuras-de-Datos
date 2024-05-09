package cola_prioridad

type heapImplementacion[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

const TAM_INICIAL = 20
const VALOR_DE_REDIMENSION = 2

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, TAM_INICIAL)
	return &heapImplementacion[T]{datos: datos, cant: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	copia := make([]T, len(arreglo)+TAM_INICIAL)
	copy(copia, arreglo)
	heapify[T](copia, funcion_cmp)
	return &heapImplementacion[T]{datos: copia, cant: len(arreglo), cmp: funcion_cmp}
}

func redimensionarHeap[T any](heap *heapImplementacion[T]) {
	if len(heap.datos) == heap.cant {
		vectorAuxiliar(heap, len(heap.datos)*VALOR_DE_REDIMENSION)
	} else if heap.cant*4 <= len(heap.datos) && (len(heap.datos)/VALOR_DE_REDIMENSION > TAM_INICIAL) {
		vectorAuxiliar(heap, len(heap.datos)/VALOR_DE_REDIMENSION)
	}
}

func vectorAuxiliar[T any](heap *heapImplementacion[T], tam int) {
	aux := make([]T, tam)
	copy(aux, heap.datos)
	heap.datos = aux
}

func (heap *heapImplementacion[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *heapImplementacion[T]) Encolar(dato T) {
	redimensionarHeap[T](heap)
	heap.datos[heap.cant] = dato

	upheap[T](heap.datos, heap.cmp, heap.cant)
	heap.cant++
}

func (heap *heapImplementacion[T]) VerMax() T {
	heap.chequearCant()
	return heap.datos[0]
}

func (heap *heapImplementacion[T]) Desencolar() T {
	heap.chequearCant()
	redimensionarHeap[T](heap)
	dato := heap.datos[0]
	heap.datos[0], heap.datos[heap.cant-1] = heap.datos[heap.cant-1], heap.datos[0]
	downheap[T](heap.datos, heap.cmp, heap.cant-1, 0)
	heap.cant--
	return dato
}

func (heap *heapImplementacion[T]) Cantidad() int {
	return heap.cant
}

func (heap *heapImplementacion[T]) chequearCant() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func upheap[T any](arr []T, cmp func(T, T) int, index int) {
	if index == 0 {
		return
	}

	pos_padre := (index - 1) / 2

	if cmp(arr[index], arr[pos_padre]) > 0 {
		arr[index], arr[pos_padre] = arr[pos_padre], arr[index]
		upheap[T](arr, cmp, pos_padre)
	}
}

func downheap[T any](arr []T, cmp func(T, T) int, cant, index int) {
	if index >= cant {
		return
	}

	pos_hij_izq := 2*index + 1
	pos_hij_der := 2*index + 2
	mayor := index

	if pos_hij_izq < cant && cmp(arr[pos_hij_izq], arr[mayor]) > 0 {
		mayor = pos_hij_izq
	}
	if pos_hij_der < cant && cmp(arr[pos_hij_der], arr[mayor]) > 0 {
		mayor = pos_hij_der
	}
	if mayor != index {
		arr[index], arr[mayor] = arr[mayor], arr[index]
		downheap[T](arr, cmp, cant, mayor)
	}
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	if len(arr) == 0 {
		return
	}
	for i := len(arr) - 1; i >= 0; i-- {
		downheap[T](arr, cmp, len(arr), i)
	}
}

func HeapSort[T any](arr []T, cmp func(T, T) int) {
	heapify(arr, cmp)
	for j := len(arr) - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		downheap[T](arr, cmp, j, 0)
	}
}
