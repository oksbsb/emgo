// +build f10x_md

package ob

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_md/mmap"
)

type OB_Periph struct {
	RDP   RRDP
	USER  RUSER
	Data0 RData0
	Data1 RData1
	WRP0  RWRP0
	WRP1  RWRP1
	WRP2  RWRP2
	WRP3  RWRP3
}

func (p *OB_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var OB = (*OB_Periph)(unsafe.Pointer(uintptr(mmap.OB_BASE)))

type RDP uint16

func (b RDP) Field(mask RDP) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDP) J(v int) RDP {
	return RDP(bits.MakeField32(v, uint32(mask)))
}

type RRDP struct{ mmio.U16 }

func (r *RRDP) Bits(mask RDP) RDP     { return RDP(r.U16.Bits(uint16(mask))) }
func (r *RRDP) StoreBits(mask, b RDP) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RRDP) SetBits(mask RDP)      { r.U16.SetBits(uint16(mask)) }
func (r *RRDP) ClearBits(mask RDP)    { r.U16.ClearBits(uint16(mask)) }
func (r *RRDP) Load() RDP             { return RDP(r.U16.Load()) }
func (r *RRDP) Store(b RDP)           { r.U16.Store(uint16(b)) }

type RMRDP struct{ mmio.UM16 }

func (rm RMRDP) Load() RDP   { return RDP(rm.UM16.Load()) }
func (rm RMRDP) Store(b RDP) { rm.UM16.Store(uint16(b)) }

type USER uint16

func (b USER) Field(mask USER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask USER) J(v int) USER {
	return USER(bits.MakeField32(v, uint32(mask)))
}

type RUSER struct{ mmio.U16 }

func (r *RUSER) Bits(mask USER) USER    { return USER(r.U16.Bits(uint16(mask))) }
func (r *RUSER) StoreBits(mask, b USER) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RUSER) SetBits(mask USER)      { r.U16.SetBits(uint16(mask)) }
func (r *RUSER) ClearBits(mask USER)    { r.U16.ClearBits(uint16(mask)) }
func (r *RUSER) Load() USER             { return USER(r.U16.Load()) }
func (r *RUSER) Store(b USER)           { r.U16.Store(uint16(b)) }

type RMUSER struct{ mmio.UM16 }

func (rm RMUSER) Load() USER   { return USER(rm.UM16.Load()) }
func (rm RMUSER) Store(b USER) { rm.UM16.Store(uint16(b)) }

type Data0 uint16

func (b Data0) Field(mask Data0) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask Data0) J(v int) Data0 {
	return Data0(bits.MakeField32(v, uint32(mask)))
}

type RData0 struct{ mmio.U16 }

func (r *RData0) Bits(mask Data0) Data0   { return Data0(r.U16.Bits(uint16(mask))) }
func (r *RData0) StoreBits(mask, b Data0) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RData0) SetBits(mask Data0)      { r.U16.SetBits(uint16(mask)) }
func (r *RData0) ClearBits(mask Data0)    { r.U16.ClearBits(uint16(mask)) }
func (r *RData0) Load() Data0             { return Data0(r.U16.Load()) }
func (r *RData0) Store(b Data0)           { r.U16.Store(uint16(b)) }

type RMData0 struct{ mmio.UM16 }

func (rm RMData0) Load() Data0   { return Data0(rm.UM16.Load()) }
func (rm RMData0) Store(b Data0) { rm.UM16.Store(uint16(b)) }

type Data1 uint16

func (b Data1) Field(mask Data1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask Data1) J(v int) Data1 {
	return Data1(bits.MakeField32(v, uint32(mask)))
}

type RData1 struct{ mmio.U16 }

func (r *RData1) Bits(mask Data1) Data1   { return Data1(r.U16.Bits(uint16(mask))) }
func (r *RData1) StoreBits(mask, b Data1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RData1) SetBits(mask Data1)      { r.U16.SetBits(uint16(mask)) }
func (r *RData1) ClearBits(mask Data1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RData1) Load() Data1             { return Data1(r.U16.Load()) }
func (r *RData1) Store(b Data1)           { r.U16.Store(uint16(b)) }

type RMData1 struct{ mmio.UM16 }

func (rm RMData1) Load() Data1   { return Data1(rm.UM16.Load()) }
func (rm RMData1) Store(b Data1) { rm.UM16.Store(uint16(b)) }

type WRP0 uint16

func (b WRP0) Field(mask WRP0) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP0) J(v int) WRP0 {
	return WRP0(bits.MakeField32(v, uint32(mask)))
}

type RWRP0 struct{ mmio.U16 }

func (r *RWRP0) Bits(mask WRP0) WRP0    { return WRP0(r.U16.Bits(uint16(mask))) }
func (r *RWRP0) StoreBits(mask, b WRP0) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RWRP0) SetBits(mask WRP0)      { r.U16.SetBits(uint16(mask)) }
func (r *RWRP0) ClearBits(mask WRP0)    { r.U16.ClearBits(uint16(mask)) }
func (r *RWRP0) Load() WRP0             { return WRP0(r.U16.Load()) }
func (r *RWRP0) Store(b WRP0)           { r.U16.Store(uint16(b)) }

type RMWRP0 struct{ mmio.UM16 }

func (rm RMWRP0) Load() WRP0   { return WRP0(rm.UM16.Load()) }
func (rm RMWRP0) Store(b WRP0) { rm.UM16.Store(uint16(b)) }

type WRP1 uint16

func (b WRP1) Field(mask WRP1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1) J(v int) WRP1 {
	return WRP1(bits.MakeField32(v, uint32(mask)))
}

type RWRP1 struct{ mmio.U16 }

func (r *RWRP1) Bits(mask WRP1) WRP1    { return WRP1(r.U16.Bits(uint16(mask))) }
func (r *RWRP1) StoreBits(mask, b WRP1) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RWRP1) SetBits(mask WRP1)      { r.U16.SetBits(uint16(mask)) }
func (r *RWRP1) ClearBits(mask WRP1)    { r.U16.ClearBits(uint16(mask)) }
func (r *RWRP1) Load() WRP1             { return WRP1(r.U16.Load()) }
func (r *RWRP1) Store(b WRP1)           { r.U16.Store(uint16(b)) }

type RMWRP1 struct{ mmio.UM16 }

func (rm RMWRP1) Load() WRP1   { return WRP1(rm.UM16.Load()) }
func (rm RMWRP1) Store(b WRP1) { rm.UM16.Store(uint16(b)) }

type WRP2 uint16

func (b WRP2) Field(mask WRP2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP2) J(v int) WRP2 {
	return WRP2(bits.MakeField32(v, uint32(mask)))
}

type RWRP2 struct{ mmio.U16 }

func (r *RWRP2) Bits(mask WRP2) WRP2    { return WRP2(r.U16.Bits(uint16(mask))) }
func (r *RWRP2) StoreBits(mask, b WRP2) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RWRP2) SetBits(mask WRP2)      { r.U16.SetBits(uint16(mask)) }
func (r *RWRP2) ClearBits(mask WRP2)    { r.U16.ClearBits(uint16(mask)) }
func (r *RWRP2) Load() WRP2             { return WRP2(r.U16.Load()) }
func (r *RWRP2) Store(b WRP2)           { r.U16.Store(uint16(b)) }

type RMWRP2 struct{ mmio.UM16 }

func (rm RMWRP2) Load() WRP2   { return WRP2(rm.UM16.Load()) }
func (rm RMWRP2) Store(b WRP2) { rm.UM16.Store(uint16(b)) }

type WRP3 uint16

func (b WRP3) Field(mask WRP3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP3) J(v int) WRP3 {
	return WRP3(bits.MakeField32(v, uint32(mask)))
}

type RWRP3 struct{ mmio.U16 }

func (r *RWRP3) Bits(mask WRP3) WRP3    { return WRP3(r.U16.Bits(uint16(mask))) }
func (r *RWRP3) StoreBits(mask, b WRP3) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *RWRP3) SetBits(mask WRP3)      { r.U16.SetBits(uint16(mask)) }
func (r *RWRP3) ClearBits(mask WRP3)    { r.U16.ClearBits(uint16(mask)) }
func (r *RWRP3) Load() WRP3             { return WRP3(r.U16.Load()) }
func (r *RWRP3) Store(b WRP3)           { r.U16.Store(uint16(b)) }

type RMWRP3 struct{ mmio.UM16 }

func (rm RMWRP3) Load() WRP3   { return WRP3(rm.UM16.Load()) }
func (rm RMWRP3) Store(b WRP3) { rm.UM16.Store(uint16(b)) }
