package sdmmc

import (
	"rtos"
	"sync/fence"

	"sdcard"

	"stm32/hal/exti"
	"stm32/hal/gpio"
)

// setClock don't use PwrSave mode because continous clock is required during
// initialization and multiple block write (and maybe more cases). Using
// PwrSave mode seems to be impractical. Low power application should disable
// the whole peripheral if not used.
func setClock(p *Periph, freqhz int) {
	var (
		clkdiv int
		cfg    BusClock
	)
	busWidth, _ := p.BusClock()
	busWidth &= BusWidth
	if freqhz > 0 {
		// BUG: This code assumes 48 MHz SDMMCCLK.
		cfg = ClkEna
		clkdiv = (48e6+freqhz-1)/freqhz - 2
	}
	if clkdiv < 0 {
		clkdiv = 0
		cfg |= ClkByp
	}
	p.SetBusClock(cfg|busWidth, clkdiv)
	p.SetDataTimeout(uint(freqhz)) // ≈ 1s
}

func setBusWidth(p *Periph, width sdcard.BusWidth) {
	if width > sdcard.Bus8 {
		panic("sdmmc: bad bus width")
	}
	cfg, clkdiv := p.BusClock()
	cfg = cfg&^BusWidth | BusClock(width*3>>2)<<3
	p.SetBusClock(cfg, clkdiv)
}

func wait(d0 gpio.Pin, done *rtos.EventFlag, deadline int64) bool {
	if d0.Load() != 0 {
		return true // Fast path.
	}
	done.Reset(0)
	l := exti.Lines(d0.Mask())
	l.ClearPending()
	fence.W() // Order writes to normal and I/O memory.
	l.EnableIRQ()
	if d0.Load() != 0 {
		l.DisableIRQ()
		return true
	}
	return done.Wait(1, deadline)
}

func busyISR(d0 gpio.Pin, done *rtos.EventFlag) {
	exti.Lines(d0.Mask()).DisableIRQ()
	done.Signal(1)
}