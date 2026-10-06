#!/usr/bin/env bash

set -eu
#
# Uso: ./pruebas.sh <RUTA-PROGRAMA>
#
# Corre cada prueba (archivos *.test / *_in) contra el programa y valida la
# salida con validar.py. La validación acepta cualquier camino mínimo válido,
# no una salida exacta.

PROGRAMA="$1"
BASENAME=$(basename "$PROGRAMA")

VALIDAR="./validar.py"

RET=0
OUT=$(mktemp)
trap "rm -f $OUT" EXIT

echo "Ejecución de pruebas de $BASENAME:"
echo ""

for t in *.test; do
    b=${t%.test}
    echo -n "Prueba $b: $(< $t)... "

    ret=0
    "$PROGRAMA" <"${b}_in" >"$OUT" 2>/dev/null || ret=$?

    if [[ $ret -ne 0 ]]; then
        echo "el programa abortó con código $ret."
        RET=1
        continue
    fi

    if python3 "$VALIDAR" "${b}_in" "$OUT"; then
        echo "OK."
    else
        RET=1
    fi
done

echo ""
if [[ $RET -ne 0 ]]; then
    echo "Hubo pruebas con errores."
else
    echo "Todas las pruebas pasaron."
fi

exit $RET
