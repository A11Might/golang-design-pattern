package adapter

import "fmt"

// Banner 被适配(adaptee)
type Banner struct {
	str string
}

func NewBanner(str string) *Banner {
	return &Banner{str: str}
}

func (b *Banner) PrintWithParen() {
	fmt.Printf("(%s)\n", b.str)
}

func (b *Banner) PrintWithAster() {
	fmt.Printf("*%s*\n", b.str)
}

// Print 对象(target)
type Print interface {
	PrintWeak()
	PrintStrong()
}

// PrintBanner 适配(adapter)
type PrintBanner struct {
	Banner
}

func NewPrintBanner(str string) *PrintBanner {
	return &PrintBanner{Banner: *NewBanner(str)}
}

func (p *PrintBanner) PrintWeak() {
	p.PrintWithParen()
}

func (p *PrintBanner) PrintStrong() {
	p.PrintWithAster()
}
