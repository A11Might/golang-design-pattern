package templatemethod

func ExampleTemplateMethod() {
	d1 := NewCharDisplay('H')
	d2 := NewStringDisplay("hello, world.")
	d3 := NewStringDisplay("你好，世界。")
	d1.Show()
	d2.Show()
	d3.Show()

	// Output:
	// <<HHHHH>>
	// +-------------+
	// |hello, world.|
	// |hello, world.|
	// |hello, world.|
	// |hello, world.|
	// |hello, world.|
	// +-------------+
	// +------------------+
	// |你好，世界。|
	// |你好，世界。|
	// |你好，世界。|
	// |你好，世界。|
	// |你好，世界。|
	// +------------------+
}
