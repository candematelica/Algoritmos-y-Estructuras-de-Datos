package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ABB = []int{12500, 25000, 50000, 100000, 200000, 400000}

func cmpEnteros(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func TestABBVacio(t *testing.T) {
	t.Log("Comprueba que ABB vacio no tiene claves")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestABBUnElemento(t *testing.T) {
	t.Log("Comprueba que ABB con un elemento tiene esa clave, unicamente")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("B") })
}

func TestABBGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))

	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestReemplazoDatoABB(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(clave, "miau")
	abb.Guardar(clave2, "guau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "miau", abb.Obtener(clave))
	require.EqualValues(t, "guau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}

func TestReemplazoVariosDatos(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean correctos")

	abb := TDADiccionario.CrearABB[int, int](cmpEnteros)
	for i := 0; i < 500; i++ {
		abb.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		abb.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = abb.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestABBBorrarUnElemento(t *testing.T) {
	t.Log("Comprueba que al borrar un ABB con un elemento, este quede vacio")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	abb.Borrar("A")
	require.EqualValues(t, 0, abb.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestABBBorrarSinHijos(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se borran los que no tengan hijos, revisando que en todo momento " +
		"el ABB se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)

	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[1]) })

	require.True(t, abb.Pertenece(claves[0]))
}

func TestABBBorrarConUnHijo(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se borran los que tengan un hijo, revisando que en todo momento " +
		"el ABB se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)

	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[0]) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[2]))
}

func TestABBBorrarConDosHijos(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el ABB, y se borra el que tenga dos hijos, revisando que en todo momento " +
		"el ABB se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Oveja"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	valor4 := "mee"
	claves := []string{clave1, clave2, clave3, clave4}
	valores := []string{valor1, valor2, valor3, valor4}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)

	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	abb.Guardar(claves[3], valores[3])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[2]) })

	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[3]))
}

func TestABBGuardarBorrado(t *testing.T) {
	t.Log("Borra un elemento del ABB y vuelve a insertar el mismo elemento, comprobando que se inserte " +
		"correctamente")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)

	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.EqualValues(t, 3, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
}

func TestABBConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	abb := TDADiccionario.CrearABB[int, string](cmpEnteros)
	clave := 10
	valor := "Gatito"

	abb.Guardar(clave, valor)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
	require.EqualValues(t, valor, abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}

func TestABBValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	abb := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBGuardarYBorrarRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces.")

	abb := TDADiccionario.CrearABB[int, int](cmpEnteros)
	for i := 0; i < 1000; i++ {
		abb.Guardar(i, i)
		require.True(t, abb.Pertenece(i))
		abb.Borrar(i)
		require.False(t, abb.Pertenece(i))
	}
}

func TestABBIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(claves[0], "")
	abb.Guardar(claves[1], "")
	abb.Guardar(claves[2], "")

	var suma int
	abb.Iterar(func(clave string, dato string) bool {
		if strings.Compare(clave, claves[1]) == 0 {
			return false
		}
		suma++
		return true
	})

	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, 1, suma)
}

func TestABBIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestABBIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	abb.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	abb := TDADiccionario.CrearABB[int, int](cmpEnteros)

	claves := make([]int, n)
	valores := make([]int, n)
	elementos := make([]int, n)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		elementos[i] = i
	}
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		elementos[i], elementos[j] = elementos[j], elementos[i]
	}
	for i := 0; i < n; i++ {
		valores[i] = elementos[i]
		claves[i] = elementos[i]
		abb.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	for i := 0; i < n; i++ {
		require.True(b, abb.Pertenece(claves[i]))
		require.EqualValues(b, valores[i], abb.Obtener(claves[i]))
	}

	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		borrado := abb.Borrar(claves[i])
		require.EqualValues(b, valores[i], borrado)
		require.False(b, abb.Pertenece(claves[i]), "Borrar muchos elementos no funciona correctamente")
	}

	require.EqualValues(b, 0, abb.Cantidad(), "Borrar muchos elementos no funciona correctamente")
}

func BenchmarkABB(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}

func TestIterarABBVacio(t *testing.T) {
	t.Log("Iterar sobre ABB vacio es simplemente tenerlo al final")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un ABB, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al ABB. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	iter := abb.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, clave1, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, clave2, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, clave3, tercero)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	abb.Guardar(claves[0], "")
	abb.Guardar(claves[1], "")
	abb.Guardar(claves[2], "")

	abb.Iterador()
	iter2 := abb.Iterador()
	iter2.Siguiente()
	segundoIter2, _ := iter2.VerActual()
	iter3 := abb.Iterador()
	primeroIter1, _ := iter3.VerActual()
	iter3.Siguiente()
	segundoIter1, _ := iter3.VerActual()
	iter3.Siguiente()
	terceroIter1, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primeroIter1, segundoIter1)
	require.NotEqualValues(t, terceroIter1, segundoIter1)
	require.NotEqualValues(t, primeroIter1, terceroIter1)
	require.EqualValues(t, segundoIter1, segundoIter2)
}

func TestVolumenIteradorInterno(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	abb := TDADiccionario.CrearABB[int, int](cmpEnteros)
	claves := make([]int, 10000)
	valores := make([]int, 10000)
	elementos := make([]int, 10000)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < 10000; i++ {
		elementos[i] = i
	}
	for i := 0; i < 10000; i++ {
		j := rand.Intn(10000)
		elementos[i], elementos[j] = elementos[j], elementos[i]
	}
	for i := 0; i < 10000; i++ {
		valores[i] = elementos[i]
		claves[i] = elementos[i]
		abb.Guardar(claves[i], valores[i])
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	abb.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestABBIteradorInternoRangosSinRangos(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno " +
		"por rangos si no se especifican los rangos")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(clave1, "miau")
	abb.Guardar(clave2, "guau")
	abb.Guardar(clave3, "moo")
	require.EqualValues(t, 3, abb.Cantidad())

	var i int
	abb.IterarRango(nil, nil, func(clave string, dato string) bool {
		switch i {
		case 0:
			require.EqualValues(t, clave, clave1)
			require.EqualValues(t, "miau", abb.Obtener(clave1))
		case 1:
			require.EqualValues(t, clave, clave2)
			require.EqualValues(t, "guau", abb.Obtener(clave2))
		case 2:
			require.EqualValues(t, clave, clave3)
			require.EqualValues(t, "moo", abb.Obtener(clave3))
		}
		i++
		return true
	})
	require.EqualValues(t, i, 3)
}

func TestIterarRangosABBSinDesde(t *testing.T) {
	t.Log("Valida que las claves sean recorridas desde el comienzo hasta un rango final (y una única vez) " +
		"con el iterador interno por rangos")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(clave1, "miau")
	abb.Guardar(clave2, "guau")
	abb.Guardar(clave3, "moo")
	require.EqualValues(t, 3, abb.Cantidad())

	var i int
	abb.IterarRango(nil, &clave2, func(clave string, dato string) bool {
		switch i {
		case 0:
			require.EqualValues(t, clave, clave1)
			require.EqualValues(t, "miau", abb.Obtener(clave1))
		case 1:
			require.EqualValues(t, clave, clave2)
			require.EqualValues(t, "guau", abb.Obtener(clave2))
		}
		i++
		return true
	})
	require.EqualValues(t, i, 2)
}

func TestIterarRangosABBSinHasta(t *testing.T) {
	t.Log("Valida que las claves sean recorridas desde un cierto valor hasta el final (y una única vez) " +
		"con el iterador interno por rangos")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(clave1, "miau")
	abb.Guardar(clave2, "guau")
	abb.Guardar(clave3, "moo")
	require.EqualValues(t, 3, abb.Cantidad())

	var i int
	abb.IterarRango(&clave2, nil, func(clave string, dato string) bool {
		switch i {
		case 0:
			require.EqualValues(t, clave2, clave)
			require.EqualValues(t, "guau", abb.Obtener(clave2))
		case 1:
			require.EqualValues(t, clave3, clave)
			require.EqualValues(t, "moo", abb.Obtener(clave3))
		}
		i++
		return true
	})
	require.EqualValues(t, i, 2)
}

func TestABBIteradorInternoRangosConRangos(t *testing.T) {
	t.Log("Valida que todas las claves pertenecientes a un rango sean recorridas (y una única vez) " +
		"con el iterador interno por rangos")
	clave1 := "Gato"
	clave2 := "Oveja"
	clave3 := "Perro"
	clave4 := "Vaca"
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(clave1, "miau")
	abb.Guardar(clave2, "mee")
	abb.Guardar(clave3, "guau")
	abb.Guardar(clave4, "moo")
	require.EqualValues(t, 4, abb.Cantidad())

	var i int
	abb.IterarRango(&clave2, &clave3, func(clave string, dato string) bool {
		switch i {
		case 0:
			require.EqualValues(t, clave, clave2)
			require.EqualValues(t, "mee", abb.Obtener(clave2))
		case 1:
			require.EqualValues(t, clave, clave3)
			require.EqualValues(t, "guau", abb.Obtener(clave3))
		}
		i++
		return true
	})
	require.EqualValues(t, i, 2)
}

func TestABBIteradorInternoRangosConCondicionDeCorte(t *testing.T) {
	t.Log("Valida que las claves pertenecientes a un rango sean recorridas (y una única vez) " +
		"hasta la condicion de corte con el iterador interno por rangos")
	clave1 := "Comadreja"
	clave2 := "Gato"
	clave3 := "Ave"
	clave4 := "Oveja"
	clave5 := "Perro"
	clave6 := "Vaca"
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar(clave1, 0)
	abb.Guardar(clave2, 1)
	abb.Guardar(clave3, 2)
	abb.Guardar(clave4, 3)
	abb.Guardar(clave5, 4)
	abb.Guardar(clave6, 5)
	require.EqualValues(t, 6, abb.Cantidad())

	var i int
	abb.IterarRango(&clave2, &clave5, func(clave string, dato int) bool {
		if dato == 3 {
			return false
		}
		i++
		return true
	})
	require.EqualValues(t, i, 1)
}

func TestABBIteradorInternoRangosValores(t *testing.T) {
	t.Log("Valida que los datos dentro del rango sean recorridos correctamente (y una única vez) " +
		"con el iterador interno por rangos")
	clave1 := "Gato"
	clave2 := "Oveja"
	clave3 := "Perro"
	clave4 := "Vaca"

	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)

	var suma int
	abb.IterarRango(&clave2, &clave3, func(clave string, dato int) bool {
		suma += dato
		return true
	})

	require.EqualValues(t, 5, suma)
}

func TestIterarRangosABBVacio(t *testing.T) {
	t.Log("Iterar por rangos sobre ABB vacio es simplemente tenerlo al final")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorRangosSinRangos(t *testing.T) {
	t.Log("Guardamos 4 valores en un ABB, e iteramos por rangos (sin pasar rangos) validando que las claves sean todas diferentes " +
		"pero pertenecientes al ABB. El iterador se deberìa comportar como un iterador externo sin rangos. Además los valores de " +
		"VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Perro"
	clave2 := "Vaca"
	clave3 := "Oveja"
	clave4 := "Gato"
	valor1 := "guau"
	valor2 := "moo"
	valor3 := "mee"
	valor4 := "miau"
	claves := []string{clave1, clave2, clave3, clave4}
	valores := []string{valor1, valor2, valor3, valor4}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	abb.Guardar(claves[3], valores[3])
	iter := abb.IteradorRango(nil, nil)

	primero, _ := iter.VerActual()
	require.EqualValues(t, clave4, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, clave3, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	tercero, _ := iter.VerActual()
	require.EqualValues(t, clave1, tercero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, clave2, cuarto)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorRangosSinDesde(t *testing.T) {
	t.Log("Guardamos 4 valores en un ABB, e iteramos por rangos (sin desde) validando que las claves sean todas diferentes " +
		"pero pertenecientes al ABB. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Perro"
	clave2 := "Vaca"
	clave3 := "Oveja"
	clave4 := "Gato"
	valor1 := "guau"
	valor2 := "moo"
	valor3 := "mee"
	valor4 := "miau"
	claves := []string{clave1, clave2, clave3, clave4}
	valores := []string{valor1, valor2, valor3, valor4}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	abb.Guardar(claves[3], valores[3])
	iter := abb.IteradorRango(nil, &clave3)

	primero, _ := iter.VerActual()
	require.EqualValues(t, clave4, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, clave3, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorRangosSinHasta(t *testing.T) {
	t.Log("Guardamos 4 valores en un ABB, e iteramos por rangos (sin hasta) validando que las claves sean todas diferentes " +
		"pero pertenecientes al ABB. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Perro"
	clave2 := "Vaca"
	clave3 := "Oveja"
	clave4 := "Gato"
	valor1 := "guau"
	valor2 := "moo"
	valor3 := "mee"
	valor4 := "miau"
	claves := []string{clave1, clave2, clave3, clave4}
	valores := []string{valor1, valor2, valor3, valor4}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	abb.Guardar(claves[3], valores[3])
	iter := abb.IteradorRango(&clave1, nil)

	primero, _ := iter.VerActual()
	require.EqualValues(t, clave1, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, clave2, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorRangosConRangos(t *testing.T) {
	t.Log("Guardamos 4 valores en un ABB, e iteramos por rangos validando que las claves sean todas diferentes " +
		"pero pertenecientes al ABB. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Perro"
	clave2 := "Vaca"
	clave3 := "Oveja"
	clave4 := "Gato"
	valor1 := "guau"
	valor2 := "moo"
	valor3 := "mee"
	valor4 := "miau"
	claves := []string{clave1, clave2, clave3, clave4}
	valores := []string{valor1, valor2, valor3, valor4}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	abb.Guardar(claves[3], valores[3])
	iter := abb.IteradorRango(&clave3, &clave1)

	primero, _ := iter.VerActual()
	require.EqualValues(t, clave3, primero)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	segundo, _ := iter.VerActual()
	require.EqualValues(t, clave1, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func ejecutarPruebasVolumenIteradorABB(b *testing.B, n int) {
	abb := TDADiccionario.CrearABB[int, int](cmpEnteros)

	claves := make([]int, n)
	valores := make([]int, n)
	elementos := make([]int, n)

	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		elementos[i] = i
	}
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		elementos[i], elementos[j] = elementos[j], elementos[i]
	}
	for i := 0; i < n; i++ {
		valores[i] = elementos[i]
		claves[i] = elementos[i]
		abb.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	iter := abb.Iterador()

	var i int
	for i < n {
		require.True(b, iter.HaySiguiente(), "Iteracion en volumen no funciona correctamente")
		clave, _ := iter.VerActual()
		require.EqualValues(b, clave, i)
		iter.Siguiente()
		i++
	}

	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")
}

func BenchmarkIteradorABB(b *testing.B) {
	b.Log("Prueba de stress del Iterador del ABB. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos por rangos todos sin problemas. Se ejecuta cada prueba b.N" +
		" veces para generar un benchmark")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorABB(b, n)
			}
		})
	}
}
