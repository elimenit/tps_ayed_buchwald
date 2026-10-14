package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	volumen                = 10000
	mensajeListaVacia      = "La lista esta vacia"
	mensajeIteradorTermino = "El iterador termino de iterar"
)

// Los tests estan siendo implementados lo mas unitarios posibles.

// La pila se encuentra vacia
func TestEstaVacia(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	require.True(t, l.EstaVacia())

	l.InsertarPrimero(10)
	l.BorrarPrimero()
	require.True(t, l.EstaVacia())
}

func TestVolumenEstaVacia(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := range volumen {
		l.InsertarPrimero(i)
		require.False(t, l.EstaVacia())
	}

	for i := 0; i < volumen-1; i++ {
		l.BorrarPrimero()
		require.False(t, l.EstaVacia())
	}
	l.BorrarPrimero()
	require.True(t, l.EstaVacia())
}

/*
Mira que el primer elemento siga siendo el primero.
*/
func TestVerPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(10)
	require.Equal(t, lista.VerPrimero(), 10)

	lista.InsertarPrimero(100)
	require.Equal(t, lista.VerPrimero(), 100)
}

// VerPrimero: no cambia despues agregar muchos datos
func TestVolumenVerPrimero(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < volumen; i++ {
		l.InsertarPrimero(i)
		require.Equal(t, l.VerPrimero(), i)
	}
}

// Panique cuando la lista esta vacia.
func TestPanicVerPrimero(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	require.Panics(t, func() { l.VerPrimero() }, mensajeListaVacia)

	l.InsertarPrimero(10)
	l.BorrarPrimero()
	require.Panics(t, func() { l.VerPrimero() }, mensajeListaVacia)

	for i := 0; i < volumen; i++ {
		require.Panics(t, func() { l.VerPrimero() })
	}
}

// Agrega elementos
func TestInsertarPrimero(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	l.InsertarPrimero(10)
	require.Equal(t, l.VerPrimero(), 10)

	l.InsertarPrimero(20)
	require.Equal(t, l.VerPrimero(), 20)

	l.InsertarUltimo(50)
	require.Equal(t, l.VerUltimo(), 50)

}

// Estabilidad de elementos agregados
func TestVolumenInsertarPrimero(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := range volumen {
		l.InsertarPrimero(i)
	}
	for i := range volumen {
		require.Equal(t, l.BorrarPrimero(), volumen-i-1)
	}
}

// Estabilidad de agregar elementos al final.
func TestInsertarUltimo(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarUltimo(10)
	require.Equal(t, l.VerUltimo(), 10)

	l.InsertarUltimo(100)
	require.Equal(t, l.VerUltimo(), 100)
}

// Estabilidad de insertar al final
func TestVolumenInsertarUltimo(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := range volumen {
		l.InsertarUltimo(i)
		require.Equal(t, l.VerUltimo(), i)
	}
}

// Estabilidad del primer elemento
func TestBorrarPrimero(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarPrimero(10)
	require.Equal(t, l.BorrarPrimero(), 10)

	l.InsertarPrimero(15)
	require.Equal(t, l.BorrarPrimero(), 15)

	l.InsertarPrimero(15)
	l.InsertarPrimero(10)

	require.Equal(t, l.BorrarPrimero(), 10)
	require.Equal(t, l.BorrarPrimero(), 15)
}

// Estabilidad de eliminar primero de muchos elementos
func TestVolumenBorrarPrimero(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := range volumen {
		l.InsertarPrimero(i)
		require.Equal(t, l.BorrarPrimero(), i)
	}

	for i := 0; i < volumen; i++ {
		l.InsertarPrimero(i)
	}
	for i := 1; i < volumen; i++ {
		require.Equal(t, l.BorrarPrimero(), volumen-i)
	}
	require.Equal(t, l.BorrarPrimero(), 0)
}

// Paniquea cuando la lista esta vacia
func TestPanicBorrarPrimero(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	require.Panics(t, func() { l.BorrarPrimero() })

	l.InsertarPrimero(10)
	l.BorrarPrimero()
	require.Panics(t, func() { l.BorrarPrimero() })

	l.InsertarPrimero(15)
	l.InsertarPrimero(10)
	l.BorrarPrimero()
	l.BorrarPrimero()
	require.Panics(t, func() { l.BorrarPrimero() })
}

// Paniquea despues de tener muchos elementos.
func TestPanicVolumenBorrarprimero(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < volumen; i++ {
		l.InsertarPrimero(i)
	}
	for i := 1; i < volumen; i++ {
		l.BorrarPrimero()
	}
	l.BorrarPrimero()
	require.Panics(t, func() { l.BorrarPrimero() }, mensajeListaVacia)
}

func TestVerUltimo(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarUltimo(10)
	require.Equal(t, l.VerUltimo(), 10)

	l.InsertarUltimo(15)
	l.InsertarUltimo(14)
	require.Equal(t, l.VerUltimo(), 14)

	l.BorrarPrimero()
	require.Equal(t, l.VerUltimo(), 14)
}

func TestVolumenVerUltimo(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	var ultimo int = 10000

	for i := range volumen {
		l.InsertarUltimo(i)
	}

	l.InsertarUltimo(ultimo)
	require.Equal(t, l.VerUltimo(), ultimo)

	for i := 0; i < volumen; i++ {
		l.BorrarPrimero()
	}
	require.Equal(t, l.VerUltimo(), ultimo)
}

// Paniquea cuando la lista esta vacia
func TestPanicVerUltimo(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	require.Panics(t, func() { l.VerUltimo() }, mensajeListaVacia)

	l.InsertarUltimo(10)
	l.InsertarUltimo(9)
	l.BorrarPrimero()
	l.BorrarPrimero()
	require.Panics(t, func() { l.VerUltimo() }, mensajeListaVacia)

}

func TestPanicVolumenVerUltimo(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := range volumen {
		l.InsertarUltimo(i)
	}

	for i := 0; i < volumen; i++ {
		l.BorrarPrimero()
	}

	require.Panics(t, func() { l.VerUltimo() }, mensajeListaVacia)

}

// Insercion o eliminacion de elementos aumenta o disminuye el largo.
func TestLargoVacio(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	require.True(t, l.EstaVacia())

	l.InsertarPrimero(10)
	require.Equal(t, l.Largo(), 1)

	l.InsertarUltimo(15)
	require.Equal(t, l.Largo(), 2)

	l.BorrarPrimero()
	l.BorrarPrimero()
	require.Equal(t, l.Largo(), 0)

	l.InsertarPrimero(15)
	require.Equal(t, l.Largo(), 1)

	l.InsertarPrimero(14)
	require.Equal(t, l.Largo(), 2)

	// lista Vacia
	l.BorrarPrimero()
	l.BorrarPrimero()
	require.Equal(t, l.Largo(), 0)

	l.InsertarUltimo(144)
	require.Equal(t, l.Largo(), 1)

	l.InsertarUltimo(10)
	require.Equal(t, l.Largo(), 2)

	l.BorrarPrimero()
	l.BorrarPrimero()

	require.Equal(t, l.Largo(), 0)

}
func TestLargo(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	l.InsertarPrimero(10)
	require.Equal(t, l.Largo(), 1)

	l.InsertarUltimo(100)
	require.Equal(t, l.Largo(), 2)

	l.BorrarPrimero()
	require.Equal(t, l.Largo(), 1)

	l.BorrarPrimero()
	require.Equal(t, l.Largo(), 0)
}

// El largo es estable segun la cantidad de elementos.
func TestLargoVolumen(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	var cantidad int = 0

	for i := range volumen {
		l.InsertarPrimero(i)
		cantidad++
		require.Equal(t, l.Largo(), cantidad)
	}

	for i := 0; i < volumen; i++ {
		l.BorrarPrimero()
		cantidad--
		require.Equal(t, l.Largo(), cantidad)
	}
	require.Equal(t, l.Largo(), 0)
}

// Validar la estabilidad de elementos
func TestIterarIgual(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)

	var obtenidos []int

	lista.Iterar(func(elemento int) bool {
		obtenidos = append(obtenidos, elemento)
		return true
	})

	esperados := []int{10, 20, 30}

	for i, e := range obtenidos {
		require.Equal(t, e, esperados[i])
	}
}

func TestIterarDuplicado(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarPrimero(10)
	l.InsertarUltimo(20)
	l.InsertarUltimo(30)

	var obtenidos []int
	var esperados = []int{20, 40, 60}

	l.Iterar(func(e int) bool {
		obtenidos = append(obtenidos, e*2)
		return true
	})

	for i, e := range obtenidos {
		require.Equal(t, e, esperados[i])
	}
}

// Iterador Externo
func TestIteradorExternoVacio(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	require.False(t, l.Iterador().HayAlgoMas())

	l.InsertarPrimero(10)
	require.True(t, l.Iterador().HayAlgoMas())

	l.InsertarUltimo(15)
	require.True(t, l.Iterador().HayAlgoMas())

	l.BorrarPrimero()
	require.True(t, l.Iterador().HayAlgoMas())

	l.BorrarPrimero()
	require.False(t, l.Iterador().HayAlgoMas())
}

// Paniquea cuando la lista tiene 0 elementos
func TestIteradorPanicVerActual(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	require.Panics(t, func() { l.Iterador().VerActual() }, mensajeIteradorTermino)

	l.InsertarPrimero(10)
	l.BorrarPrimero()
	require.Panics(t, func() { l.Iterador().VerActual() }, mensajeIteradorTermino)

}

// Iterador deberia mostrar el dato actual
func TestIteradorVerActual(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	l.InsertarPrimero(10)
	l.InsertarUltimo(20)
	ie := l.Iterador()
	require.Equal(t, ie.VerActual(), 10)

	ie.Avanzar()
	require.Equal(t, ie.VerActual(), 20)

	l.BorrarPrimero()
	require.Equal(t, l.Iterador().VerActual(), 20)

	// lista vacia
	l.BorrarPrimero()
	l.InsertarPrimero(10)
	l.InsertarPrimero(20)
	ie = l.Iterador()
	require.Equal(t, ie.VerActual(), 20)

	ie.Avanzar()
	require.Equal(t, ie.VerActual(), 10)

	// Lista Vacia
	l.BorrarPrimero()
	l.BorrarPrimero()
	l.InsertarUltimo(20)
	l.InsertarUltimo(15)
	ie = l.Iterador()
	require.Equal(t, ie.VerActual(), 20)

	ie.Avanzar()
	require.Equal(t, ie.VerActual(), 15)

	l.BorrarPrimero()
	ie = l.Iterador()
	require.Equal(t, ie.VerActual(), 15)
}

func TestIteradorPanicVerActualVolumen(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := range volumen {
		l.InsertarPrimero(i)
	}
	ie := l.Iterador()
	for i := volumen - 1; i > 0; i-- {
		require.Equal(t, ie.VerActual(), i)
		ie.Avanzar()
	}

}

func TestIteradorVacioAvanzar(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	ie := l.Iterador()

	require.Panics(t, func() { ie.Avanzar() })
}
func TestIteradorAvanzar(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	l.InsertarPrimero(10)
	l.InsertarUltimo(25)
	ie := l.Iterador()
	require.Equal(t, ie.VerActual(), 10)
	ie.Avanzar()
	require.Equal(t, ie.VerActual(), 25)
}

func TestIteradorPanicAvanzar(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	l.InsertarPrimero(10)
	ie := l.Iterador()
	ie.Avanzar()
	require.Panics(t, func() { ie.Avanzar() })
}

func TestIteradorPanicAvanzarVolumen(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	for i := range volumen {
		l.InsertarUltimo(i)
	}

	ie := l.Iterador()
	for i := 0; i < volumen; i++ {
		ie.Avanzar()
	}
	require.Panics(t, func() { ie.Avanzar() })
}

func TestIteradorInsertarVacio(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	ie := l.Iterador()
	ie.Insertar(5)
	require.Equal(t, ie.VerActual(), 5)
}

func TestIteradorInsertarVolumen(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	ie := l.Iterador()

	for i := range volumen {
		ie.Insertar(i)
		require.Equal(t, ie.VerActual(), i)
	}
}

func TestIteradorBorrarPanic(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	ie := l.Iterador()

	require.Panics(t, func() { ie.Borrar() })

	ie.Insertar(10)
	ie.Borrar()
	require.Panics(t, func() { ie.Borrar() })
}

func TestIteradorBorrarVacio(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	ie := l.Iterador()
	ie.Insertar(10)
	ie.Borrar()
	require.False(t, ie.HayAlgoMas())
}
func TestIteradorBorrarVacioVolumen(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	ie := l.Iterador()

	for i := range volumen {
		ie.Insertar(i)
	}

	for i := volumen; i > 0; i-- {
		ie.Borrar()
	}
	require.False(t, ie.HayAlgoMas())
}

func TestIteradorBorrarVolumenPanic(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	ie := l.Iterador()

	for i := range volumen {
		ie.Insertar(i)
	}

	for i := volumen; i > 0; i-- {
		ie.Borrar()
	}
	require.Panics(t, func() { ie.Borrar() }, mensajeIteradorTermino)
}
