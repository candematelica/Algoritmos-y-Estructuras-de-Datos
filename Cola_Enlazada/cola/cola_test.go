package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

// Se pueda crear una cola vacía, y esta se comporta como tal.
func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

// Se puedan Encolar elementos, que al Desencolarlos se mantenga el invariante de cola (que esta es LIFO).
// Probar con elementos diferentes, y ver que salgan en el orden deseado.
func TestComportamientoFIFO(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(7)
	cola.Encolar(5)
	cola.Encolar(3)
	cola.Encolar(1)
	require.EqualValues(t, 7, cola.Desencolar())
	require.EqualValues(t, 5, cola.VerPrimero())
	require.EqualValues(t, 5, cola.Desencolar())
	require.EqualValues(t, 3, cola.VerPrimero())
	require.EqualValues(t, 3, cola.Desencolar())
	require.EqualValues(t, 1, cola.VerPrimero())
	require.False(t, cola.EstaVacia())

}

// Prueba de volumen: Se pueden Encolar muchos elementos (1000, 10000 elementos, o el volumen que corresponda):
// hacer crecer la cola, y Desencolar elementos hasta que esté vacía, comprobando que siempre cumpla el invariante.
// Recordar no Encolar siempre lo mismo, validar que se cumpla siempre que el PrimeroVerPrimero de la cola sea el correcto paso a paso,
// y que el nuevo PrimeroVerPrimero después de cada Desencolar también sea el correcto.
const VOL = 10000

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	for i := 0; i <= VOL; i++ {
		cola.Encolar(i)
	}

	require.False(t, cola.EstaVacia())

	for i := 0; i <= VOL; i++ {

		require.EqualValues(t, i, cola.Desencolar())

	}

	require.True(t, cola.EstaVacia())

}

// Condición de borde: comprobar que al Desencolar hasta que está vacía hace que la cola se comporte como recién creada.
func TestBordeVacio(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(1)
	cola.Encolar(510)
	cola.Encolar(8)
	cola.Encolar(6)

	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 510, cola.Desencolar())
	require.EqualValues(t, 8, cola.Desencolar())
	require.EqualValues(t, 6, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}

// Condición de borde: las acciones de Desencolar y ver_PrimeroVerPrimero en una cola recién creada son inválidas.
func TestcolaRecienCreada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	// siento que ya fue considerado este caso en un test anterior
}

// Condición de borde: la acción de esta_vacía en una cola recién creada es verdadero.
func TestcolaVacia2(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	// siento que ya fue considerado este caso en un test anterior
}

// Condición de borde: las acciones de Desencolar y ver_PrimeroVerPrimero en una cola a la que se le apiló y desapiló hasta estar vacía son inválidas.
func BordeDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(600)
	cola.Encolar(89)
	cola.Encolar(0)
	cola.Encolar(198)
	require.EqualValues(t, 198, cola.Desencolar())
	require.EqualValues(t, 0, cola.Desencolar())
	require.EqualValues(t, 89, cola.Desencolar())
	require.EqualValues(t, 600, cola.Desencolar())

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	// siento que ya fue considerado este caso en un test anterior
}

// Probar Encolar diferentes tipos de datos: probar con una cola de enteros, con una cola de cadenas, etc…
func TestDistintosTiposEnteroa(t *testing.T) {
	//ints
	colaint := TDACola.CrearColaEnlazada[int]()

	colaint.Encolar(3)
	colaint.Encolar(2)

	require.EqualValues(t, 3, colaint.VerPrimero())
	require.EqualValues(t, 3, colaint.Desencolar())
	require.EqualValues(t, 2, colaint.VerPrimero())
	require.False(t, colaint.EstaVacia())
	require.EqualValues(t, 2, colaint.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaint.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaint.Desencolar() })
	require.True(t, colaint.EstaVacia())
}

func TestDistintosTiposString(t *testing.T) {
	//strings
	colastring := TDACola.CrearColaEnlazada[string]()
	colastring.Encolar("como estas")
	colastring.Encolar("hola")

	require.EqualValues(t, "como estas", colastring.VerPrimero())
	require.EqualValues(t, "como estas", colastring.Desencolar())
	require.EqualValues(t, "hola", colastring.VerPrimero())
	require.False(t, colastring.EstaVacia())
	require.EqualValues(t, "hola", colastring.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colastring.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colastring.Desencolar() })
	require.True(t, colastring.EstaVacia())
}

func TestDistintosTiposFloats(t *testing.T) {
	// foats
	colafloat := TDACola.CrearColaEnlazada[float32]()
	colafloat.Encolar(3.43)
	colafloat.Encolar(2.3343)

	require.EqualValues(t, 3.43, colafloat.VerPrimero())
	require.EqualValues(t, 3.43, colafloat.Desencolar())
	require.EqualValues(t, 2.3343, colafloat.VerPrimero())
	require.False(t, colafloat.EstaVacia())
	require.EqualValues(t, 2.3343, colafloat.Desencolar())
	require.PanicsWithValue(t, "La cola esta vacia", func() { colafloat.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { colafloat.Desencolar() })
	require.True(t, colafloat.EstaVacia())
}

// consultar si hay que testear tipos de datos compuestos
func TestUnElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	cola.Encolar(3)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 3, cola.VerPrimero())
	require.EqualValues(t, 3, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

}
