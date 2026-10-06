package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	volumen = 10000
)

// Valida que al instanciar este vacia
func TestEstaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia(), "Debe validar que la cola este vacia")

	cola.Encolar(10)
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "Debe validar que la cola este vacia")
}

func TestEncolarDesencolarVacia(t *testing.T) {
	c := TDACola.CrearColaEnlazada[string]()
	c.Encolar("example")
	c.Desencolar()
	require.True(t, c.EstaVacia())

	c.Encolar("Gopher")
	c.Encolar("Google")
	c.Encolar("AWS")
	c.Encolar("Azure")
	c.Encolar("Google Cloud")

	require.Equal(t, "Gopher", c.Desencolar(), "NO se desencolo Gopher")
	require.Equal(t, "Google", c.Desencolar(), "No se desencolo Google")
	require.Equal(t, "AWS", c.Desencolar(), "No se desencolo AWS")
	require.Equal(t, "Azure", c.Desencolar(), "No se desencolo Azure")
	require.Equal(t, "Google Cloud", c.Desencolar(), "NO se desencolo Google Cloud")
	require.True(t, c.EstaVacia())
}

func TestEncolarDesencolarPanic(t *testing.T) {
	c := TDACola.CrearColaEnlazada[int]()

	require.Panics(t, func() { c.Desencolar() })

	for i := range volumen {
		c.Encolar(i)
	}
	for i := 0; i < volumen; i++ {
		c.Desencolar()
	}
	require.Panics(t, func() { c.Desencolar() })
}

// Valida que las primitivas Encolar y Desencolar funcionen correctamente
func TestEncolarYDesencolar(t *testing.T) {
	c := TDACola.CrearColaEnlazada[string]()
	cola := TDACola.CrearColaEnlazada[int]()

	c.Encolar("Gopher")
	c.Encolar("Google")
	c.Encolar("AWS")
	c.Encolar("Azure")
	c.Encolar("Google Cloud")

	require.Equal(t, "Gopher", c.Desencolar(), "NO se desencolo Gopher")
	require.Equal(t, "Google", c.Desencolar(), "No se desencolo Google")
	require.Equal(t, "AWS", c.Desencolar(), "No se desencolo AWS")
	require.Equal(t, "Azure", c.Desencolar(), "No se desencolo Azure")
	require.Equal(t, "Google Cloud", c.Desencolar(), "NO se desencolo Google Cloud")

	// Volumen
	for i := range volumen {
		cola.Encolar(i)
	}
	for i := 0; i < volumen; i++ {
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
}

// Testea que la funcion panickee cuando desencola la cola vacia
func TestVerPrimeroPanic(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.Panics(t, func() { cola.VerPrimero() })

	cola.Encolar(1)
	cola.Desencolar()
	require.Panics(t, func() { cola.VerPrimero() })

	// Volumen
	for i := 0; i < volumen; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < volumen; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.Panics(t, func() { cola.VerPrimero() })
}

// Testea que la primitiva VerPrimero apunte siempre al primer elemento
func TestVerPrimero(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	// One Element
	cola.Encolar(10)

	require.Equal(t, 10, cola.VerPrimero())

	cola.Desencolar()

	for i := 0; i < volumen; i++ {
		cola.Encolar(i)

		require.Equal(t, 0, cola.VerPrimero())
	}

	for i := 0; i < volumen; i++ {
		require.Equal(t, i, cola.VerPrimero())

		cola.Desencolar()
	}
}
