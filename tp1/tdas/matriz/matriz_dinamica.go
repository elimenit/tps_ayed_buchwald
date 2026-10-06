package matriz

type matriz_dinamica[T any] struct {
	grilla   [][]T
	filas    int
	columnas int
}

func CrearMatriz[T any](f int, c int) Matriz[T] {
	g := make([][]T, f)
	for i := range g {
		g[i] = make([]T, c)
	}
	return &matriz_dinamica[T]{
		grilla:   g,
		filas:    f,
		columnas: c,
	}
}

func (m *matriz_dinamica[T]) Guardar(f int, c int, valor T) {
	if f > m.filas || f < 0 || c > m.columnas || c < 0 {
		panic("Direcciones Invalidas de filas y columnas!")
	}
	m.grilla[f][c] = valor
}

func (m *matriz_dinamica[T]) Obtener(f int, c int) T {

	return m.grilla[f][c]
}

func (m *matriz_dinamica[T]) Dimensiones() (int, int) {
	return m.filas, m.columnas
}
