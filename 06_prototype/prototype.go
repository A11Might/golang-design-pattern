package prototype

import "fmt"

// Product 原型(prototype)
type Product interface {
	Use(s string)
	CreateClone() Product
}

// Manager 使用者(client)
type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{
		showcase: make(map[string]Product),
	}
}

func (m *Manager) Register(name string, proto Product) {
	m.showcase[name] = proto
}

func (m *Manager) Create(protoName string) Product {
	p := m.showcase[protoName]
	return p.CreateClone()
}

// MessageBox 具体的原型(concrete prototype)
type MessageBox struct {
	decoChar rune
}

func NewMessageBox(decoChar rune) *MessageBox {
	return &MessageBox{
		decoChar: decoChar,
	}
}

func (m *MessageBox) Use(s string) {
	for i := 0; i < len(s)+4; i++ {
		fmt.Printf("%c", m.decoChar)
	}
	fmt.Println()
	fmt.Printf("%c%s%c\n", m.decoChar, s, m.decoChar)
	for i := 0; i < len(s)+4; i++ {
		fmt.Printf("%c", m.decoChar)
	}
	fmt.Println()
}

func (m *MessageBox) CreateClone() Product {
	return &MessageBox{
		decoChar: m.decoChar,
	}
}

// UnderlinePen 具体的原型(concrete prototype)
type UnderlinePen struct {
	ulChar rune
}

func NewUnderlinePen(ulChar rune) *UnderlinePen {
	return &UnderlinePen{
		ulChar: ulChar,
	}
}

func (u *UnderlinePen) Use(s string) {
	fmt.Printf("\"%s\"\n", s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", u.ulChar)
	}
	fmt.Println()
}

func (u *UnderlinePen) CreateClone() Product {
	return &UnderlinePen{
		ulChar: u.ulChar,
	}
}
