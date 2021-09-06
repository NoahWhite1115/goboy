package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	//display := flag.Bool("display", false, "display an SDL interface")

	cart := cartridgeHandler("./blargs/01-special.gb")
	mmu := newMMUNoBios(cart)
	//mmu.loadBios(BOOTROM)
	cpu := newCPUNoBios(mmu)

	//should this run in a go routine?
	//ppu := newPPU(mmu, display)

	for {
		cpu.runCommand(true)

		if mmu.readByte(0xff02) == 0x81 {
			c := mmu.readByte(0xff01)
			fmt.Printf("%d", c)
			mmu.setByte(0xff02, 0x0)
		}
	}
}
