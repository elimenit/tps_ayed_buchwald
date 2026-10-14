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

	// Ver actual devuelve el elemento apuntado por el iterador. Si el iterador ya termino de 
	// recorrer la lista, lanza un panic advirtiendo que "El iterador termino de iterar".
	VerActual() T

	// HayAlgoMas devuelve true si el iterador se encuentra sobre un elemento de la lista, o 
	// false si ya termino de recorrerla.
	HayAlgoMas() bool

	// Avanzar mueve el iterador al siguiente elemento de la lista. Si el iterador ya termino
	// de recorrerla, lanza un panic advirtiendo "El iterador termino de iterar".
	Avanzar()

	// Insertar agrega un nuevo elemento a la lista en la posicion actual del iterador.
	// Luego de la accion el iterador queda apuntando al elemento insertado.
	Insertar(T)

	// Borrar elimina el elemento actual del iterador y devuelve su dato. Luego de la accion, el
	// el elemento queda apuntando al siguiente elemento.
	// Precondicion: el iterador no debe haber terminado de recorrer la lista. En caso contrario
	// lanza un panic advirtiendo "El iterador termino de iterar".
	Borrar() T
}

