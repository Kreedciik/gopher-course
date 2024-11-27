package coffeemachine

import "fmt"

const (
	maxWaterLevel float64 = 5.0
	maxMilkLevel  float64 = 5.0
)

var ErrOutOfWater = fmt.Errorf("maximum water level is reached")
var ErrNoWater = fmt.Errorf("the water level is low")
var ErrOutOfMilk = fmt.Errorf("maximum milk level is reached")
var ErrNoMilk = fmt.Errorf("the milk level is low")

type CoffeeMachine struct {
	amountOfWater float64
	amountOfMilk  float64
}

func (m *CoffeeMachine) AddWater(w float64) (float64, error) {
	if w > maxWaterLevel || m.amountOfWater-w < 0 {
		return -1, ErrOutOfWater
	}
	m.amountOfWater += w
	return m.amountOfWater, nil
}

func (c *CoffeeMachine) AddMilk(m float64) (float64, error) {
	if m > maxMilkLevel || c.amountOfMilk-m < 0 {
		return -1, ErrOutOfMilk
	}
	c.amountOfMilk += m
	return c.amountOfMilk, nil
}

func (m *CoffeeMachine) makeCoffee() error {
	if m.amountOfWater < 2 {
		return ErrNoWater
	}

	if m.amountOfMilk < 2 {
		return ErrNoMilk
	}

	m.amountOfWater -= 2
	m.amountOfMilk -= 2
	return nil
}

func InitCoffeeMachine() CoffeeMachine {
	return CoffeeMachine{0, 0}
}
