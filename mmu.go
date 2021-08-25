package main

//TODO: reimplement mmu. Current version is for test only
type MMU struct {
	bios []uint8
	addr []uint8
}

func newMMU() *MMU {
	mmu := new(MMU)
	mmu.bios = BOOTROM
	mmu.addr = make([]uint8, 0x10000)
	return mmu
}

func (mmu *MMU) readByte(address uint16) uint8 {
	return mmu.bios[address]
}

func (mmu *MMU) setByte(address uint16, data uint8) {
	mmu.addr[address] = data
}

/*
func (mmu *MMU) addToMMU(offset, size int) {
	//check to ensure it's in range

	//add new offset to slice if null
	if len(mmu.offsets) == 0 {
		mmu.offsets = append(mmu.offsets)
	} else {
		for i, val := range mmu.offsets {
			//if i >
		}
	}

}*/
