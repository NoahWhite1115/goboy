package main

import "testing"

func setUpCPU(bios []uint8) *CPU {
	mmu := newMMU()
	mmu.loadBios(bios)
	cpu := newCPU(mmu)

	return cpu
}

//0x00
func TestNOP(t *testing.T) {
	opcodes := []uint8{0x00}
	cpu := setUpCPU(opcodes)

	ticks := cpu.runCommand(false)

	if ticks != 4 {
		t.Errorf("NOP ticks set wrong.")
	}
}

//0x01
func TestLDBCu16(t *testing.T) {
	opcodes := []uint8{0x01, 0xcc, 0xdd}
	cpu := setUpCPU(opcodes)

	ticks := cpu.runCommand(false)

	if ticks != 12 && cpu.register.BC() != 0xddcc {
		t.Errorf("LD BC u16 has errored.")
	}
}

//0x02
func TestLDBCMA(t *testing.T) {
	//set A = 0xac, BC = 0xddcc
	opcodes := []uint8{0x3e, 0xac, 0x01, 0xcc, 0xdd, 0x02}
	cpu := setUpCPU(opcodes)

	cpu.runCommand(false)
	cpu.runCommand(false)
	ticks := cpu.runCommand(false)

	if ticks != 8 && cpu.mmu.readByte(0xddcc) == 0xac {
		t.Errorf("LD (BC) A has errored.")
	}
}

//0x03
func TestINCBC(t *testing.T) {
	opcodes := []uint8{0x03}
	cpu := setUpCPU(opcodes)

	ticks := cpu.runCommand(false)

	if ticks != 8 && cpu.register.BC() == 0x0001 {
		t.Errorf("LD (BC) A has errored.")
	}
}
