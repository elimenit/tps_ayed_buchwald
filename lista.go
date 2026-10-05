package lista

type Lista[T any] interface {

	// Esta vacia devuelve Verdadero en caso de que la lista este vacia, False en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta el elemento pasado por parametro en el primer lugar de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta el elemento pasado por parametro en el ultimo lugar de la lista.
	InsertarUltimo(T)

	// BorrarPrimero extrae y devuelve el primer elemento de la lista. Si esta vacia, entra en panico 
	// con el mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el primer valor de la lista. Si esta vacia, entra en panico 
	// con el mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el ultimo valor de la lista. Si esta vacia, entra en panico 
	// con el mensaje "La lista esta vacia"
	VerUltimo() T

	// Largo obtiene la cantidad de elementos almacenados en la lista.
	Largo() int

	// Iterar recorre la lista desde el primer elemento hasta el ultimo, aplicando 
	// la funcion de visitar a cada dato. La iteracion termina al recorrer la lista completa
	// o cuando visitar devuelva false
	Iterar(visitar func(T) bool)

	// Iterador crea y devuelve un iterador externo ubicado sobre el primer elemento de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	//
	VerActual() T

	//
	HayAlgoMas() bool

	//
	Avanzar()

	//
	Insertar(T)

	//
	Borrar() T
}

