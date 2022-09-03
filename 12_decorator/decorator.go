package decorator

import (
	"fmt"
	"strings"
)

// Display 组件(component)
type Display struct {
	DisplayImpl
}

func (d *Display) Show() {
	for i := 0; i < d.GetRows(); i++ {
		fmt.Println(d.GetRowText(i))
	}
}

type DisplayImpl interface {
	GetColumns() int
	GetRows() int
	GetRowText(int) string
}

// StringDisplay 具体组件(concrete component)
type StringDisplay struct {
	str string
}

func NewStringDisplay(str string) *Display {
	return &Display{
		DisplayImpl: &StringDisplay{
			str: str,
		},
	}
}

func (s *StringDisplay) GetColumns() int {
	return len(s.str)
}

func (s *StringDisplay) GetRows() int {
	return 1
}

func (s *StringDisplay) GetRowText(row int) string {
	if row == 0 {
		return s.str
	} else {
		return ""
	}
}

// Border 装饰物(decorator)

// SideBorder 具体的装饰物(concrete decorator)
type SideBorder struct {
	borderChar rune
	display    *Display
}

func NewSideBorder(display *Display, borderChar rune) *Display {
	return &Display{
		DisplayImpl: &SideBorder{
			borderChar: borderChar,
			display:    display,
		},
	}
}

func (s *SideBorder) GetColumns() int {
	return 1 + s.display.GetColumns() + 1
}

func (s *SideBorder) GetRows() int {
	return s.display.GetRows()
}

func (s *SideBorder) GetRowText(row int) string {
	return fmt.Sprintf("%c%s%c", s.borderChar, s.display.GetRowText(row), s.borderChar)
}

// FullBorder 具体的装饰物(concrete decorator)
type FullBorder struct {
	display *Display
}

func NewFullBorder(display *Display) *Display {
	return &Display{
		DisplayImpl: &FullBorder{
			display: display,
		},
	}
}

func (f *FullBorder) GetColumns() int {
	return 1 + f.display.GetColumns() + 1
}

func (f *FullBorder) GetRows() int {
	return 1 + f.display.GetRows() + 1
}

func (f *FullBorder) GetRowText(row int) string {
	if row == 0 ||
		row == (f.display.GetRows()+1) {
		return fmt.Sprintf("+%s+", f.makeLine('-', f.display.GetColumns()))
	} else {
		return fmt.Sprintf("|%s|", f.display.GetRowText(row-1))
	}
}

func (f *FullBorder) makeLine(ch rune, count int) string {
	buf := new(strings.Builder)
	for i := 0; i < count; i++ {
		buf.WriteRune(ch)
	}
	return buf.String()
}
