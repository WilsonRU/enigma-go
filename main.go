package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	enigma := NewEnigma()

	fmt.Println("Setup")
	enigma.printSetup()

	fmt.Println("----")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		setup := scanner.Text()
		for _, char := range setup {
			fmt.Print(enigma.encode(string(char)))
		}
		fmt.Println("\n--")
		enigma.reset()
	}
}
