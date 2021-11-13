package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("PWD")

	fmt.Printf("Well done, you are running app in a %s folder\n\n", name)
}
