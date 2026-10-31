package cola_prioridad

type heap [T any] struct {
	datos []T
	cantidad int
	cmp func(T, T) int
}

const _CAPACIDAD_INICIAL int = 10
const _CANT_AUMENTO_REDUCCION int = 2
const _PROP_MIN_LIBRE int = 4

func downHeap(arreglo []T, pos_act int, cmp func(T, T) int){
	hijoIz := pos_act*2 + 1
	hijoDer := pos_act*2 + 2

	mayor := pos_act

	if hijoIz < len(arreglo) && cmp(arreglo[hijoIz], arreglo[hijoDer])>0{
		mayor = hijoIz
	}

	if hijoDer < len(arreglo) && cmp(arreglo[hijoIz], arreglo[hijoDer])<0{
		mayor = hijoDer
	}
	if mayor != pos_act {
		arreglo[pos_act], arreglo[mayor] = arreglo[mayor], arreglo[pos_act]
		downHeap(arreglo, mayor, cmp)
	}
}

func upHeap(arreglo []T, pos_act, cmp func(T, T) int){
	if pos_act = 0 {
		return
	}
	padre := (pos_act-1) / 2
	if cmp(arreglo[padre], arreglo[pos_act]) < 0{
		arreglo[padre], arreglo[pos_act] = arreglo[pos_act], arreglo[padre]
		upHeap(arreglo, padre, cmp func(T, T) int)
	}
}

func (h * heap[T any]) redimensionar() {
	if h.cantidad == cap(h.datos) {
		nueva := make([]T, h.cantidad, cap(h.datos)*_CANT_AUMENTO_REDUCCION)
		copy(nueva, h.datos)
		h.datos = nueva
	}
	if h.cantidad*_PROP_MIN_LIBRE <= cap(h.datos) && cap(h.datos)/_CANT_AUMENTO_REDUCCION >= _CAPACIDAD_INICIAL {
		nueva := make([]T, h.cantidad, cap(h.datos)/_CANT_AUMENTO_REDUCCION)
		copy(nueva, h.datos)
		h.datos = nueva
	}
}

func CrearHeap[T any](funcion_cmp funcion(T, T) int) ColaPrioridad[T]{
	return &ColaPrioridad[T]{
		datos: make([]T, 0, _CAPACIDAD_INICIAL),
		cantidad: 0,
		cmp: funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T]{
	for i := len(arreglo)/2 - 1; i >= 0; i-- {
		downHeap(arreglo, i, funcion_cmp)
	}
	return &ColaPrioridad[T]{
		datos: arreglo,
		cantidad: len(arreglo),
		cmp: funcion_cmp,
	}
}

func (h *heap[T any]) EstaVacia() bool { 
	return h.cantidad == 0 
}

func (h *heap[T any]) Encolar(elemento T) {
	h.redimensionar()
	if h.cantidad == len(h.datos){
		h.datos = h.datos[:h.cantidad+1]
	}
	h.datos[h.cantidad] = elemento
	h.cantidad++
	h.upHeap(h.datos, h.cantidad-1, h.cmp)
}

func (h *heap[T any]) VerMax() T {
	if h.EstaVacia(){
		panic("El heap esta vacio")
	}
	return h.arreglo[0]
}

func (h *heap[T any]) Desencolar() T {
	if h.cantidad == 0{
		panic("La cola esta vacia")
	}
	maximo := h.datos[0]
	h.datos[0], h.datos[h.cantidad-1] = h.datos[h.cantidad-1], h.datos[0]
	h.cantidad--
	if h.cantidad > 0{
		h.downHeap(h.datos, 0, h.cmp)
	}
	h.redimensionar()
	return maximo
}

func (h *heap[T any]) Cantidad() int { 
	return h.cantidad 
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int){
	
}
