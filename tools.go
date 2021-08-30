package main

func toWord(msb, lsb uint8) uint16 {
	word := uint16(msb)<<8 | uint16(lsb)
	return word
}

func getUpper8(val uint16) uint8 {
	return uint8(val >> 8)
}

func getLower8(val uint16) uint8 {
	return uint8(val & 0xff)
}
