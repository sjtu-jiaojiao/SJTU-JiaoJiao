package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
)

const srvCount = 6

var srvName = [...]string{
	"consul",
	"micro api",
	"micro web",
	"document",
	"goconvey",
	"realize",
}

// 0 for not call, 1 for not run, 2 for running
var status [srvCount]int

var command [srvCount]*exec.Cmd

var arg = [srvCount][]string{
	{"consul", "agent", "-ui", "-bind=127.0.0.1", "-dev"},
	{"micro", "--registry=consul", "api", "--handler=http"},
	{"micro", "--registry=consul", "web"},
	{"go", "run", "./doc/main.go", "--registry=consul"},
	{"goconvey", "-port", "8400", "-launchBrowser=false", "-poll=5s"},
	{"realize", "start"},
}

var fun [srvCount]func()

var isExit bool

var mutex sync.Mutex

func initService() {
	fun[0] = func() {
		time.Sleep(time.Duration(5) * time.Second)
		err := exec.Command("consul", "kv", "import", "@consul.json").Run()
		if err != nil {
			fmt.Println("Import consul config failed")
		}
	}
}

func main() {
	fmt.Println("Welcome to SJTU-JJ Command Console!")
	fmt.Println("-----------------------------------")
	var c string
	initService()
	showStatus()
	for {
		showCmd()
		_, _ = fmt.Scanf("%v", &c)
		switch c {
		case "h":
			fmt.Println("h: Print help information")
			fmt.Println("e: Exit")
			fmt.Println("r: Restart the whole project")
			for i, n := range srvName {
				fmt.Printf("%v: Restart %v\n", i, n)
			}
		case "e":
			isExit = true
			for _, v := range command {
				if v != nil {
					kill(v)
				}
			}
			os.Exit(0)
		case "r":
			go start(0, fun[0], arg[0]...)
			time.Sleep(time.Duration(6) * time.Second)
			for i := 1; i < srvCount; i++ {
				go start(i, fun[i], arg[i]...)
			}
		case "s":
			showStatus()
		default:
			i, err := strconv.Atoi(c)
			if err == nil {
				go start(i, fun[i], arg[i]...)
			} else {
				fmt.Println("Command not found!")
			}
		}
	}
}

func start(i int, f func(), arg ...string) {
	if status[i] != 0 {
		kill(command[i])
		return
	}

	for {
		fmt.Println("\rRestarting", srvName[i]+"...")
		command[i] = exec.Command(arg[0], arg[1:]...)
		setpid(command[i])

		err := command[i].Start()
		if err != nil {
			fmt.Println("\r\n" + err.Error())
			fmt.Print("\r\nStart ", srvName[i], " error! Retry in 3 seconds...\n")
			time.Sleep(time.Duration(3) * time.Second)
			continue
		}

		if f != nil {
			f()
		}

		status[i] = 2
		fmt.Print("\r"+srvName[i], "\t\t", colorPrint(status[i]), "\n")
		showCmd()
		err = command[i].Wait()
		status[i] = 1
		if err != nil {
			fmt.Print("\r\n"+srvName[i], " terminated!")
		}
		if isExit {
			break
		}
	}
}

func colorPrint(status int) string {
	if status == 2 {
		return color.GreenString("Running")
	}
	return color.RedString("Stopped")
}

func showCmd() {
	fmt.Print("\rcmd> ")
}

func showStatus() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("\n\t***Current Status***")
	for i, n := range srvName {
		fmt.Printf("%v: %v\t\t%v\n", i, n, colorPrint(status[i]))
	}
}
