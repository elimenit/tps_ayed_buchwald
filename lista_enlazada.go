package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	largo int
	primero *nodoLista[T]
	ultimo *nodoLista[T]
}

type iteradorLista[T any] struct {
	lista *listaEnlazada[T]
	anterior *nodo[T]
	actual *nodo[T]
}

func nodoCrear[T any](dato T) *nodoLista[T] {
	return &nodolista[T]{
		dato: dato,
		prox: nil,
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T] {
		largo: 0,
		primero: nil,
		ultimo: nil,
	}
}

func (l *listaEnlazada[T]) inicializarLista(nuevo *nodoLista[T]) {
	l.primero = nuevo
	l.ultimo = nuevo
} 

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevo := nodoCrear(valor)

	if l.EstaVacia() {
		l.inicializarLista(nuevo)
	} else {
		nuevo.prox = l.primero
		l.primero = nuevo
		l.largo++
	}
}

func (l *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevo := nodoCrear(valor)

	if l.EstaVacia() {
		l.inicializarLista(nuevo)
	} else {
		l.ultimo.prox = nuevo
		l.ultimo = nuevo
		l.largo++
	}
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	primero := l.primero
	l.primero = primero.prox
	l.largo--

	if l.primero == nil {
		l.ultimo = nil
	}
	return primero.valor
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.prox
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T] {
		lista: l,
		anterior: nil,
		actual: l.primero
	}
}

func (it *iteradorLista[T]) HayAlgoMas() bool {
	return it.actual != nil
}

func (it *iteradorLista[T]) VerActual() T {
	if !it.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return it.actual.dato
}

func (it *iteradorLista[T]) Avanzar() {
	if !it.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	it.anterior = it.actual
	it.actual = it.actual.prox
}

func (it *iteradorLista[T]) Insertar(valor T) {
	if it.anterior == nil {
		it.lista.InsertarPrimero(valor)
		it.actual = it.lista.primero
		return
	}
	nuevo := nodoCrear(valor)

	nuevo.prox = it.actual
	it.anterior.prox = nuevo
	
	if it.actual == nil {
		it.lista.ultimo = nuevo
	}
	
	it.actual = nuevo
	it.lista.largo++
}

func (it *iteradorLista[T]) Borrar() T {
	if !it.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	borrado := it.actual.dato

	if it.anterior == nil {
		it.lista.primero = it.actual.prox
		it.actual = it.actual.prox

		if it.lista.primero == nil {
			it.lista.ultimo == nil
		}
	} else {
		it.anterior.prox = it.actual.prox

		if it.actual == it.lista.ultimo {
			it.lista.ultimo = it.anterior
		}
		it.actual = it.actual.prox
	}
	it.lista.largo--
	return borrado
}