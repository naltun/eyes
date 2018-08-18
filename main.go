package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

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
	fmt.Println("  ____")
	fmt.Println(" |  __|")
	fmt.Println(" | |__ _   _  ___  ___")
	fmt.Println(" |  __| | | |/ _ \\/ __|")
	fmt.Println(" | |__| |_| |  __/\\__ \\ ")
	fmt.Println(" \\____/\\__, |\\___||___/ v0.1beta")
	fmt.Println("        __/ |")
	fmt.Printf("       |____/\n\n")
}

func menu() {
	fmt.Println("1.  Whois Lookup")
	fmt.Println("2.  DNS Lookup + Cloudflare Detector")
	fmt.Println("3.  Zone Transfer")
	fmt.Println("4.  Port Scan")
	fmt.Println("5.  HTTP Header Grabber")
	fmt.Println("6.  Honeypot Detector")
	fmt.Println("7.  Robots.txt Scanner")
	fmt.Println("8.  Link Grabber")
	fmt.Println("9.  IP Location Finder")
	fmt.Println("10. Traceroute")
	fmt.Println("11. Domain-to-IP Lookup")
	fmt.Println("12. About program")
	fmt.Println("13. Exit program")
}

func eyes() {
	fmt.Print("What do you want to do? ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		fmt.Print("Enter a domain or IP address: ")
		fmt.Scanln(&target)
		apiUrl := "http://api.hackertarget.com/whois/?q=" + target
		fmt.Println(curlReq(apiUrl))
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
		fmt.Println("**This feature requires being looked at.**")
		fmt.Print("Enter a domain or IP address: ")
		fmt.Scanln(&target)
		apiUrl := "http://api.hackertarget.com/httpheaders/?q" + target
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
	fmt.Println("")
	menu()
	fmt.Println("")
	eyes()
}

func main() {
	banner()
	menu()
	eyes()
}
