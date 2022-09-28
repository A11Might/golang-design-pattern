package command

import "fmt"

// ExampleCommand 请求者(client)、发动者(invoker)
func ExampleCommand() {
	drawCanvas := NewDrawCanvas()

	fmt.Println("鼠标拖拽绘制点")
	mouseDragged(drawCanvas, 1, 1)
	mouseDragged(drawCanvas, 1, 2)
	mouseDragged(drawCanvas, 2, 2)
	mouseDragged(drawCanvas, 1, 2)
	mouseDragged(drawCanvas, 1, 1)

	fmt.Printf("\n删除已绘制的点\n")

	fmt.Printf("\n绘制历史记录\n")
	drawCanvas.Paint()

	clear(drawCanvas)
	drawCanvas.Paint()

	// Output:
	// 鼠标拖拽绘制点
	// draw x:1, y:1
	// draw x:1, y:2
	// draw x:2, y:2
	// draw x:1, y:2
	// draw x:1, y:1
	//
	// 删除已绘制的点
	//
	// 绘制历史记录
	// draw x:1, y:1
	// draw x:1, y:2
	// draw x:2, y:2
	// draw x:1, y:2
	// draw x:1, y:1
}

func mouseDragged(drawable Drawable, x, y int) {
	cmd := NewDrawCommand(drawable, x, y)
	drawable.(*DrawCanvas).history.Append(cmd)
	cmd.Execute()
}

func clear(drawable Drawable) {
	drawable.(*DrawCanvas).history.Clear()
}
