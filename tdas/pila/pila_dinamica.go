package pila

/* Definición del struct pila proporcionado por la cátedra. */
const (
	capacidadInicial = 1
	factorAumento    = 2
	factorReduccion  = 2
	margenPerdida    = factorReduccion * factorReduccion
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, capacidadInicial),
		cantidad: 0,
	}
}
func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	for i := 0; i < nuevaCapacidad/2; i++ {
		nuevosDatos[i] = pila.datos[i]
	}
	pila.datos = nuevosDatos
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	// Si el arreglo está lleno, duplicamos su capacidad
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * factorAumento)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	pila.cantidad -= 1
	dato := pila.datos[pila.cantidad]

	if pila.cantidad > 0 && pila.cantidad*margenPerdida <= cap(pila.datos) && cap(pila.datos) > capacidadInicial {
		pila.redimensionar(cap(pila.datos) / factorReduccion)
	}
	return dato
}
