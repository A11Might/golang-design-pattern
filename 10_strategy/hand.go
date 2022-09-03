package strategy

const (
	HandValueGUU = 0 // 表示石头的值
	HandValueCHO = 1 // 表示剪刀的值
	HandValuePAA = 2 // 表示布的值
)

var (
	// 表示猜拳中三种手势的实例
	hands = []*Hand{
		NewHand(HandValueGUU),
		NewHand(HandValueCHO),
		NewHand(HandValuePAA),
	}
	// 表示猜拳中手势对应的字符串
	names = []string{
		"石头",
		"剪刀",
		"布",
	}
)

type Hand struct {
	handValue int // 猜拳中手势的值
}

func NewHand(handValue int) *Hand {
	return &Hand{
		handValue: handValue,
	}
}

// GetHand 根据手势的值获取其对应的实例
func GetHand(handValue int) *Hand {
	return hands[handValue]
}

// IsStrongerThan 如果 h 胜了 hand 则返回 true
func (h *Hand) IsStrongerThan(hand *Hand) bool {
	return h.fight(hand) == 1
}

// IsWeakerThan 如果 h 输给了 hand 则返回 true
func (h *Hand) IsWeakerThan(hand *Hand) bool {
	return h.fight(hand) == -1
}

// 计分：平 0，胜 1，负 -1
func (h *Hand) fight(hand *Hand) int {
	if h == hand {
		return 0
	} else if (h.handValue+1)%3 == hand.handValue {
		return 1
	} else {
		return -1
	}
}

// ToString 转换为手势值所对应的字符串
func (h *Hand) ToString() string {
	return names[h.handValue]
}
