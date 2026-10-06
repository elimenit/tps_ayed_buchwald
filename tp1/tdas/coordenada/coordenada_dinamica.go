package coordenada

type coordenada_dinamica struct {
	fila    int
	columna int
}

func CrearCoordenada(fila int, columna int) Coordenada {
	return &coordenada_dinamica{
		fila:    fila,
		columna: columna,
	}
}

func (cd *coordenada_dinamica) ObtenerFila() int {
	return cd.fila
}
func (cd *coordenada_dinamica) ObtenerColumna() int {
	return cd.columna
}
func (cd *coordenada_dinamica) ActualizarFila(nuevoValor int) {
	cd.fila = nuevoValor
}
func (cd *coordenada_dinamica) ActualizarColumna(nuevoValor int) {
	cd.columna = nuevoValor
}
