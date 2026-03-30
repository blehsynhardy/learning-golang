package main

import "fmt"

// panic is a built-in function that stops the normal execution of a program and begins panicking. When a function panics, it immediately stops executing and starts unwinding the stack, looking for a deferred function that can recover from the panic. If no such function is found, the program will crash and print a stack trace.

// recover is a built-in function that allows a deferred function to catch a panic and prevent the program from crashing. When a function panics, it looks for a deferred function that can recover from the panic. If such a function is found, it will be called with the value passed to panic as an argument. The recover function can then handle the panic and allow the program to continue executing.


func mightPanic(mightPanic bool) {
	if mightPanic {
		panic("Something went wrong")
	}

	fmt.Print("this function executed without panic")
}

func recoverableFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	
	mightPanic(true)
}


func main() {
	recoverableFunction()

}