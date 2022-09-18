package main

import "fmt"

type IFoo interface {
	GetFooField() string
	SetFooField(string)
	AbstractMethod()
}

// 抽象类(abstract class)
type Foo struct {
	IFoo

	fooField string
}

func (f *Foo) Method() {
	f.SetFooField("foo")
	fmt.Println("foo field:", f.GetFooField())
	f.AbstractMethod()
	fmt.Println("call Foo Method")
}

// 具体子类(concrete subclass)
type XxFoo struct {
	*Foo

	xxFooField string
}

func NewXxFoo(fooField, xxFooField string) *XxFoo {
	xxFoo := &XxFoo{
		Foo: &Foo{
			fooField: fooField,
		},
		xxFooField: xxFooField,
	}
	xxFoo.Foo.IFoo = xxFoo
	return xxFoo
}

func (x *XxFoo) GetFooField() string {
	return x.fooField
}

func (x *XxFoo) SetFooField(field string) {
	x.fooField = field
}

func (x *XxFoo) AbstractMethod() {
	fmt.Println("call xxFoo AbstractMethod")
}

func main() {
	xxFoo := NewXxFoo("origin foo", "origin xx foo")

	fmt.Println("=====获取字段值=====")
	fmt.Println("field:", xxFoo.GetFooField())
	fmt.Println("=====设置字段值=====")
	xxFoo.SetFooField("foooooo")
	fmt.Println("field:", xxFoo.GetFooField())
	fmt.Println("=====使用子类方法=====")
	xxFoo.AbstractMethod()
	fmt.Println("=====使用父类方法=====")
	xxFoo.Method()

	// Output:
	// =====获取字段值=====
	// field: origin foo
	// =====设置字段值=====
	// field: foooooo
	// =====使用子类方法=====
	// call xxFoo AbstractMethod
	// =====使用父类方法=====
	// foo field: foo
	// call xxFoo AbstractMethod
	// call Foo Method
}
