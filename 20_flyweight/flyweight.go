package flyweight

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// BigChar 享元(flyweight)
type BigChar struct {
	charName rune
	fontData string
}

func NewBigChar(charName rune) *BigChar {
	file, err := os.Open(fmt.Sprintf("%s%c%s", "big", charName, ".txt"))
	if err != nil {
		return &BigChar{
			charName: '?',
		}
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return &BigChar{
		charName: charName,
		fontData: string(content),
	}
}

func (b *BigChar) Print() {
	fmt.Print(b.fontData)
}

// 单例模式
var bigCharFactory = &BigCharFactory{
	pool: make(map[rune]*BigChar),
}

func GetInstance() *BigCharFactory {
	return bigCharFactory
}

// BigCharFactory 享元工厂(flyweight factory)
type BigCharFactory struct {
	pool map[rune]*BigChar
	sync.Mutex
}

func (b *BigCharFactory) GetBigChar(charName rune) *BigChar {
	b.Lock()
	defer b.Unlock()
	bc, ok := b.pool[charName]
	if !ok {
		bc = NewBigChar(charName)
		b.pool[charName] = bc
	}
	return bc
}

// BigString 请求者(client)
type BigString struct {
	bigChars []*BigChar
}

func NewBigString(str string) *BigString {
	bigChars := make([]*BigChar, len(str))
	factory := GetInstance()
	for i := 0; i < len(str); i++ {
		bigChars[i] = factory.GetBigChar(rune(str[i]))
	}
	return &BigString{
		bigChars: bigChars,
	}
}

func (b *BigString) Print() {
	for i := range b.bigChars {
		b.bigChars[i].Print()
	}
}
