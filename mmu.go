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

//TODO: split out into own file
func newMMUNoBios(cart *Cartridge) *MMU {
	mmu := new(MMU)

	mmu.addToMMU(0x00, 0xFF)            //restart + interrupts?
	mmu.addExisting(0x100, cart.header) //cartridge header
	mmu.addExisting(0x150, cart.rom0)   //cart rom 0
	mmu.addToMMU(0x4000, 0x7FFF)        //cart rom 1-xx
	mmu.addToMMU(0x8000, 0x97FF)        //char ram
	mmu.addToMMU(0x9800, 0x9BFF)        //BG map data 1
	mmu.addToMMU(0x9C00, 0x9FFF)        //BG map data 2
	mmu.addToMMU(0xA000, 0xBFFF)        //Cart ram
	mmu.addToMMU(0xC000, 0xCFFF)        //internal ram 1
	mmu.addToMMU(0xD000, 0xDFFF)        //internal ram 2
	mmu.addToMMU(0xFE00, 0xFE9F)        //OAM
	mmu.addToMMU(0xFF00, 0xFF7F)        //hardware regs
	mmu.addToMMU(0xFF80, 0xFFFE)        //High RAM
	mmu.addToMMU(0xFFFF, 0xFFFF)        //Interrupt reg

	//use the values provided by pandocs to set the gb state without running bios
	mmu.setByte(0xff05, 0x00)
	mmu.setByte(0xff06, 0x00)
	mmu.setByte(0xff07, 0x00)
	mmu.setByte(0xff10, 0x80)
	mmu.setByte(0xff11, 0xbf)
	mmu.setByte(0xff12, 0xf3)
	mmu.setByte(0xff14, 0xbf)
	mmu.setByte(0xff16, 0x3f)
	mmu.setByte(0xff17, 0x00)
	mmu.setByte(0xff19, 0xbf)
	mmu.setByte(0xff1a, 0x7f)
	mmu.setByte(0xff1b, 0xff)
	mmu.setByte(0xff1c, 0x9f)
	mmu.setByte(0xff1e, 0xbf)
	mmu.setByte(0xff20, 0xff)
	mmu.setByte(0xff21, 0x00)
	mmu.setByte(0xff22, 0x00)
	mmu.setByte(0xff23, 0xbf)
	mmu.setByte(0xff24, 0x77)
	mmu.setByte(0xff25, 0xf3)
	mmu.setByte(0xff26, 0xF1)
	mmu.setByte(0xff40, 0x91)
	mmu.setByte(0xff42, 0x00)
	mmu.setByte(0xff43, 0x00)
	mmu.setByte(0xff45, 0x00)
	mmu.setByte(0xff47, 0xFC)
	mmu.setByte(0xff48, 0xFF)
	mmu.setByte(0xff49, 0xFF)
	mmu.setByte(0xff4a, 0x00)
	mmu.setByte(0xff4b, 0x00)
	mmu.setByte(0xffff, 0x00)

	mmu.biosDone = true

	return mmu
}

func (mmu *MMU) loadBios(bios []uint8) {
	mmu.bios = bios
	mmu.biosDone = false
}

func (mmu *MMU) readByte(address uint16) uint8 {
	if address < 0xff && mmu.biosDone == false {
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

//TODO: better way to handle cart rom?
func (mmu *MMU) addExisting(offset uint16, data []uint8) {
	mmu.offsets = append(mmu.offsets, offset)
	mmu.addr = append(mmu.addr, data)
}
