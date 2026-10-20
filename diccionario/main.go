package main

import TDAHash "tdas/diccionario"

func main() {
	hash := TDAHash.CrearHash[int, string]()
	hash.Guardar(10, "Hola")
	hash.Borrar(10)
}
