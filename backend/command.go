package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Status struct {
	Consul   bool
	Micro    bool
	Document bool
	Goconvey bool
	Realize  bool
}

var status Status
var printFlag chan bool = make(chan bool, 10)

func main() {
	reader := bufio.NewReader(os.Stdin)
	go showStatus()
	fmt.Println("Welcome to SJTU-jj Command Console!")
	fmt.Println("-----------------------------------")

	printFlag <- true
	for {
		text, _ := reader.ReadString('\n')
		switch text[0] {
		case 'H':
			fmt.Println("H: Print help information")
			fmt.Println("P: Restart the whole project")
			fmt.Println("C: Restart consul")
			fmt.Println("A: Restart micro api")
			fmt.Println("W: Restart micro web")
			fmt.Println("D: Restart document")
			fmt.Println("G: Restart goconvey test")
			fmt.Println("R: Restart realize auto-tool")
			printFlag <- true
		case 'P':
		case 'C':
			go startConsul()
		case 'A':
		case 'W':
		case 'D':
		case 'G':
		case 'R':
		}
	}
}

func startConsul() {
	for {
		time.Sleep(time.Duration(1) * time.Second)

		cmd := exec.Command("consul", "agent", "-ui", "-bind=127.0.0.1", "-dev")

		err := cmd.Start()
		if err != nil {
			fmt.Println("start %q error.", cmd)
			continue
		}
		time.Sleep(time.Duration(5) * time.Second)

		cmd2 := exec.Command("consul", "kv", "import", "@consul.json")
		err = cmd2.Run()
		if err != nil {
			fmt.Println("import consul failed")
		}

		fmt.Println("\nConsul Restart!")
		status.Consul = true
		printFlag <- true
		err = cmd.Wait()
		if err != nil {
			fmt.Println("\nConsul Terminated!")
			status.Consul = false
			printFlag <- true
		}
	}
}

func showStatus() {
	for {
		if <-printFlag {
			fmt.Println("***Current Status***")
			fmt.Println("consul \t\t", status.Consul)
			fmt.Println("micro \t\t", status.Micro)
			fmt.Println("document \t", status.Document)
			fmt.Println("goconvey \t", status.Goconvey)
			fmt.Println("realize \t", status.Realize)
			fmt.Print("Please enter command(H for help):")
		}
	}
}
