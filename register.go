package main

type Register struct {
	a, b, c, d, e, h, l uint8
	sp, pc              uint16
	ime                 bool
	f                   *Flags
}

//constructor
func newRegister() *Register {
	reg := new(Register)
	reg.f = newFlags()

	return reg
}

//sort registers by pairings

func (r *Register) AF() uint16 {
	return uint16(r.a)<<8 | uint16(r.f.flagsToByte())
}

func (r *Register) setAF(af uint16) {
	r.a = getUpper8(af)
	r.f.flagsFromByte(getLower8(af))
}

func (r *Register) BC() uint16 {
	return uint16(r.b)<<8 | uint16(r.c)
}

func (r *Register) setBC(bc uint16) {
	r.b = getUpper8(bc)
	r.c = getLower8(bc)
}

func (r *Register) DE() uint16 {
	return uint16(r.d)<<8 | uint16(r.e)
}

func (r *Register) setDE(de uint16) {
	r.d = getUpper8(de)
	r.e = getLower8(de)
}

func (r *Register) HL() uint16 {
	return uint16(r.h)<<8 | uint16(r.l)
}

func (r *Register) setHL(hl uint16) {
	r.h = getUpper8(hl)
	r.l = getLower8(hl)
}

func (r *Register) incPC() {
	r.pc = (r.pc + 1) & 0xffff
}

func (r *Register) decSP() {
	r.sp = (r.sp - 1) & 0xffff
}

func (r *Register) incSP() {
	r.sp = (r.sp + 1) & 0xffff
}

func getUpper8(val uint16) uint8 {
	return uint8(val >> 8)
}

func getLower8(val uint16) uint8 {
	return uint8(val & 0xff)
}
