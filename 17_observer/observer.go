package observer

import (
	"fmt"
	"math/rand"
)

type INumberGenerator interface {
	GetObserver() []Observer
	SetObserver([]Observer)
	GetNumber() int
	Execute()
}

// NumberGenerator 观察对象(subject)
type NumberGenerator struct {
	INumberGenerator

	observers []Observer
}

func (n *NumberGenerator) AddObserver(observer Observer) {
	n.SetObserver(append(n.GetObserver(), observer))
}

func (n *NumberGenerator) DeleteObserver(observer Observer) {
	var observers []Observer
	for _, o := range n.observers {
		if o != observer {
			observers = append(observers, o)
		}
	}
	n.observers = observers
}

func (n *NumberGenerator) NotifyObservers() {
	for _, o := range n.observers {
		o.Update(n)
	}
}

// RandomNumberGenerator 具体的观察对象(concrete subject)
type RandomNumberGenerator struct {
	NumberGenerator

	number int
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	rng := &RandomNumberGenerator{
		NumberGenerator: NumberGenerator{
			observers: make([]Observer, 0),
		},
	}
	rng.NumberGenerator.INumberGenerator = rng
	return rng
}

func (n *RandomNumberGenerator) GetObserver() []Observer {
	return n.observers
}

func (n *RandomNumberGenerator) SetObserver(observers []Observer) {
	n.observers = observers
}

func (n *RandomNumberGenerator) GetNumber() int {
	return n.number
}

func (n *RandomNumberGenerator) Execute() {
	for i := 0; i < 20; i++ {
		n.number = rand.Intn(10)
		n.NotifyObservers()
	}
}

// Observer 观察者(observer)
type Observer interface {
	Update(generator INumberGenerator)
}

// DigitObserver 具体的观察者(concrete observer)
type DigitObserver struct {
}

func NewDigitObserver() Observer {
	return &DigitObserver{}
}

func (d *DigitObserver) Update(generator INumberGenerator) {
	fmt.Printf("DigitObserver:%d\n", generator.GetNumber())
}

// GraphObserver 具体的观察者(concrete observer)
type GraphObserver struct {
}

func NewGraphObserver() Observer {
	return &GraphObserver{}
}

func (g *GraphObserver) Update(generator INumberGenerator) {
	fmt.Print("GraphObserver:")
	count := generator.GetNumber()
	for i := 0; i < count; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
