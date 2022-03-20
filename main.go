package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Emyrk/float-breakdown/breakdown"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}

	f, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, usage())
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	_, _ = fmt.Fprintf(os.Stdout, "Value to breakdown: ")
	_, _ = color.New(color.FgGreen).Fprintf(os.Stdout, "%f\n", f)

	b := breakdown.Float32(float32(f))
	_, _ = fmt.Fprintln(os.Stdout, b.String())

}

func usage() string {
	return fmt.Sprintf("float-breakdown <float32>\n\tExample: breakdown 52.123")
}
