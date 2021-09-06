package main

import (
	"log"
	"os"
)

type Cartridge struct {
	header, rom0 []uint8
}

func cartridgeHandler(filename string) *Cartridge {
	cart := new(Cartridge)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't read file")
	}

	defer f.Close()
	header := make([]uint8, 0x50)
	f.Read(header)

	cartRom0 := make([]uint8, 0x3EAF)
	f.Read(cartRom0)

	cart.header = header
	cart.rom0 = cartRom0

	return cart
}
