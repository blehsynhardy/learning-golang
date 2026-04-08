package color

import (
	. "fmt"
	"strconv"
	"strings"
)

type Color struct {
	Value int
}

var (
	Black   = Color{Value: 1}
	Red     = Color{Value: 2}
	Green   = Color{Value: 3}
	Yellow  = Color{Value: 4}
	Blue    = Color{Value: 5}
	Magenta = Color{Value: 6}
	Cyan    = Color{Value: 7}
	White   = Color{Value: 8}
)

func Text(text string, colors ...Color) string {

	if len(colors) == 0 {
		return text
	}

	var colorCodes []string

	for _, color := range colors {
		colorCodes = append(colorCodes,strconv.Itoa(color.Value))
	}

	return Sprintf("\033[%sm%s\033[0m", strings.Join(colorCodes, ";"), text)

}
