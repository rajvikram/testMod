package main

import (
	"fmt"
	"os"

	"../mymod"
)

func main() {
	fmt.Println("Running main")

	mymod.Init("This is a test string")
	os.Exit(1)
}
