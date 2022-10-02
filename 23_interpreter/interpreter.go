package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

// Node 抽象表达式(abstract expression)
type Node interface {
	Parse(*Context)
	ToString() string
}

// ProgramNode 非终结符表达式(non-terminal expression)
// <program> ::= program <command list>
type ProgramNode struct {
	commandListNode Node
}

func (p *ProgramNode) Parse(context *Context) {
	context.SkipToken("program")
	p.commandListNode = new(CommandListNode)
	p.commandListNode.Parse(context)
}

func (p *ProgramNode) ToString() string {
	return fmt.Sprintf("[program %s]", p.commandListNode.ToString())
}

// CommandListNode 非终结符表达式(non-terminal expression)
// <command list> ::= <command>* end
type CommandListNode struct {
	list []Node
}

func (c *CommandListNode) Parse(context *Context) {
	for {
		if len(context.CurrentToken()) == 0 {
			panic("Missing 'end'")
		} else if context.CurrentToken() == "end" {
			context.SkipToken("end")
			break
		} else {
			commandNode := new(CommandNode)
			commandNode.Parse(context)
			c.list = append(c.list, commandNode)
		}
	}
}

func (c *CommandListNode) ToString() string {
	var nodeStringList []string
	for _, node := range c.list {
		nodeStringList = append(nodeStringList, node.ToString())
	}

	return fmt.Sprintf("[%s]", strings.Join(nodeStringList, ", "))
}

// CommandNode 非终结符表达式(non-terminal expression)
// <command> ::= <repeat command> | <primitive command>
type CommandNode struct {
	node Node
}

func (c *CommandNode) Parse(context *Context) {
	if context.CurrentToken() == "repeat" {
		c.node = new(RepeatCommandNode)
		c.node.Parse(context)
	} else {
		c.node = new(PrimitiveCommandNode)
		c.node.Parse(context)
	}
}

func (c *CommandNode) ToString() string {
	return c.node.ToString()
}

// RepeatCommandNode 非终结符表达式(non-terminal expression)
// <repeat command> ::= repeat <number> <command list>
type RepeatCommandNode struct {
	number          int
	commandListNode Node
}

func (r *RepeatCommandNode) Parse(context *Context) {
	context.SkipToken("repeat")
	r.number = context.CurrentNumber()
	context.NextToken()
	r.commandListNode = new(CommandListNode)
	r.commandListNode.Parse(context)
}

func (r *RepeatCommandNode) ToString() string {
	return fmt.Sprintf("[repeat %d %s]", r.number, r.commandListNode.ToString())
}

// PrimitiveCommandNode 终结符表达式(terminal expression)
// <primitive command> ::= go | right | left
type PrimitiveCommandNode struct {
	name string
}

func (r *PrimitiveCommandNode) Parse(context *Context) {
	r.name = context.CurrentToken()
	context.SkipToken(r.name)
	if r.name != "go" &&
		r.name != "right" &&
		r.name != "left" {
		panic(fmt.Sprintf("%s is undefined", r.name))
	}
}

func (r *PrimitiveCommandNode) ToString() string {
	return r.name
}

// Context 上下文(context)
type Context struct {
	tokens       []string
	index        int
	currentToken string
}

func NewContext(text string) *Context {
	tokens := strings.Split(text, " ")
	context := &Context{
		tokens:       tokens,
		index:        0,
		currentToken: tokens[0],
	}
	return context
}

func (c *Context) NextToken() string {
	if c.index+1 < len(c.tokens) {
		c.index++
		c.currentToken = c.tokens[c.index]
	} else {
		c.currentToken = ""
	}
	return c.currentToken
}

func (c *Context) CurrentToken() string {
	return c.currentToken
}

func (c *Context) SkipToken(token string) {
	if token != c.currentToken {
		panic(fmt.Sprintf("Warning: %s is expected, but %s is found.", token, c.currentToken))
	}
	c.NextToken()
}

func (c *Context) CurrentNumber() int {
	number, err := strconv.Atoi(c.currentToken)
	if err != nil {
		panic(fmt.Sprintf("Warning: %v", err.Error()))
	}
	return number
}
