package cola_prioridad_test

import (
	"math/rand"
	"strings"
	TDAColaDePrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const GRAN_TAMANIO = 10000

func cmpEnteros(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func TestColaDePrioridadVacia(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene datos.")
	heap := TDAColaDePrioridad.CrearHeap[int](cmpEnteros)

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolar(t *testing.T) {
	t.Log("Prueba que encolar un elemento al heap funcione correctamente.")
	heap := TDAColaDePrioridad.CrearHeap[int](cmpEnteros)

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(2)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 2, heap.VerMax())
}

func TestDesencolar(t *testing.T) {
	t.Log("Prueba que desencolar un elemento de un heap con un unico elemento funcione " +
		"correctamente y que se comporte como vacio.")
	heap := TDAColaDePrioridad.CrearHeap[int](cmpEnteros)

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(2)
	heap.Desencolar()

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarYDesencolarVariosElementosNumericos(t *testing.T) {
	t.Log("Prueba que encolar y desencolar varios elementos numericos funcione correctamente " +
		"y que el heap resultante cumpla las propiedades de heap.")
	heap := TDAColaDePrioridad.CrearHeap[int](cmpEnteros)

	heap.Encolar(2)
	heap.Encolar(7)
	heap.Encolar(-1)
	heap.Encolar(11)
	require.EqualValues(t, 4, heap.Cantidad())

	require.EqualValues(t, 11, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 7, heap.VerMax(), "Desencolar cambia el maximo")

	heap.Encolar(5)
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 7, heap.VerMax(), "Encolar un elemento menor al maximo no cambia "+
		"el maximo")

	heap.Desencolar()
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 5, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, 2, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, -1, heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 0, heap.Cantidad())

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarYDesencolarVariosElementosStrings(t *testing.T) {
	t.Log("Prueba que encolar y desencolar varios elementos de tipo string funcione " +
		"correctamente y que el heap resultante cumpla las propiedades de heap.")
	heap := TDAColaDePrioridad.CrearHeap[string](strings.Compare)

	heap.Encolar("Salame")
	heap.Encolar("Queso")
	heap.Encolar("Jamon")
	require.EqualValues(t, 3, heap.Cantidad())

	require.EqualValues(t, "Salame", heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, "Queso", heap.VerMax(), "Desencolar cambia el maximo")

	heap.Encolar("Coca")
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, "Queso", heap.VerMax(), "Encolar un elemento menor al maximo no cambia "+
		"el maximo")

	heap.Desencolar()
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, "Jamon", heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, "Coca", heap.VerMax())
	heap.Desencolar()
	require.EqualValues(t, 0, heap.Cantidad())

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestVolumen(t *testing.T) {
	t.Log("Prueba que encolar y desencolar muchos elementos funcione correctamente " +
		"y que el heap resultante cumpla las propiedades de heap.")
	heap := TDAColaDePrioridad.CrearHeap[int](cmpEnteros)

	for i := 0; i < GRAN_TAMANIO; i++ {
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	require.EqualValues(t, GRAN_TAMANIO, heap.Cantidad())

	for i := GRAN_TAMANIO - 1; i >= 0; i-- {
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i+1, heap.Cantidad())
		heap.Desencolar()
	}

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCrearHeapAPartirDeArregloNoModificaElArreglo(t *testing.T) {
	t.Log("Prueba que crear un heap a partir de un arreglo no modifica el arreglo original")
	arr := make([]int, 5)
	arr[0] = 6
	arr[1] = -2
	arr[2] = 3
	arr[3] = -11
	arr[4] = 17

	heap := TDAColaDePrioridad.CrearHeapArr[int](arr, cmpEnteros)

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 6, arr[0])
	require.EqualValues(t, -2, arr[1])
	require.EqualValues(t, 3, arr[2])
	require.EqualValues(t, -11, arr[3])
	require.EqualValues(t, 17, arr[4])
}

func TestCrearHeapAPartirDeArreglo(t *testing.T) {
	t.Log("Prueba crear un heap a partir de un arreglo que no cumpla las propiedades del " +
		"heap. Al crear el heap, este tiene que cumplir las propiedades.")
	arr := []int{2, 6, 9, 12, 0, -2, -8, 3, 17, 21, 4, -5, 22, 19}

	heap := TDAColaDePrioridad.CrearHeapArr[int](arr, cmpEnteros)

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 14, heap.Cantidad())
	require.EqualValues(t, 22, heap.VerMax())

	heap.Encolar(-1)
	require.EqualValues(t, 15, heap.Cantidad())
	require.EqualValues(t, 22, heap.VerMax())

	heap.Desencolar()
	heap.Desencolar()

	require.EqualValues(t, 13, heap.Cantidad())
	require.EqualValues(t, 19, heap.VerMax())

	heap.Encolar(22)
	require.EqualValues(t, 14, heap.Cantidad())
	require.EqualValues(t, 22, heap.VerMax())

	heap.Encolar(-1)
	require.EqualValues(t, 15, heap.Cantidad())
	require.EqualValues(t, 22, heap.VerMax())

	for !heap.EstaVacia() {
		heap.Desencolar()
	}

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCrearHeapAPartirDeArregloDeStrings(t *testing.T) {
	t.Log("Prueba crear un heap a partir de un arreglo de strings que no cumpla las propiedades del " +
		"heap. Al crear el heap, este tiene que cumplir las propiedades.")
	arr := []string{"Lechuga", "Tomate", "Zanahoria", "Palta", "Manzana"}

	heap := TDAColaDePrioridad.CrearHeapArr[string](arr, strings.Compare)

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.Cantidad())
	require.EqualValues(t, "Zanahoria", heap.VerMax())

	heap.Encolar("Aceitunas")
	require.EqualValues(t, 6, heap.Cantidad())
	require.EqualValues(t, "Zanahoria", heap.VerMax())

	heap.Desencolar()
	heap.Desencolar()

	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, "Palta", heap.VerMax())

	for !heap.EstaVacia() {
		heap.Desencolar()
	}

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestVolumenCrearHeapAPartirDeArreglo(t *testing.T) {
	t.Log("Prueba de volumen de crear un heap a partir de un arreglo que no cumpla las propiedades del " +
		"heap. Al crear el heap, este tiene que cumplir las propiedades.")
	arr := make([]int, GRAN_TAMANIO)

	for i := 0; i < GRAN_TAMANIO; i++ {
		arr[i] = i
	}
	for i := 0; i < GRAN_TAMANIO; i++ {
		j := rand.Intn(GRAN_TAMANIO)
		arr[i], arr[j] = arr[j], arr[i]
	}

	heap := TDAColaDePrioridad.CrearHeapArr[int](arr, cmpEnteros)

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, GRAN_TAMANIO, heap.Cantidad())
	require.EqualValues(t, 9999, heap.VerMax())

	heap.Encolar(10000)
	require.EqualValues(t, GRAN_TAMANIO+1, heap.Cantidad())
	require.EqualValues(t, 10000, heap.VerMax())

	heap.Desencolar()

	for i := 0; i < GRAN_TAMANIO/2; i++ {
		heap.Desencolar()
	}

	for i := GRAN_TAMANIO; i > GRAN_TAMANIO/2; i-- {
		heap.Encolar(i - 1)
	}

	require.EqualValues(t, 9999, heap.VerMax())
	require.EqualValues(t, GRAN_TAMANIO, heap.Cantidad())

	i := GRAN_TAMANIO - 1
	for i >= 0 {
		require.EqualValues(t, i, heap.VerMax())
		heap.Desencolar()
		i--
	}

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSortElementosNumericos(t *testing.T) {
	t.Log("Prueba que el metodo de ordenamiento de HeapSort funcione correctamente con " +
		"elementos numericos")
	arr := make([]int, 4)
	arr[0] = 4
	arr[1] = -1
	arr[2] = 8
	arr[3] = 0

	TDAColaDePrioridad.HeapSort[int](arr, cmpEnteros)
	require.EqualValues(t, -1, arr[0])
	require.EqualValues(t, 0, arr[1])
	require.EqualValues(t, 4, arr[2])
	require.EqualValues(t, 8, arr[3])
}

func TestHeapSortElementosStrings(t *testing.T) {
	t.Log("Prueba que el metodo de ordenamiento de HeapSort funcione correctamente con " +
		"elementos de tipo string")
	arr := make([]string, 4)
	arr[0] = "Empanadas"
	arr[1] = "Tortilla"
	arr[2] = "Guiso"
	arr[3] = "Sopa"

	TDAColaDePrioridad.HeapSort[string](arr, strings.Compare)
	require.EqualValues(t, "Empanadas", arr[0])
	require.EqualValues(t, "Guiso", arr[1])
	require.EqualValues(t, "Sopa", arr[2])
	require.EqualValues(t, "Tortilla", arr[3])
}
