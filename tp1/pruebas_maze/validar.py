#!/usr/bin/python3
import os
import sys

ERROR = "ERROR"

# Desplazamiento (fila, columna) de cada dirección.
DIRECCIONES = {
    "ARRIBA": (-1, 0),
    "ABAJO": (1, 0),
    "IZQUIERDA": (0, -1),
    "DERECHA": (0, 1),
}
PARED = "#"
START = "S"
EXIT = "E"


class ErrorValidacion(Exception):
    """Se levanta cuando la salida del alumno es incorrecta."""


def leer_laberintos(ruta):
    with open(ruta) as f:
        lineas = f.read().split("\n")

    laberintos = []
    i = 0
    n = len(lineas)
    while i < n:
        # Saltear lineas vacias entre laberintos o al final del archivo
        if lineas[i].strip() == "":
            i += 1
            continue
        partes = lineas[i].split()
        filas, cols = int(partes[0]), int(partes[1])
        i += 1
        grilla = lineas[i:i + filas]
        i += filas
        inicio = fin = None
        for r, fila in enumerate(grilla):
            for c, celda in enumerate(fila):
                if celda == START:
                    inicio = (r, c)
                elif celda == EXIT:
                    fin = (r, c)
        laberintos.append((grilla, inicio, fin, filas, cols))
    return laberintos


def leer_esperados(ruta_entrada):
    nombre = os.path.basename(ruta_entrada)
    if nombre.endswith("_in"):
        nombre = nombre[:-len("_in")]

    ruta = os.path.join(os.path.dirname(ruta_entrada) or ".", "esperados.txt")
    if not os.path.exists(ruta):
        raise ErrorValidacion(
            f"no se encuentra el archivo de esperados {ruta!r}")

    with open(ruta) as f:
        for linea in f:
            campos = linea.split()
            if not campos or campos[0] != nombre:
                continue
            return [None if c == ERROR else int(c) for c in campos[1:]]

    raise ErrorValidacion(
        f"no hay una línea para la prueba {nombre!r} en {ruta!r}")


def transitable(grilla, filas, cols, r, c):
    return 0 <= r < filas and 0 <= c < cols and grilla[r][c] != PARED


def validar_camino(grilla, inicio, fin, filas, cols, distancia, pasos):
    if len(pasos) != distancia:
        raise ErrorValidacion(
            f"la cantidad de pasos ({len(pasos)}) no coincide con la longitud "
            f"esperada ({distancia})")

    r, c = inicio
    for i, paso in enumerate(pasos):
        if paso not in DIRECCIONES:
            raise ErrorValidacion(
                f"paso {i + 1} invalido: {paso!r} (se esperaba ARRIBA, ABAJO, "
                f"IZQUIERDA o DERECHA)")
        dr, dc = DIRECCIONES[paso]
        r, c = r + dr, c + dc
        if not transitable(grilla, filas, cols, r, c):
            raise ErrorValidacion(
                f"el paso {i + 1} ({paso}) lleva a una celda intransitable o "
                f"fuera del laberinto: ({r}, {c})")

    if (r, c) != fin:
        raise ErrorValidacion(
            f"el camino termina en ({r}, {c}) y no en la celda de llegada {fin}")


def validar(ruta_entrada, ruta_salida):
    laberintos = leer_laberintos(ruta_entrada)
    esperados = leer_esperados(ruta_entrada)

    if len(esperados) != len(laberintos):
        raise ErrorValidacion(
            f"el archivo de esperados tiene {len(esperados)} longitud(es) pero "
            f"la entrada tiene {len(laberintos)} laberinto(s)")

    with open(ruta_salida) as f:
        # Conservamos líneas para poder distinguir la línea de pasos (que
        # podría estar vacía) de la ausencia de línea.
        lineas = f.read().split("\n")
    # split("\n") deja un "" final si el archivo termina en salto de línea.
    if lineas and lineas[-1] == "":
        lineas.pop()

    pos = 0
    for idx, (grilla, inicio, fin, filas, cols) in enumerate(laberintos, start=1):
        distancia = esperados[idx - 1]

        if pos >= len(lineas):
            raise ErrorValidacion(
                f"laberinto {idx}: faltan lineas en la salida")
        primera = lineas[pos].strip()
        pos += 1

        if distancia is None:
            if primera != ERROR:
                raise ErrorValidacion(
                    f"laberinto {idx}: no existe camino, se esperaba {ERROR!r} "
                    f"y se obtuvo {primera!r}")
            continue

        if primera == ERROR:
            raise ErrorValidacion(
                f"laberinto {idx}: se obtuvo {ERROR!r} pero existe un camino de "
                f"longitud {distancia}")
        try:
            declarada = int(primera)
        except ValueError:
            raise ErrorValidacion(
                f"laberinto {idx}: se esperaba un numero entero y se obtuvo "
                f"{primera!r}")
        if declarada != distancia:
            raise ErrorValidacion(
                f"laberinto {idx}: longitud del camino incorrecta. Se esperaba "
                f"{distancia} y se obtuvo {declarada}")

        # Debe haber una línea con los pasos.
        if pos >= len(lineas):
            raise ErrorValidacion(
                f"laberinto {idx}: falta la linea con los pasos del camino")
        pasos = lineas[pos].split()
        pos += 1
        validar_camino(grilla, inicio, fin, filas, cols, distancia, pasos)

    if pos != len(lineas):
        raise ErrorValidacion(
            f"la salida tiene {len(lineas) - pos} linea(s) de mas al final")


def main():
    if len(sys.argv) != 3:
        print(f"Uso: {sys.argv[0]} <ENTRADA> <SALIDA_ALUMNO>", file=sys.stderr)
        return 2
    try:
        validar(sys.argv[1], sys.argv[2])
    except ErrorValidacion as e:
        print(f"\tError: {e}")
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
