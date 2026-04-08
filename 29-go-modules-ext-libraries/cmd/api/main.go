package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	data := [][]string{
		{"Package", "Version", "Status"},
		{"tablewriter", "v0.0.5", "legacy"},
		{"tablewriter", "v1.1.4", "latest"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(data[0])
	table.AppendBulk(data[1:])
	table.Render()
}