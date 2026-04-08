package main

import (
	"errors"
	"fmt"
)

type ConfigItems struct {
	Key   string
	Value any
	Isset bool
}


/*
The String() method is defined on the ConfigItems struct. It uses fmt.Sprintf to format the string representation of the struct's fields. When you print an instance of ConfigItems, this String() method will be called automatically to provide a human-readable output of the struct's contents.

%v => default format for the value
%+v => adds field names to the output of the value
%#v => adds field names and type information to the output of the value
%T => type of the value
%q => double-quoted string format for the value
%s => string format for the key
%t => boolean format for the Isset field
%d => decimal format for integers 
%f => floating-point format for float fields e.g %f => 123.456000 \n %0.2f => 123.46, 2 decimal places %.2f => 123.46, 2 decimal places with leading zeros 
%p => pointer address format for the value 
%e => scientific notation format for float fields 
%x => hexadecimal format for integers 
%o => octal format for integers 
%% => literal percent sign, no value is consumed
*/

// String implements the fmt.Stringer interface for the ConfigItems struct.
// It returns a formatted string representation of the configuration item.
func (c ConfigItems) String() string {
	return fmt.Sprintf("Key: %s, Value: %v, Isset: %t", c.Key, c.Value, c.Isset)
}

// main demonstrates various string formatting techniques in Go using the fmt package,
// including Sprintf for string creation and Printf with different verbs.
func main() {


	appName := "MyApp"
	version := 1.2
	port := 8080
	isEnabled := true

	// fmt.Sprintf formats a string and returns it without printing to console
	status := fmt.Sprintf("Application: %s, Version: %.2f, Port: %d, Enabled: %t", appName, version, port, isEnabled)

	fmt.Println(status)


	item1 := ConfigItems{
		Key:   "DatabaseURL",
		Value: "postgres://user:password@localhost:5432/mydb",
		Isset: true,
	}

	item2 := ConfigItems{
		Key:   "ServerPort",
		Value: 8080,
		Isset: true,
	}

	item3 := ConfigItems{
		Key:   "DebugMode",
		Value: false,
		Isset: true,
	}

	fmt.Printf ("Item1: (%%v) %v\n", item1)
	fmt.Printf ("Item2: (%%+v) %+v\n", item2)
	fmt.Printf ("Item3: (%%#v) %#v\n", item3)


	err := errors.New("An error occurred while processing the configuration")
	fmt.Printf("Error: %d: %w\n", port,  err)

}