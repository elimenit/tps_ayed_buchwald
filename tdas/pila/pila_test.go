package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	pila1 := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila1.EstaVacia())

	pila2 := TDAPila.CrearPilaDinamica[bool]()
	require.True(t, pila2.EstaVacia())

}

func TestApilarYDesapilar(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[string]()

	p.Apilar("Gopher")
	p.Apilar("Google")

	require.False(t, p.EstaVacia(), "La pila no debería estar vacia si tiene elementos apilados!")

	require.Equal(t, "Google", p.VerTope(), "El tope esperado era 'Google'!")
	require.Equal(t, "Google", p.Desapilar(), "Se esperaba desapilar 'Google'!")
	require.Equal(t, "Gopher", p.Desapilar(), "Se esperaba desapilar 'Gopher'!")

	require.True(t, p.EstaVacia(), "La pila debería estar vacia!")
}

func TestApilaYDesapilarVacio(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.Panics(t, func() { pila.Desapilar() })

	pila.Apilar(10)
	pila.Desapilar()

	require.Panics(t, func() { pila.Desapilar() })

}
func TestVolumenRedimension(t *testing.T) {
	p := TDAPila.CrearPilaDinamica[int]()
	volumen := 10000

	for i := 0; i < volumen; i++ {
		p.Apilar(i)
	}

	for i := volumen - 1; i >= 0; i-- {
		require.Equal(t, i, p.VerTope(), "Fallo en VerTope al probar con volumen")
		require.Equal(t, i, p.Desapilar(), "Fallo en Desapilar al probar con volumen")
	}

	require.True(t, p.EstaVacia(), "La pila debería estar vacia")
	require.Panics(t, func() { p.Desapilar() })
}

func TestVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.Panics(t, func() { pila.VerTope() })

	pila.Apilar(10)
	pila.Desapilar()

	require.Panics(t, func() { pila.VerTope() })
}
