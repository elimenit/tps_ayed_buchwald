package diccionario

import (
	"tdas/pila"
)

type nodoArbol[K comparable, V any] struct {
	clave     K
	dato      V
	izquierda *nodoArbol[K, V]
	derecha   *nodoArbol[K, V]
}

type abb[K comparable, V any] struct {
	raiz     *nodoArbol[K, V]
	cantidad int
	comparar func(K, K) int
}

type iterABB[K comparable, V any] struct {
	pila     pila.Pila[*nodoArbol[K, V]]
	desde    *K
	hasta    *K
	comparar func(K, K) int
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoArbol[K, V] {
	return &nodoArbol[K, V]{
		clave: clave,
		dato:  dato,
	}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:     nil,
		cantidad: 0,
		comparar: funcion_cmp,
	}
}

func (a *abb[K, V]) guardarRecursiva(actual *nodoArbol[K, V], nuevoNodo *nodoArbol[K, V]) {
	comparacion := a.comparar(nuevoNodo.clave, actual.clave)
	if comparacion == 0 {
		actual.dato = nuevoNodo.dato
		return
	}

	if comparacion < 0 {
		if actual.izquierda == nil {
			actual.izquierda = nuevoNodo
			a.cantidad++
			return
		}
		a.guardarRecursiva(actual.izquierda, nuevoNodo)
		return
	}

	if comparacion > 0 {
		if actual.derecha == nil {
			actual.derecha = nuevoNodo
			a.cantidad++
			return
		}
		a.guardarRecursiva(actual.derecha, nuevoNodo)
		return
	}
}

func (a *abb[K, V]) buscarRecursiva(actual *nodoArbol[K, V], clave K) *nodoArbol[K, V] {
	if actual == nil {
		return nil
	}
	comparacion := a.comparar(clave, actual.clave)

	if comparacion == 0 {
		return actual
	} else if comparacion < 0 {
		return a.buscarRecursiva(actual.izquierda, clave)
	}
	return a.buscarRecursiva(actual.derecha, clave)

}

func (a *abb[K, V]) borrarNodo(actual *nodoArbol[K, V], clave K) *nodoArbol[K, V] {
	if actual == nil {
		return nil
	}
	comparacion := a.comparar(clave, actual.clave)

	if comparacion < 0 {
		actual.izquierda = a.borrarNodo(actual.izquierda, clave)
		return actual
	}

	if comparacion > 0 {
		actual.derecha = a.borrarNodo(actual.derecha, clave)
		return actual
	}

	if actual.izquierda == nil {
		return actual.derecha
	}

	if actual.derecha == nil {
		return actual.izquierda
	}

	siguiente := actual.derecha
	for siguiente.izquierda != nil {
		siguiente = siguiente.izquierda
	}
	actual.clave = siguiente.clave
	actual.dato = siguiente.dato

	actual.derecha = a.borrarNodo(actual.derecha, siguiente.clave)
	return actual
}

func iterarRecursiva[K comparable, V any](nodo *nodoArbol[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool, comparar func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || comparar(nodo.clave, *desde) > 0 {
		if !iterarRecursiva(nodo.izquierda, desde, hasta, visitar, comparar) {
			return false
		}
	}
	if (desde == nil || comparar(nodo.clave, *desde) >= 0) && (hasta == nil || comparar(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || comparar(nodo.clave, *hasta) < 0 {
		if !iterarRecursiva(nodo.derecha, desde, hasta, visitar, comparar) {
			return false
		}
	}
	return true
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	nuevoNodo := crearNodo(clave, dato)
	if a.raiz == nil {
		a.raiz = nuevoNodo
		a.cantidad++
		return
	}
	a.guardarRecursiva(a.raiz, nuevoNodo)
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	return a.buscarRecursiva(a.raiz, clave) != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	resultado := a.buscarRecursiva(a.raiz, clave)
	if resultado == nil {
		panic("La clave no pertenece al diccionario")
	}
	return resultado.dato
}

func (a *abb[K, V]) Borrar(clave K) V {
	nodo := a.buscarRecursiva(a.raiz, clave)

	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	dato := nodo.dato

	a.raiz = a.borrarNodo(a.raiz, clave)

	a.cantidad--

	return dato
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterarRecursiva(a.raiz, nil, nil, visitar, a.comparar)
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iterABB[K, V]{
		pila:     pila.CrearPilaDinamica[*nodoArbol[K, V]](),
		desde:    nil,
		hasta:    nil,
		comparar: a.comparar,
	}
	actual := a.raiz
	for actual != nil {
		iterador.pila.Apilar(actual)
		actual = actual.izquierda
	}
	return iterador
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRecursiva(a.raiz, desde, hasta, visitar, a.comparar)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := &iterABB[K, V]{
		pila:     pila.CrearPilaDinamica[*nodoArbol[K, V]](),
		desde:    desde,
		hasta:    hasta,
		comparar: a.comparar,
	}
	actual := a.raiz
	for actual != nil {
		if desde != nil && a.comparar(actual.clave, *desde) < 0 {
			actual = actual.derecha
		} else {
			iterador.pila.Apilar(actual)
			actual = actual.izquierda
		}
	}
	return iterador
}

func (iter *iterABB[K, V]) HayAlgoMas() bool {
	if iter.pila.EstaVacia() {
		return false
	}
	if iter.hasta == nil {
		return true
	}
	actual := iter.pila.VerTope()
	return iter.comparar(actual.clave, *iter.hasta) <= 0
}

func (iter *iterABB[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pila.VerTope()
	return actual.clave, actual.dato
}

func (iter *iterABB[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pila.Desapilar()
	if actual.derecha == nil {
		return
	}
	auxiliar := actual.derecha
	for auxiliar != nil {
		iter.pila.Apilar(auxiliar)
		auxiliar = auxiliar.izquierda
	}
}
