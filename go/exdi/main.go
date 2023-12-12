package main

import (
	"log"
)

func fnWithoutDI() {
	log.Println("No DI")
}

type Logger interface {
	Println(s string)
}

type myLogger struct{}

func (ml myLogger) Println(s string) {
	log.Println(s)
}

func fnWithDI(logger Logger) {
	logger.Println("Fn with DI")
}

type DIStruct struct {
	logger Logger
}

var gblVar DIStruct = DIStruct{myLogger{}}

func fnBeforeDI() {
	gblVar.logger.Println("Function using the variable (that could have been DI)")
}

func (dis DIStruct) structMethodWithDI() {
	dis.logger.Println("Struct method with DI")
}

func main() {
	fnWithoutDI()               // No DI
	fnWithDI(myLogger{})        // Fn with DI
	fnBeforeDI()                // Function using the variable (that could have been DI)
	gblVar.structMethodWithDI() // Struct method with DI
}
