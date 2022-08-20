package singleton

import "fmt"

func ExampleSingleton() {
	fmt.Println("Start.")
	obj1 := GetInstance()
	obj2 := GetInstance()
	if obj1 == obj2 {
		fmt.Println("obj1 和 ojb2 是相同的实例。")
	} else {
		fmt.Println("obj1 和 ojb2 是不同的实例。")
	}
	fmt.Println("End.")

	// Output:
	// Start.
	// 生成了一个实例
	// obj1 和 ojb2 是相同的实例。
	// End.
}
