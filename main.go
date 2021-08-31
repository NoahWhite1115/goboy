package main

import (
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

	mmu := newMMU()
	mmu.loadBios(BOOTROM)
	cpu := newCPU(mmu)

	//should this run in a go routine?
	//ppu := newPPU(mmu, display)

	for {
		cpu.runCommand(true)
	}
}
