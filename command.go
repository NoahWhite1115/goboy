package main

import (
	"log"
	"strconv"
)

type Command struct {
	ticks   int
	argsLen int
	op      func(*Register, *MMU, []uint8)
	desc    string
}

func newCommand(ticks int, argsLen int, desc string, op func(*Register, *MMU, []uint8)) *Command {
	cmd := new(Command)
	cmd.ticks = ticks
	cmd.argsLen = argsLen
	cmd.desc = desc
	cmd.op = op

	return cmd
}

//TODO: move this to it's own file
func toWord(msb, lsb uint8) uint16 {
	word := uint16(msb)<<8 | uint16(lsb)
	return word
}

func (cmd *Command) printCmd(a []uint8) {
	out := cmd.desc
	for _, v := range a {
		out += " "
		out += strconv.Itoa(int(v))
	}
	log.Println(out)
}

func makeOpCodes() []*Command {
	opCodeArr := make([]*Command, 0x1ff)

	opCodeArr[0x00] = newCommand(4, 0, "NOP",
		func(r *Register, m *MMU, a []uint8) {
		})

	opCodeArr[0x01] = newCommand(12, 2, "LD BC, nn",
		func(r *Register, m *MMU, a []uint8) {
			r.setBC(toWord(a[1], a[0]))
		})

	opCodeArr[0x02] = newCommand(8, 0, "LD (BC), A",
		func(r *Register, m *MMU, a []uint8) {
			m.setByte(r.BC(), r.a)
		})

	opCodeArr[0x03] = newCommand(8, 0, "INC BC",
		func(r *Register, m *MMU, a []uint8) {
			r.setBC(r.BC() + 1)
		})

	opCodeArr[0x04] = newCommand(4, 0, "INC B",
		func(r *Register, m *MMU, a []uint8) {
			r.b = inc(r.b, r.f)
		})

	opCodeArr[0x05] = newCommand(4, 0, "DEC B",
		func(r *Register, m *MMU, a []uint8) {
			r.b = dec(r.b, r.f)
		})

	opCodeArr[0x06] = newCommand(8, 1, "LD B, n",
		func(r *Register, m *MMU, a []uint8) {
			r.b = a[0]
		})

	opCodeArr[0x0a] = newCommand(8, 0, "LD A,(BC)",
		func(r *Register, m *MMU, a []uint8) {
			r.a = m.readByte(r.BC())
		})

	opCodeArr[0x0c] = newCommand(4, 0, "INC C",
		func(r *Register, m *MMU, a []uint8) {
			r.c = inc(r.c, r.f)
		})

	opCodeArr[0x0d] = newCommand(4, 0, "DEC C",
		func(r *Register, m *MMU, a []uint8) {
			r.c = dec(r.c, r.f)
		})

	opCodeArr[0x0e] = newCommand(8, 1, "LD C, n",
		func(r *Register, m *MMU, a []uint8) {
			r.c = a[0]
		})

	opCodeArr[0x11] = newCommand(12, 2, "LD DE, nn",
		func(r *Register, m *MMU, a []uint8) {
			r.setDE(toWord(a[1], a[0]))
		})

	opCodeArr[0x12] = newCommand(8, 0, "LD (DE), A",
		func(r *Register, m *MMU, a []uint8) {
			m.setByte(r.DE(), r.a)
		})

	opCodeArr[0x13] = newCommand(8, 0, "INC DE",
		func(r *Register, m *MMU, a []uint8) {
			r.setDE(r.DE() + 1)
		})

	opCodeArr[0x16] = newCommand(8, 1, "LD D, n",
		func(r *Register, m *MMU, a []uint8) {
			r.d = a[0]
		})

	opCodeArr[0x18] = newCommand(8, 1, "JR n",
		func(r *Register, m *MMU, a []uint8) {
			r.pc = r.pc + uint16(a[0])
		})

	opCodeArr[0x1a] = newCommand(8, 0, "LD A,(DE)",
		func(r *Register, m *MMU, a []uint8) {
			r.a = m.readByte(r.DE())
		})

	opCodeArr[0x1e] = newCommand(8, 1, "LD E, n",
		func(r *Register, m *MMU, a []uint8) {
			r.e = a[0]
		})

	//TODO: set ticks to be variable
	opCodeArr[0x20] = newCommand(8, 1, "JR NZ, n",
		func(r *Register, m *MMU, a []uint8) {
			if !r.f.getZero() {
				r.pc = r.pc + uint16(a[0])
			}
		})

	opCodeArr[0x21] = newCommand(12, 2, "LD HL, nn",
		func(r *Register, m *MMU, a []uint8) {
			r.setHL(toWord(a[1], a[0]))
		})

	opCodeArr[0x22] = newCommand(8, 0, "LD (HL+), A",
		func(r *Register, m *MMU, a []uint8) {
			r.setHL(r.HL() + 1)
			m.setByte(r.HL(), r.a)
		})

	opCodeArr[0x23] = newCommand(8, 0, "INC HL",
		func(r *Register, m *MMU, a []uint8) {
			r.setHL(r.HL() + 1)
		})

	opCodeArr[0x26] = newCommand(8, 1, "LD H, n",
		func(r *Register, m *MMU, a []uint8) {
			r.h = a[0]
		})

	//TODO: set ticks to be variable
	opCodeArr[0x28] = newCommand(8, 1, "JR Z, n",
		func(r *Register, m *MMU, a []uint8) {
			if r.f.getZero() {
				r.pc = r.pc + uint16(a[0])
			}
		})

	opCodeArr[0x2e] = newCommand(8, 1, "LD L, n",
		func(r *Register, m *MMU, a []uint8) {
			r.e = a[0]
		})

	opCodeArr[0x31] = newCommand(12, 2, "LD SP, nn",
		func(r *Register, m *MMU, a []uint8) {
			r.sp = toWord(a[1], a[0])
		})

	opCodeArr[0x32] = newCommand(8, 0, "LD (HL-), A",
		func(r *Register, m *MMU, a []uint8) {
			r.setHL(r.HL() - 1)
			m.setByte(r.HL(), r.a)
		})

	opCodeArr[0x3d] = newCommand(4, 0, "DEC A",
		func(r *Register, m *MMU, a []uint8) {
			r.a = dec(r.a, r.f)
		})

	//TODO: check the ticks on this
	opCodeArr[0x3e] = newCommand(8, 1, "LD A, u8",
		func(r *Register, m *MMU, a []uint8) {
			r.a = a[0]
		})

	for i, v := range [6]uint8{'b', 'c', 'd', 'e', 'h', 'l'} {

		opCodeArr[0x40+i] = newCommand(4, 0, "LD B, "+string(v),
			func(r *Register, m *MMU, a []uint8) {
				val := r.getByByte(v)
				r.b = val
			})

		opCodeArr[0x48+i] = newCommand(4, 0, "LD C, "+string(v),
			func(r *Register, m *MMU, a []uint8) {
				val := r.getByByte(v)
				r.c = val
			})

		opCodeArr[0x50+i] = newCommand(4, 0, "LD D, "+string(v),
			func(r *Register, m *MMU, a []uint8) {
				val := r.getByByte(v)
				r.d = val
			})

		opCodeArr[0x58+i] = newCommand(4, 0, "LD E, "+string(v),
			func(r *Register, m *MMU, a []uint8) {
				val := r.getByByte(v)
				r.e = val
			})

		opCodeArr[0x60+i] = newCommand(4, 0, "LD H, "+string(v),
			func(r *Register, m *MMU, a []uint8) {
				val := r.getByByte(v)
				r.h = val
			})

		opCodeArr[0x68+i] = newCommand(4, 0, "LD L, "+string(v),
			func(r *Register, m *MMU, a []uint8) {
				val := r.getByByte(v)
				r.l = val
			})
	}

	//TODO: Merge these into for loop above?
	opCodeArr[0x47] = newCommand(4, 0, "LD D, A",
		func(r *Register, m *MMU, a []uint8) {
			r.d = r.a
		})

	opCodeArr[0x57] = newCommand(4, 0, "LD D, A",
		func(r *Register, m *MMU, a []uint8) {
			r.d = r.a
		})

	opCodeArr[0x67] = newCommand(4, 0, "LD H, A",
		func(r *Register, m *MMU, a []uint8) {
			r.h = r.a
		})

	opCodeArr[0x77] = newCommand(8, 0, "LD (HL), A",
		func(r *Register, m *MMU, a []uint8) {
			m.setByte(r.HL(), r.a)
		})

	opCodeArr[0x7b] = newCommand(4, 0, "LD A, E",
		func(r *Register, m *MMU, a []uint8) {
			r.a = r.e
		})

	opCodeArr[0xaf] = newCommand(4, 0, "XOR A, A",
		func(r *Register, m *MMU, a []uint8) {
			r.a = xor(r.a, r.a, r.f)
		})

	opCodeArr[0xcd] = newCommand(24, 2, "CALL u16",
		func(r *Register, m *MMU, a []uint8) {
			call(r, m, toWord(a[1], a[0]))
		})

	opCodeArr[0xe0] = newCommand(8, 1, "LD (FF00+u8), A",
		func(r *Register, m *MMU, a []uint8) {
			m.setByte(uint16(0xff00)+uint16(a[0]), r.a)
		})

	opCodeArr[0xe2] = newCommand(8, 0, "LD (FF00+C), A",
		func(r *Register, m *MMU, a []uint8) {
			m.setByte(uint16(0xff00)+uint16(r.c), r.a)
		})

	opCodeArr[0xea] = newCommand(16, 2, "LD (u16), A",
		func(r *Register, m *MMU, a []uint8) {
			m.setByte(toWord(a[1], a[0]), r.a)
		})

	opCodeArr[0xf0] = newCommand(8, 1, "LD A, (FF00+u8)",
		func(r *Register, m *MMU, a []uint8) {
			r.a = m.readByte(uint16(0xff00) + uint16(a[0]))
		})

	opCodeArr[0xfe] = newCommand(8, 1, "CP A, u8",
		func(r *Register, m *MMU, a []uint8) {
			sub(r.a, a[0], r.f)
		})

	//build bitwise ops
	for bit := uint8(0); bit < 8; bit += 1 {
		//A register bit ops
		opCodeArr[0x100+0x47+0x08*uint16(bit)] = newCommand(8, 0, "BIT "+strconv.Itoa(int(bit))+", A",
			func(r *Register, m *MMU, a []uint8) {
				bitCheck(r.a, bit, r.f)
			})

		//B register bit ops
		opCodeArr[0x100+0x40+0x08*uint16(bit)] = newCommand(8, 0, "BIT "+strconv.Itoa(int(bit))+", B",
			func(r *Register, m *MMU, a []uint8) {
				bitCheck(r.b, bit, r.f)
			})

		//C register bit ops
		opCodeArr[0x100+0x41+0x08*uint16(bit)] = newCommand(8, 0, "BIT "+strconv.Itoa(int(bit))+", C",
			func(r *Register, m *MMU, a []uint8) {
				bitCheck(r.c, bit, r.f)
			})

		//D register bit ops
		opCodeArr[0x100+0x42+0x08*uint16(bit)] = newCommand(8, 0, "BIT "+strconv.Itoa(int(bit))+", D",
			func(r *Register, m *MMU, a []uint8) {
				bitCheck(r.d, bit, r.f)
			})

		//E register bit ops
		opCodeArr[0x100+0x43+0x08*uint16(bit)] = newCommand(8, 0, "BIT "+strconv.Itoa(int(bit))+", E",
			func(r *Register, m *MMU, a []uint8) {
				bitCheck(r.e, bit, r.f)
			})

		//H register bit ops
		opCodeArr[0x100+0x44+0x08*uint16(bit)] = newCommand(8, 0, "BIT "+strconv.Itoa(int(bit))+", H",
			func(r *Register, m *MMU, a []uint8) {
				bitCheck(r.h, bit, r.f)
			})

		//L register bit ops
		opCodeArr[0x100+0x45+0x08*uint16(bit)] = newCommand(8, 0, "BIT "+strconv.Itoa(int(bit))+", L",
			func(r *Register, m *MMU, a []uint8) {
				bitCheck(r.l, bit, r.f)
			})

		//HL register bit ops
		opCodeArr[0x100+0x46+0x08*uint16(bit)] = newCommand(12, 0, "BIT "+strconv.Itoa(int(bit))+", HL",
			func(r *Register, m *MMU, a []uint8) {
				bitCheck(m.readByte(r.HL()), bit, r.f)
			})

	}

	return opCodeArr
}

//cp and sub use the same cmd, but cp just sets the flags, not the A reg
func sub(i1, i2 uint8, f *Flags) uint8 {
	f.setZero(i1-i2 == 0)
	f.setN(true)
	f.setH((0x0f & i2) > (0x0f & i1))
	f.setCarry(i2 > i1)
	return i1 - i2
}

func xor(i1, i2 uint8, f *Flags) uint8 {
	var result uint8 = i1 ^ i2

	f.setZero(result == 0)
	f.setH(false)
	f.setN(false)
	f.setCarry(false)

	return result
}

//8 bit check
func bitCheck(val, bit uint8, f *Flags) {
	f.setH(true)
	f.setN(false)
	if bit < 8 {
		f.setZero(val&(1<<bit) != 0)
	}
}

func inc(val uint8, f *Flags) uint8 {
	newVal := (val + 1)
	f.setZero(newVal == 0)
	f.setN(false)
	f.setH((0x0f & newVal) < (0x0f & val))
	return newVal
}

func dec(val uint8, f *Flags) uint8 {
	newVal := (val - 1)
	f.setZero(newVal == 0)
	f.setN(true)
	f.setH((0x0f & newVal) == 0)
	return newVal
}

func push(r *Register, m *MMU, val uint16) {
	h := uint8(val >> 4)
	l := uint8(val & 0xff)

	r.decSP()
	m.setByte(r.sp, h)
	r.decSP()
	m.setByte(r.sp, l)
}

func call(r *Register, m *MMU, val uint16) {

	push(r, m, r.pc+2)
	r.pc = val
}
