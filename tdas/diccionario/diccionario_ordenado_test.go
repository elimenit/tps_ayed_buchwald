package diccionario_test

import (
	TDADict "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

// ABB
const (
	_VOLUMEN = 1000
)

// Funcion para comparar dos claves
func cmpInt(a int, b int) int {
	return a - b
}

// Crear un ABB no deberia ser un problema
func TestABBVacio(t *testing.T) {

	a := TDADict.CrearABB[int, int](cmpInt)
	require.NotNil(t, a)

	require.Equal(t, a.Cantidad(), 0)
}

// ABB guarda los elementos
func TestABBGuardarUnElemento(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	a.Guardar(1, 10)

	require.Equal(t, a.Cantidad(), 1)
	require.True(t, a.Pertenece(1))
	require.Equal(t, a.Obtener(1), 10)
}

// EL ABB debe guardar muchos elementos
func TestABBGuardarVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}
	require.Equal(t, a.Cantidad(), _VOLUMEN)

	// First Key = 0; Last Key = _VOLUMEN -1
	require.True(t, a.Pertenece(0))
	require.True(t, a.Pertenece(_VOLUMEN-1))
	require.Equal(t, a.Obtener(0), 1)
	require.Equal(t, a.Obtener(_VOLUMEN-1), _VOLUMEN)
}

// Un elemento pertenece al ABB
func TestABBPertenece(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	a.Guardar(0, 10)
	require.True(t, a.Pertenece(0))

	a.Guardar(1, 11)
	require.True(t, a.Pertenece(1))
}

// Agregar muchos datos si pertenecen al ABB
func TestABBPerteneceVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for e := range _VOLUMEN {
		a.Guardar(e, e+1)
		require.True(t, a.Pertenece(e))
	}
	require.True(t, a.Pertenece(0))
}

// Obtiene un elemento existente
func TestABBObtener(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	// Basic
	a.Guardar(1, 10)
	require.Equal(t, a.Obtener(1), 10)

	a.Guardar(2, 15)
	require.Equal(t, a.Obtener(2), 15)

	// Equal dato and not equal keys
	a.Guardar(3, 10)
	a.Guardar(0, 15)
	require.Equal(t, a.Obtener(3), 10)
	require.Equal(t, a.Obtener(0), 15)
}

// Agrega muchos elementos
func TestABBObtenerVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
		require.Equal(t, a.Obtener(i), i+1)
	}

	// First & End is in ABB
	require.Equal(t, a.Obtener(0), 1)
	require.Equal(t, a.Obtener(_VOLUMEN-1), _VOLUMEN)
}

// Si no existe la clave debe paniquear
func TestABBObtenerPanic(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	require.Panics(t, func() { a.Obtener(1) })
	require.Panics(t, func() { a.Obtener(_VOLUMEN) })

	a.Guardar(10, 100)
	require.NotPanics(t, func() { a.Obtener(10) })
}

func TestABBObtenerPanicVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
		require.Panics(t, func() { a.Obtener(i + 1) })
	}

	for i := range _VOLUMEN {
		a.Borrar(i)
		require.Panics(t, func() { a.Obtener(i) })
	}
}

// Borra un elemento
func TestABBBorrar(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	a.Guardar(10, 11)
	a.Guardar(8, 0)

	require.Equal(t, a.Borrar(10), 11)
	require.Equal(t, a.Borrar(8), 0)

	// El elemento no deberia seguir apareciendo
	require.Panics(t, func() { a.Borrar(10) })
	require.Panics(t, func() { a.Borrar(8) })
}

func TestABBBorrarVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}
	for i := range _VOLUMEN {
		require.Equal(t, a.Borrar(i), i+1)
	}
}

// Verifica que el elemento se haya eliminado
func TestABBBorrarPanic(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	require.Panics(t, func() { a.Borrar(10) })
	require.Panics(t, func() { a.Borrar(9) })

	a.Guardar(10, 11)
	a.Borrar(10)
	require.Panics(t, func() { a.Borrar(10) })
}
func TestABBBorrarPanicVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
		require.NotPanics(t, func() { a.Borrar(i) })
	}

	for i := range _VOLUMEN {
		require.Panics(t, func() { a.Borrar(i) })
	}

}
func TestABBCantidad(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	a.Guardar(1, 10)
	require.Equal(t, a.Cantidad(), 1)

	a.Guardar(2, 12)
	require.Equal(t, a.Cantidad(), 2)
}

func TestABBCantidadVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
		require.Equal(t, a.Cantidad(), i+1)
	}
}
func TestABBCantidadBorrar(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	a.Guardar(1, 11)
	a.Guardar(2, 1)
	require.Equal(t, a.Cantidad(), 2)

	a.Borrar(1)
	require.Equal(t, a.Cantidad(), 1)

	a.Borrar(2)
	require.Equal(t, a.Cantidad(), 0)
}

func TestABBCantidadBorrarVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
		require.Equal(t, a.Cantidad(), i+1)
	}
	for i := range _VOLUMEN {
		a.Borrar(i)
	}
	require.Equal(t, a.Cantidad(), 0)
}

func TestABBObtenerCantidadBorrar(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	a.Guardar(2, 10)
	require.Equal(t, a.Cantidad(), 1)
	require.NotPanics(t, func() { a.Obtener(2) })

	a.Guardar(10, 23)
	require.Equal(t, a.Cantidad(), 2)
	require.NotPanics(t, func() { a.Obtener(2) })

	require.Equal(t, a.Borrar(2), 10)
	require.Equal(t, a.Obtener(10), 23)
	require.Equal(t, a.Cantidad(), 1)

}
func TestABBObtenerCantidadBorrarVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}
	for i := range _VOLUMEN {
		a.Borrar(i)
	}
	require.Equal(t, a.Cantidad(), 0)
}

func TestABBIterarVacio(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	a.Iterar(func(clave int, valor int) bool {
		return true
	})
}
func TestABBIterarElementosOrdenados(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	// lista := [1, 2, 3, 4, 5, 6, 7, 8]
	a.Guardar(1, 10)
	a.Guardar(2, 10)
	a.Guardar(3, 10)
	a.Guardar(4, 10)
	a.Guardar(5, 10)
	a.Guardar(6, 10)
	a.Guardar(7, 10)
	a.Guardar(8, 10)

	var claveVisitadas []int
	a.Iterar(func(clave int, valor int) bool {
		claveVisitadas = append(claveVisitadas, clave)
		return true
	})
	require.Equal(t, len(claveVisitadas), 8)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, claveVisitadas)
}

func TestABBIterarElementosOrdenadosVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	var expected_keys []int
	for i := range _VOLUMEN {
		a.Guardar(i, 10)
		expected_keys = append(expected_keys, i)
	}
	require.Equal(t, a.Cantidad(), _VOLUMEN)
	require.Equal(t, len(expected_keys), _VOLUMEN)

	var clavesVisitadas []int

	a.Iterar(func(clave, dato int) bool {
		clavesVisitadas = append(clavesVisitadas, clave)
		return true
	})
	// valida el orden
	require.Equal(t, clavesVisitadas, expected_keys)
}
func TestABBIterarElementosNoOrdenados(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)

	// lista := []int {5, 4, 3, 2, 1}
	var clavesVisitadas []int
	a.Guardar(5, 10)
	a.Guardar(4, 10)
	a.Guardar(3, 10)
	a.Guardar(2, 10)
	a.Guardar(1, 10)

	a.Iterar(func(clave, dato int) bool {
		clavesVisitadas = append(clavesVisitadas, clave)
		return true
	})

	require.Equal(t, a.Cantidad(), 5)
	require.Equal(t, []int{1, 2, 3, 4, 5}, clavesVisitadas)

}

func TestABBIterarElementosNoOrdenadosVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	for i := _VOLUMEN; i > 0; i-- {
		a.Guardar(i, 10)
	}
	var clavesVisitadas []int

	a.Iterar(func(clave, dato int) bool {
		clavesVisitadas = append(clavesVisitadas, clave)
		return true
	})
	require.Equal(t, len(clavesVisitadas), _VOLUMEN)
}

func TestABBIterador(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	a.Guardar(1, 10)
	a.Guardar(2, 12)

	it := a.Iterador()
	key, value := it.VerActual()
	require.Equal(t, key, 1)
	require.Equal(t, value, 10)
	it.Avanzar()
	key, value = it.VerActual()
	require.Equal(t, key, 2)
	require.Equal(t, value, 12)

}
func TestABBIteradorPanic(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	a.Guardar(1, 10)
	a.Guardar(2, 12)

	it := a.Iterador()
	require.NotPanics(t, func() { it.Avanzar() })
	require.NotPanics(t, func() { it.Avanzar() })

}
func TestABBIteradorVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}

	i := 0
	for it := a.Iterador(); it.HayAlgoMas(); it.Avanzar() {
		key, value := it.VerActual()
		require.Equal(t, key, i)
		require.Equal(t, value, i+1)
		i++
	}

}

func TestABBIteradorVerActualVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}
	it := a.Iterador()

	for i := 0; i < _VOLUMEN-1; i++ {
		it.Avanzar()
	}
	key, value := it.VerActual()
	require.Equal(t, key, _VOLUMEN-1)
	require.Equal(t, value, _VOLUMEN)
}
func TestABBIteradorRango(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	a.Guardar(1, 1)
	a.Guardar(2, 3)
	a.Guardar(3, 2)
	a.Guardar(4, 5)
	limSup := 4
	limInf := 2
	it := a.IteradorRango(&limInf, &limSup)

	key, _ := it.VerActual()
	require.Equal(t, key, limInf)
	it.Avanzar()

	key, _ = it.VerActual()
	require.Equal(t, key, 3)
	it.Avanzar()

	key, _ = it.VerActual()
	require.Equal(t, key, limSup)
}

func TestABBIteradorRangoUnico(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	a.Guardar(10, 11)
	limInf := 10
	limSup := 10

	it := a.IteradorRango(&limInf, &limSup)
	key, value := it.VerActual()
	require.Equal(t, key, 10)
	require.Equal(t, value, 11)
}

// Rangos Inferior inexistente
func TestABBIteradorRangoInferiorInexistente(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	a.Guardar(10, 12)
	a.Guardar(4, 12)
	limInf := 0
	limSup := 3
	it := a.IteradorRango(&limInf, &limSup)
	require.Panics(t, func() { it.VerActual() })
}
func TestABBIteradorRangoSuperiorInexistente(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	a.Guardar(1, 2)
	a.Guardar(2, 3)
	a.Guardar(3, 4)
	a.Guardar(4, 5)
	a.Guardar(5, 6)

	limInf := 6
	limSup := 100
	it := a.IteradorRango(&limInf, &limSup)

	require.Panics(t, func() { it.VerActual() })
}
func TestABBIteradorRangoInferiorInexistenteVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := -_VOLUMEN
	limSup := -1

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}

	it := a.IteradorRango(&limInf, &limSup)
	require.Panics(t, func() { it.VerActual() })
}
func TestABBIteradorRangoSuperiorInexistenteVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := _VOLUMEN + 1
	limSup := _VOLUMEN + 4

	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}

	it := a.IteradorRango(&limInf, &limSup)

	require.Panics(t, func() { it.VerActual() })
}

func TestABBIterarRango(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := 0
	limSup := 5
	var keys []int

	// 1, 2, 3, 4, 5
	a.Guardar(1, 10)
	a.Guardar(2, 10)
	a.Guardar(3, 10)
	a.Guardar(4, 10)
	a.Guardar(5, 10)

	a.IterarRango(&limInf, &limSup, func(clave, dato int) bool {
		keys = append(keys, clave)
		return true
	})
	require.Equal(t, len(keys), 5)
	require.Equal(t, keys, []int{1, 2, 3, 4, 5})

}
func TestABBIterarRangoVacio(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := 0
	limSup := 0
	var keys []int
	a.IterarRango(&limInf, &limSup, func(clave, dato int) bool {
		keys = append(keys, clave)
		return true
	})
	require.Equal(t, len(keys), 0)
	require.Equal(t, keys, []int(nil))
}
func TestABBIterarRangoNoOrdenado(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := 0
	limSup := 5
	var keys []int

	a.Guardar(5, 8)
	a.Guardar(4, 8)
	a.Guardar(3, 8)
	a.Guardar(2, 8)
	a.Guardar(1, 8)
	a.IterarRango(&limInf, &limSup, func(clave, dato int) bool {
		keys = append(keys, clave)
		return true
	})
	require.Equal(t, len(keys), 5)
	require.Equal(t, keys, []int{1, 2, 3, 4, 5})
}
func TestABBIterarRangoInferiorInexistente(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := 1
	limSup := 10
	var keys []int
	a.Guardar(11, 2)
	a.Guardar(51, 2)
	a.IterarRango(&limInf, &limSup, func(clave, dato int) bool {
		keys = append(keys, clave)
		return true
	})
	require.Equal(t, len(keys), 0)
	require.Equal(t, keys, []int(nil))
}
func TestABBIterarRangoSuperiorInexistente(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := 50
	limSup := 100
	var keys []int
	a.Guardar(1, 4)
	a.Guardar(10, 15)
	a.Guardar(8, 9)
	a.Guardar(7, 6)
	a.Guardar(5, 6)
	a.IterarRango(&limInf, &limSup, func(clave, dato int) bool {
		keys = append(keys, clave)
		return true
	})
	require.Equal(t, len(keys), 0)
	require.Equal(t, keys, []int(nil))
}
func TestABBIterarRangoInferiorVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := -10
	limSup := -5
	var keys []int
	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}
	a.IterarRango(&limInf, &limSup, func(clave, dato int) bool {
		keys = append(keys, clave)
		return true
	})
	require.Equal(t, len(keys), 0)
	require.Equal(t, keys, []int(nil))
}
func TestABBIterarRangoSuperiorVolumen(t *testing.T) {
	a := TDADict.CrearABB[int, int](cmpInt)
	limInf := _VOLUMEN + 1
	limSup := _VOLUMEN + 100
	var keys []int
	for i := range _VOLUMEN {
		a.Guardar(i, i+1)
	}
	a.IterarRango(&limInf, &limSup, func(clave, dato int) bool {
		keys = append(keys, clave)
		return true
	})
	require.Equal(t, len(keys), 0)
	require.Equal(t, keys, []int(nil))
}
