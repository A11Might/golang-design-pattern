package builder

import "fmt"

// ExampleBuilder 使用者(client)
func ExampleBuilder() {
	textBuilder := NewTextBuilder()
	director := NewDirector(textBuilder)
	director.Construct()
	result := textBuilder.GetResult()
	fmt.Println(result)

	htmlBuilder := NewHTMLBuilder()
	director = NewDirector(htmlBuilder)
	director.Construct()
	filename := htmlBuilder.GetResult()
	fmt.Printf("%s文件编写完成。\n", filename)

	// Output:
	// ====================
	// 『Greeting』
	//
	// ■从早上至下午
	//
	//     .早上好。
	//     .下午好。
	//
	// ■晚上
	//
	//     .晚上好。
	//     .晚安。
	//     .再见。
	//
	// ====================
	//
	// Greeting.html文件编写完成。
}
