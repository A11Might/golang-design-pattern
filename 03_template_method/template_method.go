package templatemethod

import "fmt"

type DisplayIface interface {
	Open()
	Print()
	Close()
}

// Display 抽象类(abstract class)
type Display struct {
	DisplayIface
}

func (d *Display) Display() {
	d.Open()
	for i := 0; i < 5; i++ {
		d.Print()
	}
	d.Close()
}

// CharDisplay 具体类(concrete class)
type CharDisplay struct {
	char rune
}

func NewCharDisplay(char rune) *CharDisplay {
	return &CharDisplay{
		char: char,
	}
}

func (c *CharDisplay) Open() {
	fmt.Print("<<")
}

func (c *CharDisplay) Print() {
	fmt.Printf("%c", c.char)
}

func (c *CharDisplay) Close() {
	fmt.Println(">>")
}

// StringDisplay 具体类(concrete class)
type StringDisplay struct {
	str   string
	width int
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{
		str:   str,
		width: len(str),
	}
}

func (s *StringDisplay) Open() {
	s.printLine()
}

func (s *StringDisplay) Print() {
	fmt.Printf("|%s|\n", s.str)
}

func (s *StringDisplay) Close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
