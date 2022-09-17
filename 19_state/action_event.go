package state

type ActionEvent string

const (
	ButtonUse   ActionEvent = "buttonUse"
	ButtonAlarm ActionEvent = "buttonAlarm"
	ButtonPhone ActionEvent = "buttonPhone"
	ButtonExit  ActionEvent = "buttonExit"
)
