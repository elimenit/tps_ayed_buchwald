package hash

import (
	"fmt"
	"hash/fnv"
)

const (
	capacidadInicial    = 10
	proporcionModificar = 2
	cargaMax            = 0.7
	cargaMin            = 0.25
)

const (
	VACIO = iota
	OCUPADO
	BORRADO
)

type elemento[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type hash[K comparable, V any] struct {
	tabla        []elemento[K, V]
	cantidad     int
	cantBorrados int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hash[K, V]{tabla: make([]elemento[K, V], capacidadInicial)}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func funcionHashing[K comparable](clave K) uint64 {
	h := fnv.New64a()
	h.Write(convertirABytes(clave))
	return h.Sum64()
}

func (h *hash[K, V]) obtenerPosicion(clave K) int {
	hasheado := funcionHashing(clave)
	return int(hasheado % uint64(len(h.tabla)))
}

//devuelvo una posicion de la tabla y un booleano indicando si la clave ya existia
func (h *hash[K, V]) buscarPosicionClave(clave K) (int, bool) {
	posicion := h.obtenerPosicion(clave)
	primerBorrado := -1

	for visitadas := 0; visitadas < len(h.tabla); visitadas++ {
		if h.tabla[posicion].estado == BORRADO && primerBorrado == -1 {
			primerBorrado = posicion
		}

		if h.tabla[posicion].estado == OCUPADO && h.tabla[posicion].clave == clave {
			return posicion, true
		}

		if h.tabla[posicion].estado == VACIO {
			if primerBorrado != -1 {
				return primerBorrado, false
			}
			return posicion, false
		}

		posicion = (posicion + 1) % len(h.tabla)
	}
	return primerBorrado, false
}

func (h *hash[K, V]) redimensionar(nuevaCapacidad int) {
	tablaAnterior := h.tabla
	h.tabla = make([]elemento[K, V], nuevaCapacidad)
	h.cantidad = 0
	h.cantBorrados = 0
	for _, elemento := range tablaAnterior {
		if elemento.estado == OCUPADO {
			h.Guardar(elemento.clave, elemento.dato)
		}
	}
}

func (h *hash[K, V]) Guardar(clave K, dato V) {
	factorCarga := float64(h.cantidad+h.cantBorrados) / float64(len(h.tabla))

	if factorCarga >= cargaMax {
		h.redimensionar(len(h.tabla) * proporcionModificar)
	}

	posicion, encontrada := h.buscarPosicionClave(clave)

	if encontrada {
		h.tabla[posicion].dato = dato
		return
	}

	if h.tabla[posicion].estado == BORRADO {
		h.cantBorrados--
	}

	h.tabla[posicion] = elemento[K, V]{
		clave:  clave,
		dato:   dato,
		estado: OCUPADO,
	}
	h.cantidad++
}

func (h *hash[K, V]) Pertenece(clave K) bool {
	_, encontrada := h.buscarPosicionClave(clave)
	return encontrada
}

func (h *hash[K, V]) Obtener(clave K) V {
	posicion, encontrada := h.buscarPosicionClave(clave)
	if !encontrada {
		panic("La clave no pertenece al diccionario")
	}
	return h.tabla[posicion].dato
}

func (h *hash[K, V]) Borrar(clave K) V {
	posicion, encontrada := h.buscarPosicionClave(clave)
	if !encontrada {
		panic("La clave no pertenece al diccionario")
	}
	dato := h.tabla[posicion].dato
	h.tabla[posicion].estado = BORRADO
	h.cantBorrados++
	h.cantidad--

	factorCarga := float64(h.cantidad) / float64(len(h.tabla))
	if factorCarga <= cargaMin && len(h.tabla) > capacidadInicial {
		h.redimensionar(len(h.tabla) / proporcionModificar)
	}
	return dato
}

func (h *hash[K, V]) Cantidad() int {
	return h.cantidad
}

//func Iterar(func(clave K, dato V) bool) {}

//func Iterador() IteradorLista[K, V] {}

//func (it *IterDiccionario[K, V]) HayAlgoMas() bool {}

//func VerActual() (K, V) {}

//func Avanzar() {}
