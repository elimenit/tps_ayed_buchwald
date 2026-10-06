package matriz

type Matriz[T any] interface {

	// Obtiene el elemento en la posicion de la fila f y columna c.
	Obtener(f int, c int) T

	// Asigna el valor en la posicion de fila f y columna c.
	Guardar(f int, c int, valor T)

	// Devuleve las dimensiones de la Matriz (fila, columna).
	Dimensiones() (int, int)
}
