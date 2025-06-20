.section .data
msg1: .ascii "21"
msg2: .ascii "10"
msg3: .ascii "20"
msg4: .ascii "8"
msg5: .ascii "0"

.section .text
.global _start

_start:

    // print() -> "21"
    MOV X0, #1
    LDR X1, =msg1
    MOV X2, #2
    MOV X8, #64
    SVC #0
    // print() -> "10"
    MOV X0, #1
    LDR X1, =msg2
    MOV X2, #2
    MOV X8, #64
    SVC #0
    // print() -> "20"
    MOV X0, #1
    LDR X1, =msg3
    MOV X2, #2
    MOV X8, #64
    SVC #0
    // print() -> "8"
    MOV X0, #1
    LDR X1, =msg4
    MOV X2, #1
    MOV X8, #64
    SVC #0
    // print() -> "0"
    MOV X0, #1
    LDR X1, =msg5
    MOV X2, #1
    MOV X8, #64
    SVC #0

    // Salida del programa
    MOV X0, #0
    MOV X8, #93
    SVC #0
