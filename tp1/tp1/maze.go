package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"maze/tdas/cola"
	"maze/tdas/coordenada"
	"maze/tdas/matriz"
)

// desplazamiento representa un movimiento en el laberinto
type desplazamiento struct {
	df, dc int
	nombre string
}

// En Go no se pueden declarar slices/arrays como "const".
// La buena práctica es usar variables a nivel de paquete (no exportadas) y no mutarlas.
var movimientosPosibles = [...]desplazamiento{
	{0, 1, "DERECHA"},
	{1, 0, "ABAJO"},
	{-1, 0, "ARRIBA"},
	{0, -1, "IZQUIERDA"},
}

// leerLaberintos lee la entrada desde STDIN y devuelve un slice de matrices ya completadas.
func leerLaberintos() ([]matriz.Matriz[rune], error) {
	var listaLaberintos []matriz.Matriz[rune]
	scanner := bufio.NewScanner(os.Stdin)

	leyendoDimensiones := true
	var numeroFilas, numeroColumnas, filaActual int
	var laberintoActual matriz.Matriz[rune]

	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}

		if leyendoDimensiones {
			var err error
			numeroFilas, numeroColumnas, err = parsearDimensiones(linea)
			if err != nil {
				return nil, err
			}

			laberintoActual = matriz.CrearMatriz[rune](numeroFilas, numeroColumnas)
			filaActual = 0
			leyendoDimensiones = false

		} else {
			for c, char := range linea {
				if c < numeroColumnas {
					laberintoActual.Guardar(filaActual, c, char)
				}
			}
			filaActual++

			if filaActual == numeroFilas {
				listaLaberintos = append(listaLaberintos, laberintoActual)
				leyendoDimensiones = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error leyendo la entrada estándar: %w", err)
	}

	return listaLaberintos, nil
}

// parsearDimensiones extrae las filas y columnas de una línea de texto.
func parsearDimensiones(linea string) (int, int, error) {
	partes := strings.Fields(linea)
	if len(partes) != 2 {
		return 0, 0, fmt.Errorf("formato de dimensiones incorrecto: %s", linea)
	}

	filas, err := strconv.Atoi(partes[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error al convertir filas: %w", err)
	}

	columnas, err := strconv.Atoi(partes[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error al convertir columnas: %w", err)
	}

	return filas, columnas, nil
}

// encontrarExtremos busca la posición inicial (S) y final (E) en el laberinto.
func encontrarExtremos(grid matriz.Matriz[rune]) (coordenada.Coordenada, coordenada.Coordenada, error) {
	filas, columnas := grid.Dimensiones()
	start := coordenada.CrearCoordenada(-1, -1)
	end := coordenada.CrearCoordenada(-1, -1)
	encontroS, encontroE := false, false

	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			val := grid.Obtener(i, j)
			if val == 'S' {
				start = coordenada.CrearCoordenada(i, j)
				encontroS = true
			} else if val == 'E' {
				end = coordenada.CrearCoordenada(i, j)
				encontroE = true
			}
		}
	}

	if !encontroS || !encontroE {
		return start, end, fmt.Errorf("faltan los puntos de inicio (S) o fin (E) en el laberinto")
	}

	return start, end, nil
}

// bfs encuentra la ruta más corta usando Búsqueda en Anchura.
func bfs(grid matriz.Matriz[rune], start coordenada.Coordenada, end coordenada.Coordenada) []coordenada.Coordenada {
	filas, columnas := grid.Dimensiones()
	visitado := matriz.CrearMatriz[bool](filas, columnas)
	padres := matriz.CrearMatriz[coordenada.Coordenada](filas, columnas)

	// Inicializar la matriz de padres
	posicionInvalida := coordenada.CrearCoordenada(-1, -1)
	for i := 0; i < filas; i++ {
		for j := 0; j < columnas; j++ {
			padres.Guardar(i, j, posicionInvalida)
		}
	}

	queue := cola.CrearColaEnlazada[coordenada.Coordenada]()
	queue.Encolar(start)
	visitado.Guardar(start.ObtenerFila(), start.ObtenerColumna(), true)

	for !queue.EstaVacia() {
		actual := queue.Desencolar()

		if actual == end {
			break
		}

		for _, mov := range movimientosPosibles {
			nf, nc := actual.ObtenerFila()+mov.df, actual.ObtenerColumna()+mov.dc
			siguiente := coordenada.CrearCoordenada(nf, nc)

			if fueraDeLimites(nf, nc, filas, columnas) ||
				grid.Obtener(nf, nc) == '#' ||
				visitado.Obtener(nf, nc) {
				continue
			}

			visitado.Guardar(nf, nc, true)
			padres.Guardar(nf, nc, actual)
			queue.Encolar(siguiente)
		}
	}

	return reconstruirCamino(padres, start, end, visitado)
}

// fueraDeLimites es una función auxiliar que verifica si una coordenada pertenece al tablero.
func fueraDeLimites(f, c, filas, columnas int) bool {
	return f < 0 || f >= filas || c < 0 || c >= columnas
}

// reconstruirCamino toma la matriz de padres y la invierte para devolver la ruta secuencial.
func reconstruirCamino(padres matriz.Matriz[coordenada.Coordenada], start, end coordenada.Coordenada, visitado matriz.Matriz[bool]) []coordenada.Coordenada {
	if !visitado.Obtener(end.ObtenerFila(), end.ObtenerColumna()) {
		return nil
	}

	var path []coordenada.Coordenada
	actual := end
	posicionInvalida := coordenada.CrearCoordenada(-1, -1)

	for actual != posicionInvalida {
		path = append(path, actual)
		if actual == start {
			break
		}
		actual = padres.Obtener(actual.ObtenerFila(), actual.ObtenerColumna())
	}

	// Invertir el camino ya que fue construido desde el fin hacia el inicio
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// convertirAMovimientos mapea una lista de coordenadas a sus instrucciones de texto correspondientes.
func convertirAMovimientos(camino []coordenada.Coordenada) []string {
	var movimientosTexto []string

	for i := 0; i < len(camino)-1; i++ {
		actual := camino[i]
		siguiente := camino[i+1]

		df := siguiente.ObtenerFila() - actual.ObtenerFila()
		dc := siguiente.ObtenerColumna() - actual.ObtenerColumna()

		for _, mov := range movimientosPosibles {
			if mov.df == df && mov.dc == dc {
				movimientosTexto = append(movimientosTexto, mov.nombre)
				break
			}
		}
	}
	return movimientosTexto
}

// resolverLaberinto orquesta la búsqueda y la impresión de los movimientos para un único laberinto.
func resolverLaberinto(grid matriz.Matriz[rune]) {
	start, end, err := encontrarExtremos(grid)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	path := bfs(grid, start, end)
	if path == nil {
		fmt.Println("ERROR")
		return
	}

	movimientos := convertirAMovimientos(path)
	fmt.Println(len(movimientos)) // Cantidad de pasos = transiciones (len(path) - 1)

	if len(movimientos) > 0 {
		fmt.Println(strings.Join(movimientos, " "))
	}
}

func main() {
	laberintos, err := leerLaberintos()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error al procesar laberintos:", err)
		os.Exit(1)
	}

	for _, laberinto := range laberintos {
		resolverLaberinto(laberinto)
	}
}
