package main

import "fmt"

type Component interface {
	Describe() string
	GetCount() int
}

type AppleDecorator struct {
	Component
	Type string
	Num  int
}

func (apple *AppleDecorator) Describe() string {
	return fmt.Sprintf("%s, %s", apple.Component.Describe(), apple.Type)
}

func (apple *AppleDecorator) GetCount() int {
	return apple.Component.GetCount() + apple.Num
}

func NewAppleDecorator(c Component, t string, n int) Component {
	return &AppleDecorator{c, t, n}
}

type Fruit struct {
	Count       int
	Description string
}

func (f *Fruit) Describe() string {
	return f.Description
}

func (f *Fruit) GetCount() int {
	return f.Count
}

func main() {
	apple := NewAppleDecorator(&Fruit{
		Count:       1,
		Description: "foo",
	}, "bar", 2)

	fmt.Println(apple.GetCount())
	fmt.Println(apple.Describe())
}
