package cpu

import (
	"fmt"
)

const (
	MemSize = 65536 // 64KB memory

	// Opcodes
	NOP  = 0x00
	LDAI = 0x01 // LDA #imm
	LDBI = 0x02 // LDB #imm
	ADD  = 0x03 // ADD A, B
	SUB  = 0x04 // SUB A, B
	JMP  = 0x05 // JMP addr
	LDA  = 0x06 // LDA addr
	STA  = 0x07 // STA addr
	CMP  = 0x08 // CMP A, B
	JE   = 0x09 // Jump if A == B
	JNE  = 0x0A // Jump if A != B
	OUT  = 0x0B // OUT A to port
	IN   = 0x0C // IN from port to A
	HLT  = 0xFF
)

// CPU represents a simple 8-bit CPU.
type CPU struct {
	A, B     byte     // Registers
	PC       uint16   // Program counter
	ZeroFlag bool     // Result of last CMP
	Mem      [MemSize]byte // RAM
	IO       [256]byte     // I/O ports
	Halted   bool
}

// New creates and returns a new CPU.
func New() *CPU {
	return &CPU{}
}

// LoadProgram loads a program into memory starting at address 0.
func (c *CPU) LoadProgram(program []byte) {
	copy(c.Mem[:], program)
}

// Step executes a single instruction.
func (c *CPU) Step() error {
	op := c.fetch()

	switch op {
	case NOP:
		// Do nothing

	case LDAI:
		c.A = c.fetch()

	case LDBI:
		c.B = c.fetch()

	case ADD:
		c.A += c.B

	case SUB:
		c.A -= c.B

	case JMP:
		addr := c.fetchWord()
		c.PC = addr

	case LDA:
		addr := c.fetchWord()
		c.A = c.Mem[addr]

	case STA:
		addr := c.fetchWord()
		c.Mem[addr] = c.A

	case CMP:
		c.ZeroFlag = (c.A == c.B)

	case JE:
		addr := c.fetchWord()
		if c.ZeroFlag {
			c.PC = addr
		}

	case JNE:
		addr := c.fetchWord()
		if !c.ZeroFlag {
			c.PC = addr
		}

	case OUT:
		port := c.fetch()
		c.IO[port] = c.A

	case IN:
		port := c.fetch()
		c.A = c.IO[port]

	case HLT:
		c.Halted = true

	default:
		return fmt.Errorf("unknown opcode: 0x%X", op)
	}

	return nil
}

// Run starts the CPU execution loop.
func (c *CPU) Run() error {
	for !c.Halted {
		if err := c.Step(); err != nil {
			return err
		}
	}
	return nil
}

// fetch gets the next byte from memory and advances PC.
func (c *CPU) fetch() byte {
	b := c.Mem[c.PC]
	c.PC++
	return b
}

// fetchWord reads a 16-bit address (little-endian).
func (c *CPU) fetchWord() uint16 {
	low := uint16(c.fetch())
	high := uint16(c.fetch())
	return low | (high << 8)
}

