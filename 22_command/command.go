package command

import "fmt"

// Command 命令(command)
type Command interface {
	Execute()
}

// MacroCommand 具体的命令(concrete command)
type MacroCommand struct {
	commands *Stack
}

func NewMacroCommand() *MacroCommand {
	return &MacroCommand{
		commands: NewStack(),
	}
}

func (m *MacroCommand) Execute() {
	for _, command := range *m.commands {
		command.Execute()
	}
}

func (m *MacroCommand) Append(cmd Command) {
	if cmd != m {
		m.commands.Push(cmd)
	}
}

func (m *MacroCommand) Undo() {
	if !m.commands.IsEmpty() {
		m.commands.Pop()
	}
}

func (m *MacroCommand) Clear() {
	m.commands.Clear()
}

// DrawCommand 具体的命令(concrete command)
type DrawCommand struct {
	drawable Drawable
	x        int
	y        int
}

func NewDrawCommand(drawable Drawable, x, y int) *DrawCommand {
	return &DrawCommand{
		drawable: drawable,
		x:        x,
		y:        y,
	}
}

func (d *DrawCommand) Execute() {
	d.drawable.Draw(d.x, d.y)
}

type Drawable interface {
	Draw(x, y int)
}

// DrawCanvas 接收者(receiver)、发动者(invoker)
type DrawCanvas struct {
	history *MacroCommand
}

func NewDrawCanvas() *DrawCanvas {
	return &DrawCanvas{
		history: NewMacroCommand(),
	}
}

func (d *DrawCanvas) Paint() {
	d.history.Execute()
}

func (d *DrawCanvas) Draw(x, y int) {
	fmt.Printf("draw x:%d, y:%d\n", x, y)
}
