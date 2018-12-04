package utils

import (
	"fmt"
	"os"
)

var output *os.File

func Print(a ...interface{}) (int, error) {
	return fmt.Fprint(output, a...)
}

func Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(output, format, a...)
}

func Println(a ...interface{}) (int, error) {
	return fmt.Fprintln(output, a...)
}

func OutputOpen() {
	ofile, err := os.Create("./output")
	if err != nil {
		panic(err)
	}
	output = ofile
}

func OutputClose() {
	output.Close()
}
