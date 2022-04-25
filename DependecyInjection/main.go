package main

import (
	"fmt"
)

// interface 1
type Logger interface {
	Log(message string)
}

// interface 2 

type Decorator interface {
	Decorate(message string) string
}


// use the type func(...)... with signature to transfer the specific func (input parameter) to meet the interface (return function)
//!!! IMPORTANT
type LoggerAdapter func(message string)

func (la LoggerAdapter) Log(message string) {
	la(message)
}

// define the Decorator directly to interface 

type SimpleDecorator struct {
	message string
}

func (sd SimpleDecorator) Decorate(message string) string{
	return fmt.Sprintf("%s %s!", sd.message, message)
}

func LogStandard(message string) {
	fmt.Println(message)
}
// defined the concrete instance

type SimpleInstance struct{
	l Logger
	d Decorator
}

func (s SimpleInstance) process(message string) {
	s.l.Log(s.d.Decorate(message))
}

func newSimpleInstance(l Logger, d Decorator) SimpleInstance{
	return SimpleInstance{
		l: l,
		d: d,
	}
}

// another Logger method 
func LogNonStandard(message string) {
	fmt.Printf("%s is not print in standard\n", message)
}

func printWithInjection(log func(message string)) {
		// create a adapter and "inject" a method
	la := LoggerAdapter(log)
	// create a static decorator
	sd := SimpleDecorator{
		message: "Hello",
	}
	// create a concreate Instance
	nsi := newSimpleInstance(la, sd)
	// run process method
	nsi.process("Jimmy")
}

func main() {
	// print with different Logging method
	printWithInjection(LogStandard)
	printWithInjection(LogNonStandard)
	
	printWithInjection(func(message string){
		fmt.Println("This is a anonymous printing method : " + message)
	})
}