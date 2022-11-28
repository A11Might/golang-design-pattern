package templatemethod

import "fmt"

type IDisplay interface {
	Open()
	Print()
	Close()
}

// Display 抽象类(abstract class)
type Display struct {
	i IDisplay
}

func (d *Display) Show() {
	d.i.Open()
	for i := 0; i < 5; i++ {
		d.i.Print()
	}
	d.i.Close()
}

// CharDisplay 具体类(concrete class)
type CharDisplay struct {
	Display

	char rune
}

func NewCharDisplay(char rune) *CharDisplay {
	cd := &CharDisplay{
		char: char,
	}
	cd.i = cd
	return cd
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
	Display

	str   string
	width int
}

func NewStringDisplay(str string) *StringDisplay {
	sd := &StringDisplay{
		str:   str,
		width: len(str),
	}
	sd.i = sd
	return sd
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
