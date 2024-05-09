package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const GRANTAMANIO = 10000

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo(), "El largo de una lista vacia es 0")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Hay un panic al querer borrar el primer elemento de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Hay un panic al querer ver el primer elemento de una lista vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Hay un panic al querer ver el ultimo elemento de una lista vacia")
}

func TestInsertarYBorrarVariosElementosNumericos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	require.EqualValues(t, 3, lista.VerPrimero(), "En una lista vacia, los elementos numericos se insertan correctamente al principio")
	require.EqualValues(t, 3, lista.VerUltimo(), "En una lista vacia, los elementos numericos se insertan correctamente al principio")
	require.EqualValues(t, 1, lista.Largo(), "Luego de insertar un elemento, el largo aumenta")
	lista.InsertarUltimo(-1)
	require.EqualValues(t, 3, lista.VerPrimero())
	require.EqualValues(t, -1, lista.VerUltimo(), "Los elementos numericos se insertan correctamente al final")
	require.EqualValues(t, 2, lista.Largo())
	lista.InsertarPrimero(0)
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, -1, lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())

	require.False(t, lista.EstaVacia())

	lista.BorrarPrimero()
	require.EqualValues(t, 3, lista.VerPrimero(), "Los elementos numericos se borran correctamente")
	require.EqualValues(t, -1, lista.VerUltimo(), "Los elementos numericos se borran correctamente")
	require.EqualValues(t, 2, lista.Largo(), "Al borrar un elemento, el largo disminuye")
	lista.BorrarPrimero()
	require.EqualValues(t, -1, lista.VerPrimero())
	require.EqualValues(t, -1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
}

func TestInsertarYBorrarVariosElementosDeTipoString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("elem1")
	require.EqualValues(t, "elem1", lista.VerPrimero(), "En una lista vacia, los elementos de tipo string se insertan correctamente al principio")
	require.EqualValues(t, "elem1", lista.VerUltimo(), "En una lista vacia, los elementos de tipo string se insertan correctamente al principio")
	require.EqualValues(t, 1, lista.Largo(), "Luego de insertar un elemento, el largo aumenta")
	lista.InsertarUltimo("elem2")
	require.EqualValues(t, "elem1", lista.VerPrimero())
	require.EqualValues(t, "elem2", lista.VerUltimo(), "Los elementos de tipo string se insertan correctamente al final")
	require.EqualValues(t, 2, lista.Largo())
	lista.InsertarPrimero("elem3")
	require.EqualValues(t, "elem3", lista.VerPrimero())
	require.EqualValues(t, "elem2", lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())

	require.False(t, lista.EstaVacia())

	lista.BorrarPrimero()
	require.EqualValues(t, "elem1", lista.VerPrimero(), "Los elementos de tipo string se borran correctamente")
	require.EqualValues(t, "elem2", lista.VerUltimo(), "Los elementos de tipo string se borran correctamente")
	require.EqualValues(t, 2, lista.Largo(), "Al borrar un elemento, el largo disminuye")
	lista.BorrarPrimero()
	require.EqualValues(t, "elem2", lista.VerPrimero())
	require.EqualValues(t, "elem2", lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
}

func TestListaConUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(7)
	require.EqualValues(t, 7, lista.VerPrimero(), "En una lista vacia, los elementos se insertan correctamente al final")
	require.EqualValues(t, 7, lista.VerUltimo(), "En una lista vacia, los elementos se insertan correctamente al final")
	require.EqualValues(t, 1, lista.Largo())

	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia(), "Al borrar un elemento de una lista de largo 1, esta se comporta como vacia")
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}

func TestIteradorExternoIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(2)
	lista.InsertarUltimo(-1)
	lista.InsertarPrimero(5)
	lista.InsertarUltimo(7)
	require.EqualValues(t, 4, lista.Largo())
	require.False(t, lista.EstaVacia())

	i := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		switch i {
		case 0:
			require.EqualValues(t, 5, iter.VerActual(), "caso 0")
		case 1:
			require.EqualValues(t, 2, iter.VerActual(), "caso 1")
		case 2:
			require.EqualValues(t, -1, iter.VerActual(), "caso 2")
		case 3:
			require.EqualValues(t, 7, iter.VerActual(), "caso 3")
		}
		i++
	}
}

func TestIteradorExternoBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(3)
	lista.InsertarUltimo(-1)
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(9)
	require.EqualValues(t, 4, lista.Largo())
	require.False(t, lista.EstaVacia())

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Borrar()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIteradorExternoInsertarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	i := 0
	iter := lista.Iterador()
	iter.Insertar(5)
	iter.Insertar(4)
	iter.Insertar(3)
	iter.Insertar(2)
	iter.Insertar(1)
	iter.Insertar(0)

	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		switch i {
		case 0:
			require.EqualValues(t, 0, iter.VerActual())
		case 1:
			require.EqualValues(t, 1, iter.VerActual())
		case 2:
			require.EqualValues(t, 2, iter.VerActual())
		case 3:
			require.EqualValues(t, 3, iter.VerActual())
		case 4:
			require.EqualValues(t, 4, iter.VerActual())
		case 5:
			require.EqualValues(t, 5, iter.VerActual())
		}
		i++
	}
}

func TestIteradorExternoListaDeUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	iter := lista.Iterador()
	iter.Insertar(5)

	require.EqualValues(t, 5, iter.VerActual())
	require.EqualValues(t, 5, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.False(t, lista.EstaVacia())
	elem_borrado := iter.Borrar()
	require.EqualValues(t, 5, elem_borrado)
	require.True(t, lista.EstaVacia())
}

func TestIteradorExternoVolumenInsertar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	iter := lista.Iterador()

	for i := 0; i < GRANTAMANIO; i++ {
		iter.Insertar(i)
	}

	for i := GRANTAMANIO - 1; i > 0; i-- {
		require.EqualValues(t, i, iter.VerActual())
		iter.Siguiente()
	}
}

func TestIteradorExternoVolumenBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < GRANTAMANIO; i++ {
		lista.InsertarPrimero(i)
	}

	i := GRANTAMANIO - 1
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		require.EqualValues(t, i, iter.VerActual())
		iter.Borrar()
		i--
	}
}

func TestIteradorInternoConCondicionDeCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	lista.InsertarUltimo(-1)
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(9)
	require.False(t, lista.EstaVacia())

	var suma int
	lista.Iterar(func(dato int) bool {
		if dato == -1 {
			return false
		}
		suma = suma + dato
		return true
	})
	require.EqualValues(t, 3, suma, "El iterador interno funciona")
}

func TestIteradorInternoListaCompleta(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(3)
	lista.InsertarUltimo(-1)
	lista.InsertarPrimero(0)
	lista.InsertarUltimo(9)
	require.False(t, lista.EstaVacia())

	var suma int
	lista.Iterar(func(dato int) bool {
		suma = suma + dato
		return true
	})
	require.EqualValues(t, 11, suma, "Se puede iterar toda la lista correctamente con el iterador interno")
}

func TestIteradorInternoVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := GRANTAMANIO; i > 0; i-- {
		lista.InsertarPrimero(i)
	}

	var elem int
	i := 1
	lista.Iterar(func(dato int) bool {
		elem = dato
		require.EqualValues(t, i, dato)
		i++
		return true
	})
	require.EqualValues(t, GRANTAMANIO, elem)
}

func TestVolumenPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < GRANTAMANIO; i++ {
		lista.InsertarPrimero(i)
	}
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 9999, lista.VerPrimero())
	require.EqualValues(t, 0, lista.VerUltimo())
	require.EqualValues(t, GRANTAMANIO, lista.Largo())
	for i := GRANTAMANIO - 1; i >= 0; i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 0, lista.VerUltimo())
		require.EqualValues(t, i+1, lista.Largo())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}
func TestVolumenUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < GRANTAMANIO; i++ {
		lista.InsertarUltimo(i)
	}
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 9999, lista.VerUltimo())
	require.EqualValues(t, GRANTAMANIO, lista.Largo())
	for i := 0; i < GRANTAMANIO; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, 9999, lista.VerUltimo())
		require.EqualValues(t, GRANTAMANIO-i, lista.Largo())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
}
