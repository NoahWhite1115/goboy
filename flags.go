package main

type Flags struct {
	//flagarr
	//0: Carry
	//1: BCD
	//2: BCD
	//3: Zero
	flagarr []bool
}

func newFlags() *Flags {
	flags := new(Flags)
	flags.flagarr = make([]bool, 4)

	return flags
}

func (f *Flags) setCarry(in bool) {
	f.flagarr[0] = in
}

func (f *Flags) getCarry() bool {
	return (f.flagarr[0])
}

func (f *Flags) setZero(in bool) {
	f.flagarr[3] = in
}

func (f *Flags) getZero() bool {
	return (f.flagarr[3])
}

func (f *Flags) flagsToByte() uint8 {
	var out uint8 = 0

	for i, v := range f.flagarr {
		if v {
			out |= (1<<i + 4)
		}
	}

	return out
}

func (f *Flags) flagsFromByte(in uint8) {
	for i := 0; i < 4; i++ {
		if (in>>(i+4) | 1) == 0 {
			f.flagarr[i] = false
		} else {
			f.flagarr[i] = true
		}
	}
}
