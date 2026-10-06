package coordenada

type Coordenada interface {
	ObtenerFila() int
	ObtenerColumna() int

	ActualizarFila(nuevoValor int)
	ActualizarColumna(nuevoValor int)
}
