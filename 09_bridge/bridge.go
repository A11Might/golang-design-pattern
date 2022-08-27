package bridge

import "fmt"

type IDisplay interface {
	DisplaySth()
}

// Display 抽象部分(abstraction)
type Display struct {
	DisplayImpl
}

func NewDisplay(impl DisplayImpl) IDisplay {
	return &Display{
		DisplayImpl: impl,
	}
}

func (d *Display) Open() {
	d.RawOpen()
}

func (d *Display) Print() {
	d.RawPrint()
}

func (d *Display) Close() {
	d.RawClose()
}

func (d *Display) DisplaySth() {
	d.Open()
	d.Print()
	d.Close()
}

// CountDisplay 改善后的抽象部分(refined abstraction)
type CountDisplay struct {
	*Display
}

func NewCountDisplay(impl DisplayImpl) IDisplay {
	return &CountDisplay{
		Display: NewDisplay(impl).(*Display),
	}
}

func (c *CountDisplay) MultiDisplaySth(times int) {
	c.Open()
	for i := 0; i < times; i++ {
		c.Print()
	}
	c.Close()
}

// DisplayImpl 实现者(implementor)
type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

// StringDisplayImpl 具体实现者(concrete implementor)
type StringDisplayImpl struct {
	str   string
	width int
}

func NewStringDisplayImpl(str string) DisplayImpl {
	return &StringDisplayImpl{
		str:   str,
		width: len(str),
	}
}

func (s *StringDisplayImpl) RawOpen() {
	s.printLine()
}
func (s *StringDisplayImpl) RawPrint() {
	fmt.Println("|" + s.str + "|")
}

func (s *StringDisplayImpl) RawClose() {
	s.printLine()
}

func (s *StringDisplayImpl) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
