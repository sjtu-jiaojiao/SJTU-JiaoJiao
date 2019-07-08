package main

import (
	"fmt"
	"os"
	"strings"
	"consul"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to SJTU-jj Command Console!")
	fmt.Println("-----------------------------------")
	for{
		fmt.Println("Please enter command(H for help):")
		text,_ := reader.ReadRune()
		switch text{
		case 'H':
			fmt.Println("H: Print help information")
			fmt.Println("P: Restart the whole project")
			fmt.Println("C: Restart consul")
			fmt.Println("M: Restart micro")
			fmt.Println("D: Restart document")
			fmt.Println("G: Restart goconvey test")
			fmt.Println("R: Restart realize auto-tool")
		case 'P':
		case 'C':
		case 'M':
		case 'D':
		case 'G':
		case 'R':
		}

	}
}