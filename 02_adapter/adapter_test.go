package adapter

// ExampleAdapter 请求者(client)
func ExampleAdapter() {
	var p Print
	p = NewPrintBanner("Hello")
	p.PrintWeak()
	p.PrintStrong()

	// Output:
	// (Hello)
	// *Hello*
}
