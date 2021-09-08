package main

import (
	"log"
	"os"
)

type Cartridge struct {
	header, rom0, rom1 []uint8
}

func cartridgeHandler(filename string) *Cartridge {
	cart := new(Cartridge)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Can't read file")
	}

	defer f.Close()
	throwaway := make([]uint8, 0x100)
	f.Read(throwaway)
	header := make([]uint8, 0x50)
	f.Read(header)

	cartRom0 := make([]uint8, 0x3EB0)
	f.Read(cartRom0)

	cartRom1 := make([]uint8, 0x4000)
	f.Read(cartRom1)

	cart.header = header
	cart.rom0 = cartRom0
	cart.rom1 = cartRom1

	return cart
}
