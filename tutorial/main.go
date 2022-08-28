package main

import "fmt"

// 抽象类(abstract class)
type Foo struct {
	FooImpl

	fooField string
}

func (f *Foo) Method() {
	f.SetFooField("foo")
	fmt.Println("foo field:", f.GetFooField())
	f.AbstractMethod()
	fmt.Println("call Foo Method")
}

type FooImpl interface {
	GetFooField() string
	SetFooField(string)
	AbstractMethod()
}

// 具体子类(concrete subclass)
type xxFoo struct {
	*Foo

	xxFooField string
}

func (x *xxFoo) GetFooField() string {
	return x.fooField
}

func (x *xxFoo) SetFooField(field string) {
	x.fooField = field
}

func (x *xxFoo) AbstractMethod() {
	fmt.Println("call xxFoo AbstractMethod")
}

func main() {
	foo := Foo{
		FooImpl: &xxFoo{
			// 初始化字段值
			Foo: &Foo{
				fooField: "origin foo",
			},
			xxFooField: "origin xx foo",
		},
	}

	fmt.Println("=====获取字段值=====")
	fmt.Println("field:", foo.GetFooField())
	fmt.Println("=====设置字段值=====")
	foo.SetFooField("foooooo")
	fmt.Println("field:", foo.GetFooField())
	fmt.Println("=====使用子类方法=====")
	foo.AbstractMethod()
	fmt.Println("=====使用父类方法=====")
	foo.Method()

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
