package proxy

import "fmt"

// ExampleProxy 请求者(client)
func ExampleProxy() {
	p := NewPrinterProxy("Alice")
	fmt.Printf("现在的名字是%s。\n", p.GetPrinterName())
	p.SetPrinterName("Bob")
	fmt.Printf("现在的名字是%s。\n", p.GetPrinterName())
	p.Print("Hello, world.")

	// Output:
	// 现在的名字是Alice。
	// 现在的名字是Bob。
	// 正在生成 Printer 的实例(Bob).....结束。
	// === Bob ===
	// Hello, world.
}
