package main

import (
	"fmt"
	"log"
	"os"
)

type pizza struct {
	price int
}

func (p *pizza) getPrice() int {
	return p.price
}

func loggerDecoratorGetPrice(f func() int, l *log.Logger) func() int {
	return func() int {
		l.Println("calling getPrice()")
		return f()
	}
}

func main() {
	myLogger := log.New(os.Stdout, "", 0)
	pizza := pizza{price: 5}

	decorateGet := loggerDecoratorGetPrice(pizza.getPrice, myLogger)()
	fmt.Println(decorateGet)
}

//type somedata struct {
//	s string
//}
//
//func (d *somedata) printSomething(prefix string) {
//	fmt.Println(prefix, d.s)
//}
//
//// DoPrintSomething is a exported function
//type DoPrintSomething func(string)
//
//func decorate(f DoPrintSomething) DoPrintSomething {
//	return func(s string) { // the return must match the full DoPrintSomething function signature
//		f(s)
//	}
//}
//
//func main() {
//
//	d := somedata{s: "i'm not decorated\n"}
//	d.printSomething(">>")
//
//	e := somedata{s: "i now am decorated... woohoo!"}
//	decorate(e.printSomething)("==>") // <-- notice the ending () .. this is where function args would be going (we'll be doing that later)
//
//}
