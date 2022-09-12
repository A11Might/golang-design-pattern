package mediator

import "fmt"

// Mediator 中介者(mediator)
type Mediator interface {
	CreateColleagues()
	ColleagueChanged()
}

// LoginFrame 具体的中介者(concrete mediator)
type LoginFrame struct {
	checkGuest   *ColleagueCheckbox
	checkLogin   *ColleagueCheckbox
	textUser     *ColleagueTextField
	textPass     *ColleagueTextField
	buttonOk     *ColleagueButton
	buttonCancel *ColleagueButton
	title        string
}

func NewLoginFrame(title string) *LoginFrame {
	lf := &LoginFrame{
		title: title,
	}
	// 生成各个 Colleague
	lf.CreateColleagues()
	// 设置初始状态
	lf.ColleagueChanged()
	return lf
}

func (l *LoginFrame) CreateColleagues() {
	l.checkGuest = NewColleagueCheckbox("Guest", true)
	l.checkLogin = NewColleagueCheckbox("Login", false)
	l.textUser = NewColleagueTextField("", 10)
	l.textPass = NewColleagueTextField("", 10)
	l.buttonOk = NewColleagueButton("OK")
	l.buttonCancel = NewColleagueButton("Cancel")

	l.checkGuest.SetMediator(l)
	l.checkLogin.SetMediator(l)
	l.textUser.SetMediator(l)
	l.textPass.SetMediator(l)
	l.buttonOk.SetMediator(l)
	l.buttonCancel.SetMediator(l)
}

func (l *LoginFrame) ColleagueChanged() {
	if l.checkGuest.state {
		l.textUser.SetColleagueEnable(false)
		l.textPass.SetColleagueEnable(false)
		l.buttonOk.SetColleagueEnable(true)
	} else {
		l.textUser.SetColleagueEnable(true)
		l.userPassChanged()
	}
}

func (l *LoginFrame) userPassChanged() {
	if len(l.textUser.text) > 0 {
		l.textPass.SetColleagueEnable(true)
		if len(l.textPass.text) > 0 {
			l.buttonOk.SetColleagueEnable(true)
		} else {
			l.buttonOk.SetColleagueEnable(false)
		}
	} else {
		l.textPass.SetColleagueEnable(false)
		l.buttonOk.SetColleagueEnable(false)
	}
}

func (l *LoginFrame) Show() {
	fmt.Println(l.title)
	if l.checkGuest.state {
		fmt.Print("Guest: [x], ")
	} else {
		fmt.Print("Guest: [], ")
	}
	if l.checkLogin.state {
		fmt.Print("Login: [x]\n")
	} else {
		fmt.Print("Login: []\n")
	}
	fmt.Printf("Username: [%s] ", l.textUser.text)
	if l.textUser.enabled {
		fmt.Println("enabled")
	} else {
		fmt.Println("disabled")
	}
	fmt.Printf("Password: [%s] ", l.textPass.text)
	if l.textPass.enabled {
		fmt.Println("enabled")
	} else {
		fmt.Println("disabled")
	}
	if l.buttonOk.enabled {
		fmt.Print("OK: enabled, ")
	} else {
		fmt.Print("OK: disabled, ")
	}
	if l.buttonCancel.enabled {
		fmt.Print("Cancel: enabled")
	} else {
		fmt.Print("Cancel: disabled")
	}
	fmt.Printf("\n\n")
}

// Colleague 同事(colleague)
type Colleague interface {
	SetMediator(mediator Mediator)
	SetColleagueEnable(enabled bool)
}

// ColleagueButton 具体的同事(concrete colleague)
type ColleagueButton struct {
	mediator Mediator
	caption  string
	enabled  bool
}

func NewColleagueButton(caption string) *ColleagueButton {
	return &ColleagueButton{
		caption: caption,
	}
}

func (c *ColleagueButton) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ColleagueButton) SetColleagueEnable(enabled bool) {
	c.enabled = enabled
}

// ColleagueTextField 具体的同事(concrete colleague)
type ColleagueTextField struct {
	mediator   Mediator
	text       string
	columns    int
	enabled    bool
	background string
}

func NewColleagueTextField(text string, columns int) *ColleagueTextField {
	return &ColleagueTextField{
		text:    text,
		columns: columns,
	}
}

func (c *ColleagueTextField) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ColleagueTextField) SetColleagueEnable(enabled bool) {
	c.enabled = enabled
	if enabled {
		c.background = "white"
	} else {
		c.background = "lightGray"
	}
}

func (c *ColleagueTextField) TextValueChanged(text string) {
	c.text = text
	c.mediator.ColleagueChanged()
}

// ColleagueCheckbox 具体的同事(concrete colleague)
type ColleagueCheckbox struct {
	mediator Mediator
	caption  string
	state    bool
	enabled  bool
}

func NewColleagueCheckbox(caption string, state bool) *ColleagueCheckbox {
	return &ColleagueCheckbox{
		caption: caption,
		state:   state,
	}
}

func (c *ColleagueCheckbox) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ColleagueCheckbox) SetColleagueEnable(enabled bool) {
	c.enabled = enabled
}

func (c *ColleagueCheckbox) ItemStateChanged(state bool) {
	c.state = state
	c.mediator.ColleagueChanged()
}
