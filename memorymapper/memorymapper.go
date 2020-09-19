package memorymapper

import (
	"log"
	"oasis_vm/memory"
)

type Region struct {
	device memory.MemoryInterface
	start  int
	end    int
	remap  bool
}

type MemoryMapper struct {
	regions []Region
}

func NewMemoryMapper() *MemoryMapper {
	return &MemoryMapper{regions: []Region{}}
}

func (mp *MemoryMapper) MapDevice(device memory.MemoryInterface, start int, end int, remap bool) func() {
	region := Region{
		device,
		start,
		end,
		remap,
	}

	mp.regions = append([]Region{region}, mp.regions...)
	return func() {
		for i, r := range mp.regions {
			if r == region {
				s1 := mp.regions[:i]
				s2 := mp.regions[i+1:]
				mp.regions = append(s1, s2...)
			}
		}
	}
}

func (mp *MemoryMapper) findRegion(address int) *Region {
	for _, r := range mp.regions {
		if address >= r.start && address <= r.end {
			return &r
		}
	}

	log.Fatal("No memory region found for address")
	return nil
}

func (mp *MemoryMapper) GetUint16(address int) int {
	region := mp.findRegion(address)
	var finalAddress int
	if region.remap {
		finalAddress = address - region.start
	} else {
		finalAddress = address
	}
	return region.device.GetUint16(finalAddress)
}

func (mp *MemoryMapper) GetUint8(address int) byte {
	region := mp.findRegion(address)
	var finalAddress int
	if region.remap {
		finalAddress = address - region.start
	} else {
		finalAddress = address
	}
	return region.device.GetUint8(finalAddress)
}

func (mp *MemoryMapper) SetUint16(address int, value int) {
	region := mp.findRegion(address)
	var finalAddress int
	if region.remap {
		finalAddress = address - region.start
	} else {
		finalAddress = address
	}
	region.device.SetUint16(finalAddress, value)
}

func (mp *MemoryMapper) SetUint8(address int, value int) {
	region := mp.findRegion(address)
	var finalAddress int
	if region.remap {
		finalAddress = address - region.start
	} else {
		finalAddress = address
	}
	region.device.SetUint8(finalAddress, value)
}
