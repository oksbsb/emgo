package main

import (
	"rtos"
	"strconv"

	"stm32/hal/dma"
	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
	"stm32/hal/usart"
)

var tts *usart.Driver

func init() {
	system.SetupPLL(8, 1, 48/8)
	systick.Setup(2e6)

	gpio.A.EnableClock(true)
	tx := gpio.A.Pin(9)

	tx.Setup(&gpio.Config{Mode: gpio.Alt})
	tx.SetAltFunc(gpio.USART1_AF1)
	d := dma.DMA1
	d.EnableClock(true)
	tts = usart.NewDriver(usart.USART1, d.Channel(2, 0), nil, nil)
	tts.Periph().EnableClock(true)
	tts.Periph().SetBaudRate(115200)
	tts.Periph().Enable()
	tts.EnableTx()

	rtos.IRQ(irq.USART1).Enable()
	rtos.IRQ(irq.DMA1_Channel2_3).Enable()
}

func main() {
	b := -123
	strconv.WriteInt(tts, b, 10, 0, 0)
	tts.WriteString("\r\n")
	strconv.WriteInt(tts, b, 10, 6, ' ')
	tts.WriteString("\r\n")
	strconv.WriteInt(tts, b, 10, 6, '0')
	tts.WriteString("\r\n")
	strconv.WriteInt(tts, b, 10, 6, '.')
	tts.WriteString("\r\n")
	strconv.WriteInt(tts, b, 10, -6, ' ')
	tts.WriteString("\r\n")
	strconv.WriteInt(tts, b, 10, -6, '0')
	tts.WriteString("\r\n")
	strconv.WriteInt(tts, b, 10, -6, '.')
	tts.WriteString("\r\n")
}

func ttsISR() {
	tts.ISR()
}

func ttsDMAISR() {
	tts.TxDMAISR()
}

//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.USART1:          ttsISR,
	irq.DMA1_Channel2_3: ttsDMAISR,
}
