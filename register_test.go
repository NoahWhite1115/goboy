package main

import "testing"

func TestRegisterBC(t *testing.T) {
	register := newRegister()

	register.setBC(0xffcc)
	if register.b != 0xff || register.c != 0xcc || register.BC() != 0xffcc {
		t.Errorf("Register BC was incorrect, got %d %d, wanted %d", register.b, register.c, 0xffcc)
	}
}

func TestRegisterDE(t *testing.T) {
	register := newRegister()

	register.setDE(0xffcc)
	if register.d != 0xff || register.e != 0xcc || register.DE() != 0xffcc {
		t.Errorf("Register DE was incorrect, got %d %d, wanted %d", register.d, register.e, 0xffcc)
	}
}

func TestRegisterHL(t *testing.T) {
	register := newRegister()

	register.setHL(0xffcc)
	if register.h != 0xff || register.l != 0xcc || register.HL() != 0xffcc {
		t.Errorf("Register HL was incorrect, got %d %d, wanted %d", register.h, register.l, 0xffcc)
	}
}

func TestRegisterAF(t *testing.T) {
	register := newRegister()

	register.setAF(0xffc0)
	if register.a != 0xff || register.f.getCarry() != false || register.f.getZero() != true || register.AF() != 0xffc0 {
		t.Errorf("Register AF was incorrect, got %d %d, wanted %d", register.a, register.f.flagsToByte(), 0xffc0)
	}
}

func TestSetByByte(t *testing.T) {
	register := newRegister()

	register.setByByte('a', 0xff)
	register.setByByte('b', 0xfe)
	register.setByByte('h', 0x01)

	if register.a != 0xff || register.b != 0xfe || register.h != 0x01 {
		t.Errorf("Set By Byte is failing.")
	}
}

func TestGetByByte(t *testing.T) {
	register := newRegister()

	register.a = 0xff
	register.b = 0xfe
	register.h = 0x01

	if register.getByByte('a') != 0xff || register.getByByte('b') != 0xfe || register.getByByte('h') != 0x01 {
		t.Errorf("Get By Byte is failing.")
	}
}
