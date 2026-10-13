package hash

// borrar y obtener deben entrar en panico con el msj "La clave no pertenece al diccionario"
// en caso que la clave no se encuentre en el diccionario
type Diccionario[K comparable, V any] interface {

	// Guardar inserta una nueva clave con su dato asociado. Si la clave ya existe en el diccionario, actualiza su dato.
	Guardar(clave K, dato V)

	// Pertenece indica si una clave se encuentra almacenada en el diccionario.
	Pertenece(clave K) bool

	// Obtener devuelve el dato asociado a la clave indicada. Si la clave no existe en el diccionario, provoca un panic.
	Obtener(clave K) V

	// Borrar elimina la clave indicada del diccionario y devuelve su dato asociado. Si la clave no pertenece al diccionario,
	// provoca un panic.
	Borrar(clave K) V

	// Cantidad devuelve la cantidad de claves almacenadas en el diccionario.
	Cantidad() int

	//
	//Iterar(func(clave K, dato V) bool)

	//
	//Iterador() IterDiccionario[K, V]
}

type IterDiccionario[K comparable, V any] interface {
	HayAlgoMas() bool
	//de aca para abajo deben paniquear "El iterador termino de iterar"
	// si ya no quedan elementos a iterar.
	VerActual() (K, V)
	Avanzar()
}

//el iterador debe permitir recorrer todos los elementos almacenados en el hash
// sin importar el orden en el que son devueltos.
