package main

import (
	"fmt"
	"os"

	"github.com/rajvikram/testMod/mymod"
)

func main() {
	fmt.Println("Running main")

	mymod.Init("This is a test string")
	os.Exit(1)
}
