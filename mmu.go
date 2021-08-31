package main

//TODO: reimplement mmu. Current version is for test only
type MMU struct {
	bios     []uint8
	biosDone bool
	addr     [][]uint8
	offsets  []uint16
}

func newMMU() *MMU {
	mmu := new(MMU)

	mmu.addToMMU(0x00, 0xFF)     //restart + interrupts?
	mmu.addToMMU(0x100, 0x14F)   //cartridge header
	mmu.addToMMU(0x150, 0x3FFF)  //cart rom 0
	mmu.addToMMU(0x4000, 0x7FFF) //cart rom 1-xx
	mmu.addToMMU(0x8000, 0x97FF) //char ram
	mmu.addToMMU(0x9800, 0x9BFF) //BG map data 1
	mmu.addToMMU(0x9C00, 0x9FFF) //BG map data 2
	mmu.addToMMU(0xA000, 0xBFFF) //Cart ram
	mmu.addToMMU(0xC000, 0xCFFF) //internal ram 1
	mmu.addToMMU(0xD000, 0xDFFF) //internal ram 2
	mmu.addToMMU(0xFE00, 0xFE9F) //OAM
	mmu.addToMMU(0xFF00, 0xFF7F) //hardware regs
	mmu.addToMMU(0xFF80, 0xFFFE) //High RAM
	mmu.addToMMU(0xFFFF, 0xFFFF) //Interrupt reg

	return mmu
}

func (mmu *MMU) loadBios(bios []uint8) {
	mmu.bios = bios
	mmu.biosDone = false
}

func (mmu *MMU) readByte(address uint16) uint8 {
	if address < 0xff {
		return mmu.bios[address]
	}

	for i, v := range mmu.offsets {
		if v == address {
			return mmu.addr[i][0]
		} else if v > address {
			return mmu.addr[i-1][address-mmu.offsets[i-1]]
		}
	}

	//return 00 if not found
	//TODO: make this better!
	return 0x00
}

func (mmu *MMU) setByte(address uint16, data uint8) {
	for i, v := range mmu.offsets {
		if v == address {
			mmu.addr[i][0] = data
			break
		} else if v > address {
			mmu.addr[i-1][address-mmu.offsets[i-1]] = data
			break
		}
	}
}

func (mmu *MMU) addToMMU(offset, finish uint16) {
	size := finish - offset + 1
	mmu.offsets = append(mmu.offsets, offset)
	mmu.addr = append(mmu.addr, make([]uint8, size))
}
