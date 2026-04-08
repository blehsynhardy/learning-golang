package main

import (
	"fmt"
	"time"
)

func SayHello(message string, delay time.Duration)  {

	time.Sleep(delay)

	fmt.Println("say Hello", message)

}

func main() {

	fmt.Println("===========GO ROUTINES START HERE ===========")

	go SayHello("world", 2*time.Second)

	fmt.Println("===========GO ROUTINES END HERE ===========")

	//wait for the go routine to finish
	time.Sleep(3*time.Second)



	//we need a fix to wait for the go routine to finish before exiting the main function, otherwise the program will exit before the go routine has a chance to run. We can use a WaitGroup to wait for the go routine to finish.

	//check the next example for the fix using WaitGroup or section

}