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
func newCPU() *CPU {
	cpu := new(CPU)
	cpu.register = newRegister()
	cpu.mmu = newMMU()
	cpu.opcodes = makeOpCodes()

	return cpu
}

func (cpu *CPU) runCommand() int {
	//TODO: interrupt handler

	//get next instruction code
	pc := cpu.register.pc
	opcode := uint16(cpu.mmu.readByte(pc))
	pc++

	//get command from list using code
	var command *Command

	if opcode == 0xcb {
		opcode = uint16(cpu.mmu.readByte(pc))
		pc++
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
		args[i] = cpu.mmu.readByte(pc)
		pc++
	}

	//run the command
	command.printCmd(args)
	command.op(cpu.register, cpu.mmu, args)

	//update program counter
	cpu.register.pc = pc

	return 1
}
