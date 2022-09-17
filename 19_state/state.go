package state

import "fmt"

// Context 上下文(context)
type Context interface {
	SetClock(hour int)
	ChangeState(state State)
	CallSecurityCenter(msg string)
	RecordLog(msg string)
}

// SafeFrame 具体的上下文(concrete context)
type SafeFrame struct {
	state State
}

func NewSafeFrame(state State) *SafeFrame {
	return &SafeFrame{
		state: state,
	}
}

func (s *SafeFrame) ActionPerformed(e ActionEvent) {
	switch e {
	case ButtonUse:
		s.state.DoUse(s)
	case ButtonAlarm:
		s.state.DoAlarm(s)
	case ButtonPhone:
		s.state.DoPhone(s)
	case ButtonExit:
		return
	default:
		fmt.Println("?")
	}
}

func (s *SafeFrame) SetClock(hour int) {
	if hour < 10 {
		fmt.Printf("现在时间是 0%d:00\n", hour)
	} else {
		fmt.Printf("现在时间是 %d:00\n", hour)
	}
	s.state.DoClock(s, hour)
}

func (s *SafeFrame) ChangeState(state State) {
	fmt.Printf("从%s状态变成了%s状态。\n", s.state.ToString(), state.ToString())
	s.state = state
}

func (s *SafeFrame) CallSecurityCenter(msg string) {
	fmt.Printf("call! %s\n", msg)
}

func (s *SafeFrame) RecordLog(msg string) {
	fmt.Printf("record ... %s\n", msg)
}

// State 状态(state)
type State interface {
	DoClock(Context, int)
	DoUse(Context)
	DoAlarm(Context)
	DoPhone(Context)
	ToString() string
}

// 单例模式
var dayState = &DayState{}

func GetDayState() *DayState {
	return dayState
}

// DayState 具体的状态(concrete state)
type DayState struct {
}

func (d *DayState) DoClock(context Context, hour int) {
	if hour < 9 || 17 <= hour {
		context.ChangeState(GetNightState()) // 状态迁移
	}
}

func (d *DayState) DoUse(context Context) {
	context.RecordLog("使用金库（白天）")
}

func (d *DayState) DoAlarm(context Context) {
	context.CallSecurityCenter("按下警铃（白天）")
}

func (d *DayState) DoPhone(context Context) {
	context.CallSecurityCenter("正常通话（白天）")
}

func (d *DayState) ToString() string {
	return "[白天]"
}

// 单例模式
var nightState = &NightState{}

func GetNightState() *NightState {
	return nightState
}

// NightState 具体的状态(concrete state)
type NightState struct {
}

func (d *NightState) DoClock(context Context, hour int) {
	if 9 <= hour || hour < 17 {
		context.ChangeState(GetDayState()) // 状态迁移
	}
}

func (d *NightState) DoUse(context Context) {
	context.RecordLog("紧急：晚上使用金库！")
}

func (d *NightState) DoAlarm(context Context) {
	context.CallSecurityCenter("按下警铃（晚上）")
}

func (d *NightState) DoPhone(context Context) {
	context.CallSecurityCenter("晚上通话录音")
}

func (d *NightState) ToString() string {
	return "[晚上]"
}
