# Astra

A simple fantasy console emulator featuring an 8-bit CPU implemented in Go.  
It includes registers, memory, and a basic instruction set to run simple programs.

---

## Features

- 8-bit CPU with:
  - Registers: A, B (8-bit)
  - Program Counter (PC), Stack Pointer (SP) (16-bit)
  - Flags: Zero, Carry
- 64KB Memory
- Simple instruction set with:
  - NOP, LDA, LDB, ADD, JMP, HLT
- Step-by-step CPU execution cycle

---

## Getting Started

### Prerequisites

- Go 1.18+

---

- ## Example Program

Loads two numbers, adds them, and halts:

| Instruction | Description         |
|-------------|---------------------|
| LDA 10      | Load 10 into A      |
| LDB 20      | Load 20 into B      |
| ADD A, B    | A = A + B           |
| HLT         | Halt execution      |

Expected output:
```
A register after addition: 30
```

## CPU Instruction Set (Opcodes)

| Opcode | Instruction | Description               |
|--------|-------------|---------------------------|
| 0x00   | NOP         | No operation              |
| 0x01   | LDA #imm    | Load immediate into A     |
| 0x02   | LDB #imm    | Load immediate into B     |
| 0x03   | ADD A, B    | Add B to A                |
| 0x04   | JMP addr    | Jump to address           |
| 0xFF   | HLT         | Halt CPU                  |

---

## License

This project is licensed under the MIT License.
