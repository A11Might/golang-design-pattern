package strategy

import (
	"fmt"
	"math/rand"
)

// Play 上下文(context)
type Play struct {
	name      string
	strategy  Strategy
	winCount  int
	loseCount int
	gameCount int
}

func NewPlay(name string, strategy Strategy) *Play {
	return &Play{
		name:     name,
		strategy: strategy,
	}
}

func (p *Play) NextHand() *Hand {
	return p.strategy.NextHand()
}

func (p *Play) Win() {
	p.strategy.Study(true)
	p.winCount++
	p.gameCount++
}

func (p *Play) Lose() {
	p.strategy.Study(false)
	p.loseCount++
	p.gameCount++
}

func (p *Play) Even() {
	p.gameCount++
}

func (p *Play) ToString() string {
	return fmt.Sprintf("[%s]:%d games, %d win, %d lose]", p.name, p.gameCount, p.winCount, p.loseCount)
}

// Strategy 策略
type Strategy interface {
	NextHand() *Hand
	Study(win bool)
}

// WinningStrategy 具体策略(concrete strategy)
type WinningStrategy struct {
	rand     *rand.Rand
	won      bool
	prevHand *Hand
}

func NewWinningStrategy(seed int64) *WinningStrategy {
	return &WinningStrategy{
		rand: rand.New(rand.NewSource(seed)),
	}
}

func (w *WinningStrategy) NextHand() *Hand {
	if !w.won {
		w.prevHand = GetHand(w.rand.Intn(3))
	}
	return w.prevHand
}

func (w *WinningStrategy) Study(win bool) {
	w.won = win
}

// ProbStrategy 具体策略(concrete strategy)
type ProbStrategy struct {
	rand             *rand.Rand
	prevHandValue    int
	currentHandValue int
	history          [3][3]int
}

func NewProbStrategy(seed int64) *ProbStrategy {
	return &ProbStrategy{
		rand: rand.New(rand.NewSource(seed)),
		history: [3][3]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}
}

func (p *ProbStrategy) NextHand() *Hand {
	bet := p.rand.Intn(p.getSum(p.currentHandValue))
	handValue := 0
	if bet < p.history[p.currentHandValue][0] {
		handValue = 0
	} else if bet < p.history[p.currentHandValue][0]+p.history[p.currentHandValue][1] {
		handValue = 1
	} else {
		handValue = 2
	}
	p.prevHandValue = p.currentHandValue
	p.currentHandValue = handValue
	return GetHand(handValue)
}

func (p *ProbStrategy) getSum(hv int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += p.history[hv][i]
	}
	return sum
}

func (p *ProbStrategy) Study(win bool) {
	if win {
		p.history[p.prevHandValue][p.currentHandValue]++
	} else {
		p.history[p.prevHandValue][(p.currentHandValue+1)%3]++
		p.history[p.prevHandValue][(p.currentHandValue+2)%3]++
	}
}
