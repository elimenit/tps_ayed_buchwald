package cola

type nodo[T any] struct {
	dato T
	sig  *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
	}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) verTope() T {
	return cola.ultimo.dato
}

// Si la cola esta vacia lanza panic.
func (cola *colaEnlazada[T]) lanzarPanicColavacia() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := crearNodo(dato)

	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.sig = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	cola.lanzarPanicColavacia()

	dato := cola.primero.dato

	cola.primero = cola.primero.sig

	if cola.primero == nil {
		cola.ultimo = nil
	}
	return dato
}

func (cola *colaEnlazada[T]) VerPrimero() T {

	cola.lanzarPanicColavacia()
	return cola.primero.dato
}

func crearNodo[T any](dato T) *nodo[T] {
	return &nodo[T]{
		dato: dato,
		sig:  nil,
	}
}
