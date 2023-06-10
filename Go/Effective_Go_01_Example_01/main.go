package main

import "fmt"

type T struct {
	name  string // name of the object
	value string // its value
}

func (t *T) SetName(name string) *T {
	t.name = name
	return t
}

func (t *T) SetValue(value string) *T {
	t.value = value
	return t
}

func (t *T) ToString() string {
	return fmt.Sprintf("%s,%s", t.name, t.value)
}

func main() {

	t := T{}

	// 不可以在调用其他函数的点之前换行
	t.SetName("Hello,")
		.SetValue("World!")
	// 点之后可以
	t.SetName("Hello,").
		SetValue("World!")
	// 不可以将括号分割
	t.SetName( 
		"Hello,").SetValue(" 
		World!") 
	fmt.Println(t)
}
