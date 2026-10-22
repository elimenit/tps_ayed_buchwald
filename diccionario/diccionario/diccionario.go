package diccionario

import (
	"fmt"
	"hash/fnv"
)

const (
	_CAPACIDADINICIAL    = 10
	_PROPORCIONMODIFICAR = 2
	_CARGAMAX            = 0.7
	_CARGAMIN            = 0.25
)

const (
	_VACIO = iota
	_OCUPADO
	_BORRADO
)

// Estructura del iterador externo
type iteradorHash[K comparable, V any] struct {
	hash     *hash[K, V]
	posicion int
	vistos   int
}

type elemento[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type hash[K comparable, V any] struct {
	tabla         []elemento[K, V]
	cantidad      int
	_cantBorrados int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hash[K, V]{
		tabla:         make([]elemento[K, V], _CAPACIDADINICIAL),
		cantidad:      0,
		_cantBorrados: 0,
	}
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

func (h *hash[K, V]) buscarPosicionClave(clave K) (int, bool) {
	posicion := h.obtenerPosicion(clave)
	_primerBorrado := -1

	for visitadas := 0; visitadas < len(h.tabla); visitadas++ {
		if h.tabla[posicion].estado == _BORRADO && _primerBorrado == -1 {
			_primerBorrado = posicion
		}

		if h.tabla[posicion].estado == _OCUPADO && h.tabla[posicion].clave == clave {
			return posicion, true
		}

		if h.tabla[posicion].estado == _VACIO {
			if _primerBorrado != -1 {
				return _primerBorrado, false
			}
			return posicion, false
		}

		posicion = (posicion + 1) % len(h.tabla)
	}
	return _primerBorrado, false
}

func (h *hash[K, V]) redimensionar(nuevaCapacidad int) {
	tablaAnterior := h.tabla
	h.tabla = make([]elemento[K, V], nuevaCapacidad)
	h.cantidad = 0
	h._cantBorrados = 0
	for _, elemento := range tablaAnterior {
		if elemento.estado == _OCUPADO {
			h.Guardar(elemento.clave, elemento.dato)
		}
	}
}

func (h *hash[K, V]) Guardar(clave K, dato V) {
	factorCarga := float64(h.cantidad+h._cantBorrados) / float64(len(h.tabla))

	if factorCarga >= _CARGAMAX {
		h.redimensionar(len(h.tabla) * _PROPORCIONMODIFICAR)
	}

	posicion, encontrada := h.buscarPosicionClave(clave)

	if encontrada {
		h.tabla[posicion].dato = dato
		return
	}

	if h.tabla[posicion].estado == _BORRADO {
		h._cantBorrados--
	}

	h.tabla[posicion] = elemento[K, V]{
		clave:  clave,
		dato:   dato,
		estado: _OCUPADO,
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
	h.tabla[posicion].estado = _BORRADO
	h._cantBorrados++
	h.cantidad--

	factorCarga := float64(h.cantidad) / float64(len(h.tabla))
	if factorCarga <= _CARGAMIN && len(h.tabla) > _CAPACIDADINICIAL {
		h.redimensionar(len(h.tabla) / _PROPORCIONMODIFICAR)
	}
	return dato
}

func (h *hash[K, V]) Cantidad() int {
	return h.cantidad
}

// Iterador interno
func (h *hash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, elem := range h.tabla {
		if elem.estado == _OCUPADO {
			if !visitar(elem.clave, elem.dato) {
				break
			}
		}
	}
}

// Creación del iterador externo
func (h *hash[K, V]) Iterador() IterDiccionario[K, V] {
	it := &iteradorHash[K, V]{
		hash:     h,
		posicion: 0,
		vistos:   0,
	}
	it.buscarSiguienteOcupado()
	return it
}

func (it *iteradorHash[K, V]) buscarSiguienteOcupado() {
	for it.posicion < len(it.hash.tabla) && it.hash.tabla[it.posicion].estado != _OCUPADO {
		it.posicion++
	}
}

func (it *iteradorHash[K, V]) HayAlgoMas() bool {
	return it.vistos < it.hash.cantidad
}

func (it *iteradorHash[K, V]) VerActual() (K, V) {
	if !it.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return it.hash.tabla[it.posicion].clave, it.hash.tabla[it.posicion].dato
}

func (it *iteradorHash[K, V]) Avanzar() {
	if !it.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	it.posicion++
	it.vistos++
	it.buscarSiguienteOcupado()
}
