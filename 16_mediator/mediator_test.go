package mediator

import "fmt"

func ExampleMediator() {
	lf := NewLoginFrame("Mediator Sample")
	lf.Show()

	fmt.Println("选择作为用户登录")
	lf.checkGuest.ItemStateChanged(false)
	lf.checkLogin.ItemStateChanged(true)
	lf.Show()

	fmt.Println("在用户框输入字符")
	lf.textUser.TextValueChanged("Huqihang")
	lf.Show()

	fmt.Println("又输入了密码") 
	lf.textPass.TextValueChanged("*")
	lf.Show()

	fmt.Println("选择作为游客访问")
	lf.checkGuest.ItemStateChanged(true)
	lf.checkLogin.ItemStateChanged(false)
	lf.Show()

	fmt.Println("输入了密码，但删除了用户名")
	lf.checkGuest.ItemStateChanged(false)
	lf.checkLogin.ItemStateChanged(true)
	lf.textUser.TextValueChanged("")
	lf.Show()

	// Output:
	// Mediator Sample
	// Guest: [x], Login: []
	// Username: [] disabled
	// Password: [] disabled
	// OK: enabled, Cancel: disabled
	//
	// 选择作为用户登录
	// Mediator Sample
	// Guest: [], Login: [x]
	// Username: [] enabled
	// Password: [] disabled
	// OK: disabled, Cancel: disabled
	//
	// 在用户框输入字符
	// Mediator Sample
	// Guest: [], Login: [x]
	// Username: [Huqihang] enabled
	// Password: [] enabled
	// OK: disabled, Cancel: disabled
	//
	// 又输入了密码
	// Mediator Sample
	// Guest: [], Login: [x]
	// Username: [Huqihang] enabled
	// Password: [*] enabled
	// OK: enabled, Cancel: disabled
	//
	// 选择作为游客访问
	// Mediator Sample
	// Guest: [x], Login: []
	// Username: [Huqihang] disabled
	// Password: [*] disabled
	// OK: enabled, Cancel: disabled
	//
	// 输入了密码，但删除了用户名
	// Mediator Sample
	// Guest: [], Login: [x]
	// Username: [] enabled
	// Password: [*] disabled
	// OK: disabled, Cancel: disabled
}
