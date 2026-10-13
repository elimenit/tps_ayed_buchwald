package main

import (
	"bufio"
	"errors"
	"fmt"
	"maze/tdas/cola"
	"os"
	"strconv"
	"strings"
)

// Constantes para evitar literales en el código
const (
	Pared  = '#'
	Inicio = 'S'
	Fin    = 'E'

	MovAbajo     = "ABAJO"
	MovArriba    = "ARRIBA"
	MovDerecha   = "DERECHA"
	MovIzquierda = "IZQUIERDA"
	MsjError     = "ERROR"
)

var direcciones = []struct {
	df, dc int
}{
	{0, 1},  // Derecha
	{1, 0},  // Abajo
	{-1, 0}, // Arriba
	{0, -1}, // Izquierda
}

type Coordenada struct {
	f, c int
}

// leerLaberintos lee por entrada estándar y devuelve un slice de matrices bidimensionales ([][]rune)
func leerLaberintos() [][][]rune {
	var listaLaberintos [][][]rune
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}

		dimensiones := strings.Fields(linea)
		filas, _ := strconv.Atoi(dimensiones[0])

		var laberinto [][]rune
		for i := 0; i < filas; i++ {
			scanner.Scan()
			filaTexto := strings.TrimSpace(scanner.Text())
			laberinto = append(laberinto, []rune(filaTexto))
		}

		listaLaberintos = append(listaLaberintos, laberinto)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Hubo un error al leer STDIN: %v", err))
	}

	return listaLaberintos
}

// BuscarCamino orquesta la búsqueda: ubica los extremos y ejecuta BFS
func BuscarCamino(laberinto [][]rune) ([]Coordenada, error) {
	inicio, fin, err := encontrarExtremos(laberinto)
	if err != nil {
		return nil, err
	}

	return bfs(laberinto, inicio, fin)
}

// encontrarExtremos recorre el laberinto para devolver la posición S y E
func encontrarExtremos(laberinto [][]rune) (Coordenada, Coordenada, error) {
	inicio := Coordenada{-1, -1}
	fin := Coordenada{-1, -1}

	for f, fila := range laberinto {
		for c, valor := range fila {
			if valor == Inicio {
				inicio = Coordenada{f, c}
			} else if valor == Fin {
				fin = Coordenada{f, c}
			}
		}
	}

	if inicio.f == -1 || fin.f == -1 {
		return inicio, fin, errors.New("no se encontró inicio o fin en el laberinto")
	}

	return inicio, fin, nil
}

// bfs ejecuta la búsqueda en anchura pura, devuelve el camino de Coordenadas
func bfs(laberinto [][]rune, inicio Coordenada, fin Coordenada) ([]Coordenada, error) {
	filas := len(laberinto)
	columnas := len(laberinto[0])

	visitados := make([][]bool, filas)
	padres := make([][]Coordenada, filas)

	for i := 0; i < filas; i++ {
		visitados[i] = make([]bool, columnas)
		padres[i] = make([]Coordenada, columnas)
		for j := 0; j < columnas; j++ {
			padres[i][j] = Coordenada{-1, -1}
		}
	}

	q := cola.CrearColaEnlazada[Coordenada]()
	q.Encolar(inicio)
	visitados[inicio.f][inicio.c] = true

	for !q.EstaVacia() {
		actual := q.Desencolar()

		if actual == fin {
			break
		}

		for _, dir := range direcciones {
			vecino := Coordenada{actual.f + dir.df, actual.c + dir.dc}

			if esPosicionValida(laberinto, vecino, filas, columnas) && !visitados[vecino.f][vecino.c] {
				visitados[vecino.f][vecino.c] = true
				padres[vecino.f][vecino.c] = actual
				q.Encolar(vecino)
			}
		}
	}

	if !visitados[fin.f][fin.c] {
		return nil, errors.New("no hay camino posible")
	}

	return reconstruirCamino(padres, inicio, fin), nil
}

func esPosicionValida(laberinto [][]rune, pos Coordenada, filas, columnas int) bool {
	if pos.f < 0 || pos.f >= filas || pos.c < 0 || pos.c >= columnas {
		return false
	}
	if laberinto[pos.f][pos.c] == Pared {
		return false
	}
	return true
}

// reconstruirCamino genera el slice de coordenadas desde el inicio hasta el fin
func reconstruirCamino(padres [][]Coordenada, inicio, fin Coordenada) []Coordenada {
	var caminoInvertido []Coordenada
	actual := fin
	coordenadaVacia := Coordenada{-1, -1}

	for actual != coordenadaVacia {
		caminoInvertido = append(caminoInvertido, actual)
		if actual == inicio {
			break
		}
		actual = padres[actual.f][actual.c]
	}

	// Invertir el slice
	camino := make([]Coordenada, len(caminoInvertido))
	for i, j := 0, len(caminoInvertido)-1; i < len(caminoInvertido); i, j = i+1, j-1 {
		camino[i] = caminoInvertido[j]
	}

	return camino
}

// convertirAMovimientos traduce los saltos entre coordenadas consecutivas a un slice de strings direccionales
func convertirAMovimientos(camino []Coordenada) []string {
	var movimientos []string

	for i := 0; i < len(camino)-1; i++ {
		actual := camino[i]
		siguiente := camino[i+1]

		df := siguiente.f - actual.f
		dc := siguiente.c - actual.c

		if df == 1 {
			movimientos = append(movimientos, MovAbajo)
		} else if df == -1 {
			movimientos = append(movimientos, MovArriba)
		} else if dc == 1 {
			movimientos = append(movimientos, MovDerecha)
		} else if dc == -1 {
			movimientos = append(movimientos, MovIzquierda)
		}
	}
	return movimientos
}
func main() {
	laberintos := leerLaberintos()

	for _, laberinto := range laberintos {
		camino, err := BuscarCamino(laberinto)

		if err != nil {
			fmt.Println(MsjError)
		} else {
			movimientos := convertirAMovimientos(camino)
			fmt.Println(len(movimientos))
			fmt.Println(strings.Join(movimientos, " "))
		}
	}
}
