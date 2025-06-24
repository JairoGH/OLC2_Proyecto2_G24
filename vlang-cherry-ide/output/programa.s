.section .data
    .align 3    // alinea dobles a 8 bytes
    msg_1: .asciz "HOLA"
    msg_2: .asciz "MUNDO"
    msg_3: .asciz "V"
    msg_4: .asciz "lang"
    float_const_5: .double 2.100000
    float_const_6: .double 2.500000
    float_const_7: .double 3.140000
    float_const_8: .double 3.140000
    float_const_9: .double 2.500000
    float_const_10: .double 3.700000
    msg_11: .asciz "Hola"
    msg_12: .asciz "Hola"
    msg_13: .asciz "ABC"
    msg_14: .asciz "XYZ"
    msg_15: .asciz "Test"
    msg_16: .asciz "Prueba"
    float_const_17: .double 5.500000
    float_const_18: .double 6.700000
    float_const_19: .double 10.000000
    float_const_20: .double 4.200000
    msg_21: .asciz "🔹 DECLARACIÓN DE SLICES:\n"
    msg_22: .asciz "hola"
    msg_23: .asciz "mundo"
    msg_24: .asciz "Numeros originales: "
    msg_25: .asciz "Palabras originales: "
    msg_26: .asciz "🔹 FUNCIÓN LEN (longitud):\n"
    msg_27: .asciz "No. Numeros:"
    msg_28: .asciz "No. Palabras: "
    msg_29: .asciz "🔹 ACCESO POR ÍNDICE:\n"
    msg_30: .asciz "Elemento en Posicion: "
    msg_31: .asciz "🔹 ASIGNACIÓN POR ÍNDICE:\n"
    msg_32: .asciz "Después de numeros[1] = 99: "
    msg_33: .asciz "🔹 FUNCIÓN APPEND (agregar elementos):\n"
    msg_34: .asciz "Después de append(numeros, 77): "
    msg_35: .asciz " Cantidas de Elementos "
    msg_36: .asciz "🔹 FUNCIÓN INDEXOF (buscar elementos):\n"
    msg_37: .asciz "Posicion de No. 99 : "
    msg_38: .asciz "Posicion de No. 999 : "
    msg_39: .asciz "🔹 FUNCIÓN JOIN (concatenar con separador):\n"
    msg_40: .asciz " "
    msg_41: .asciz ""
    msg_42: .asciz "JUNTAR PALABRAS"
    msg_43: .asciz ", "
    msg_44: .asciz ""
    msg_45: .asciz "SEPARAR PALABRAS POR ',': "
    msg_46: .asciz "_"
    msg_47: .asciz ""
    msg_48: .asciz "SEPARAR PALABRAS POR '_': "
    msg_49: .asciz "\n🔹 RESUMEN FINAL:\n"
    msg_50: .asciz "Numeros finales: "
    msg_51: .asciz " (longitud: "
    msg_52: .asciz ")"
    msg_53: .asciz "Palabras finales: "
    msg_54: .asciz " (longitud: "
    msg_55: .asciz ")"
    msg_56: .asciz "================== Sentencias de Control de Flujo  ========="
    msg_57: .asciz "================== IF  ================== "
    msg_58: .asciz "Condición compleja 1 verdadera"
    msg_59: .asciz "Condición compleja 2 verdadera"
    msg_60: .asciz "Condición compleja 3 verdadera"
    msg_61: .asciz "Ninguna condición se cumplió"
    msg_62: .asciz "================== EJEMPLO 2  ================"
    msg_63: .asciz "Operaciones aritméticas 1"
    msg_64: .asciz "Operaciones aritméticas 2"
    msg_65: .asciz "Operaciones aritméticas 3"
    msg_66: .asciz "Ninguna operación aritmética se cumplió"
    msg_67: .asciz "================== EJEMPLO3  ========="
    msg_68: .asciz "Juan"
    msg_69: .asciz ""
    msg_70: .asciz "Juan"
    msg_71: .asciz "Usuario válido"
    msg_72: .asciz "Usuario inactivo"
    msg_73: .asciz "Datos incompletos"
    msg_74: .asciz "Estado desconocido"
    msg_75: .asciz "================== EJEMPLO4  ========="
    msg_76: .asciz "Nivel válido"
    msg_77: .asciz "Puntos suficientes"
    msg_78: .asciz "Jugador activo"
    msg_79: .asciz "¡Jugador experto!"
    msg_80: .asciz "Jugador intermedio"
    msg_81: .asciz "Sin vidas"
    msg_82: .asciz "Puntos insuficientes"
    msg_83: .asciz "Nivel inválido"
    msg_84: .asciz "================== EJEMPLO5  ========="
    msg_85: .asciz "Excelente - A+"
    msg_86: .asciz "Muy bueno - A"
    msg_87: .asciz "Bueno - B+"
    msg_88: .asciz "Satisfactorio - B"
    msg_89: .asciz "Regular - C+"
    msg_90: .asciz "Suficiente - C"
    msg_91: .asciz "Insuficiente - D+"
    msg_92: .asciz "Deficiente - D"
    msg_93: .asciz "Reprobado - F"
    msg_94: .asciz "================== EJEMPLO6  ========="
    msg_95: .asciz "Matemáticas avanzadas 1"
    msg_96: .asciz "Matemáticas avanzadas 2"
    msg_97: .asciz "Cálculos incorrectos"
    msg_98: .asciz "=========== SWITCH-CASE ============="
    msg_99: .asciz "======================================="
    msg_100: .asciz "=========== EJEMPLO 1: CALCULADORA BÁSICA ============="
    msg_101: .asciz "*"
    msg_102: .asciz "+"
    msg_103: .asciz "-"
    msg_104: .asciz "*"
    msg_105: .asciz "/"
    msg_106: .asciz "Suma:"
    msg_107: .asciz "Resta:"
    msg_108: .asciz "Multiplicación:"
    msg_109: .asciz "División:"
    msg_110: .asciz "Operación no válida"
    msg_111: .asciz "=========== EJEMPLO 2: DÍAS DE LA SEMANA ============="
    msg_112: .asciz "Lunes - Inicio de semana laboral"
    msg_113: .asciz "Martes - Día productivo"
    msg_114: .asciz "Miércoles - Mitad de semana"
    msg_115: .asciz "Jueves - Casi fin de semana"
    msg_116: .asciz "Viernes - ¡Por fin!"
    msg_117: .asciz "Sábado - Día de descanso"
    msg_118: .asciz "Domingo - Día familiar"
    msg_119: .asciz "Día inválido"
    msg_120: .asciz "=========== EJEMPLO 3: CALIFICACIONES ============="
    msg_121: .asciz "B"
    msg_122: .asciz "A"
    msg_123: .asciz "B"
    msg_124: .asciz "C"
    msg_125: .asciz "D"
    msg_126: .asciz "F"
    msg_127: .asciz "Excelente - 90-100 puntos"
    msg_128: .asciz "¡Felicitaciones!"
    msg_129: .asciz "Muy bien - 80-89 puntos"
    msg_130: .asciz "Buen trabajo"
    msg_131: .asciz "Regular - 70-79 puntos"
    msg_132: .asciz "Puedes mejorar"
    msg_133: .asciz "Necesita mejora - 60-69 puntos"
    msg_134: .asciz "Estudia más"
    msg_135: .asciz "Reprobado - Menos de 60 puntos"
    msg_136: .asciz "Necesitas estudiar mucho más"
    msg_137: .asciz "Calificación no válida"
    msg_138: .asciz "=========== EJEMPLO 4: ESTADO DE SISTEMA ============="
    msg_139: .asciz "Sistema ACTIVO"
    msg_140: .asciz "Modo DEBUG habilitado"
    msg_141: .asciz "Modo PRODUCCIÓN"
    msg_142: .asciz "Sistema INACTIVO"
    msg_143: .asciz "Iniciando secuencia de arranque..."
    msg_144: .asciz "=========== EJEMPLO 5: MENÚ DE OPCIONES ============="
    msg_145: .asciz "configuracion"
    msg_146: .asciz "inicio"
    msg_147: .asciz "perfil"
    msg_148: .asciz "configuracion"
    msg_149: .asciz "ayuda"
    msg_150: .asciz "salir"
    msg_151: .asciz "=== PANTALLA PRINCIPAL ==="
    msg_152: .asciz "Bienvenido al sistema"
    msg_153: .asciz "=== PERFIL DE USUARIO ==="
    msg_154: .asciz "Nombre: Juan Pérez"
    msg_155: .asciz "Email: juan@email.com"
    msg_156: .asciz "=== CONFIGURACIÓN ==="
    msg_157: .asciz "1. Cambiar idioma"
    msg_158: .asciz "2. Configurar notificaciones"
    msg_159: .asciz "3. Privacidad"
    msg_160: .asciz "=== CENTRO DE AYUDA ==="
    msg_161: .asciz "¿En qué podemos ayudarte?"
    msg_162: .asciz "=== CERRANDO SESIÓN ==="
    msg_163: .asciz "¡Hasta luego!"
    msg_164: .asciz "Opción no reconocida"
    msg_165: .asciz "Usa: inicio, perfil, configuracion, ayuda, salir"
    msg_166: .asciz "=========== EJEMPLO 6: TIPOS DE ARCHIVOS ============="
    msg_167: .asciz "pdf"
    msg_168: .asciz "txt"
    msg_169: .asciz "doc"
    msg_170: .asciz "pdf"
    msg_171: .asciz "jpg"
    msg_172: .asciz "png"
    msg_173: .asciz "mp3"
    msg_174: .asciz "mp4"
    msg_175: .asciz "Archivo de texto plano"
    msg_176: .asciz "Editor: Bloc de notas"
    msg_177: .asciz "Documento de Word"
    msg_178: .asciz "Editor: Microsoft Word"
    msg_179: .asciz "Documento PDF"
    msg_180: .asciz "Visor: Adobe Reader"
    msg_181: .asciz "Imagen JPEG"
    msg_182: .asciz "Visor: Galería de fotos"
    msg_183: .asciz "Imagen PNG"
    msg_184: .asciz "Visor: Galería de fotos"
    msg_185: .asciz "Archivo de audio"
    msg_186: .asciz "Reproductor: Reproductor de música"
    msg_187: .asciz "Archivo de video"
    msg_188: .asciz "Reproductor: Reproductor de video"
    msg_189: .asciz "Tipo de archivo no reconocido"
    msg_190: .asciz "Abrir con aplicación predeterminada"
    msg_191: .asciz "=========== EJEMPLO 7: NIVELES DE JUEGO ============="
    msg_192: .asciz "NIVEL 1: Tutorial"
    msg_193: .asciz "Aprendiendo los controles básicos"
    msg_194: .asciz "NIVEL 2: Principiante"
    msg_195: .asciz "Primeros desafíos"
    msg_196: .asciz "NIVEL 3: Intermedio"
    msg_197: .asciz "La dificultad aumenta"
    msg_198: .asciz "NIVEL 4: Avanzado"
    msg_199: .asciz "Solo para expertos"
    msg_200: .asciz "NIVEL 5: Maestro"
    msg_201: .asciz "El desafío final"
    msg_202: .asciz "Nivel no disponible"
    msg_203: .asciz "Puntos obtenidos:"
    msg_204: .asciz "=========== EJEMPLO 8: CÓDIGOS DE ERROR ============="
    msg_205: .asciz "OK - Solicitud exitosa"
    msg_206: .asciz "Bad Request - Solicitud incorrecta"
    msg_207: .asciz "Unauthorized - No autorizado"
    msg_208: .asciz "Forbidden - Acceso prohibido"
    msg_209: .asciz "Not Found - Recurso no encontrado"
    msg_210: .asciz "Verifica la URL"
    msg_211: .asciz "Internal Server Error - Error del servidor"
    msg_212: .asciz "Contacta al administrador"
    msg_213: .asciz "Service Unavailable - Servicio no disponible"
    msg_214: .asciz "Intenta más tarde"
    msg_215: .asciz "Código de error desconocido"
    msg_216: .asciz "=========== EJEMPLO 9: COMANDOS DE TERMINAL ============="
    msg_217: .asciz "ls"
    msg_218: .asciz "ls"
    msg_219: .asciz "pwd"
    msg_220: .asciz "cd"
    msg_221: .asciz "mkdir"
    msg_222: .asciz "rm"
    msg_223: .asciz "help"
    msg_224: .asciz "exit"
    msg_225: .asciz "Listando archivos del directorio:"
    msg_226: .asciz "archivo1.txt"
    msg_227: .asciz "archivo2.pdf"
    msg_228: .asciz "carpeta1/"
    msg_229: .asciz "Directorio actual:"
    msg_230: .asciz "/home/usuario/documentos"
    msg_231: .asciz "Cambiar directorio"
    msg_232: .asciz "Uso: cd <directorio>"
    msg_233: .asciz "Crear directorio"
    msg_234: .asciz "Uso: mkdir <nombre>"
    msg_235: .asciz "Eliminar archivo"
    msg_236: .asciz "¡CUIDADO! Esta acción es irreversible"
    msg_237: .asciz "Comandos disponibles:"
    msg_238: .asciz "ls, pwd, cd, mkdir, rm, help, exit"
    msg_239: .asciz "Cerrando terminal..."
    msg_240: .asciz "Comando no reconocido"
    msg_241: .asciz "Escribe 'help' para ver comandos disponibles"
    msg_242: .asciz "=========== EJEMPLO 10: SIN DEFAULT ============="
    msg_243: .asciz "verde"
    msg_244: .asciz "rojo"
    msg_245: .asciz "azul"
    msg_246: .asciz "amarillo"
    msg_247: .asciz "verde"
    msg_248: .asciz "negro"
    msg_249: .asciz "blanco"
    msg_250: .asciz "Color cálido y energético"
    msg_251: .asciz "Color frío y tranquilo"
    msg_252: .asciz "Color brillante y alegre"
    msg_253: .asciz "Color de la naturaleza"
    msg_254: .asciz "Ausencia de color"
    msg_255: .asciz "Presencia de todos los colores"
    msg_256: .asciz "Análisis de color completado"
    msg_257: .asciz "=========== EJEMPLO 11: SWITCH CON VARIABLES ============="
    msg_258: .asciz ""
    msg_259: .asciz "Recién nacido"
    msg_260: .asciz "Bebé"
    msg_261: .asciz "Niño pequeño"
    msg_262: .asciz "Adolescente"
    msg_263: .asciz "Joven adulto"
    msg_264: .asciz "Adulto"
    msg_265: .asciz "En la flor de la vida"
    msg_266: .asciz "Adulto mayor"
    msg_267: .asciz "Edad no categorizada"
    msg_268: .asciz "Categoría de Edad:"
    msg_269: .asciz "=========== EJEMPLO 12: SWITCH ANIDADO COMPLEJO ============="
    msg_270: .asciz "admin"
    msg_271: .asciz "escribir"
    msg_272: .asciz "admin"
    msg_273: .asciz "editor"
    msg_274: .asciz "usuario"
    msg_275: .asciz "Usuario ADMINISTRADOR"
    msg_276: .asciz "leer"
    msg_277: .asciz "escribir"
    msg_278: .asciz "eliminar"
    msg_279: .asciz "✅ Permiso de lectura concedido"
    msg_280: .asciz "✅ Permiso de escritura concedido"
    msg_281: .asciz "✅ Permiso de eliminación concedido"
    msg_282: .asciz "❌ Permiso no reconocido"
    msg_283: .asciz "Usuario EDITOR"
    msg_284: .asciz "leer"
    msg_285: .asciz "escribir"
    msg_286: .asciz "eliminar"
    msg_287: .asciz "✅ Permiso de lectura concedido"
    msg_288: .asciz "✅ Permiso de escritura concedido"
    msg_289: .asciz "❌ Permiso de eliminación DENEGADO"
    msg_290: .asciz "❌ Permiso no reconocido"
    msg_291: .asciz "Usuario ESTÁNDAR"
    msg_292: .asciz "leer"
    msg_293: .asciz "escribir"
    msg_294: .asciz "eliminar"
    msg_295: .asciz "✅ Permiso de lectura concedido"
    msg_296: .asciz "❌ Permiso de escritura DENEGADO"
    msg_297: .asciz "❌ Permiso de eliminación DENEGADO"
    msg_298: .asciz "❌ Permiso no reconocido"
    msg_299: .asciz "❌ Tipo de usuario no válido"
    msg_300: .asciz "=========== FIN DE EJEMPLOS ============="
    msg_301: .asciz "¡Switch-case implementado exitosamente!"
buffer_int: .skip 32
buffer_float: .skip 64
buffer_string: .skip 512
temp_buffer: .skip 256
msg_nl: .asciz "\n"
msg_menos: .asciz "-"
msg_punto: .asciz "."
const_100: .double 100.0
str_empty: .asciz ""

.section .text
.global _start

_start:
// === DEBUG: EXPRESIÓN BINARIA + ===
// Operador: +, Tipos: string + string
// === CONCATENACIÓN DE STRINGS ===
    adr x9, msg_1
    adr x10, msg_2
    adr x11, buffer_string
// Concatenar strings: x11 = x9 + x10
    mov x0, x9
    mov x1, x10
    mov x2, x11
    bl concat_strings
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR STRING  ===
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA + ===
// Operador: +, Tipos: string + string
// === CONCATENACIÓN DE STRINGS ===
    adr x9, msg_3
    adr x10, msg_4
    adr x11, buffer_string
// Concatenar strings: x11 = x9 + x10
    mov x0, x9
    mov x1, x10
    mov x2, x11
    bl concat_strings
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR STRING  ===
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA + ===
// === DEBUG: EXPRESIÓN BINARIA * ===
// === DEBUG: EXPRESIÓN BINARIA - ===
// Operador: -, Tipos: int - int
    mov x9, #4
    mov x10, #2
    sub x11, x9, x10
// === FIN EXPRESIÓN BINARIA ===
// Operador: *, Tipos: int * int
    mov x12, #20
    mov x13, x11
    mul x14, x12, x13
// === FIN EXPRESIÓN BINARIA ===
// Operador: +, Tipos: int + int
    mov x15, #1
    mov x16, x14
    add x17, x15, x16
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR INT  ===
    mov x0, x17
    bl print_int
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA - ===
// === DEBUG: EXPRESIÓN BINARIA * ===
// Operador: *, Tipos: int * int
    mov x9, #2
    mov x10, #10
    mul x11, x9, x10
// === FIN EXPRESIÓN BINARIA ===
// Operador: -, Tipos: int - float
    scvtf d0, x11
    adr x12, float_const_5
    ldr d1, [x12]
    fsub d2, d0, d1
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR FLOAT  ===
    fmov d0, d2
    bl print_float
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA / ===
// Operador: /, Tipos: int / float
    mov x9, #20
    scvtf d0, x9
    adr x10, float_const_6
    ldr d1, [x10]
    fdiv d2, d0, d1
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR FLOAT  ===
    fmov d0, d2
    bl print_float
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA - ===
// Operador: -, Tipos: int - int
    mov x9, #10
    mov x10, #2
    sub x11, x9, x10
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR INT  ===
    mov x0, x11
    bl print_int
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA % ===
// Operador: %, Tipos: int % int
    mov x9, #11
    mov x10, #2
    udiv x11, x9, x10
    msub x11, x11, x10, x9
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR INT  ===
    mov x0, x11
    bl print_int
    bl print_newline
// === LITERAL BOOL: true ===
// === IMPRIMIR BOOL  ===
    mov x9, #1
    mov x0, x9
    bl print_bool
    bl print_newline
// === LITERAL BOOL: false ===
// === IMPRIMIR BOOL  ===
    mov x9, #0
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA == ===
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x10, #5
    mov x11, #5
    cmp x10, x11
    cset x9, eq
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA != ===
// Operador: !=, Tipos: int != int
// === OPERACIÓN RELACIONAL != ===
    mov x10, #5
    mov x11, #3
    cmp x10, x11
    cset x9, ne
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA == ===
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x10, #10
    mov x11, #3
    cmp x10, x11
    cset x9, eq
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA == ===
// Operador: ==, Tipos: float == float
// === OPERACIÓN RELACIONAL == ===
    adr x10, float_const_7
    ldr d0, [x10]
    adr x11, float_const_8
    ldr d1, [x11]
    fcmp d0, d1
    cset x9, eq
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA != ===
// Operador: !=, Tipos: float != float
// === OPERACIÓN RELACIONAL != ===
    adr x10, float_const_9
    ldr d0, [x10]
    adr x11, float_const_10
    ldr d1, [x11]
    fcmp d0, d1
    cset x9, ne
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA == ===
// Operador: ==, Tipos: string == string
// === OPERACIÓN RELACIONAL == ===
    adr x10, msg_11
    adr x11, msg_12
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x9, x0
    cmp x9, #0
    cset x9, eq
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA != ===
// Operador: !=, Tipos: string != string
// === OPERACIÓN RELACIONAL != ===
    adr x10, msg_13
    adr x11, msg_14
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x9, x0
    cmp x9, #0
    cset x9, ne
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA == ===
// Operador: ==, Tipos: string == string
// === OPERACIÓN RELACIONAL == ===
    adr x10, msg_15
    adr x11, msg_16
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x9, x0
    cmp x9, #0
    cset x9, eq
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA == ===
// === LITERAL BOOL: true ===
// === LITERAL BOOL: true ===
// Operador: ==, Tipos: bool == bool
// === OPERACIÓN RELACIONAL == ===
    mov x10, #1
    mov x11, #1
    cmp x10, x11
    cset x9, eq
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA != ===
// === LITERAL BOOL: false ===
// === LITERAL BOOL: true ===
// Operador: !=, Tipos: bool != bool
// === OPERACIÓN RELACIONAL != ===
    mov x10, #0
    mov x11, #1
    cmp x10, x11
    cset x9, ne
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA < ===
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x10, #50
    mov x11, #10
    cmp x10, x11
    cset x9, lt
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA > ===
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x10, #15
    mov x11, #10
    cmp x10, x11
    cset x9, gt
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA <= ===
// Operador: <=, Tipos: int <= int
// === OPERACIÓN RELACIONAL <= ===
    mov x10, #10
    mov x11, #10
    cmp x10, x11
    cset x9, le
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA >= ===
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x10, #10
    mov x11, #5
    cmp x10, x11
    cset x9, ge
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA < ===
// Operador: <, Tipos: int < float
// === OPERACIÓN RELACIONAL < ===
    mov x10, #5
    scvtf d0, x10
    adr x11, float_const_17
    ldr d1, [x11]
    fcmp d0, d1
    cset x9, lt
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA > ===
// Operador: >, Tipos: float > int
// === OPERACIÓN RELACIONAL > ===
    adr x10, float_const_18
    ldr d0, [x10]
    mov x11, #6
    scvtf d1, x11
    fcmp d0, d1
    cset x9, gt
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA >= ===
// Operador: >=, Tipos: int >= float
// === OPERACIÓN RELACIONAL >= ===
    mov x10, #10
    scvtf d0, x10
    adr x11, float_const_19
    ldr d1, [x11]
    fcmp d0, d1
    cset x9, ge
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA <= ===
// Operador: <=, Tipos: float <= int
// === OPERACIÓN RELACIONAL <= ===
    adr x10, float_const_20
    ldr d0, [x10]
    mov x11, #5
    scvtf d1, x11
    fcmp d0, d1
    cset x9, le
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA && ===
// === LITERAL BOOL: true ===
// === LITERAL BOOL: true ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x10, #1
    cmp x10, #0
    beq .Land_false_0
    mov x11, #1
    and x9, x10, x11
    b .Land_end_1
.Land_false_0:
    mov x9, #0
.Land_end_1:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA && ===
// === LITERAL BOOL: true ===
// === LITERAL BOOL: false ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x10, #1
    cmp x10, #0
    beq .Land_false_2
    mov x11, #0
    and x9, x10, x11
    b .Land_end_3
.Land_false_2:
    mov x9, #0
.Land_end_3:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA && ===
// === LITERAL BOOL: false ===
// === LITERAL BOOL: true ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x10, #0
    cmp x10, #0
    beq .Land_false_4
    mov x11, #1
    and x9, x10, x11
    b .Land_end_5
.Land_false_4:
    mov x9, #0
.Land_end_5:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA && ===
// === LITERAL BOOL: false ===
// === LITERAL BOOL: false ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x10, #0
    cmp x10, #0
    beq .Land_false_6
    mov x11, #0
    and x9, x10, x11
    b .Land_end_7
.Land_false_6:
    mov x9, #0
.Land_end_7:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA || ===
// === LITERAL BOOL: true ===
// === LITERAL BOOL: true ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x10, #1
    cmp x10, #1
    beq .Lor_true_8
    mov x11, #1
    orr x9, x10, x11
    b .Lor_end_9
.Lor_true_8:
    mov x9, #1
.Lor_end_9:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA || ===
// === LITERAL BOOL: true ===
// === LITERAL BOOL: false ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x10, #1
    cmp x10, #1
    beq .Lor_true_10
    mov x11, #0
    orr x9, x10, x11
    b .Lor_end_11
.Lor_true_10:
    mov x9, #1
.Lor_end_11:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA || ===
// === LITERAL BOOL: false ===
// === LITERAL BOOL: true ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x10, #0
    cmp x10, #1
    beq .Lor_true_12
    mov x11, #1
    orr x9, x10, x11
    b .Lor_end_13
.Lor_true_12:
    mov x9, #1
.Lor_end_13:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA || ===
// === LITERAL BOOL: false ===
// === LITERAL BOOL: false ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x10, #0
    cmp x10, #1
    beq .Lor_true_14
    mov x11, #0
    orr x9, x10, x11
    b .Lor_end_15
.Lor_true_14:
    mov x9, #1
.Lor_end_15:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === INICIO EXPRESIÓN UNARIA ===
// === LITERAL BOOL: true ===
// Operador unario: !
// === NEGACIÓN LÓGICA ! ===
    mov x9, #1
    eor x10, x9, #1
// === FIN EXPRESIÓN UNARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x10
    bl print_bool
    bl print_newline
    bl print_newline
// === INICIO EXPRESIÓN UNARIA ===
// === LITERAL BOOL: false ===
// Operador unario: !
// === NEGACIÓN LÓGICA ! ===
    mov x9, #0
    eor x10, x9, #1
// === FIN EXPRESIÓN UNARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x10
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x10, #5
    mov x11, #3
    cmp x10, x11
    cset x9, gt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x13, #2
    mov x14, #4
    cmp x13, x14
    cset x12, lt
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x16, x9
    cmp x16, #0
    beq .Land_false_16
    mov x17, x12
    and x15, x16, x17
    b .Land_end_17
.Land_false_16:
    mov x15, #0
.Land_end_17:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x15
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA || ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x10, #5
    mov x11, #3
    cmp x10, x11
    cset x9, lt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x13, #2
    mov x14, #4
    cmp x13, x14
    cset x12, lt
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x16, x9
    cmp x16, #1
    beq .Lor_true_18
    mov x17, x12
    orr x15, x16, x17
    b .Lor_end_19
.Lor_true_18:
    mov x15, #1
.Lor_end_19:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x15
    bl print_bool
    bl print_newline
    bl print_newline
// === INICIO EXPRESIÓN UNARIA ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x10, #5
    mov x11, #10
    cmp x10, x11
    cset x9, gt
// === FIN EXPRESIÓN BINARIA ===
// Operador unario: !
// === NEGACIÓN LÓGICA ! ===
    mov x12, x9
    eor x13, x12, #1
// === FIN EXPRESIÓN UNARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x13
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA && ===
// === LITERAL BOOL: false ===
// === LITERAL BOOL: true ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x10, #0
    cmp x10, #0
    beq .Land_false_20
    mov x11, #1
    and x9, x10, x11
    b .Land_end_21
.Land_false_20:
    mov x9, #0
.Land_end_21:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA || ===
// === LITERAL BOOL: true ===
// === LITERAL BOOL: false ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x10, #1
    cmp x10, #1
    beq .Lor_true_22
    mov x11, #0
    orr x9, x10, x11
    b .Lor_end_23
.Lor_true_22:
    mov x9, #1
.Lor_end_23:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x9
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x10, #5
    mov x11, #3
    cmp x10, x11
    cset x9, gt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x13, #2
    mov x14, #4
    cmp x13, x14
    cset x12, lt
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x16, x9
    cmp x16, #0
    beq .Land_false_24
    mov x17, x12
    and x15, x16, x17
    b .Land_end_25
.Land_false_24:
    mov x15, #0
.Land_end_25:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x15
    bl print_bool
    bl print_newline
    bl print_newline
// === DEBUG: EXPRESIÓN BINARIA || ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x10, #5
    mov x11, #3
    cmp x10, x11
    cset x9, lt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x13, #2
    mov x14, #4
    cmp x13, x14
    cset x12, lt
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x16, x9
    cmp x16, #1
    beq .Lor_true_26
    mov x17, x12
    orr x15, x16, x17
    b .Lor_end_27
.Lor_true_26:
    mov x15, #1
.Lor_end_27:
// === FIN EXPRESIÓN BINARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x15
    bl print_bool
    bl print_newline
    bl print_newline
// === INICIO EXPRESIÓN UNARIA ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x10, #5
    mov x11, #10
    cmp x10, x11
    cset x9, gt
// === FIN EXPRESIÓN BINARIA ===
// Operador unario: !
// === NEGACIÓN LÓGICA ! ===
    mov x12, x9
    eor x13, x12, #1
// === FIN EXPRESIÓN UNARIA ===
// === IMPRIMIR BOOL  ===
    mov x0, x13
    bl print_bool
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_21
    mov x0, x9
    bl print_string
    bl print_newline
// === DECLARACIÓN SLICE numeros: []int ===
// Ajustar stack para slices: 256 bytes
    sub sp, sp, #256
// Elemento 0 (literal): 10
    mov x0, #10
    str x0, [sp, #0]
// Elemento 1 (literal): 20
    mov x0, #20
    str x0, [sp, #8]
// Elemento 2 (literal): 30
    mov x0, #30
    str x0, [sp, #16]
// === DECLARACIÓN SLICE palabras: []string ===
// Elemento 0 (literal): hola
    adr x0, msg_22
    str x0, [sp, #24]
// Elemento 1 (literal): mundo
    adr x0, msg_23
    str x0, [sp, #32]
// === IMPRIMIR STRING  ===
    adr x9, msg_24
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #0]    // cargar int[0]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #8]    // cargar int[1]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #16]    // cargar int[2]
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_25
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #24]    // cargar string[0]
    bl print_string
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #32]    // cargar string[1]
    bl print_string
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_26
    mov x0, x9
    bl print_string
    bl print_newline
// === FUNCIÓN len(numeros) ===
// Slice numeros tiene 3 elementos
    mov x9, #3
// === IMPRIMIR STRING  ===
    adr x10, msg_27
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === FUNCIÓN len(palabras) ===
// Slice palabras tiene 2 elementos
    mov x9, #2
// === IMPRIMIR STRING  ===
    adr x10, msg_28
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_29
    mov x0, x9
    bl print_string
    bl print_newline
// === ACCESO POR ÍNDICE numeros[1] ===
    ldr x9, [sp, #8]    // cargar numeros[1]
// === IMPRIMIR STRING  ===
    adr x10, msg_30
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_31
    mov x0, x9
    bl print_string
    bl print_newline
// === ACCESO POR ÍNDICE numeros[1] ===
    ldr x9, [sp, #8]    // cargar numeros[1]
// === ASIGNACIÓN POR ÍNDICE numeros[1] = 99 ===
// Asignación directa al elemento 1
    mov x10, #99
    str x10, [sp, #8]    // numeros[1] = 99
// Asignación a slice completada exitosamente
// === ACCESO POR ÍNDICE numeros[0] ===
    ldr x9, [sp, #0]    // cargar numeros[0]
// === IMPRIMIR STRING  ===
    adr x10, msg_32
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_33
    mov x0, x9
    bl print_string
    bl print_newline
// === FUNCIÓN append(numeros, 77) ===
// Expandiendo slice numeros de 3 a 4 elementos
// Copiando 3 elementos existentes
    ldr x18, [sp, #0]    // cargar elemento 0 original
    str x18, [sp, #40]    // guardar elemento 0 en nueva posición
    ldr x18, [sp, #8]    // cargar elemento 1 original
    str x18, [sp, #48]    // guardar elemento 1 en nueva posición
    ldr x18, [sp, #16]    // cargar elemento 2 original
    str x18, [sp, #56]    // guardar elemento 2 en nueva posición
// Agregando nuevo elemento en offset 64
    mov x0, #77
    str x0, [sp, #64]
// append completado: numeros ahora tiene 4 elementos
// === IMPRIMIR STRING  ===
    adr x9, msg_34
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #40]    // cargar int[0]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #48]    // cargar int[1]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #56]    // cargar int[2]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #64]    // cargar int[3]
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === FUNCIÓN len(numeros) ===
// Slice numeros tiene 4 elementos
    mov x9, #4
// === IMPRIMIR STRING  ===
    adr x10, msg_35
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_36
    mov x0, x9
    bl print_string
    bl print_newline
// === FUNCIÓN indexOf(numeros, 99) ===
// Buscando 99 en slice numeros (tamaño: 4)
    mov x12, #99
    mov x9, #0
indexOf_loop_numeros_29:
    cmp x9, #4
    bge indexOf_not_found_numeros_31
    add x10, sp, #40
    ldr x10, [x10, x9, lsl #3]
    cmp x10, x12
    beq indexOf_found_numeros_30
    add x9, x9, #1
    b indexOf_loop_numeros_29
indexOf_found_numeros_30:
    mov x11, x9
    b indexOf_end_numeros_32
indexOf_not_found_numeros_31:
    mov x11, #0
    sub x11, x11, #1
indexOf_end_numeros_32:
// indexOf completado, resultado en x11
// === IMPRIMIR STRING  ===
    adr x9, msg_37
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x11
    bl print_int
    bl print_newline
// === FUNCIÓN indexOf(numeros, 999) ===
// Buscando 999 en slice numeros (tamaño: 4)
    mov x12, #999
    mov x9, #0
indexOf_loop_numeros_33:
    cmp x9, #4
    bge indexOf_not_found_numeros_35
    add x10, sp, #40
    ldr x10, [x10, x9, lsl #3]
    cmp x10, x12
    beq indexOf_found_numeros_34
    add x9, x9, #1
    b indexOf_loop_numeros_33
indexOf_found_numeros_34:
    mov x11, x9
    b indexOf_end_numeros_36
indexOf_not_found_numeros_35:
    mov x11, #0
    sub x11, x11, #1
indexOf_end_numeros_36:
// indexOf completado, resultado en x11
// === IMPRIMIR STRING  ===
    adr x9, msg_38
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x11
    bl print_int
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_39
    mov x0, x9
    bl print_string
    bl print_newline
// === FUNCIÓN join(palabras, " ") ===
// Uniendo 2 elementos con separador
// Inicializar registros para join
    adr x11, msg_40
    adr x12, buffer_string
    adr x13, temp_buffer
    ldr x10, [sp, #24]
    mov x0, x10
    adr x1, msg_41
    mov x2, x12
    bl concat_strings
    mov x9, #1
join_loop_palabras_37:
    cmp x9, #2
    bge join_end_palabras_38
    mov x0, x12
    mov x1, x11
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x15, sp, #24
    ldr x10, [x15, x9, lsl #3]
    mov x0, x12
    mov x1, x10
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x9, x9, #1
    b join_loop_palabras_37
join_end_palabras_38:
// join completado, resultado en x12
// === IMPRIMIR STRING  ===
    adr x9, msg_42
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x12
    bl print_string
    bl print_newline
// === FUNCIÓN join(palabras, ", ") ===
// Uniendo 2 elementos con separador
// Inicializar registros para join
    adr x11, msg_43
    adr x12, buffer_string
    adr x13, temp_buffer
    ldr x10, [sp, #24]
    mov x0, x10
    adr x1, msg_44
    mov x2, x12
    bl concat_strings
    mov x9, #1
join_loop_palabras_39:
    cmp x9, #2
    bge join_end_palabras_40
    mov x0, x12
    mov x1, x11
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x15, sp, #24
    ldr x10, [x15, x9, lsl #3]
    mov x0, x12
    mov x1, x10
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x9, x9, #1
    b join_loop_palabras_39
join_end_palabras_40:
// join completado, resultado en x12
// === IMPRIMIR STRING  ===
    adr x9, msg_45
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x12
    bl print_string
    bl print_newline
// === FUNCIÓN join(palabras, "_") ===
// Uniendo 2 elementos con separador
// Inicializar registros para join
    adr x11, msg_46
    adr x12, buffer_string
    adr x13, temp_buffer
    ldr x10, [sp, #24]
    mov x0, x10
    adr x1, msg_47
    mov x2, x12
    bl concat_strings
    mov x9, #1
join_loop_palabras_41:
    cmp x9, #2
    bge join_end_palabras_42
    mov x0, x12
    mov x1, x11
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x15, sp, #24
    ldr x10, [x15, x9, lsl #3]
    mov x0, x12
    mov x1, x10
    mov x2, x13
    bl concat_strings
    mov x14, x12
    mov x12, x13
    mov x13, x14
    add x9, x9, #1
    b join_loop_palabras_41
join_end_palabras_42:
// join completado, resultado en x12
// === IMPRIMIR STRING  ===
    adr x9, msg_48
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    mov x0, x12
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_49
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_50
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #40]    // cargar int[0]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #48]    // cargar int[1]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #56]    // cargar int[2]
    bl print_int
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #64]    // cargar int[3]
    bl print_int
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === FUNCIÓN len(numeros) ===
// Slice numeros tiene 4 elementos
    mov x9, #4
// === IMPRIMIR STRING  ===
    adr x10, msg_51
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_52
    mov x0, x11
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_53
    mov x0, x9
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR SLICE_NAME  ===
    mov x0, #91          // '['
    bl print_char
    ldr x0, [sp, #24]    // cargar string[0]
    bl print_string
    mov x0, #44          // ','
    bl print_char
    mov x0, #32          // ' '
    bl print_char
    ldr x0, [sp, #32]    // cargar string[1]
    bl print_string
    mov x0, #93          // ']'
    bl print_char
    bl print_newline
// === FUNCIÓN len(palabras) ===
// Slice palabras tiene 2 elementos
    mov x9, #2
// === IMPRIMIR STRING  ===
    adr x10, msg_54
    mov x0, x10
    bl print_string
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    mov x0, #32          // ' ' (espacio)
    bl print_char
// === IMPRIMIR STRING  ===
    adr x11, msg_55
    mov x0, x11
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_56
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_57
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: a ===
// === DECLARAR VARIABLE: a int ===
    sub sp, sp, #8  // Reservar espacio para a
    mov x9, #10
    str x9, [sp]
// Variable a declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut a := ? (int)
// === DECLARAR VARIABLE MUT: b ===
// === DECLARAR VARIABLE: b int ===
    sub sp, sp, #8  // Reservar espacio para b
    mov x9, #5
    str x9, [sp]
// Variable b declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut b := ? (int)
// === DECLARAR VARIABLE MUT: c ===
// === DECLARAR VARIABLE: c int ===
    sub sp, sp, #8  // Reservar espacio para c
    mov x9, #15
    str x9, [sp]
// Variable c declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut c := ? (int)
// === DECLARAR VARIABLE MUT: d ===
// === DECLARAR VARIABLE: d int ===
    sub sp, sp, #8  // Reservar espacio para d
    mov x9, #0
    str x9, [sp]
// Variable d declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut d := ? (int)
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// === ACCESO A VARIABLE: a ===
// === CARGAR VARIABLE: a (offset: 24) ===
    ldr x9, [sp, #24]
// Variable cargada exitosamente: a
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x11, x9
    mov x12, #0
    cmp x11, x12
    cset x10, gt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// === ACCESO A VARIABLE: b ===
// === CARGAR VARIABLE: b (offset: 16) ===
    ldr x13, [sp, #16]
// Variable cargada exitosamente: b
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x15, x13
    mov x16, #0
    cmp x15, x16
    cset x14, gt
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x18, x10
    cmp x18, #0
    beq .Land_false_43
    mov x19, x14
    and x17, x18, x19
    b .Land_end_44
.Land_false_43:
    mov x17, #0
.Land_end_44:
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA || ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// === ACCESO A VARIABLE: c ===
// === CARGAR VARIABLE: c (offset: 8) ===
    ldr x20, [sp, #8]
// Variable cargada exitosamente: c
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x22, x20
    mov x23, #10
    cmp x22, x23
    cset x21, gt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA == ===
// === ACCESO A VARIABLE: d ===
// === CARGAR VARIABLE: d (offset: 0) ===
    ldr x24, [sp, #0]
// Variable cargada exitosamente: d
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x26, x24
    mov x27, #0
    cmp x26, x27
    cset x25, eq
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x29, x21
    cmp x29, #1
    beq .Lor_true_45
    mov x30, x25
    orr x28, x29, x30
    b .Lor_end_46
.Lor_true_45:
    mov x28, #1
.Lor_end_46:
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
// ADVERTENCIA: Reciclando registros temporales
    mov x10, x17
    cmp x10, #0
    beq .Land_false_47
    mov x11, x28
    and x9, x10, x11
    b .Land_end_48
.Land_false_47:
    mov x9, #0
.Land_end_48:
// === FIN EXPRESIÓN BINARIA ===
    cmp x9, #0
    beq .Lskip_branch_49
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x12, msg_58
    mov x0, x12
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_42
.Lskip_branch_49:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #2 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA || ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// === ACCESO A VARIABLE: a ===
// === CARGAR VARIABLE: a (offset: 24) ===
    ldr x13, [sp, #24]
// Variable cargada exitosamente: a
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x15, x13
    mov x16, #0
    cmp x15, x16
    cset x14, lt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// === ACCESO A VARIABLE: b ===
// === CARGAR VARIABLE: b (offset: 16) ===
    ldr x17, [sp, #16]
// Variable cargada exitosamente: b
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x19, x17
    mov x20, #0
    cmp x19, x20
    cset x18, lt
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x22, x14
    cmp x22, #1
    beq .Lor_true_50
    mov x23, x18
    orr x21, x22, x23
    b .Lor_end_51
.Lor_true_50:
    mov x21, #1
.Lor_end_51:
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// === ACCESO A VARIABLE: c ===
// === CARGAR VARIABLE: c (offset: 8) ===
    ldr x24, [sp, #8]
// Variable cargada exitosamente: c
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x26, x24
    mov x27, #5
    cmp x26, x27
    cset x25, lt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// === ACCESO A VARIABLE: d ===
// === CARGAR VARIABLE: d (offset: 0) ===
    ldr x28, [sp, #0]
// Variable cargada exitosamente: d
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
// ADVERTENCIA: Reciclando registros temporales
    mov x30, x28
    mov x9, #0
    cmp x30, x9
    cset x29, gt
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x11, x25
    cmp x11, #0
    beq .Land_false_52
    mov x12, x29
    and x10, x11, x12
    b .Land_end_53
.Land_false_52:
    mov x10, #0
.Land_end_53:
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x14, x21
    cmp x14, #0
    beq .Land_false_54
    mov x15, x10
    and x13, x14, x15
    b .Land_end_55
.Land_false_54:
    mov x13, #0
.Land_end_55:
// === FIN EXPRESIÓN BINARIA ===
    cmp x13, #0
    beq .Lskip_branch_56
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x16, msg_59
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_42
.Lskip_branch_56:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #3 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === INICIO EXPRESIÓN UNARIA ===
// === DEBUG: EXPRESIÓN BINARIA == ===
// === ACCESO A VARIABLE: a ===
// === CARGAR VARIABLE: a (offset: 24) ===
    ldr x17, [sp, #24]
// Variable cargada exitosamente: a
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x19, x17
    mov x20, #0
    cmp x19, x20
    cset x18, eq
// === FIN EXPRESIÓN BINARIA ===
// Operador unario: !
// === NEGACIÓN LÓGICA ! ===
    mov x21, x18
    eor x22, x21, #1
// === FIN EXPRESIÓN UNARIA ===
// === DEBUG: EXPRESIÓN BINARIA || ===
// === DEBUG: EXPRESIÓN BINARIA != ===
// === ACCESO A VARIABLE: b ===
// === CARGAR VARIABLE: b (offset: 16) ===
    ldr x23, [sp, #16]
// Variable cargada exitosamente: b
// Operador: !=, Tipos: int != int
// === OPERACIÓN RELACIONAL != ===
    mov x25, x23
    mov x26, #5
    cmp x25, x26
    cset x24, ne
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: c ===
// === CARGAR VARIABLE: c (offset: 8) ===
    ldr x27, [sp, #8]
// Variable cargada exitosamente: c
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x29, x27
    mov x30, #15
    cmp x29, x30
    cset x28, ge
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
// ADVERTENCIA: Reciclando registros temporales
    mov x10, x24
    cmp x10, #1
    beq .Lor_true_57
    mov x11, x28
    orr x9, x10, x11
    b .Lor_end_58
.Lor_true_57:
    mov x9, #1
.Lor_end_58:
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x13, x22
    cmp x13, #0
    beq .Land_false_59
    mov x14, x9
    and x12, x13, x14
    b .Land_end_60
.Land_false_59:
    mov x12, #0
.Land_end_60:
// === FIN EXPRESIÓN BINARIA ===
    cmp x12, #0
    beq .Lskip_branch_61
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x15, msg_60
    mov x0, x15
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_42
.Lskip_branch_61:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x16, msg_61
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_42:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_62
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: x ===
// === DECLARAR VARIABLE: x int ===
    sub sp, sp, #8  // Reservar espacio para x
    mov x9, #20
    str x9, [sp]
// Variable x declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut x := ? (int)
// === DECLARAR VARIABLE MUT: y ===
// === DECLARAR VARIABLE: y int ===
    sub sp, sp, #8  // Reservar espacio para y
    mov x9, #15
    str x9, [sp]
// Variable y declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut y := ? (int)
// === DECLARAR VARIABLE MUT: z ===
// === DECLARAR VARIABLE: z int ===
    sub sp, sp, #8  // Reservar espacio para z
    mov x9, #10
    str x9, [sp]
// Variable z declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut z := ? (int)
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// === DEBUG: EXPRESIÓN BINARIA + ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x9, [sp, #16]
// Variable cargada exitosamente: x
// === ACCESO A VARIABLE: y ===
// === CARGAR VARIABLE: y (offset: 8) ===
    ldr x10, [sp, #8]
// Variable cargada exitosamente: y
// Operador: +, Tipos: int + int
    mov x11, x9
    mov x12, x10
    add x13, x11, x12
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA * ===
// === ACCESO A VARIABLE: z ===
// === CARGAR VARIABLE: z (offset: 0) ===
    ldr x14, [sp, #0]
// Variable cargada exitosamente: z
// Operador: *, Tipos: int * int
    mov x15, x14
    mov x16, #2
    mul x17, x15, x16
// === FIN EXPRESIÓN BINARIA ===
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x19, x13
    mov x20, x17
    cmp x19, x20
    cset x18, gt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA <= ===
// === DEBUG: EXPRESIÓN BINARIA - ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x21, [sp, #16]
// Variable cargada exitosamente: x
// === ACCESO A VARIABLE: y ===
// === CARGAR VARIABLE: y (offset: 8) ===
    ldr x22, [sp, #8]
// Variable cargada exitosamente: y
// Operador: -, Tipos: int - int
    mov x23, x21
    mov x24, x22
    sub x25, x23, x24
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA / ===
// === ACCESO A VARIABLE: z ===
// === CARGAR VARIABLE: z (offset: 0) ===
    ldr x26, [sp, #0]
// Variable cargada exitosamente: z
// Operador: /, Tipos: int / int
    mov x27, x26
    mov x28, #2
    udiv x29, x27, x28
// === FIN EXPRESIÓN BINARIA ===
// Operador: <=, Tipos: int <= int
// === OPERACIÓN RELACIONAL <= ===
// ADVERTENCIA: Reciclando registros temporales
    mov x9, x25
    mov x10, x29
    cmp x9, x10
    cset x30, le
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x12, x18
    cmp x12, #0
    beq .Land_false_63
    mov x13, x30
    and x11, x12, x13
    b .Land_end_64
.Land_false_63:
    mov x11, #0
.Land_end_64:
// === FIN EXPRESIÓN BINARIA ===
    cmp x11, #0
    beq .Lskip_branch_65
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x14, msg_63
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_62
.Lskip_branch_65:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #2 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA || ===
// === DEBUG: EXPRESIÓN BINARIA == ===
// === DEBUG: EXPRESIÓN BINARIA % ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x15, [sp, #16]
// Variable cargada exitosamente: x
// Operador: %, Tipos: int % int
    mov x16, x15
    mov x17, #3
    udiv x18, x16, x17
    msub x18, x18, x17, x16
// === FIN EXPRESIÓN BINARIA ===
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x20, x18
    mov x21, #0
    cmp x20, x21
    cset x19, eq
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === DEBUG: EXPRESIÓN BINARIA + ===
// === ACCESO A VARIABLE: y ===
// === CARGAR VARIABLE: y (offset: 8) ===
    ldr x22, [sp, #8]
// Variable cargada exitosamente: y
// === ACCESO A VARIABLE: z ===
// === CARGAR VARIABLE: z (offset: 0) ===
    ldr x23, [sp, #0]
// Variable cargada exitosamente: z
// Operador: +, Tipos: int + int
    mov x24, x22
    mov x25, x23
    add x26, x24, x25
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA * ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x27, [sp, #16]
// Variable cargada exitosamente: x
// Operador: *, Tipos: int * int
    mov x28, x27
    mov x29, #2
    mul x30, x28, x29
// === FIN EXPRESIÓN BINARIA ===
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
// ADVERTENCIA: Reciclando registros temporales
    mov x10, x26
    mov x11, x30
    cmp x10, x11
    cset x9, ge
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x13, x19
    cmp x13, #1
    beq .Lor_true_66
    mov x14, x9
    orr x12, x13, x14
    b .Lor_end_67
.Lor_true_66:
    mov x12, #1
.Lor_end_67:
// === FIN EXPRESIÓN BINARIA ===
    cmp x12, #0
    beq .Lskip_branch_68
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x15, msg_64
    mov x0, x15
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_62
.Lskip_branch_68:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #3 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// === DEBUG: EXPRESIÓN BINARIA / ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x16, [sp, #16]
// Variable cargada exitosamente: x
// Operador: /, Tipos: int / int
    mov x17, x16
    mov x18, #2
    udiv x19, x17, x18
// === FIN EXPRESIÓN BINARIA ===
// === ACCESO A VARIABLE: y ===
// === CARGAR VARIABLE: y (offset: 8) ===
    ldr x20, [sp, #8]
// Variable cargada exitosamente: y
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x22, x19
    mov x23, x20
    cmp x22, x23
    cset x21, gt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA < ===
// === DEBUG: EXPRESIÓN BINARIA * ===
// === ACCESO A VARIABLE: z ===
// === CARGAR VARIABLE: z (offset: 0) ===
    ldr x24, [sp, #0]
// Variable cargada exitosamente: z
// Operador: *, Tipos: int * int
    mov x25, x24
    mov x26, #3
    mul x27, x25, x26
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA + ===
// === ACCESO A VARIABLE: x ===
// === CARGAR VARIABLE: x (offset: 16) ===
    ldr x28, [sp, #16]
// Variable cargada exitosamente: x
// === ACCESO A VARIABLE: y ===
// === CARGAR VARIABLE: y (offset: 8) ===
    ldr x29, [sp, #8]
// Variable cargada exitosamente: y
// Operador: +, Tipos: int + int
// ADVERTENCIA: Reciclando registros temporales
    mov x30, x28
    mov x9, x29
    add x10, x30, x9
// === FIN EXPRESIÓN BINARIA ===
// Operador: <, Tipos: int < int
// === OPERACIÓN RELACIONAL < ===
    mov x12, x27
    mov x13, x10
    cmp x12, x13
    cset x11, lt
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x15, x21
    cmp x15, #0
    beq .Land_false_69
    mov x16, x11
    and x14, x15, x16
    b .Land_end_70
.Land_false_69:
    mov x14, #0
.Land_end_70:
// === FIN EXPRESIÓN BINARIA ===
    cmp x14, #0
    beq .Lskip_branch_71
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x17, msg_65
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_62
.Lskip_branch_71:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x18, msg_66
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_62:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_67
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: nombre ===
// === DECLARAR VARIABLE: nombre string ===
    sub sp, sp, #8  // Reservar espacio para nombre
    adr x9, msg_68
    str x9, [sp]
// Variable nombre declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut nombre := ? (string)
// === DECLARAR VARIABLE MUT: edad ===
// === DECLARAR VARIABLE: edad int ===
    sub sp, sp, #8  // Reservar espacio para edad
    mov x9, #25
    str x9, [sp]
// Variable edad declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut edad := ? (int)
// === DECLARAR VARIABLE MUT: activo ===
// === LITERAL BOOL: true ===
// === DECLARAR VARIABLE: activo bool ===
    sub sp, sp, #8  // Reservar espacio para activo
    mov x9, #1
    str x9, [sp]
// Variable activo declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut activo := ? (bool)
// === DECLARAR VARIABLE MUT: vacio ===
// === DECLARAR VARIABLE: vacio string ===
    sub sp, sp, #8  // Reservar espacio para vacio
    adr x9, msg_69
    str x9, [sp]
// Variable vacio declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut vacio := ? (string)
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA == ===
// === ACCESO A VARIABLE: nombre ===
// === CARGAR VARIABLE: nombre (offset: 24) ===
    ldr x9, [sp, #24]
// Variable cargada exitosamente: nombre
// Operador: ==, Tipos: string == string
// === OPERACIÓN RELACIONAL == ===
    mov x11, x9
    adr x12, msg_70
    mov x0, x11
    mov x1, x12
    bl strcmp
    mov x10, x0
    cmp x10, #0
    cset x10, eq
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: edad ===
// === CARGAR VARIABLE: edad (offset: 16) ===
    ldr x13, [sp, #16]
// Variable cargada exitosamente: edad
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x15, x13
    mov x16, #18
    cmp x15, x16
    cset x14, ge
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x18, x10
    cmp x18, #0
    beq .Land_false_73
    mov x19, x14
    and x17, x18, x19
    b .Land_end_74
.Land_false_73:
    mov x17, #0
.Land_end_74:
// === FIN EXPRESIÓN BINARIA ===
// === ACCESO A VARIABLE: activo ===
// === CARGAR VARIABLE: activo (offset: 8) ===
    ldr x20, [sp, #8]
// Variable cargada exitosamente: activo
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x22, x17
    cmp x22, #0
    beq .Land_false_75
    mov x23, x20
    and x21, x22, x23
    b .Land_end_76
.Land_false_75:
    mov x21, #0
.Land_end_76:
// === FIN EXPRESIÓN BINARIA ===
    cmp x21, #0
    beq .Lskip_branch_77
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x24, msg_71
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_72
.Lskip_branch_77:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #2 ---
// Reseteando contadores de registros por límite
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA != ===
// === ACCESO A VARIABLE: nombre ===
// === CARGAR VARIABLE: nombre (offset: 24) ===
    ldr x9, [sp, #24]
// Variable cargada exitosamente: nombre
// === ACCESO A VARIABLE: vacio ===
// === CARGAR VARIABLE: vacio (offset: 0) ===
    ldr x10, [sp, #0]
// Variable cargada exitosamente: vacio
// Operador: !=, Tipos: string != string
// === OPERACIÓN RELACIONAL != ===
    mov x12, x9
    mov x13, x10
    mov x0, x12
    mov x1, x13
    bl strcmp
    mov x11, x0
    cmp x11, #0
    cset x11, ne
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// === ACCESO A VARIABLE: edad ===
// === CARGAR VARIABLE: edad (offset: 16) ===
    ldr x14, [sp, #16]
// Variable cargada exitosamente: edad
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x16, x14
    mov x17, #0
    cmp x16, x17
    cset x15, gt
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x19, x11
    cmp x19, #0
    beq .Land_false_78
    mov x20, x15
    and x18, x19, x20
    b .Land_end_79
.Land_false_78:
    mov x18, #0
.Land_end_79:
// === FIN EXPRESIÓN BINARIA ===
// === INICIO EXPRESIÓN UNARIA ===
// === ACCESO A VARIABLE: activo ===
// === CARGAR VARIABLE: activo (offset: 8) ===
    ldr x21, [sp, #8]
// Variable cargada exitosamente: activo
// Operador unario: !
// === NEGACIÓN LÓGICA ! ===
    mov x22, x21
    eor x23, x22, #1
// === FIN EXPRESIÓN UNARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x25, x18
    cmp x25, #0
    beq .Land_false_80
    mov x26, x23
    and x24, x25, x26
    b .Land_end_81
.Land_false_80:
    mov x24, #0
.Land_end_81:
// === FIN EXPRESIÓN BINARIA ===
    cmp x24, #0
    beq .Lskip_branch_82
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x27, msg_72
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_72
.Lskip_branch_82:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #3 ---
// Reseteando contadores de registros por límite
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA || ===
// === DEBUG: EXPRESIÓN BINARIA == ===
// === ACCESO A VARIABLE: nombre ===
// === CARGAR VARIABLE: nombre (offset: 24) ===
    ldr x9, [sp, #24]
// Variable cargada exitosamente: nombre
// === ACCESO A VARIABLE: vacio ===
// === CARGAR VARIABLE: vacio (offset: 0) ===
    ldr x10, [sp, #0]
// Variable cargada exitosamente: vacio
// Operador: ==, Tipos: string == string
// === OPERACIÓN RELACIONAL == ===
    mov x12, x9
    mov x13, x10
    mov x0, x12
    mov x1, x13
    bl strcmp
    mov x11, x0
    cmp x11, #0
    cset x11, eq
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA <= ===
// === ACCESO A VARIABLE: edad ===
// === CARGAR VARIABLE: edad (offset: 16) ===
    ldr x14, [sp, #16]
// Variable cargada exitosamente: edad
// Operador: <=, Tipos: int <= int
// === OPERACIÓN RELACIONAL <= ===
    mov x16, x14
    mov x17, #0
    cmp x16, x17
    cset x15, le
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
    mov x19, x11
    cmp x19, #1
    beq .Lor_true_83
    mov x20, x15
    orr x18, x19, x20
    b .Lor_end_84
.Lor_true_83:
    mov x18, #1
.Lor_end_84:
// === FIN EXPRESIÓN BINARIA ===
    cmp x18, #0
    beq .Lskip_branch_85
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x21, msg_73
    mov x0, x21
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_72
.Lskip_branch_85:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x22, msg_74
    mov x0, x22
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_72:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_75
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: nivel ===
// === DECLARAR VARIABLE: nivel int ===
    sub sp, sp, #8  // Reservar espacio para nivel
    mov x9, #5
    str x9, [sp]
// Variable nivel declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut nivel := ? (int)
// === DECLARAR VARIABLE MUT: puntos ===
// === DECLARAR VARIABLE: puntos int ===
    sub sp, sp, #8  // Reservar espacio para puntos
    mov x9, #100
    str x9, [sp]
// Variable puntos declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut puntos := ? (int)
// === DECLARAR VARIABLE MUT: vidas ===
// === DECLARAR VARIABLE: vidas int ===
    sub sp, sp, #8  // Reservar espacio para vidas
    mov x9, #3
    str x9, [sp]
// Variable vidas declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut vidas := ? (int)
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA > ===
// === ACCESO A VARIABLE: nivel ===
// === CARGAR VARIABLE: nivel (offset: 16) ===
    ldr x9, [sp, #16]
// Variable cargada exitosamente: nivel
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x11, x9
    mov x12, #0
    cmp x11, x12
    cset x10, gt
// === FIN EXPRESIÓN BINARIA ===
    cmp x10, #0
    beq .Lskip_branch_87
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x13, msg_76
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: puntos ===
// === CARGAR VARIABLE: puntos (offset: 8) ===
    ldr x14, [sp, #8]
// Variable cargada exitosamente: puntos
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x16, x14
    mov x17, #100
    cmp x16, x17
    cset x15, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x15, #0
    beq .Lskip_branch_89
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x18, msg_77
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA > ===
// === ACCESO A VARIABLE: vidas ===
// === CARGAR VARIABLE: vidas (offset: 0) ===
    ldr x19, [sp, #0]
// Variable cargada exitosamente: vidas
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x21, x19
    mov x22, #0
    cmp x21, x22
    cset x20, gt
// === FIN EXPRESIÓN BINARIA ===
    cmp x20, #0
    beq .Lskip_branch_91
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x23, msg_78
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nivel ===
// === CARGAR VARIABLE: nivel (offset: 16) ===
    ldr x24, [sp, #16]
// Variable cargada exitosamente: nivel
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x26, x24
    mov x27, #5
    cmp x26, x27
    cset x25, ge
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: puntos ===
// === CARGAR VARIABLE: puntos (offset: 8) ===
    ldr x28, [sp, #8]
// Variable cargada exitosamente: puntos
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
// ADVERTENCIA: Reciclando registros temporales
    mov x30, x28
    mov x9, #100
    cmp x30, x9
    cset x29, ge
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x11, x25
    cmp x11, #0
    beq .Land_false_93
    mov x12, x29
    and x10, x11, x12
    b .Land_end_94
.Land_false_93:
    mov x10, #0
.Land_end_94:
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: vidas ===
// === CARGAR VARIABLE: vidas (offset: 0) ===
    ldr x13, [sp, #0]
// Variable cargada exitosamente: vidas
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x15, x13
    mov x16, #3
    cmp x15, x16
    cset x14, ge
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x18, x10
    cmp x18, #0
    beq .Land_false_95
    mov x19, x14
    and x17, x18, x19
    b .Land_end_96
.Land_false_95:
    mov x17, #0
.Land_end_96:
// === FIN EXPRESIÓN BINARIA ===
    cmp x17, #0
    beq .Lskip_branch_97
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x20, msg_79
    mov x0, x20
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_92
.Lskip_branch_97:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x21, msg_80
    mov x0, x21
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_92:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// Pop scope
    b .Lif_final_90
.Lskip_branch_91:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x22, msg_81
    mov x0, x22
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_90:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// Pop scope
    b .Lif_final_88
.Lskip_branch_89:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x23, msg_82
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_88:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// Pop scope
    b .Lif_final_86
.Lskip_branch_87:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x24, msg_83
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_86:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_84
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: nota ===
// === DECLARAR VARIABLE: nota int ===
    sub sp, sp, #8  // Reservar espacio para nota
    mov x9, #85
    str x9, [sp]
// Variable nota declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut nota := ? (int)
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nota ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: nota
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x11, x9
    mov x12, #95
    cmp x11, x12
    cset x10, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x10, #0
    beq .Lskip_branch_99
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x13, msg_85
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_98
.Lskip_branch_99:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #2 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nota ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x14, [sp, #0]
// Variable cargada exitosamente: nota
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x16, x14
    mov x17, #90
    cmp x16, x17
    cset x15, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x15, #0
    beq .Lskip_branch_100
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x18, msg_86
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_98
.Lskip_branch_100:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #3 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nota ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x19, [sp, #0]
// Variable cargada exitosamente: nota
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x21, x19
    mov x22, #85
    cmp x21, x22
    cset x20, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x20, #0
    beq .Lskip_branch_101
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x23, msg_87
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_98
.Lskip_branch_101:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #4 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nota ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x24, [sp, #0]
// Variable cargada exitosamente: nota
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x26, x24
    mov x27, #80
    cmp x26, x27
    cset x25, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x25, #0
    beq .Lskip_branch_102
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x28, msg_88
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_98
.Lskip_branch_102:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #5 ---
// Reseteando contadores de registros por límite
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nota ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: nota
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x11, x9
    mov x12, #75
    cmp x11, x12
    cset x10, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x10, #0
    beq .Lskip_branch_103
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x13, msg_89
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_98
.Lskip_branch_103:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #6 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nota ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x14, [sp, #0]
// Variable cargada exitosamente: nota
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x16, x14
    mov x17, #70
    cmp x16, x17
    cset x15, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x15, #0
    beq .Lskip_branch_104
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x18, msg_90
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_98
.Lskip_branch_104:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #7 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nota ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x19, [sp, #0]
// Variable cargada exitosamente: nota
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x21, x19
    mov x22, #65
    cmp x21, x22
    cset x20, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x20, #0
    beq .Lskip_branch_105
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x23, msg_91
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_98
.Lskip_branch_105:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #8 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: nota ===
// === CARGAR VARIABLE: nota (offset: 0) ===
    ldr x24, [sp, #0]
// Variable cargada exitosamente: nota
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x26, x24
    mov x27, #60
    cmp x26, x27
    cset x25, ge
// === FIN EXPRESIÓN BINARIA ===
    cmp x25, #0
    beq .Lskip_branch_106
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x28, msg_92
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_98
.Lskip_branch_106:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x29, msg_93
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_98:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_94
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE MUT: base ===
// === DECLARAR VARIABLE: base int ===
    sub sp, sp, #8  // Reservar espacio para base
    mov x9, #4
    str x9, [sp]
// Variable base declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut base := ? (int)
// === DECLARAR VARIABLE MUT: exponente ===
// === DECLARAR VARIABLE: exponente int ===
    sub sp, sp, #8  // Reservar espacio para exponente
    mov x9, #2
    str x9, [sp]
// Variable exponente declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut exponente := ? (int)
// === DECLARAR VARIABLE MUT: resultado ===
// === DEBUG: EXPRESIÓN BINARIA * ===
// === ACCESO A VARIABLE: base ===
// === CARGAR VARIABLE: base (offset: 8) ===
    ldr x9, [sp, #8]
// Variable cargada exitosamente: base
// === ACCESO A VARIABLE: base ===
// === CARGAR VARIABLE: base (offset: 8) ===
    ldr x10, [sp, #8]
// Variable cargada exitosamente: base
// Operador: *, Tipos: int * int
    mov x11, x9
    mov x12, x10
    mul x13, x11, x12
// === FIN EXPRESIÓN BINARIA ===
// === DECLARAR VARIABLE: resultado int ===
    sub sp, sp, #8  // Reservar espacio para resultado
    str x13, [sp]
// Variable resultado declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut resultado := ? (int)
// === DECLARAR VARIABLE MUT: factorial ===
// === DECLARAR VARIABLE: factorial int ===
    sub sp, sp, #8  // Reservar espacio para factorial
    mov x9, #6
    str x9, [sp]
// Variable factorial declarada en [sp] (offset actual: 0)
// Declaración mut inferida: mut factorial := ? (int)
// === INICIO ESTRUCTURA IF-ELSE IF-ELSE ===
// --- Evaluando Branch #1 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA && ===
// === DEBUG: EXPRESIÓN BINARIA > ===
// === ACCESO A VARIABLE: resultado ===
// === CARGAR VARIABLE: resultado (offset: 8) ===
    ldr x9, [sp, #8]
// Variable cargada exitosamente: resultado
// Operador: >, Tipos: int > int
// === OPERACIÓN RELACIONAL > ===
    mov x11, x9
    mov x12, #10
    cmp x11, x12
    cset x10, gt
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA >= ===
// === ACCESO A VARIABLE: factorial ===
// === CARGAR VARIABLE: factorial (offset: 0) ===
    ldr x13, [sp, #0]
// Variable cargada exitosamente: factorial
// Operador: >=, Tipos: int >= int
// === OPERACIÓN RELACIONAL >= ===
    mov x15, x13
    mov x16, #6
    cmp x15, x16
    cset x14, ge
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
    mov x18, x10
    cmp x18, #0
    beq .Land_false_108
    mov x19, x14
    and x17, x18, x19
    b .Land_end_109
.Land_false_108:
    mov x17, #0
.Land_end_109:
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA == ===
// === DEBUG: EXPRESIÓN BINARIA * ===
// === DEBUG: EXPRESIÓN BINARIA + ===
// === ACCESO A VARIABLE: base ===
// === CARGAR VARIABLE: base (offset: 24) ===
    ldr x20, [sp, #24]
// Variable cargada exitosamente: base
// === ACCESO A VARIABLE: exponente ===
// === CARGAR VARIABLE: exponente (offset: 16) ===
    ldr x21, [sp, #16]
// Variable cargada exitosamente: exponente
// Operador: +, Tipos: int + int
    mov x22, x20
    mov x23, x21
    add x24, x22, x23
// === FIN EXPRESIÓN BINARIA ===
// Operador: *, Tipos: int * int
    mov x25, x24
    mov x26, #2
    mul x27, x25, x26
// === FIN EXPRESIÓN BINARIA ===
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x29, x27
    mov x30, #12
    cmp x29, x30
    cset x28, eq
// === FIN EXPRESIÓN BINARIA ===
// Operador: &&, Tipos: bool && bool
// === OPERACIÓN LÓGICA && ===
// ADVERTENCIA: Reciclando registros temporales
    mov x10, x17
    cmp x10, #0
    beq .Land_false_110
    mov x11, x28
    and x9, x10, x11
    b .Land_end_111
.Land_false_110:
    mov x9, #0
.Land_end_111:
// === FIN EXPRESIÓN BINARIA ===
    cmp x9, #0
    beq .Lskip_branch_112
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x12, msg_95
    mov x0, x12
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_107
.Lskip_branch_112:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Evaluando Branch #2 ---
// >>> PROCESANDO BRANCH IF/ELSE-IF <<<
// === DEBUG: EXPRESIÓN BINARIA || ===
// === DEBUG: EXPRESIÓN BINARIA == ===
// === DEBUG: EXPRESIÓN BINARIA / ===
// === ACCESO A VARIABLE: resultado ===
// === CARGAR VARIABLE: resultado (offset: 8) ===
    ldr x13, [sp, #8]
// Variable cargada exitosamente: resultado
// === ACCESO A VARIABLE: base ===
// === CARGAR VARIABLE: base (offset: 24) ===
    ldr x14, [sp, #24]
// Variable cargada exitosamente: base
// Operador: /, Tipos: int / int
    mov x15, x13
    mov x16, x14
    udiv x17, x15, x16
// === FIN EXPRESIÓN BINARIA ===
// === ACCESO A VARIABLE: exponente ===
// === CARGAR VARIABLE: exponente (offset: 16) ===
    ldr x18, [sp, #16]
// Variable cargada exitosamente: exponente
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x20, x17
    mov x21, x18
    cmp x20, x21
    cset x19, eq
// === FIN EXPRESIÓN BINARIA ===
// === DEBUG: EXPRESIÓN BINARIA == ===
// === DEBUG: EXPRESIÓN BINARIA % ===
// === ACCESO A VARIABLE: factorial ===
// === CARGAR VARIABLE: factorial (offset: 0) ===
    ldr x22, [sp, #0]
// Variable cargada exitosamente: factorial
// Operador: %, Tipos: int % int
    mov x23, x22
    mov x24, #3
    udiv x25, x23, x24
    msub x25, x25, x24, x23
// === FIN EXPRESIÓN BINARIA ===
// Operador: ==, Tipos: int == int
// === OPERACIÓN RELACIONAL == ===
    mov x27, x25
    mov x28, #0
    cmp x27, x28
    cset x26, eq
// === FIN EXPRESIÓN BINARIA ===
// Operador: ||, Tipos: bool || bool
// === OPERACIÓN LÓGICA || ===
// ADVERTENCIA: Reciclando registros temporales
    mov x30, x19
    cmp x30, #1
    beq .Lor_true_113
    mov x9, x26
    orr x29, x30, x9
    b .Lor_end_114
.Lor_true_113:
    mov x29, #1
.Lor_end_114:
// === FIN EXPRESIÓN BINARIA ===
    cmp x29, #0
    beq .Lskip_branch_115
// === EJECUTANDO BLOQUE (runtime true) ===
// Push scope: if
// === IMPRIMIR STRING  ===
    adr x10, msg_96
    mov x0, x10
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
    b .Lif_final_107
.Lskip_branch_115:
// Condición falsa - continuar al siguiente branch
// >>> FIN PROCESAMIENTO BRANCH <<<
// --- Ejecutando ELSE ---
// === INICIO BLOQUE ELSE ===
// Push scope: else
// Ejecutando sentencia 1 del bloque else
// === IMPRIMIR STRING  ===
    adr x11, msg_97
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
// Pop scope
// === FIN BLOQUE ELSE ===
.Lif_final_107:
// === FIN ESTRUCTURA IF-ELSE IF-ELSE ===

// === IMPRIMIR STRING  ===
    adr x9, msg_98
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_99
    mov x0, x9
    bl print_string
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_100
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: operacion ===
// === DECLARAR VARIABLE: operacion string ===
    sub sp, sp, #8  // Reservar espacio para operacion
    adr x9, msg_101
    str x9, [sp]
// Variable operacion declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: Var1 ===
// === DECLARAR VARIABLE: Var1 int ===
    sub sp, sp, #8  // Reservar espacio para Var1
    mov x9, #3
    str x9, [sp]
// Variable Var1 declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: Var2 ===
// === DECLARAR VARIABLE: Var2 int ===
    sub sp, sp, #8  // Reservar espacio para Var2
    mov x9, #5
    str x9, [sp]
// Variable Var2 declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: operacion ===
// === CARGAR VARIABLE: operacion (offset: 16) ===
    ldr x9, [sp, #16]
// Variable cargada exitosamente: operacion
// Switch sobre expresión tipo: string
    mov x10, x9
// Comparación case 0
    adr x11, msg_102
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x12, x0
    cmp x12, #0
    beq .Lcase_1_0
// Comparación case 1
    adr x13, msg_103
    mov x0, x10
    mov x1, x13
    bl strcmp
    mov x14, x0
    cmp x14, #0
    beq .Lcase_1_1
// Comparación case 2
    adr x15, msg_104
    mov x0, x10
    mov x1, x15
    bl strcmp
    mov x16, x0
    cmp x16, #0
    beq .Lcase_1_2
// Comparación case 3
    adr x17, msg_105
    mov x0, x10
    mov x1, x17
    bl strcmp
    mov x18, x0
    cmp x18, #0
    beq .Lcase_1_3
    b .Lswitch_default_1
.Lcase_1_0:
// Ejecutando case 0
// === DECLARAR VARIABLE: resultado ===
// === DEBUG: EXPRESIÓN BINARIA + ===
// === ACCESO A VARIABLE: Var1 ===
// === CARGAR VARIABLE: Var1 (offset: 8) ===
    ldr x19, [sp, #8]
// Variable cargada exitosamente: Var1
// === ACCESO A VARIABLE: Var2 ===
// === CARGAR VARIABLE: Var2 (offset: 0) ===
    ldr x20, [sp, #0]
// Variable cargada exitosamente: Var2
// Operador: +, Tipos: int + int
    mov x21, x19
    mov x22, x20
    add x23, x21, x22
// === FIN EXPRESIÓN BINARIA ===
// === DECLARAR VARIABLE: resultado int ===
// Error: variable 'resultado' ya declarada en ámbito 'global'
// === IMPRIMIR STRING  ===
    adr x24, msg_106
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: resultado (offset: 32) ===
    ldr x25, [sp, #32]
// === IMPRIMIR INT  ===
    mov x0, x25
    bl print_int
    bl print_newline
    bl print_newline
    b .Lswitch_end_1
.Lcase_1_1:
// Ejecutando case 1
// === DECLARAR VARIABLE: resultado ===
// === DEBUG: EXPRESIÓN BINARIA - ===
// === ACCESO A VARIABLE: Var1 ===
// === CARGAR VARIABLE: Var1 (offset: 8) ===
    ldr x26, [sp, #8]
// Variable cargada exitosamente: Var1
// === ACCESO A VARIABLE: Var2 ===
// === CARGAR VARIABLE: Var2 (offset: 0) ===
    ldr x27, [sp, #0]
// Variable cargada exitosamente: Var2
// Operador: -, Tipos: int - int
    mov x28, x26
    mov x29, x27
    sub x30, x28, x29
// === FIN EXPRESIÓN BINARIA ===
// === DECLARAR VARIABLE: resultado int ===
// Error: variable 'resultado' ya declarada en ámbito 'global'
// === IMPRIMIR STRING  ===
// ADVERTENCIA: Reciclando registros temporales
    adr x9, msg_107
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: resultado (offset: 32) ===
    ldr x10, [sp, #32]
// === IMPRIMIR INT  ===
    mov x0, x10
    bl print_int
    bl print_newline
    bl print_newline
    b .Lswitch_end_1
.Lcase_1_2:
// Ejecutando case 2
// === DECLARAR VARIABLE: resultado ===
// === DEBUG: EXPRESIÓN BINARIA * ===
// === ACCESO A VARIABLE: Var1 ===
// === CARGAR VARIABLE: Var1 (offset: 8) ===
    ldr x11, [sp, #8]
// Variable cargada exitosamente: Var1
// === ACCESO A VARIABLE: Var2 ===
// === CARGAR VARIABLE: Var2 (offset: 0) ===
    ldr x12, [sp, #0]
// Variable cargada exitosamente: Var2
// Operador: *, Tipos: int * int
    mov x13, x11
    mov x14, x12
    mul x15, x13, x14
// === FIN EXPRESIÓN BINARIA ===
// === DECLARAR VARIABLE: resultado int ===
// Error: variable 'resultado' ya declarada en ámbito 'global'
// === IMPRIMIR STRING  ===
    adr x16, msg_108
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: resultado (offset: 32) ===
    ldr x17, [sp, #32]
// === IMPRIMIR INT  ===
    mov x0, x17
    bl print_int
    bl print_newline
    bl print_newline
    b .Lswitch_end_1
.Lcase_1_3:
// Ejecutando case 3
// === DECLARAR VARIABLE: resultado ===
// === DEBUG: EXPRESIÓN BINARIA / ===
// === ACCESO A VARIABLE: Var1 ===
// === CARGAR VARIABLE: Var1 (offset: 8) ===
    ldr x18, [sp, #8]
// Variable cargada exitosamente: Var1
// === ACCESO A VARIABLE: Var2 ===
// === CARGAR VARIABLE: Var2 (offset: 0) ===
    ldr x19, [sp, #0]
// Variable cargada exitosamente: Var2
// Operador: /, Tipos: int / int
    mov x20, x18
    mov x21, x19
    udiv x22, x20, x21
// === FIN EXPRESIÓN BINARIA ===
// === DECLARAR VARIABLE: resultado int ===
// Error: variable 'resultado' ya declarada en ámbito 'global'
// === IMPRIMIR STRING  ===
    adr x23, msg_109
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: resultado (offset: 32) ===
    ldr x24, [sp, #32]
// === IMPRIMIR INT  ===
    mov x0, x24
    bl print_int
    bl print_newline
    bl print_newline
    b .Lswitch_end_1
.Lswitch_default_1:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x25, msg_110
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_1:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_111
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: dia_numero ===
// === DECLARAR VARIABLE: dia_numero int ===
    sub sp, sp, #8  // Reservar espacio para dia_numero
    mov x9, #3
    str x9, [sp]
// Variable dia_numero declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: dia_numero ===
// === CARGAR VARIABLE: dia_numero (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: dia_numero
// Switch sobre expresión tipo: int
    mov x10, x9
// Comparación case 0
    mov x11, #1
    cmp x10, x11
    beq .Lcase_2_0
// Comparación case 1
    mov x12, #2
    cmp x10, x12
    beq .Lcase_2_1
// Comparación case 2
    mov x13, #3
    cmp x10, x13
    beq .Lcase_2_2
// Comparación case 3
    mov x14, #4
    cmp x10, x14
    beq .Lcase_2_3
// Comparación case 4
    mov x15, #5
    cmp x10, x15
    beq .Lcase_2_4
// Comparación case 5
    mov x16, #6
    cmp x10, x16
    beq .Lcase_2_5
// Comparación case 6
    mov x17, #7
    cmp x10, x17
    beq .Lcase_2_6
    b .Lswitch_default_2
.Lcase_2_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x18, msg_112
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_2
.Lcase_2_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x19, msg_113
    mov x0, x19
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_2
.Lcase_2_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x20, msg_114
    mov x0, x20
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_2
.Lcase_2_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x21, msg_115
    mov x0, x21
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_2
.Lcase_2_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x22, msg_116
    mov x0, x22
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_2
.Lcase_2_5:
// Ejecutando case 5
// === IMPRIMIR STRING  ===
    adr x23, msg_117
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_2
.Lcase_2_6:
// Ejecutando case 6
// === IMPRIMIR STRING  ===
    adr x24, msg_118
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_2
.Lswitch_default_2:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x25, msg_119
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_2:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_120
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: letra_calificacion ===
// === DECLARAR VARIABLE: letra_calificacion string ===
    sub sp, sp, #8  // Reservar espacio para letra_calificacion
    adr x9, msg_121
    str x9, [sp]
// Variable letra_calificacion declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: letra_calificacion ===
// === CARGAR VARIABLE: letra_calificacion (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: letra_calificacion
// Switch sobre expresión tipo: string
    mov x10, x9
// Comparación case 0
    adr x11, msg_122
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x12, x0
    cmp x12, #0
    beq .Lcase_3_0
// Comparación case 1
    adr x13, msg_123
    mov x0, x10
    mov x1, x13
    bl strcmp
    mov x14, x0
    cmp x14, #0
    beq .Lcase_3_1
// Comparación case 2
    adr x15, msg_124
    mov x0, x10
    mov x1, x15
    bl strcmp
    mov x16, x0
    cmp x16, #0
    beq .Lcase_3_2
// Comparación case 3
    adr x17, msg_125
    mov x0, x10
    mov x1, x17
    bl strcmp
    mov x18, x0
    cmp x18, #0
    beq .Lcase_3_3
// Comparación case 4
    adr x19, msg_126
    mov x0, x10
    mov x1, x19
    bl strcmp
    mov x20, x0
    cmp x20, #0
    beq .Lcase_3_4
    b .Lswitch_default_3
.Lcase_3_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x21, msg_127
    mov x0, x21
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x22, msg_128
    mov x0, x22
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_3
.Lcase_3_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x23, msg_129
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x24, msg_130
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_3
.Lcase_3_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x25, msg_131
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x26, msg_132
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_3
.Lcase_3_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x27, msg_133
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x28, msg_134
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_3
.Lcase_3_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x29, msg_135
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x30, msg_136
    mov x0, x30
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_3
.Lswitch_default_3:
// Ejecutando case default
// === IMPRIMIR STRING  ===
// ADVERTENCIA: Reciclando registros temporales
    adr x9, msg_137
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_3:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_138
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: sistema_activo ===
// === LITERAL BOOL: true ===
// === DECLARAR VARIABLE: sistema_activo bool ===
    sub sp, sp, #8  // Reservar espacio para sistema_activo
    mov x9, #1
    str x9, [sp]
// Variable sistema_activo declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: modo_debug ===
// === LITERAL BOOL: false ===
// === DECLARAR VARIABLE: modo_debug bool ===
    sub sp, sp, #8  // Reservar espacio para modo_debug
    mov x9, #0
    str x9, [sp]
// Variable modo_debug declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: sistema_activo ===
// === CARGAR VARIABLE: sistema_activo (offset: 8) ===
    ldr x9, [sp, #8]
// Variable cargada exitosamente: sistema_activo
// Switch sobre expresión tipo: bool
    mov x10, x9
// === LITERAL BOOL: true ===
// Comparación case 0
    mov x11, #1
    cmp x10, x11
    beq .Lcase_4_0
// === LITERAL BOOL: false ===
// Comparación case 1
    mov x12, #0
    cmp x10, x12
    beq .Lcase_4_1
    b .Lswitch_end_4
.Lcase_4_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x13, msg_139
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: modo_debug ===
// === CARGAR VARIABLE: modo_debug (offset: 0) ===
    ldr x14, [sp, #0]
// Variable cargada exitosamente: modo_debug
// Switch sobre expresión tipo: bool
    mov x15, x14
// === LITERAL BOOL: true ===
// Comparación case 0
    mov x16, #1
    cmp x15, x16
    beq .Lcase_5_0
// === LITERAL BOOL: false ===
// Comparación case 1
    mov x17, #0
    cmp x15, x17
    beq .Lcase_5_1
    b .Lswitch_end_5
.Lcase_5_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x18, msg_140
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_5
.Lcase_5_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x19, msg_141
    mov x0, x19
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_5
.Lswitch_end_5:
// === FIN SWITCH STATEMENT ===
    b .Lswitch_end_4
.Lcase_4_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x20, msg_142
    mov x0, x20
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x21, msg_143
    mov x0, x21
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_4
.Lswitch_end_4:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_144
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: opcion ===
// === DECLARAR VARIABLE: opcion string ===
    sub sp, sp, #8  // Reservar espacio para opcion
    adr x9, msg_145
    str x9, [sp]
// Variable opcion declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: opcion ===
// === CARGAR VARIABLE: opcion (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: opcion
// Switch sobre expresión tipo: string
    mov x10, x9
// Comparación case 0
    adr x11, msg_146
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x12, x0
    cmp x12, #0
    beq .Lcase_6_0
// Comparación case 1
    adr x13, msg_147
    mov x0, x10
    mov x1, x13
    bl strcmp
    mov x14, x0
    cmp x14, #0
    beq .Lcase_6_1
// Comparación case 2
    adr x15, msg_148
    mov x0, x10
    mov x1, x15
    bl strcmp
    mov x16, x0
    cmp x16, #0
    beq .Lcase_6_2
// Comparación case 3
    adr x17, msg_149
    mov x0, x10
    mov x1, x17
    bl strcmp
    mov x18, x0
    cmp x18, #0
    beq .Lcase_6_3
// Comparación case 4
    adr x19, msg_150
    mov x0, x10
    mov x1, x19
    bl strcmp
    mov x20, x0
    cmp x20, #0
    beq .Lcase_6_4
    b .Lswitch_default_6
.Lcase_6_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x21, msg_151
    mov x0, x21
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x22, msg_152
    mov x0, x22
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_6
.Lcase_6_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x23, msg_153
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x24, msg_154
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x25, msg_155
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_6
.Lcase_6_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x26, msg_156
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x27, msg_157
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x28, msg_158
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x29, msg_159
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_6
.Lcase_6_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x30, msg_160
    mov x0, x30
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
// ADVERTENCIA: Reciclando registros temporales
    adr x9, msg_161
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_6
.Lcase_6_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x10, msg_162
    mov x0, x10
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x11, msg_163
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_6
.Lswitch_default_6:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x12, msg_164
    mov x0, x12
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x13, msg_165
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_6:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_166
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: extension ===
// === DECLARAR VARIABLE: extension string ===
    sub sp, sp, #8  // Reservar espacio para extension
    adr x9, msg_167
    str x9, [sp]
// Variable extension declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: extension ===
// === CARGAR VARIABLE: extension (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: extension
// Switch sobre expresión tipo: string
    mov x10, x9
// Comparación case 0
    adr x11, msg_168
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x12, x0
    cmp x12, #0
    beq .Lcase_7_0
// Comparación case 1
    adr x13, msg_169
    mov x0, x10
    mov x1, x13
    bl strcmp
    mov x14, x0
    cmp x14, #0
    beq .Lcase_7_1
// Comparación case 2
    adr x15, msg_170
    mov x0, x10
    mov x1, x15
    bl strcmp
    mov x16, x0
    cmp x16, #0
    beq .Lcase_7_2
// Comparación case 3
    adr x17, msg_171
    mov x0, x10
    mov x1, x17
    bl strcmp
    mov x18, x0
    cmp x18, #0
    beq .Lcase_7_3
// Comparación case 4
    adr x19, msg_172
    mov x0, x10
    mov x1, x19
    bl strcmp
    mov x20, x0
    cmp x20, #0
    beq .Lcase_7_4
// Comparación case 5
    adr x21, msg_173
    mov x0, x10
    mov x1, x21
    bl strcmp
    mov x22, x0
    cmp x22, #0
    beq .Lcase_7_5
// Comparación case 6
    adr x23, msg_174
    mov x0, x10
    mov x1, x23
    bl strcmp
    mov x24, x0
    cmp x24, #0
    beq .Lcase_7_6
    b .Lswitch_default_7
.Lcase_7_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x25, msg_175
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x26, msg_176
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x27, msg_177
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x28, msg_178
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x29, msg_179
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x30, msg_180
    mov x0, x30
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
// ADVERTENCIA: Reciclando registros temporales
    adr x9, msg_181
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x10, msg_182
    mov x0, x10
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x11, msg_183
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x12, msg_184
    mov x0, x12
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_5:
// Ejecutando case 5
// === IMPRIMIR STRING  ===
    adr x13, msg_185
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x14, msg_186
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lcase_7_6:
// Ejecutando case 6
// === IMPRIMIR STRING  ===
    adr x15, msg_187
    mov x0, x15
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x16, msg_188
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_7
.Lswitch_default_7:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x17, msg_189
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x18, msg_190
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_7:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_191
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: Nivel ===
// === DECLARAR VARIABLE: Nivel int ===
    sub sp, sp, #8  // Reservar espacio para Nivel
    mov x9, #2
    str x9, [sp]
// Variable Nivel declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: Puntos ===
// === DECLARAR VARIABLE: Puntos int ===
    sub sp, sp, #8  // Reservar espacio para Puntos
    mov x9, #0
    str x9, [sp]
// Variable Puntos declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: nivel ===
// === CARGAR VARIABLE: nivel (offset: 144) ===
    ldr x9, [sp, #144]
// Variable cargada exitosamente: nivel
// Switch sobre expresión tipo: int
    mov x10, x9
// Comparación case 0
    mov x11, #1
    cmp x10, x11
    beq .Lcase_8_0
// Comparación case 1
    mov x12, #2
    cmp x10, x12
    beq .Lcase_8_1
// Comparación case 2
    mov x13, #3
    cmp x10, x13
    beq .Lcase_8_2
// Comparación case 3
    mov x14, #4
    cmp x10, x14
    beq .Lcase_8_3
// Comparación case 4
    mov x15, #5
    cmp x10, x15
    beq .Lcase_8_4
    b .Lswitch_default_8
.Lcase_8_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x16, msg_192
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x17, msg_193
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNAR VARIABLE: Puntos = int (offset: 0) ===
    mov x18, #100
    str x18, [sp, #0]
    b .Lswitch_end_8
.Lcase_8_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x19, msg_194
    mov x0, x19
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x20, msg_195
    mov x0, x20
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNAR VARIABLE: Puntos = int (offset: 0) ===
    mov x21, #250
    str x21, [sp, #0]
    b .Lswitch_end_8
.Lcase_8_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x22, msg_196
    mov x0, x22
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x23, msg_197
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNAR VARIABLE: Puntos = int (offset: 0) ===
    mov x24, #500
    str x24, [sp, #0]
    b .Lswitch_end_8
.Lcase_8_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x25, msg_198
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x26, msg_199
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNAR VARIABLE: puntos = int (offset: 136) ===
    mov x27, #1000
    str x27, [sp, #136]
    b .Lswitch_end_8
.Lcase_8_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x28, msg_200
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x29, msg_201
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNAR VARIABLE: Puntos = int (offset: 0) ===
    mov x30, #2000
    str x30, [sp, #0]
    b .Lswitch_end_8
.Lswitch_default_8:
// Ejecutando case default
// === IMPRIMIR STRING  ===
// ADVERTENCIA: Reciclando registros temporales
    adr x9, msg_202
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === ASIGNAR VARIABLE: Puntos = int (offset: 0) ===
    mov x10, #0
    str x10, [sp, #0]
.Lswitch_end_8:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_203
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: Puntos (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x9
    bl print_int
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_204
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: codigo_error ===
// === DECLARAR VARIABLE: codigo_error int ===
    sub sp, sp, #8  // Reservar espacio para codigo_error
    mov x9, #404
    str x9, [sp]
// Variable codigo_error declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: codigo_error ===
// === CARGAR VARIABLE: codigo_error (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: codigo_error
// Switch sobre expresión tipo: int
    mov x10, x9
// Comparación case 0
    mov x11, #200
    cmp x10, x11
    beq .Lcase_9_0
// Comparación case 1
    mov x12, #400
    cmp x10, x12
    beq .Lcase_9_1
// Comparación case 2
    mov x13, #401
    cmp x10, x13
    beq .Lcase_9_2
// Comparación case 3
    mov x14, #403
    cmp x10, x14
    beq .Lcase_9_3
// Comparación case 4
    mov x15, #404
    cmp x10, x15
    beq .Lcase_9_4
// Comparación case 5
    mov x16, #500
    cmp x10, x16
    beq .Lcase_9_5
// Comparación case 6
    mov x17, #503
    cmp x10, x17
    beq .Lcase_9_6
    b .Lswitch_default_9
.Lcase_9_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x18, msg_205
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_9
.Lcase_9_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x19, msg_206
    mov x0, x19
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_9
.Lcase_9_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x20, msg_207
    mov x0, x20
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_9
.Lcase_9_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x21, msg_208
    mov x0, x21
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_9
.Lcase_9_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x22, msg_209
    mov x0, x22
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x23, msg_210
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_9
.Lcase_9_5:
// Ejecutando case 5
// === IMPRIMIR STRING  ===
    adr x24, msg_211
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x25, msg_212
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_9
.Lcase_9_6:
// Ejecutando case 6
// === IMPRIMIR STRING  ===
    adr x26, msg_213
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x27, msg_214
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_9
.Lswitch_default_9:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x28, msg_215
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: codigo_error (offset: 0) ===
    ldr x29, [sp, #0]
// === IMPRIMIR INT  ===
    mov x0, x29
    bl print_int
    bl print_newline
    bl print_newline
.Lswitch_end_9:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_216
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: comando ===
// === DECLARAR VARIABLE: comando string ===
    sub sp, sp, #8  // Reservar espacio para comando
    adr x9, msg_217
    str x9, [sp]
// Variable comando declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: comando ===
// === CARGAR VARIABLE: comando (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: comando
// Switch sobre expresión tipo: string
    mov x10, x9
// Comparación case 0
    adr x11, msg_218
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x12, x0
    cmp x12, #0
    beq .Lcase_10_0
// Comparación case 1
    adr x13, msg_219
    mov x0, x10
    mov x1, x13
    bl strcmp
    mov x14, x0
    cmp x14, #0
    beq .Lcase_10_1
// Comparación case 2
    adr x15, msg_220
    mov x0, x10
    mov x1, x15
    bl strcmp
    mov x16, x0
    cmp x16, #0
    beq .Lcase_10_2
// Comparación case 3
    adr x17, msg_221
    mov x0, x10
    mov x1, x17
    bl strcmp
    mov x18, x0
    cmp x18, #0
    beq .Lcase_10_3
// Comparación case 4
    adr x19, msg_222
    mov x0, x10
    mov x1, x19
    bl strcmp
    mov x20, x0
    cmp x20, #0
    beq .Lcase_10_4
// Comparación case 5
    adr x21, msg_223
    mov x0, x10
    mov x1, x21
    bl strcmp
    mov x22, x0
    cmp x22, #0
    beq .Lcase_10_5
// Comparación case 6
    adr x23, msg_224
    mov x0, x10
    mov x1, x23
    bl strcmp
    mov x24, x0
    cmp x24, #0
    beq .Lcase_10_6
    b .Lswitch_default_10
.Lcase_10_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x25, msg_225
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x26, msg_226
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x27, msg_227
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x28, msg_228
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x29, msg_229
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x30, msg_230
    mov x0, x30
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
// ADVERTENCIA: Reciclando registros temporales
    adr x9, msg_231
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x10, msg_232
    mov x0, x10
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x11, msg_233
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x12, msg_234
    mov x0, x12
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x13, msg_235
    mov x0, x13
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x14, msg_236
    mov x0, x14
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_5:
// Ejecutando case 5
// === IMPRIMIR STRING  ===
    adr x15, msg_237
    mov x0, x15
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x16, msg_238
    mov x0, x16
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lcase_10_6:
// Ejecutando case 6
// === IMPRIMIR STRING  ===
    adr x17, msg_239
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_10
.Lswitch_default_10:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x18, msg_240
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x19, msg_241
    mov x0, x19
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_10:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_242
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: color ===
// === DECLARAR VARIABLE: color string ===
    sub sp, sp, #8  // Reservar espacio para color
    adr x9, msg_243
    str x9, [sp]
// Variable color declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: color ===
// === CARGAR VARIABLE: color (offset: 0) ===
    ldr x9, [sp, #0]
// Variable cargada exitosamente: color
// Switch sobre expresión tipo: string
    mov x10, x9
// Comparación case 0
    adr x11, msg_244
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x12, x0
    cmp x12, #0
    beq .Lcase_11_0
// Comparación case 1
    adr x13, msg_245
    mov x0, x10
    mov x1, x13
    bl strcmp
    mov x14, x0
    cmp x14, #0
    beq .Lcase_11_1
// Comparación case 2
    adr x15, msg_246
    mov x0, x10
    mov x1, x15
    bl strcmp
    mov x16, x0
    cmp x16, #0
    beq .Lcase_11_2
// Comparación case 3
    adr x17, msg_247
    mov x0, x10
    mov x1, x17
    bl strcmp
    mov x18, x0
    cmp x18, #0
    beq .Lcase_11_3
// Comparación case 4
    adr x19, msg_248
    mov x0, x10
    mov x1, x19
    bl strcmp
    mov x20, x0
    cmp x20, #0
    beq .Lcase_11_4
// Comparación case 5
    adr x21, msg_249
    mov x0, x10
    mov x1, x21
    bl strcmp
    mov x22, x0
    cmp x22, #0
    beq .Lcase_11_5
    b .Lswitch_end_11
.Lcase_11_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x23, msg_250
    mov x0, x23
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_11
.Lcase_11_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x24, msg_251
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_11
.Lcase_11_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x25, msg_252
    mov x0, x25
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_11
.Lcase_11_3:
// Ejecutando case 3
// === IMPRIMIR STRING  ===
    adr x26, msg_253
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_11
.Lcase_11_4:
// Ejecutando case 4
// === IMPRIMIR STRING  ===
    adr x27, msg_254
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_11
.Lcase_11_5:
// Ejecutando case 5
// === IMPRIMIR STRING  ===
    adr x28, msg_255
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_11
.Lswitch_end_11:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_256
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_257
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: Edad ===
// === DECLARAR VARIABLE: Edad int ===
    sub sp, sp, #8  // Reservar espacio para Edad
    mov x9, #25
    str x9, [sp]
// Variable Edad declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: categoria ===
// === DECLARAR VARIABLE: categoria string ===
    sub sp, sp, #8  // Reservar espacio para categoria
    adr x9, msg_258
    str x9, [sp]
// Variable categoria declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: Edad ===
// === CARGAR VARIABLE: Edad (offset: 8) ===
    ldr x9, [sp, #8]
// Variable cargada exitosamente: Edad
// Switch sobre expresión tipo: int
    mov x10, x9
// Comparación case 0
    mov x11, #0
    cmp x10, x11
    beq .Lcase_12_0
// Comparación case 1
    mov x12, #1
    cmp x10, x12
    beq .Lcase_12_1
// Comparación case 2
    mov x13, #2
    cmp x10, x13
    beq .Lcase_12_2
// Comparación case 3
    mov x14, #13
    cmp x10, x14
    beq .Lcase_12_3
// Comparación case 4
    mov x15, #18
    cmp x10, x15
    beq .Lcase_12_4
// Comparación case 5
    mov x16, #25
    cmp x10, x16
    beq .Lcase_12_5
// Comparación case 6
    mov x17, #65
    cmp x10, x17
    beq .Lcase_12_6
    b .Lswitch_default_12
.Lcase_12_0:
// Ejecutando case 0
// === ASIGNAR VARIABLE: categoria = string (offset: 0) ===
    adr x18, msg_259
    str x18, [sp, #0]
    b .Lswitch_end_12
.Lcase_12_1:
// Ejecutando case 1
// === ASIGNAR VARIABLE: categoria = string (offset: 0) ===
    adr x19, msg_260
    str x19, [sp, #0]
    b .Lswitch_end_12
.Lcase_12_2:
// Ejecutando case 2
// === ASIGNAR VARIABLE: categoria = string (offset: 0) ===
    adr x20, msg_261
    str x20, [sp, #0]
    b .Lswitch_end_12
.Lcase_12_3:
// Ejecutando case 3
// === ASIGNAR VARIABLE: categoria = string (offset: 0) ===
    adr x21, msg_262
    str x21, [sp, #0]
    b .Lswitch_end_12
.Lcase_12_4:
// Ejecutando case 4
// === ASIGNAR VARIABLE: categoria = string (offset: 0) ===
    adr x22, msg_263
    str x22, [sp, #0]
    b .Lswitch_end_12
.Lcase_12_5:
// Ejecutando case 5
// === ASIGNAR VARIABLE: categoria = string (offset: 0) ===
    adr x23, msg_264
    str x23, [sp, #0]
// === IMPRIMIR STRING  ===
    adr x24, msg_265
    mov x0, x24
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_12
.Lcase_12_6:
// Ejecutando case 6
// === ASIGNAR VARIABLE: categoria = string (offset: 0) ===
    adr x25, msg_266
    str x25, [sp, #0]
    b .Lswitch_end_12
.Lswitch_default_12:
// Ejecutando case default
// === ASIGNAR VARIABLE: categoria = string (offset: 0) ===
    adr x26, msg_267
    str x26, [sp, #0]
.Lswitch_end_12:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_268
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === CARGAR VARIABLE: categoria (offset: 0) ===
    ldr x9, [sp, #0]
// === IMPRIMIR STRING  ===
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_269
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === DECLARAR VARIABLE: tipo_usuario ===
// === DECLARAR VARIABLE: tipo_usuario string ===
    sub sp, sp, #8  // Reservar espacio para tipo_usuario
    adr x9, msg_270
    str x9, [sp]
// Variable tipo_usuario declarada en [sp] (offset actual: 0)
// === DECLARAR VARIABLE: permiso ===
// === DECLARAR VARIABLE: permiso string ===
    sub sp, sp, #8  // Reservar espacio para permiso
    adr x9, msg_271
    str x9, [sp]
// Variable permiso declarada en [sp] (offset actual: 0)
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: tipo_usuario ===
// === CARGAR VARIABLE: tipo_usuario (offset: 8) ===
    ldr x9, [sp, #8]
// Variable cargada exitosamente: tipo_usuario
// Switch sobre expresión tipo: string
    mov x10, x9
// Comparación case 0
    adr x11, msg_272
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x12, x0
    cmp x12, #0
    beq .Lcase_13_0
// Comparación case 1
    adr x13, msg_273
    mov x0, x10
    mov x1, x13
    bl strcmp
    mov x14, x0
    cmp x14, #0
    beq .Lcase_13_1
// Comparación case 2
    adr x15, msg_274
    mov x0, x10
    mov x1, x15
    bl strcmp
    mov x16, x0
    cmp x16, #0
    beq .Lcase_13_2
    b .Lswitch_default_13
.Lcase_13_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x17, msg_275
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: permiso ===
// === CARGAR VARIABLE: permiso (offset: 0) ===
    ldr x18, [sp, #0]
// Variable cargada exitosamente: permiso
// Switch sobre expresión tipo: string
    mov x19, x18
// Comparación case 0
    adr x20, msg_276
    mov x0, x19
    mov x1, x20
    bl strcmp
    mov x21, x0
    cmp x21, #0
    beq .Lcase_14_0
// Comparación case 1
    adr x22, msg_277
    mov x0, x19
    mov x1, x22
    bl strcmp
    mov x23, x0
    cmp x23, #0
    beq .Lcase_14_1
// Comparación case 2
    adr x24, msg_278
    mov x0, x19
    mov x1, x24
    bl strcmp
    mov x25, x0
    cmp x25, #0
    beq .Lcase_14_2
    b .Lswitch_default_14
.Lcase_14_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x26, msg_279
    mov x0, x26
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_14
.Lcase_14_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x27, msg_280
    mov x0, x27
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_14
.Lcase_14_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x28, msg_281
    mov x0, x28
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_14
.Lswitch_default_14:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x29, msg_282
    mov x0, x29
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_14:
// === FIN SWITCH STATEMENT ===
    b .Lswitch_end_13
.Lcase_13_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x30, msg_283
    mov x0, x30
    bl print_string
    bl print_newline
    bl print_newline
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: permiso ===
// === CARGAR VARIABLE: permiso (offset: 0) ===
// ADVERTENCIA: Reciclando registros temporales
    ldr x9, [sp, #0]
// Variable cargada exitosamente: permiso
// Switch sobre expresión tipo: string
    mov x10, x9
// Comparación case 0
    adr x11, msg_284
    mov x0, x10
    mov x1, x11
    bl strcmp
    mov x12, x0
    cmp x12, #0
    beq .Lcase_15_0
// Comparación case 1
    adr x13, msg_285
    mov x0, x10
    mov x1, x13
    bl strcmp
    mov x14, x0
    cmp x14, #0
    beq .Lcase_15_1
// Comparación case 2
    adr x15, msg_286
    mov x0, x10
    mov x1, x15
    bl strcmp
    mov x16, x0
    cmp x16, #0
    beq .Lcase_15_2
    b .Lswitch_default_15
.Lcase_15_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x17, msg_287
    mov x0, x17
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_15
.Lcase_15_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
    adr x18, msg_288
    mov x0, x18
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_15
.Lcase_15_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x19, msg_289
    mov x0, x19
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_15
.Lswitch_default_15:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x20, msg_290
    mov x0, x20
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_15:
// === FIN SWITCH STATEMENT ===
    b .Lswitch_end_13
.Lcase_13_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x21, msg_291
    mov x0, x21
    bl print_string
    bl print_newline
    bl print_newline
// === INICIO SWITCH STATEMENT ===
// === ACCESO A VARIABLE: permiso ===
// === CARGAR VARIABLE: permiso (offset: 0) ===
    ldr x22, [sp, #0]
// Variable cargada exitosamente: permiso
// Switch sobre expresión tipo: string
    mov x23, x22
// Comparación case 0
    adr x24, msg_292
    mov x0, x23
    mov x1, x24
    bl strcmp
    mov x25, x0
    cmp x25, #0
    beq .Lcase_16_0
// Comparación case 1
    adr x26, msg_293
    mov x0, x23
    mov x1, x26
    bl strcmp
    mov x27, x0
    cmp x27, #0
    beq .Lcase_16_1
// Comparación case 2
    adr x28, msg_294
    mov x0, x23
    mov x1, x28
    bl strcmp
    mov x29, x0
    cmp x29, #0
    beq .Lcase_16_2
    b .Lswitch_default_16
.Lcase_16_0:
// Ejecutando case 0
// === IMPRIMIR STRING  ===
    adr x30, msg_295
    mov x0, x30
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_16
.Lcase_16_1:
// Ejecutando case 1
// === IMPRIMIR STRING  ===
// ADVERTENCIA: Reciclando registros temporales
    adr x9, msg_296
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_16
.Lcase_16_2:
// Ejecutando case 2
// === IMPRIMIR STRING  ===
    adr x10, msg_297
    mov x0, x10
    bl print_string
    bl print_newline
    bl print_newline
    b .Lswitch_end_16
.Lswitch_default_16:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x11, msg_298
    mov x0, x11
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_16:
// === FIN SWITCH STATEMENT ===
    b .Lswitch_end_13
.Lswitch_default_13:
// Ejecutando case default
// === IMPRIMIR STRING  ===
    adr x12, msg_299
    mov x0, x12
    bl print_string
    bl print_newline
    bl print_newline
.Lswitch_end_13:
// === FIN SWITCH STATEMENT ===
// === IMPRIMIR STRING  ===
    adr x9, msg_300
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline
// === IMPRIMIR STRING  ===
    adr x9, msg_301
    mov x0, x9
    bl print_string
    bl print_newline
    bl print_newline

    // Limpiar stack de slices antes de salir
    add sp, sp, #256
    // Salir del programa
    mov x8, #93
    mov x0, #0
    svc 0

// --------------------------------------------------------
//            FUNCIONES AUXILIARES ARM64
// --------------------------------------------------------

print_int:
    stp   x29, x30, [sp, #-16]!   // Guardar frame pointer y link register
    mov   x29, sp
    stp   x1, x2, [sp, #-16]!     // Guardar registros que vamos a usar
    stp   x3, x4, [sp, #-16]!
    stp   x5, x6, [sp, #-16]!

    //  Usar comparación con registro zero, no inmediato
    cmp   x0, xzr                 // ¿Es negativo? (usar xzr en lugar de #0)
    bge   .Lpi_pos
    
    //  Manejar números negativos - GUARDAR x0 original
    stp   x7, x8, [sp, #-16]!     // Guardar más registros
    mov   x7, x0                  // GUARDAR valor original en x7
    
    // Imprimir signo menos
    mov   x8, #64                 // Syscall write
    ldr   x1, =msg_menos          // Imprimir "-"
    mov   x2, #1
    mov   x0, #1
    svc   0
    
    //  RESTAURAR y negar el valor original
    mov   x0, x7                  // Restaurar valor original
    neg   x0, x0                  // Hacer positivo
    ldp   x7, x8, [sp], #16       // Restaurar registros

.Lpi_pos:
    ldr   x2, =buffer_int         // Buffer para dígitos
    add   x2, x2, #32             // Empezar desde el final
    mov   x3, #0                  // Contador de dígitos
    mov   x6, #10                 // Divisor

.Lpi_loop:
    udiv  x4, x0, x6              // x4 = x0 / 10
    msub  x5, x4, x6, x0          // x5 = x0 % 10 (resto)
    add   x5, x5, #48             // Convertir a ASCII
    sub   x2, x2, #1              // Retroceder en buffer
    strb  w5, [x2]                // Guardar dígito
    mov   x0, x4                  // Siguiente iteración
    add   x3, x3, #1              // Incrementar contador
    cmp   x0, #0
    bne   .Lpi_loop

    mov   x0, #1                  // stdout
    mov   x1, x2                  // Buffer con dígitos
    mov   x2, x3                  // Cantidad de dígitos
    mov   x8, #64                 // Syscall write
    svc   0

    ldp   x5, x6, [sp], #16       // Restaurar registros
    ldp   x3, x4, [sp], #16
    ldp   x1, x2, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

print_float:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x19, x20, [sp, #-16]!   // Guardar registros
    stp   x21, x22, [sp, #-16]!

    fcvtzs  x20, d0               // Convertir parte entera
    mov     x0, x20
    bl      print_int             // Imprimir parte entera

    mov     x8, #64               // Imprimir punto decimal
    ldr     x1, =msg_punto
    mov     x2, #1
    mov     x0, #1
    svc     0

    scvtf   d1, x20               // Convertir entero a float
    fsub    d2, d0, d1            // d2 = parte fraccionaria
    mov     x2, #100
    scvtf   d1, x2                // d1 = 100.0
    fmul    d2, d2, d1            // Multiplicar por 100
    fcvtzs  x20, d2               // Convertir a entero

    cmp     x20, #0               // Valor absoluto
    bge     .Lpf_pos
    neg     x20, x20
.Lpf_pos:

    mov     x19, #10              // Extraer dígitos
    udiv    x21, x20, x19         // Decenas
    msub    x22, x21, x19, x20    // Unidades
    add     w21, w21, #48         // Convertir a ASCII
    add     w22, w22, #48

    mov     x0, x21               // Imprimir dígitos
    bl      print_char
    mov     x0, x22
    bl      print_char

    ldp   x21, x22, [sp], #16     // Restaurar registros
    ldp   x19, x20, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

print_string:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    mov   x19, x0                 // Guardar dirección del string

    bl    strlen                  // Calcular longitud
    mov   x2, x0                  // x2 = longitud
    mov   x1, x19                 // x1 = dirección string
    mov   x0, #1                  // stdout
    mov   x8, #64                 // Syscall write
    svc   0

    ldp   x29, x30, [sp], #16
    ret

print_bool:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp

    add   w0, w0, #48             // 0→'0', 1→'1' (ASCII)
    bl    print_char              // Imprimir carácter

    ldp   x29, x30, [sp], #16
    ret

print_char:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp

    strb  w0, [sp, #-1]!          // Poner carácter en stack
    mov   x0, #1                  // stdout
    mov   x1, sp                  // Dirección del carácter
    mov   x2, #1                  // 1 byte
    mov   x8, #64                 // Syscall write
    svc   0
    add   sp, sp, #1              // Limpiar stack

    ldp   x29, x30, [sp], #16
    ret

print_newline:
    mov   x0, #1                  // stdout
    ldr   x1, =msg_nl             // "\n"
    mov   x2, #1                  // 1 byte
    mov   x8, #64                 // Syscall write
    svc   0
    ret

strlen:
    mov   x1, #0                  // contador = 0
.Lstrlen_loop:
    ldrb  w2, [x0, x1]           // Cargar byte
    cmp   w2, #0                 // ¿Es '\0'?
    beq   .Lstrlen_end
    add   x1, x1, #1             // Incrementar contador
    b     .Lstrlen_loop
.Lstrlen_end:
    mov   x0, x1                 // Retornar longitud
    ret

strcmp:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x2, x3, [sp, #-16]!

strcmp_loop:
    ldrb  w2, [x0], #1            // Cargar char de string1
    ldrb  w3, [x1], #1            // Cargar char de string2
    
    cmp   w2, w3                  // Comparar caracteres
    bne   strcmp_different
    
    cmp   w2, #0                  // ¿Final de string?
    beq   strcmp_equal
    
    b     strcmp_loop
    
strcmp_equal:
    mov   x0, #0                  // Iguales = 0
    b     strcmp_end
    
strcmp_different:
    mov   x0, #1                  // Diferentes = 1
    
strcmp_end:
    ldp   x2, x3, [sp], #16       // Restaurar registros
    ldp   x29, x30, [sp], #16
    ret

concat_strings:
    stp   x29, x30, [sp, #-16]!   // Guardar frame
    mov   x29, sp
    stp   x19, x20, [sp, #-16]!   // Guardar registros
    stp   x21, x22, [sp, #-16]!

    mov   x19, x0                 // string1
    mov   x20, x1                 // string2
    mov   x21, x2                 // buffer destino
    mov   x22, #0                 // índice destino

.Lconcat_copy1:                   // Copiar string1
    ldrb  w3, [x19], #1           // Cargar byte y avanzar
    cmp   w3, #0                  // ¿Es '\0'?
    beq   .Lconcat_copy2
    strb  w3, [x21, x22]          // Guardar en destino
    add   x22, x22, #1            // Avanzar índice
    b     .Lconcat_copy1

.Lconcat_copy2:                   // Copiar string2
    ldrb  w3, [x20], #1           // Cargar byte y avanzar
    strb  w3, [x21, x22]          // Guardar (incluye '\0')
    cmp   w3, #0                  // ¿Era '\0'?
    beq   .Lconcat_end
    add   x22, x22, #1            // Avanzar índice
    b     .Lconcat_copy2

.Lconcat_end:
    mov   x0, x21                 // Retornar resultado

    ldp   x21, x22, [sp], #16     // Restaurar registros
    ldp   x19, x20, [sp], #16
    ldp   x29, x30, [sp], #16
    ret

