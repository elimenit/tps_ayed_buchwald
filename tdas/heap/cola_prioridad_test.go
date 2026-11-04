package cola_prioridad_test

import (
	"strings"
	CP "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararCadenas(cad1 string, cad2 string) int {
	return strings.Compare(cad1, cad2)
}

func compararEnteros(a int, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	heap := CP.CrearHeap(compararEnteros)

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())

	require.PanicsWithValue(t, "El heap esta vacio", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapEncolarYDesencolarBasico(t *testing.T) {
	heap := CP.CrearHeap(compararEnteros)

	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())

	heap.Encolar(20)
	require.Equal(t, 2, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())

	heap.Encolar(5)
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())

	require.Equal(t, 20, heap.Desencolar())
	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "El heap esta vacio", func() { heap.VerMax() })
}

func TestHeapCadenas(t *testing.T) {
	heap := CP.CrearHeap(compararCadenas)

	heap.Encolar("perro")
	heap.Encolar("gato")
	heap.Encolar("zapato")
	heap.Encolar("auto")

	require.Equal(t, 4, heap.Cantidad())

	require.Equal(t, "zapato", heap.Desencolar())
	require.Equal(t, "perro", heap.Desencolar())
	require.Equal(t, "gato", heap.Desencolar())
	require.Equal(t, "auto", heap.Desencolar())
}

func TestHeapVolumenYRedimension(t *testing.T) {
	heap := CP.CrearHeap(compararEnteros)
	volumen := 1000

	for i := 0; i < volumen; i++ {
		heap.Encolar(i)
	}

	require.Equal(t, volumen, heap.Cantidad())
	require.False(t, heap.EstaVacia())

	for i := volumen - 1; i >= 0; i-- {
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
}

func TestCrearHeapArr(t *testing.T) {
	arreglo := []int{10, 5, 20, 1, 15, 30, 25}

	heap := CP.CrearHeapArr(arreglo, compararEnteros)

	require.False(t, heap.EstaVacia())
	require.Equal(t, 7, heap.Cantidad())

	esperado := []int{30, 25, 20, 15, 10, 5, 1}

	for _, val := range esperado {
		require.Equal(t, val, heap.VerMax())
		require.Equal(t, val, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
}

func TestCrearHeapArrVacio(t *testing.T) {
	arreglo := []int{}
	heap := CP.CrearHeapArr(arreglo, compararEnteros)

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSort(t *testing.T) {
	arreglo := []int{10, 5, 20, 1, 15, 30, 25, -5, 100, 0}

	CP.HeapSort(arreglo, compararEnteros)

	esperado := []int{-5, 0, 1, 5, 10, 15, 20, 25, 30, 100}

	require.Equal(t, esperado, arreglo)
}

func TestHeapSortVacioYUnElemento(t *testing.T) {
	arrVacio := []int{}
	CP.HeapSort(arrVacio, compararEnteros)
	require.Equal(t, []int{}, arrVacio)

	arrUno := []int{42}
	CP.HeapSort(arrUno, compararEnteros)
	require.Equal(t, []int{42}, arrUno)
}

func TestHeapDesencolarVaciaElHeapYPermiteReencolar(t *testing.T) {
	heap := CP.CrearHeap(compararEnteros)

	heap.Encolar(10)
	heap.Encolar(20)

	heap.Desencolar()
	heap.Desencolar()

	require.True(t, heap.EstaVacia())

	heap.Encolar(50)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 50, heap.VerMax())
}
