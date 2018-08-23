package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	whois "github.com/likexian/whois-go"
)

/*

The structure of the `eyes' source code is as follows:

	Variables
		|
		v
	Functions
		|
		v
	  Main

#############
# VARIABLES #
#############

In this section, we declare variables that we will use elsewhere in the source code.

#############
# FUNCTIONS #
#############

In this section, we delcare any functions that we need. The primary function that
we will use is eyes(). This is the function that handles 1) variables and
2) functions to implement the program logic.

In the eyes() function, we use a simple switch statement to handle user input
and complete tasks. For each case, we handle the logic for accomplishing the task.

NOTE: Effort is made to keep the program API as small as possible. To aid this effort,
instead of starting with a new function, we should first consider whether we can add
the logic under a particular switch case.

The example for this can be case "1": originally written as a whoIs() function, code was
added directly under case "1" and the function was removed in order to keep the API small.

########
# MAIN #
########

In this section, we leverage our functions to handle the logic. Under main, the primary
function we will use is eyes().

Effort is made to keep the main function as small as possible.

*/

/**************
** VARIABLES  *
**************/

var choice string
var target string

/************
* FUNCTIONS *
************/

// Make a curl request to the API endpoints
func curlReq(url string) string {
	req, _ := http.NewRequest("GET", url, nil) // Wherever `<myVar>, _ :=...' is seen, the `_' is the required *second* variable
	resp, _ := http.DefaultClient.Do(req)      // Certain Golang functions require multiple assigned variables in order to be used
	body, _ := ioutil.ReadAll(resp.Body)       // Normally we would use `err' for error handling, but I have chosen to ignore this. `_' is the throwaway variable
	resp.Body.Close()
	return string(body)
}

// banner created with figlet.js (github.com/patorjk/figlet.js)
func banner() {
	fmt.Print(`
  ____
 |  __|
 | |__ _   _  ___  ___
 |  __| | | |/ _ \/ __|
 | |__| |_| |  __/\__ \
 \____/\__, |\___||___/ v0.1beta
        __/ |
       |____/

`)
}

func menu() {
	fmt.Print(`
1. Whois Lookup
2. DNS Lookup + Cloudflare Detector
3. Zone Transfer
4. Port Scan
5. HTTP Header Grabber
6. Honeypot Detector
7. Robots.txt Scanner
8. Link Grabber
9. IP Location Finder
10. Traceroute
11. Domain-to-IP Lookup
12. About Program
13. Exit Program

`)
}

func eyes() {
	fmt.Print("What do you want to do? ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		fmt.Print("Enter a domain or IP address: ")
		fmt.Scanln(&target)
		res, err := whois.Whois(target)
		if err != nil {
			panic(err)
			display()
		}
		fmt.Println(res)
		display()

	case "2":
		// fmt.Print("Enter a domain: ")
		// fmt.Scanln(&target)
		// apiUrl := "http://api.hackertarget.com/dnslookup/?q=" + target
		// fmt.Println(curlReq(apiUrl))
		fmt.Println("This feature is being worked on.")
		display()

	case "3":
		fmt.Print("Enter a domain: ")
		fmt.Scanln(&target)
		apiUrl := "http://api.hackertarget.com/zonetransfer/?q=" + target
		fmt.Println(curlReq(apiUrl))
		display()

	case "4":
		fmt.Print("Enter a domain or IP address: ")
		fmt.Scanln(&target)
		apiUrl := "http://api.hackertarget.com/nmap/?q=" + target
		fmt.Println(curlReq(apiUrl))
		display()

	case "5":
		fmt.Print("Enter a domain or IP address: ")
		fmt.Scanln(&target)
		apiUrl := "http://api.hackertarget.com/httpheaders/?q=" + target
		fmt.Println(curlReq(apiUrl))
		display()

	case "6":
		fmt.Println("This feature is being worked on.")
		display()

	case "7":
		fmt.Print("This feature makes a direct call to the target -- would you like to continue? [Y\n] ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "" {
			fmt.Print("No argument given.")
			display()
		} else if answer == "y" {
			fmt.Println("Enter domain (without protocol): ")
			fmt.Scanln(&target)
			apiUrl := "http://" + target + "/robots.txt"
			fmt.Println(curlReq(apiUrl))
			display()
		} else {
			fmt.Println("Going back to menu...")
			display()
		}

	case "8":
		fmt.Print("Enter URL (without protocol): ")
		fmt.Scanln(&target)
		apiUrl := "https://api.hackertarget.com/pagelinks/?q=http://" + target
		fmt.Println(curlReq(apiUrl))
		display()

	case "9":
		fmt.Print("Enter IP address: ")
		fmt.Scanln(&target)
		if target == "" {
			fmt.Println("No argument given.")
			display()
		}
		apiUrl := "http://ipinfo.io/" + target + "/geo"
		fmt.Println(curlReq(apiUrl))
		display()

	case "10":
		fmt.Print("Enter a domain or IP address: ")
		fmt.Scanln(&target)
		if target == "" {
			fmt.Println("No argument given.")
			display()
		}
		apiUrl := "https://api.hackertarget.com/mtr/?q=" + target
		fmt.Println(curlReq(apiUrl))
		display()

	case "11":
		fmt.Print("Enter a domain: ")
		fmt.Scanln(&target)
		if target == "" {
			fmt.Println("No argument given.")
			display()
		}
		ipAddr, _ := net.LookupIP(target)
		for _, ip := range ipAddr {
			fmt.Println(ip)
		}
		display()

	case "12":
		fmt.Println("This program was created by Noah Altunian (skeeba), and was adapted from eyes (github.com/naltun/eyes), which was")
		fmt.Println("adapted from ReconDog (github.com/UltimateHackers/ReconDog). It is licensed under the GNU GPLv2. Love your Free/Libre software.")
		display()

	case "13":
		fmt.Println("Bye")
		os.Exit(0)

	default:
		fmt.Println("Your choice is invalid.")
		display()
	}
}

func display() {
	menu()
	eyes()
}

func main() {
	banner()
	menu()
	eyes()
}
