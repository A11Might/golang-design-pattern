package proxy

import (
	"fmt"
	"sync"
	"time"
)

// Printable 主体(subject)
type Printable interface {
	SetPrinterName(string)
	GetPrinterName() string
	Print(string)
}

// Printer 实际的主体(real subject)
type Printer struct {
	name string
}

func NewPrinter(name string) *Printer {
	p := &Printer{
		name: name,
	}
	p.heavyJob(fmt.Sprintf("正在生成 Printer 的实例(%s)", name))
	return p
}

func (p *Printer) SetPrinterName(name string) {
	p.name = name
}

func (p *Printer) GetPrinterName() string {
	return p.name
}

func (p *Printer) Print(str string) {
	fmt.Printf("=== %s ===\n", p.name)
	fmt.Println(str)
}

func (p *Printer) heavyJob(msg string) {
	fmt.Print(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
	fmt.Println("结束。")
}

// PrinterProxy 代理人(proxy)
type PrinterProxy struct {
	name string
	real *Printer
	sync.Mutex
}

func NewPrinterProxy(name string) *PrinterProxy {
	return &PrinterProxy{
		name: name,
	}
}

func (p *PrinterProxy) SetPrinterName(name string) {
	p.Lock()
	defer p.Unlock()
	if p.real != nil {
		p.real.SetPrinterName(name)
	}
	p.name = name
}

func (p *PrinterProxy) GetPrinterName() string {
	return p.name
}

func (p *PrinterProxy) Print(str string) {
	p.realize()
	p.real.Print(str)
}

func (p *PrinterProxy) realize() {
	p.Lock()
	defer p.Unlock()
	if p.real == nil {
		p.real = NewPrinter(p.name)
	}
}
