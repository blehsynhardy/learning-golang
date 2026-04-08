package main

import (
	"fmt"

	"github.com/blehysnhardy/go-modules/internal/color"
)

func main() {

	fmt.Println("===========COLORER START HERE ===========")

	blackText := "this is black text"
	redText := "this is a red text"
	blueText := "this is a blue text"

	fmt.Println(color.Text(redText, color.Red))
	fmt.Println(color.Text(blueText, color.Blue))
	fmt.Println(color.Text(blackText, color.Black))
}	