package main

import "fmt"

type pizza interface {
	getPrice() int
}

type bostonPizza struct {
	price int
}

type chicagoPizza struct {
	price int
}

func newBostonPizza() *bostonPizza {
	return &bostonPizza{price: 10}
}

func newChicagoPizza() *chicagoPizza {
	return &chicagoPizza{price: 7}
}

func (c *bostonPizza) getPrice() int {
	return c.price
}

func (c *chicagoPizza) getPrice() int {
	return c.price
}

const (
	pepperoniPrice = 2
	tomatoPrice    = 1
)

type pepperoniTopping struct {
	pizza pizza
}
type tomatoTopping struct {
	pizza pizza
}

func (p *pepperoniTopping) getPrice() int {
	return p.pizza.getPrice() + pepperoniPrice
}

func (t *tomatoTopping) getPrice() int {
	return t.pizza.getPrice() + tomatoPrice
}

func main() {
	chicagoPepperoniPizza := &pepperoniTopping{
		pizza: newChicagoPizza(),
	}

	bostonDoublePepperoniTomatoPizza := &pepperoniTopping{
		pizza: &pepperoniTopping{
			pizza: &tomatoTopping{
				pizza: newBostonPizza(),
			},
		},
	}

	fmt.Printf("Price of chicago Pepperoni Pizza is %d\n",
		chicagoPepperoniPizza.getPrice())
	fmt.Printf("Price of boston DoublePepperoni Tomato Pizza is %d\n",
		bostonDoublePepperoniTomatoPizza.getPrice())
}

//func main() {
//	inheritanceChicagoPeperoni := &chicagoPeperoni{
//		pizza: newChicagoPizza(),
//	}
//	fmt.Printf("Price of chicagoPepperoniPizza is %d\n",
//		inheritanceChicagoPeperoni.getPrice())
//}
//type chicagoPeperoni struct {
//	pizza pizza
//}
//
//func (p *chicagoPeperoni) getPrice() int {
//	return p.pizza.getPrice() + pepperoniPrice
//}
