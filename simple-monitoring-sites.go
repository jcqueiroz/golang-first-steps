// Author: jcqueiroz
// Language US-USA
// Begin learnig Golang
// version:1.0

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 2
const delay = 2

func main() {

	displayIndroduction()

	for {

		displayMenu()
		command := weCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Display logs...")
			printLogs()
		case 0:
			fmt.Println("program exit...")
			os.Exit(0)
		default:
			fmt.Println("I don't recognize this command")
			os.Exit(-1)

		}
	}
}

func displayIndroduction() {
	name := "jcqueiroz"
	version := 1.1
	fmt.Println("Hello, sr.", name)
	fmt.Println("This version of the program is", version)
}

func displayMenu() {
	fmt.Println("1- Display Monitoring")
	fmt.Println("2- Display Logs")
	fmt.Println("0- Exit of Program")
}

func weCommand() int {
	var commandWrite int
	fmt.Scan(&commandWrite)
	fmt.Println("The command chosen was", commandWrite)

	fmt.Println("")

	return commandWrite

}

func startMonitoring() {
	fmt.Println("Monitoring...")
	// sites := []string{"https://random-status-code.herokuapp.com/",
	// 	"https://www.google.com/", "https://www.cloudflare.com"}

	sites := weSitesOfFiles()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}
	fmt.Println("")

}
func testSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Site loading success!")
		registerLog(site, true)

	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registerLog(site, false)

	}

}

func weSitesOfFiles() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		fmt.Println(line)

		if err == io.EOF {

			break
		}
	}

	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(file)

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site +
		" -online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(string(file))
}
