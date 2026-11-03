package cola_prioridad_test

import (
	"strings"
	CP "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararCadenas(cad1 string, cad2 string) int {
	return strings.Compare(cad1, cad2)
}

func TestHeap(t *testing.T) {
	cp := CP.CrearHeap[string](compararCadenas, "cadena", "otra")

	require.Equal(t, cp, 1)
}
