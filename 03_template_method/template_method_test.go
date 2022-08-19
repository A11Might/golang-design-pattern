package templatemethod

import "testing"

func ExampleTemplateMethod(t *testing.T) {
	d1 := Display{
		DisplayIface: NewCharDisplay('H'),
	}
	d2 := Display{
		DisplayIface: NewStringDisplay("hello, world."),
	}
	d3 := Display{
		DisplayIface: NewStringDisplay("你好，世界。"),
	}
	d1.Display()
	d2.Display()
	d3.Display()

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
