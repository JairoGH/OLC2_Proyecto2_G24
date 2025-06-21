#!/bin/bash

clear  # Limpiar consola

# Archivo fuente
SOURCE="./output/programa.s"
OBJ="programa.o"
EXEC="programa"

# Ensamblar
echo "🔧 Ensamblando $SOURCE..."
aarch64-linux-gnu-as -o $OBJ $SOURCE
if [ $? -ne 0 ]; then
    echo "❌ Error al ensamblar."
    exit 1
fi

# Linkear
echo "🔗 Linkeando $OBJ..."
aarch64-linux-gnu-ld -o $EXEC $OBJ
if [ $? -ne 0 ]; then
    echo "❌ Error al linkear."
    exit 1
fi

# Ejecutar
echo "🚀 Ejecutando con QEMU..."
qemu-aarch64 ./$EXEC
