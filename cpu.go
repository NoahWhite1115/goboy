package main

import (
	"fmt"
	"log"
	"strconv"
)

type CPU struct {
	register *Register
	mmu      *MMU
	opcodes  []*Command
}

//constructor
func newCPU(mmu *MMU) *CPU {
	cpu := new(CPU)
	cpu.register = newRegister()
	cpu.opcodes = makeOpCodes()

	cpu.mmu = mmu
	return cpu
}

func (cpu *CPU) runCommand() uint8 {
	//TODO: interrupt handler

	//get next instruction code
	opcode := uint16(cpu.mmu.readByte(cpu.register.pc))
	cpu.register.incPC()

	//get command from list using code
	var command *Command

	if opcode == 0xcb {
		opcode = uint16(cpu.mmu.readByte(cpu.register.pc))
		cpu.register.incPC()
		opcode = opcode + 0x100
		command = cpu.opcodes[opcode]
	} else {
		command = cpu.opcodes[opcode]
	}

	//maybe check if command is valid
	if command == nil {
		fmt.Printf("%x \n", opcode)
		log.Fatal("Command " + strconv.Itoa(int(opcode)) + " not found.")
	}

	//extract arguments from mem
	args := make([]uint8, command.argsLen)
	for i := 0; i < command.argsLen; i++ {
		args[i] = cpu.mmu.readByte(cpu.register.pc)
		cpu.register.incPC()
	}

	//run the command
	command.printCmd(args, cpu.register)
	command.op(cpu.register, cpu.mmu, args)

	return command.ticks
}
