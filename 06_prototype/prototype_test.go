package prototype

func ExamplePrototype() {
	// 准备
	manger := NewManager()
	uPen := NewUnderlinePen('~')
	mBox := NewMessageBox('*')
	sBox := NewMessageBox('/')
	manger.Register("strong message", uPen)
	manger.Register("warning box", mBox)
	manger.Register("slash box", sBox)

	// 生成
	p1 := manger.Create("strong message")
	p1.Use("Hello, world.")
	p2 := manger.Create("warning box")
	p2.Use("Hello, world.")
	p3 := manger.Create("slash box")
	p3.Use("Hello, world.")

	// Output:
	// "Hello, world."
	// ~~~~~~~~~~~~~
	// *****************
	// *Hello, world.*
	// *****************
	// /////////////////
	// /Hello, world./
	// /////////////////
}
