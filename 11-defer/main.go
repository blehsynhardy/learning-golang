package main

import "fmt"

func deferExample() {
	println("This will be printed first")
	defer println("this is a defereed message")
	println("This will be printed last")
}

// lifo defer example shows that the deferred functions are executed in the reverse order of their declaration. In this example, the first deferred message will be printed last, and the third deferred message will be printed first. This is because the deferred functions are executed in a last-in-first-out (LIFO) order when the surrounding function returns.
func lifoDeferExample() {
	defer println("This is the first deferred message")
	defer println("This is the second deferred message")
	defer println("This is the third deferred message")
	println("This will be printed first")
	println("This will be printed second")
	println("This will be printed third")
}

func main() {

	defer func() {
		fmt.Println("This will be printed last recoverd from panic")
	}()

	// defer is used to delay the execution of a function until the surrounding function returns. This is useful for cleaning up resources, such as closing files or database connections, after they are no longer needed.

	// defer is often used in conjunction with the panic and recover functions to handle errors and ensure that resources are properly cleaned up even in the case of an error.

	// defer can also be used to ensure that a function is always called, regardless of how the surrounding function exits. This can be useful for logging or other cleanup tasks that need to be performed regardless of the outcome of the surrounding function.

	deferExample()
	lifoDeferExample()

}