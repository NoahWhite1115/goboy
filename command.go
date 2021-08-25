package main

import "log"

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

func makeOpCodes() []*Command {
	opCodeArr := make([]*Command, 0x1ff)

	opCodeArr[0x00] = newCommand(4, 0, "NOP",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("NOP")
		})

	opCodeArr[0x01] = newCommand(12, 2, "LD BC, nn",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("LD BC, " + string(a[1]) + string(a[0]))
			r.setBC(toWord(a[1], a[0]))
		})

	opCodeArr[0x06] = newCommand(8, 1, "LD B, n",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("LD B, %x", string(a[0]))
			r.b = a[0]
		})

	opCodeArr[0x11] = newCommand(12, 2, "LD DE, nn",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("LD DE, %x", string(a[0]))
			r.setDE(toWord(a[1], a[0]))
		})

	opCodeArr[0x21] = newCommand(12, 2, "LD HL, nn",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("LD HL, %x", string(a[0]))
			r.setHL(toWord(a[1], a[0]))
		})

	opCodeArr[0x31] = newCommand(12, 2, "LD SP, nn",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("LD SP, %x", string(a[0]))
			r.sp = toWord(a[1], a[0])
		})

	opCodeArr[0x32] = newCommand(8, 0, "LD (HLD), A",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("LD (HLD), A")
			r.decHL()

			m.setByte(r.HL(), r.a)
		})

	opCodeArr[0x50] = newCommand(4, 0, "LD D, B",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("LD D, B")
			r.d = r.b
		})

	opCodeArr[0xAF] = newCommand(4, 0, "XOR A, A",
		func(r *Register, m *MMU, a []uint8) {
			log.Println("XOR A, A")
			r.a = xor(r.a, r.a, r.f)
		})

	return opCodeArr
}

func xor(i1, i2 uint8, f *Flags) uint8 {
	var result uint8 = i1 ^ i2

	//TODO: set other flags
	f.setZero(result == 0)
	f.setCarry(false)

	return result
}
