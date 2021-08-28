package main

func main() {
	cpu := newCPU()

	for {
		cpu.runCommand()
	}
}
