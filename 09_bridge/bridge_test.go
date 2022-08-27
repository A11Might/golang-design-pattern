package bridge

func ExampleBridge() {
	var (
		d1 IDisplay
		d2 IDisplay
		d3 *CountDisplay
	)
	d1 = NewDisplay(NewStringDisplayImpl("Hello, China."))
	d2 = NewCountDisplay(NewStringDisplayImpl("Hello, World."))
	d3 = NewCountDisplay(NewStringDisplayImpl("Hello, Universe.")).(*CountDisplay)
	d1.DisplaySth()
	d2.DisplaySth()
	d3.DisplaySth()
	d3.MultiDisplaySth(5)

	// Output:
	// +-------------+
	// |Hello, China.|
	// +-------------+
	// +-------------+
	// |Hello, World.|
	// +-------------+
	// +----------------+
	// |Hello, Universe.|
	// +----------------+
	// +----------------+
	// |Hello, Universe.|
	// |Hello, Universe.|
	// |Hello, Universe.|
	// |Hello, Universe.|
	// |Hello, Universe.|
	// +----------------+
}
