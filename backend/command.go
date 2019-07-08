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
	MicroApi bool
	MicroWeb bool
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
		case 'h':
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
		case 'p':
			go startConsul()
			time.Sleep(time.Duration(7) * time.Second)
			go startMicroApi()
			go startMicroWeb()
			go startServeDoc()
			go startGoConvey()
			go startRealize()
		case 'C':
		case 'c':
			go startConsul()
		case 'A':
		case 'a':
			go startMicroApi()
		case 'W':
		case 'w':
			go startMicroWeb()
		case 'D':
		case 'd':
			go startServeDoc()
		case 'G':
		case 'g':
			go startGoConvey()
		case 'R':
		case 'r':
			go startRealize()
		}
	}
}

func startConsul() {
	if status.Consul {
		fmt.Println("\nDon't Restart Module being running.")
		printFlag <- true
		return
	}

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

func startMicroApi() {
	if status.MicroApi {
		fmt.Println("\nDon't Restart Module being running.")
		printFlag <- true
		return
	}

	for {
		time.Sleep(time.Duration(1) * time.Second)

		cmd := exec.Command("micro", "--registry=consul", "api", "--handler=http")

		err := cmd.Start()
		if err != nil {
			fmt.Println("start %q error.", cmd)
			continue
		}

		fmt.Println("\nMicro api Restart!")
		status.MicroApi = true
		printFlag <- true
		err = cmd.Wait()
		if err != nil {
			fmt.Println("\nMicro api Terminated!")
			status.MicroApi = false
			printFlag <- true
		}
	}
}

func startMicroWeb() {
	if status.MicroWeb {
		fmt.Println("\nDon't Restart Module being running.")
		printFlag <- true
		return
	}

	for {
		time.Sleep(time.Duration(1) * time.Second)

		cmd := exec.Command("micro", "--registry=consul", "web")

		err := cmd.Start()
		if err != nil {
			fmt.Println("start %q error.", cmd)
			continue
		}

		fmt.Println("\nMicro web Restart!")
		status.MicroWeb = true
		printFlag <- true
		err = cmd.Wait()
		if err != nil {
			fmt.Println("\nMicro web Terminated!")
			status.MicroWeb = false
			printFlag <- true
		}
	}
}

func startServeDoc() {
	if status.Document {
		fmt.Println("\nDon't Restart Module being running.")
		printFlag <- true
		return
	}

	for {
		time.Sleep(time.Duration(1) * time.Second)

		cmd := exec.Command("go", "run", "servedoc.go", "--registry=consul")

		err := cmd.Start()
		if err != nil {
			fmt.Println("start %q error.", cmd)
			continue
		}

		fmt.Println("\nServeDoc Restart!")
		status.Document = true
		printFlag <- true
		err = cmd.Wait()
		if err != nil {
			fmt.Println("\nServeDoc Terminated!")
			status.Document = false
			printFlag <- true
		}
	}
}

func startGoConvey() {
	if status.Goconvey {
		fmt.Println("\nDon't Restart Module being running.")
		printFlag <- true
		return
	}

	for {
		time.Sleep(time.Duration(1) * time.Second)

		cmd := exec.Command("goconvey", "-port", "8400", "--launchBrowser=false")

		err := cmd.Start()
		if err != nil {
			fmt.Println("start %q error.", cmd)
			continue
		}

		fmt.Println("\nGoConvey Restart!")
		status.Goconvey = true
		printFlag <- true
		err = cmd.Wait()
		if err != nil {
			fmt.Println("\nGoConvey Terminated!")
			status.Goconvey = false
			printFlag <- true
		}
	}
}

func startRealize() {
	if status.Realize {
		fmt.Println("\nDon't Restart Module being running.")
		printFlag <- true
		return
	}

	for {
		time.Sleep(time.Duration(1) * time.Second)

		cmd := exec.Command("realize", "start")
		time.AfterFunc(600*time.Second, func() { cmd.Process.Kill() })

		err := cmd.Start()
		if err != nil {
			fmt.Println("start %q error.", cmd)
			continue
		}

		fmt.Println("\nRealize Restart!")
		status.Realize = true
		printFlag <- true
		err = cmd.Wait()
		if err != nil {
			fmt.Println("\nRealize Terminated!")
			status.Realize = false
			printFlag <- true
		}
	}
}

func showStatus() {
	for {
		if <-printFlag {
			fmt.Println("***Current Status***")
			fmt.Println("consul \t\t", status.Consul)
			fmt.Println("microApi \t", status.MicroApi)
			fmt.Println("microWeb \t", status.MicroWeb)
			fmt.Println("document \t", status.Document)
			fmt.Println("goconvey \t", status.Goconvey)
			fmt.Println("realize \t", status.Realize)
			fmt.Print("Please enter command(H for help):")
		}
	}
}
