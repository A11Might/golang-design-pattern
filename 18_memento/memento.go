package memento

import (
	"fmt"
	"math/rand"
	"strings"
)

// Gamer 原发器(originator)
type Gamer struct {
	money      int
	fruits     []string
	fruitsName []string
}

func NewGamer(money int) *Gamer {
	return &Gamer{
		money: money,
		fruitsName: []string{
			"苹果", "葡萄", "香蕉", "橘子",
		},
	}
}

func (g *Gamer) GetMoney() int {
	return g.money
}

func (g *Gamer) Bet() {
	dice := rand.Intn(6) + 1
	switch dice {
	case 1:
		g.money += 100
		fmt.Println("所持金币增加了。")
	case 2:
		g.money /= 2
		fmt.Println("所持金币减半了")
	case 6:
		f := g.getFruit()
		g.fruits = append(g.fruits, f)
		fmt.Printf("获得了水果（%s）。\n", f)
	default:
		fmt.Println("什么都没有发生。")
	}
}

func (g *Gamer) CreateMemento() *Memento {
	m := newMemento(g.money)
	for _, f := range g.fruits {
		if strings.HasPrefix(f, "好吃的") {
			m.addFruit(f)
		}
	}
	return m
}

func (g *Gamer) RestoreMemento(memento *Memento) {
	g.money = memento.GetMoney()
	g.fruits = memento.getFruit()
}

func (g *Gamer) ToString() string {
	return fmt.Sprintf("[money] = %d, fruits = %v]", g.money, g.fruits)
}

func (g *Gamer) getFruit() string {
	prefix := ""
	if rand.Intn(2) == 0 {
		prefix = "好吃的"
	}
	return prefix + g.fruitsName[rand.Intn(len(g.fruitsName))]
}

// Memento 备忘录(memento)
type Memento struct {
	money  int
	fruits []string
}

func newMemento(money int) *Memento {
	return &Memento{
		money: money,
	}
}

func (m *Memento) GetMoney() int {
	return m.money
}

func (m *Memento) addFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

func (m *Memento) getFruit() []string {
	cp := make([]string, len(m.fruits))
	_ = copy(cp, m.fruits)
	return cp
}
