package state

import "fmt"

func ExampleState() {
	dayState := GetDayState()
	frame := NewSafeFrame(dayState)

	frame.SetClock(10)
	fmt.Println("操作：使用金库")
	frame.ActionPerformed(ButtonUse)
	fmt.Println("操作：按下警铃")
	frame.ActionPerformed(ButtonAlarm)
	fmt.Println("操作：正常通话")
	frame.ActionPerformed(ButtonPhone)
	fmt.Println()

	frame.SetClock(22)
	fmt.Println("操作：使用金库")
	frame.ActionPerformed(ButtonUse)
	fmt.Println("操作：按下警铃")
	frame.ActionPerformed(ButtonAlarm)
	fmt.Println("操作：正常通话")
	frame.ActionPerformed(ButtonPhone)

	// Output:
	// 现在时间是 10:00
	// 操作：使用金库
	// record ... 使用金库（白天）
	// 操作：按下警铃
	// call! 按下警铃（白天）
	// 操作：正常通话
	// call! 正常通话（白天）
	//
	// 现在时间是 22:00
	// 从[白天]状态变成了[晚上]状态。
	// 操作：使用金库
	// record ... 紧急：晚上使用金库！
	// 操作：按下警铃
	// call! 按下警铃（晚上）
	// 操作：正常通话
	// call! 晚上通话录音
}
