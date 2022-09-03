package strategy

import (
	"fmt"
)

func ExampleStrategy() {
	player1 := NewPlay("Taro", NewWinningStrategy(314))
	player2 := NewPlay("Hana", NewProbStrategy(15))
	for i := 0; i < 10; i++ {
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()
		if nextHand1.IsStrongerThan(nextHand2) {
			fmt.Printf("Winner:%s\n", player1.ToString())
			player1.Win()
			player2.Lose()
		} else if nextHand2.IsStrongerThan(nextHand1) {
			fmt.Printf("Winner:%s\n", player2.ToString())
			player2.Win()
			player1.Lose()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}
	}
	fmt.Println("Total result:")
	fmt.Println(player1.ToString())
	fmt.Println(player2.ToString())

	// Output:
	// Winner:[Taro]:0 games, 0 win, 0 lose]
	// Winner:[Taro]:1 games, 1 win, 0 lose]
	// Even...
	// Even...
	// Winner:[Taro]:4 games, 2 win, 0 lose]
	// Even...
	// Even...
	// Winner:[Hana]:7 games, 0 win, 3 lose]
	// Winner:[Taro]:8 games, 3 win, 1 lose]
	// Even...
	// Total result:
	// [Taro]:10 games, 4 win, 1 lose]
	// [Hana]:10 games, 1 win, 4 lose]
}
