package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

// Se pueda crear una Pila vacía, y esta se comporta como tal.
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}

// Se puedan apilar elementos, que al desapilarlos se mantenga el invariante de pila (que esta es LIFO).
// Probar con elementos diferentes, y ver que salgan en el orden deseado.
func TestComportamientoLIFO(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(1)
	pila.Apilar(3)
	pila.Apilar(5)
	pila.Apilar(7)
	require.EqualValues(t, 7, pila.Desapilar())
	require.EqualValues(t, 5, pila.VerTope())
	require.EqualValues(t, 5, pila.Desapilar())
	require.EqualValues(t, 3, pila.VerTope())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 1, pila.VerTope())
	require.False(t, pila.EstaVacia())

}

// Prueba de volumen: Se pueden apilar muchos elementos (1000, 10000 elementos, o el volumen que corresponda):
// hacer crecer la pila, y desapilar elementos hasta que esté vacía, comprobando que siempre cumpla el invariante.
// Recordar no apilar siempre lo mismo, validar que se cumpla siempre que el tope de la pila sea el correcto paso a paso,
// y que el nuevo tope después de cada desapilar también sea el correcto.
const VOL = 10000

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	for i := 0; i <= VOL; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}

	require.False(t, pila.EstaVacia())

	for i := 0; i <= VOL; i++ {
		require.EqualValues(t, VOL-i, pila.VerTope())
		require.EqualValues(t, VOL-i, pila.Desapilar())

	}

	require.True(t, pila.EstaVacia())

}

// Condición de borde: comprobar que al desapilar hasta que está vacía hace que la pila se comporte como recién creada.
func TestBordeVacio(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(6)
	pila.Apilar(8)
	pila.Apilar(510)
	pila.Apilar(1)

	require.EqualValues(t, 1, pila.Desapilar())
	require.EqualValues(t, 510, pila.Desapilar())
	require.EqualValues(t, 8, pila.Desapilar())
	require.EqualValues(t, 6, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}

// Condición de borde: las acciones de desapilar y ver_tope en una pila recién creada son inválidas.
func TestPilaRecienCreada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	// siento que ya fue considerado este caso en un test anterior
}

// Condición de borde: la acción de esta_vacía en una pila recién creada es verdadero.
func TestPilaVacia2(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	// siento que ya fue considerado este caso en un test anterior
}

// Condición de borde: las acciones de desapilar y ver_tope en una pila a la que se le apiló y desapiló hasta estar vacía son inválidas.
func TestPilaApiladayDesapilada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(600)
	pila.Apilar(89)
	pila.Apilar(0)
	pila.Apilar(198)
	require.EqualValues(t, 198, pila.Desapilar())
	require.EqualValues(t, 0, pila.Desapilar())
	require.EqualValues(t, 89, pila.Desapilar())
	require.EqualValues(t, 600, pila.Desapilar())

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	// siento que ya fue considerado este caso en un test anterior
}

// Probar apilar diferentes tipos de datos: probar con una pila de enteros, con una pila de cadenas, etc…
func TestDistintosTiposEnteros(t *testing.T) {
	//ints
	pilaint := TDAPila.CrearPilaDinamica[int]()
	pilaint.Apilar(2)
	pilaint.Apilar(3)
	require.EqualValues(t, 3, pilaint.VerTope())
	require.EqualValues(t, 3, pilaint.Desapilar())
	require.EqualValues(t, 2, pilaint.VerTope())
	require.False(t, pilaint.EstaVacia())
	require.EqualValues(t, 2, pilaint.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaint.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaint.Desapilar() })
	require.True(t, pilaint.EstaVacia())
}

func TestDistintosTiposString(t *testing.T) {
	//strings
	pilastring := TDAPila.CrearPilaDinamica[string]()
	pilastring.Apilar("hola")
	pilastring.Apilar("como estas")
	require.EqualValues(t, "como estas", pilastring.VerTope())
	require.EqualValues(t, "como estas", pilastring.Desapilar())
	require.EqualValues(t, "hola", pilastring.VerTope())
	require.False(t, pilastring.EstaVacia())
	require.EqualValues(t, "hola", pilastring.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilastring.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilastring.Desapilar() })
	require.True(t, pilastring.EstaVacia())
}

func TestDistintosTiposFloats(t *testing.T) {
	// foats
	pilafloat := TDAPila.CrearPilaDinamica[float32]()
	pilafloat.Apilar(2.3343)
	pilafloat.Apilar(3.43)
	require.EqualValues(t, 3.43, pilafloat.VerTope())
	require.EqualValues(t, 3.43, pilafloat.Desapilar())
	require.EqualValues(t, 2.3343, pilafloat.VerTope())
	require.False(t, pilafloat.EstaVacia())
	require.EqualValues(t, 2.3343, pilafloat.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilafloat.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilafloat.Desapilar() })
	require.True(t, pilafloat.EstaVacia())
}

// consultar si hay que testear tipos de datos compuestos
func TestUnElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	pila.Apilar(3)
	require.False(t, pila.EstaVacia())
	require.EqualValues(t, 3, pila.VerTope())
	require.EqualValues(t, 3, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

}
