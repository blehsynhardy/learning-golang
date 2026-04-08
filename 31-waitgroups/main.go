package main

//introducing waitgroups to wait for go routines to finish before exiting the main function, otherwise the program will exit before the go routine has a chance to run. We can use a WaitGroup to wait for the go routine to finish.

import (
	"fmt"
	"sync"
	"time"
)

func SayHello(message string, delay time.Duration, wg *sync.WaitGroup)  {

	time.Sleep(delay)

	fmt.Println("say Hello", message)
	defer wg.Done()
}	


func main() {


	/*

		rule of wait group

		1. we need to add the number of go routines we want to wait for using wg.Add(n)
		2. we need to call wg.Done() at the end of each go routine to signal that the go routine has finished
		3. we need to call wg.Wait() at the end of the main function to wait for all go routines to finish before exiting the program
		4. Add outside of your go routine, and Done inside your go routine, and Wait at the end of your main function after all your go routines have been called.
		5. make sure u send  a pointer reference to the wait group in your go routine, otherwise it will not work and you will get a panic error. instead of a copy of the wait group, you will be sending a pointer reference to the wait group, so that all go routines can access the same wait group and signal when they are done.



		*/

	wg := sync.WaitGroup{}
	wg.Add(4)
	fmt.Println("===========GO ROUTINES START HERE ===========")

	go SayHello("world", 4*time.Second, &wg)
	go SayHello("jesus", 2*time.Second, &wg)
	go SayHello("resurrectiom", 1*time.Second, &wg)
	go SayHello("power", 3*time.Second, &wg)

	fmt.Println("===========GO ROUTINES END HERE ===========")

	//wait for the go routine to finish
	wg.Wait()




}