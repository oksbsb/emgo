// Peripheral: USART_Periph  Universal Synchronous Asynchronous Receiver Transmitter.
// Instances:
//  USART1  mmap.USART1_BASE
// Registers:
//  0x00 32  CR1  Control register 1.
//  0x04 32  CR2  Control register 2.
//  0x08 32  CR3  Control register 3.
//  0x0C 32  BRR  Baud rate register.
//  0x10 32  GTPR Guard time and prescaler register.
//  0x14 32  RTOR Receiver Time Out register.
//  0x18 32  RQR  Request register.
//  0x1C 32  ISR  Interrupt and status register.
//  0x20 32  ICR  Interrupt flag Clear register.
//  0x24 16  RDR  Receive Data register.
//  0x28 16  TDR  Transmit Data register.
// Import:
//  stm32/o/f030x6/mmap
package usart

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	UE     CR1 = 0x01 << 0  //+ USART Enable.
	RE     CR1 = 0x01 << 2  //+ Receiver Enable.
	TE     CR1 = 0x01 << 3  //+ Transmitter Enable.
	IDLEIE CR1 = 0x01 << 4  //+ IDLE Interrupt Enable.
	RXNEIE CR1 = 0x01 << 5  //+ RXNE Interrupt Enable.
	TCIE   CR1 = 0x01 << 6  //+ Transmission Complete Interrupt Enable.
	TXEIE  CR1 = 0x01 << 7  //+ TXE Interrupt Enable.
	PEIE   CR1 = 0x01 << 8  //+ PE Interrupt Enable.
	PS     CR1 = 0x01 << 9  //+ Parity Selection.
	PCE    CR1 = 0x01 << 10 //+ Parity Control Enable.
	WAKE   CR1 = 0x01 << 11 //+ Receiver Wakeup method.
	M      CR1 = 0x01 << 12 //+ Word Length.
	MME    CR1 = 0x01 << 13 //+ Mute Mode Enable.
	CMIE   CR1 = 0x01 << 14 //+ Character match interrupt enable.
	OVER8  CR1 = 0x01 << 15 //+ Oversampling by 8-bit or 16-bit mode.
	DEDT   CR1 = 0x1F << 16 //+ DEDT[4:0] bits (Driver Enable Deassertion Time).
	DEAT   CR1 = 0x1F << 21 //+ DEAT[4:0] bits (Driver Enable Assertion Time).
	RTOIE  CR1 = 0x01 << 26 //+ Receive Time Out interrupt enable.
	EOBIE  CR1 = 0x01 << 27 //+ End of Block interrupt enable.
)

const (
	UEn     = 0
	REn     = 2
	TEn     = 3
	IDLEIEn = 4
	RXNEIEn = 5
	TCIEn   = 6
	TXEIEn  = 7
	PEIEn   = 8
	PSn     = 9
	PCEn    = 10
	WAKEn   = 11
	Mn      = 12
	MMEn    = 13
	CMIEn   = 14
	OVER8n  = 15
	DEDTn   = 16
	DEATn   = 21
	RTOIEn  = 26
	EOBIEn  = 27
)

const (
	ADDM7    CR2 = 0x01 << 4  //+ 7-bit or 4-bit Address Detection.
	LBCL     CR2 = 0x01 << 8  //+ Last Bit Clock pulse.
	CPHA     CR2 = 0x01 << 9  //+ Clock Phase.
	CPOL     CR2 = 0x01 << 10 //+ Clock Polarity.
	CLKEN    CR2 = 0x01 << 11 //+ Clock Enable.
	STOP     CR2 = 0x03 << 12 //+ STOP[1:0] bits (STOP bits).
	SWAP     CR2 = 0x01 << 15 //+ SWAP TX/RX pins.
	RXINV    CR2 = 0x01 << 16 //+ RX pin active level inversion.
	TXINV    CR2 = 0x01 << 17 //+ TX pin active level inversion.
	DATAINV  CR2 = 0x01 << 18 //+ Binary data inversion.
	MSBFIRST CR2 = 0x01 << 19 //+ Most Significant Bit First.
	ABREN    CR2 = 0x01 << 20 //+ Auto Baud-Rate Enable.
	ABRMODE  CR2 = 0x03 << 21 //+ ABRMOD[1:0] bits (Auto Baud-Rate Mode).
	RTOEN    CR2 = 0x01 << 23 //+ Receiver Time-Out enable.
	ADD      CR2 = 0xFF << 24 //+ Address of the USART node.
)

const (
	ADDM7n    = 4
	LBCLn     = 8
	CPHAn     = 9
	CPOLn     = 10
	CLKENn    = 11
	STOPn     = 12
	SWAPn     = 15
	RXINVn    = 16
	TXINVn    = 17
	DATAINVn  = 18
	MSBFIRSTn = 19
	ABRENn    = 20
	ABRMODEn  = 21
	RTOENn    = 23
	ADDn      = 24
)

const (
	EIE    CR3 = 0x01 << 0  //+ Error Interrupt Enable.
	HDSEL  CR3 = 0x01 << 3  //+ Half-Duplex Selection.
	DMAR   CR3 = 0x01 << 6  //+ DMA Enable Receiver.
	DMAT   CR3 = 0x01 << 7  //+ DMA Enable Transmitter.
	RTSE   CR3 = 0x01 << 8  //+ RTS Enable.
	CTSE   CR3 = 0x01 << 9  //+ CTS Enable.
	CTSIE  CR3 = 0x01 << 10 //+ CTS Interrupt Enable.
	ONEBIT CR3 = 0x01 << 11 //+ One sample bit method enable.
	OVRDIS CR3 = 0x01 << 12 //+ Overrun Disable.
	DDRE   CR3 = 0x01 << 13 //+ DMA Disable on Reception Error.
	DEM    CR3 = 0x01 << 14 //+ Driver Enable Mode.
	DEP    CR3 = 0x01 << 15 //+ Driver Enable Polarity Selection.
)

const (
	EIEn    = 0
	HDSELn  = 3
	DMARn   = 6
	DMATn   = 7
	RTSEn   = 8
	CTSEn   = 9
	CTSIEn  = 10
	ONEBITn = 11
	OVRDISn = 12
	DDREn   = 13
	DEMn    = 14
	DEPn    = 15
)

const (
	DIV_FRACTION BRR = 0x0F << 0  //+ Fraction of USARTDIV.
	DIV_MANTISSA BRR = 0xFFF << 4 //+ Mantissa of USARTDIV.
)

const (
	DIV_FRACTIONn = 0
	DIV_MANTISSAn = 4
)

const (
	PSC GTPR = 0xFF << 0 //+ PSC[7:0] bits (Prescaler value).
	GT  GTPR = 0xFF << 8 //+ GT[7:0] bits (Guard time value).
)

const (
	PSCn = 0
	GTn  = 8
)

const (
	RTO  RTOR = 0xFFFFFF << 0 //+ Receiver Time Out Value.
	BLEN RTOR = 0xFF << 24    //+ Block Length.
)

const (
	RTOn  = 0
	BLENn = 24
)

const (
	ABRRQ RQR = 0x01 << 0 //+ Auto-Baud Rate Request.
	SBKRQ RQR = 0x01 << 1 //+ Send Break Request.
	MMRQ  RQR = 0x01 << 2 //+ Mute Mode Request.
	RXFRQ RQR = 0x01 << 3 //+ Receive Data flush Request.
)

const (
	ABRRQn = 0
	SBKRQn = 1
	MMRQn  = 2
	RXFRQn = 3
)

const (
	PE    ISR = 0x01 << 0  //+ Parity Error.
	FE    ISR = 0x01 << 1  //+ Framing Error.
	NE    ISR = 0x01 << 2  //+ Noise detected Flag.
	ORE   ISR = 0x01 << 3  //+ OverRun Error.
	IDLE  ISR = 0x01 << 4  //+ IDLE line detected.
	RXNE  ISR = 0x01 << 5  //+ Read Data Register Not Empty.
	TC    ISR = 0x01 << 6  //+ Transmission Complete.
	TXE   ISR = 0x01 << 7  //+ Transmit Data Register Empty.
	CTSIF ISR = 0x01 << 9  //+ CTS interrupt flag.
	CTS   ISR = 0x01 << 10 //+ CTS flag.
	RTOF  ISR = 0x01 << 11 //+ Receiver Time Out.
	ABRE  ISR = 0x01 << 14 //+ Auto-Baud Rate Error.
	ABRF  ISR = 0x01 << 15 //+ Auto-Baud Rate Flag.
	BUSY  ISR = 0x01 << 16 //+ Busy Flag.
	CMF   ISR = 0x01 << 17 //+ Character Match Flag.
	SBKF  ISR = 0x01 << 18 //+ Send Break Flag.
	RWU   ISR = 0x01 << 19 //+ Receive Wake Up from mute mode Flag.
	TEACK ISR = 0x01 << 21 //+ Transmit Enable Acknowledge Flag.
	REACK ISR = 0x01 << 22 //+ Receive Enable Acknowledge Flag.
)

const (
	PEn    = 0
	FEn    = 1
	NEn    = 2
	OREn   = 3
	IDLEn  = 4
	RXNEn  = 5
	TCn    = 6
	TXEn   = 7
	CTSIFn = 9
	CTSn   = 10
	RTOFn  = 11
	ABREn  = 14
	ABRFn  = 15
	BUSYn  = 16
	CMFn   = 17
	SBKFn  = 18
	RWUn   = 19
	TEACKn = 21
	REACKn = 22
)

const (
	PECF   ICR = 0x01 << 0  //+ Parity Error Clear Flag.
	FECF   ICR = 0x01 << 1  //+ Framing Error Clear Flag.
	NCF    ICR = 0x01 << 2  //+ Noise detected Clear Flag.
	ORECF  ICR = 0x01 << 3  //+ OverRun Error Clear Flag.
	IDLECF ICR = 0x01 << 4  //+ IDLE line detected Clear Flag.
	TCCF   ICR = 0x01 << 6  //+ Transmission Complete Clear Flag.
	CTSCF  ICR = 0x01 << 9  //+ CTS Interrupt Clear Flag.
	RTOCF  ICR = 0x01 << 11 //+ Receiver Time Out Clear Flag.
	CMCF   ICR = 0x01 << 17 //+ Character Match Clear Flag.
)

const (
	PECFn   = 0
	FECFn   = 1
	NCFn    = 2
	ORECFn  = 3
	IDLECFn = 4
	TCCFn   = 6
	CTSCFn  = 9
	RTOCFn  = 11
	CMCFn   = 17
)
