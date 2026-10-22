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

const (
	_PARED = '#'
	Inicio = 'S'
	Fin    = 'E'

	MovAbajo     = "ABAJO"
	MovArriba    = "ARRIBA"
	MovDerecha   = "DERECHA"
	MovIzquierda = "IZQUIERDA"
	MsjError     = "ERROR"
)

// Variable global
var PosicionInvalida = Coordenada{-1, -1}

var direcciones = []struct {
	df, dc int
	nombre string
}{
	{0, 1, MovDerecha},
	{1, 0, MovAbajo},
	{-1, 0, MovArriba},
	{0, -1, MovIzquierda},
}

type Coordenada struct {
	f, c int
}

func ResolverLaberinto() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}

		dimensiones := strings.Fields(linea)
		if len(dimensiones) < 2 {
			continue
		}

		filas, err1 := strconv.Atoi(dimensiones[0])
		columnas, err2 := strconv.Atoi(dimensiones[1])
		if err1 != nil || err2 != nil {
			continue
		}

		var laberinto [][]rune
		for i := 0; i < filas; i++ {
			if !scanner.Scan() {
				break
			}
			filaTexto := strings.TrimRight(scanner.Text(), "\r\n")
			runas := []rune(filaTexto)

			// Validamos que la fila cumpla con la cantidad de columnas informada
			if len(runas) < columnas {
				// Si faltan caracteres, se puede rellenar o descartar
				continue
			}

			laberinto = append(laberinto, runas[:columnas])
		}

		if len(laberinto) != filas {
			fmt.Println(MsjError)
			continue
		}

		i, f, err := encontrarExtremos(laberinto)
		if err != nil {
			fmt.Println(MsjError)
			continue
		}

		camino, err := bfs(laberinto, i, f)
		if err != nil {
			fmt.Println(MsjError)
			continue
		}

		movimientos := convertirAMovimientos(camino)
		fmt.Println(len(movimientos))
		fmt.Println(strings.Join(movimientos, " "))
	}
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
	if laberinto[pos.f][pos.c] == _PARED {
		return false
	}
	return true
}

// reconstruirCamino genera el slice de coordenadas desde el inicio hasta el fin
func reconstruirCamino(padres [][]Coordenada, inicio, fin Coordenada) []Coordenada {
	var caminoInvertido []Coordenada
	actual := fin

	for actual != PosicionInvalida {
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

		// Buscamos la dirección correspondiente
		for _, dir := range direcciones {
			if dir.df == df && dir.dc == dc {
				movimientos = append(movimientos, dir.nombre)
				break
			}
		}
	}
	return movimientos
}
