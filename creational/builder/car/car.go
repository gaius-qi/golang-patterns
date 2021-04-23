package car

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportWheels Wheels = "sports"
	SteelWheels        = "steel"
)

type Builder interface {
	Color(c Color) Builder
	Wheels(w Wheels) Builder
	TopSpeed(s Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive()
	Stop()
}

type car struct {
	Interface
	color  Color
	wheels Wheels
	speed  Speed
}

func (c *car) Drive() {
	fmt.Printf("A %s car with %s tires is going at %f ", c.color, c.wheels, c.speed)
	if c.speed == MPH {
		fmt.Println("MPH")
	} else {
		fmt.Println("KPH")
	}
}

func (c *car) Stop() {
	fmt.Println("The car stopped.")
}

type carBuilder struct {
	Builder
	c *car
}

func NewCarBuilder() Builder {
	cb := new(carBuilder)
	cb.c = new(car)
	return cb
}

func (cb *carBuilder) Color(c Color) Builder {
	cb.c.color = c
	return cb
}

func (cb *carBuilder) Wheels(w Wheels) Builder {
	cb.c.wheels = w
	return cb
}

func (cb *carBuilder) TopSpeed(s Speed) Builder {
	cb.c.speed = s
	return cb
}

func (cb *carBuilder) Build() Interface {
	return cb.c
}
