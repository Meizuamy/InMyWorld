[TOC]

# Effective GO 01

## Format

golang是一门非常简单的语言，golang设计的目的就是简化代码的复杂型，在代码的格式方面，golang会为你自动的格式化代码，不需要你来手动的去给你的代码代码格式化，例如：
```golang
type T struct {
    name string // name of the object
    value int // its value
}
```
通常我们定义一个结构体会是类似上面的代码，当你使用 `ctrl + s` 保存上面的代码，golang会帮你自动的格式化上面的代码，效果类似于下面的代码
```golang
type T struct {
	name  string // name of the object
	value string // its value
}

```
golang 默认使用制表符控制缩进，并且任何一行代码没有字数限制，如果觉得单行代码太长可以在下一行使用制表符进行缩进
```golang
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
```

golang中使用if，for, switch时，条件语句中都可以不带括号, 类似的，>> << 这些优化可以让语法更加高效
```golang
package main

import "fmt"

func main() {
	value := 1
	// if
	if value == 1 {
		fmt.Println("value equals 1!")
	}

	// switch
	for i := 0; i < value; i++ {
		fmt.Printf("value is %d\n", i)
	}

	// switch case
	switch value {
	case 1:
		fmt.Println("value equals 1!")
	case 2:
		fmt.Println("value equals 2!")
	}
    
    x := 1
	y := 1
	// >> <<
	fmt.Printf("%d", x<<8+y<<16)
}
```

