package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"flag"
)

var GREEN string = "\033[0;32m"
var NC string = "\033[0m"

//name of OS
var OS string = runtime.GOOS

func main(){
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if(len(args)< 1){
		fmt.Println("ERROR: missing url (--help)")
		os.Exit(1)
	}
	logo()
	url := args[0]
	fmt.Println("Target:", url)
	
	fmt.Print("Open Firefox to wayback machine.")
	runF(url,"w")
	
	fmt.Print("Open Firefox to the site:", url)
	runF(url,"s")	

	fmt.Print("Open Firefox using google dorking.")
	runF(url,"g")

	fmt.Println("FUZING.")
	fuff(url)

	fmt.Println("FINISH\a")
}

func fuff(url string){
	url = url + "FUZZ"

	wordlists := []string{
		"wordlist/Web-Content/common.txt",
		"wordlist/Web-Content/common-and-portuguese.txt",
		"wordlist/Web-Content/quickhits.txt",
		"wordlist/Web-Content/directory-list-2.3-medium.txt",
		"wordlist/Web-Content/directory-list-2.3-big.txt",}

	for _, wordlist := range wordlists {
		fmt.Printf("Running ffuf with wordlist: %s\n", wordlist)

		// Construct and execute the ffuf command
		cmd := exec.Command("ffuf", "-u", url, "-w", wordlist, "-fc", "302")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error running ffuf with wordlist %s: %v\n", wordlist, err)
			continue
		}

		fmt.Printf("Output for %s:\n[*] %s\n", wordlist, output)
		fmt.Println(" Done ✅\a")
	}
}

func runF(url string, tipe string){
	switch(OS){
		//LINUX
		case "linux":
			switch(tipe){
				case "g":
					urlw := "https://www.google.com/search?client=firefox-b-d&q=site%3A" + url[8:] + "+|+|+\"" + url[8:] + "\""
					cmd := exec.Command("firefox",urlw)
					err := cmd.Start()
					if err != nil {
						fmt.Println("Error:", err)
					}
					fmt.Println(" Done ✅\a")
				case "w":
					urlw := "https://web.archive.org/web/*/"+url+"*"
					cmd := exec.Command("firefox",urlw)
					err := cmd.Start()
					if err != nil {
						fmt.Println("Error:", err)
					}
					fmt.Println(" Done ✅\a")
				case "s":
					cmd := exec.Command("firefox",url)
								err := cmd.Start()
								if err != nil {
										fmt.Println("Error:", err)
								}
								fmt.Println(" Done ✅\a")
				default:
					fmt.Println("Error")
			}

		//MAC 
		case "darwin":
			switch(tipe){
				case "g":
					urlw := "https://www.google.com/search?client=firefox-b-d&q=site%3A" + url[8:] + "+|+|+\"" + url[8:] + "\""
					cmd := exec.Command("/Applications/Firefox.app/Contents/MacOS/firefox",urlw)
					err := cmd.Start()
					if err != nil {
						fmt.Println("Error:", err)
					}
					fmt.Println(" Done ✅\a")
				case "w":
					urlw := "https://web.archive.org/web/*/"+url+"*"
					cmd := exec.Command("/Applications/Firefox.app/Contents/MacOS/firefox",urlw)
					err := cmd.Start()
					if err != nil {
						fmt.Println("Error:", err)
					}
					fmt.Println(" Done ✅\a")
				case "s":
					cmd := exec.Command("/Applications/Firefox.app/Contents/MacOS/firefox",url)
								err := cmd.Start()
								if err != nil {
										fmt.Println("Error:", err)
								}
								fmt.Println(" Done ✅\a")
				default:
					fmt.Println("Error")
			}
		default:
			fmt.Println("MAC or LINUX")
			os.Exit(0)
	}
}

func logo(){

	var d string = `
†'########:::::'##:::'##::::'##:::'#####::: †​
† ##.... ##::'####::: ###::'###::'##.. ##:: †
† ##:::: ##::.. ##::: ####'####:'##:::: ##: †
† ########::::: ##::: ## ### ##: ##:::: ##: †
† ##.... ##:::: ##::: ##. #: ##: ##:::: ##: †
† ##:::: ##:::: ##::: ##:.:: ##:. ##:: ##:: †
† ########:::'######: ##:::: ##::. #####::: †
†........::::......::..:::::..::::.....:::: †

`
	fmt.Print(GREEN)
	fmt.Println(d)
	fmt.Print(NC)
}

func usage() {
	logo()
    fmt.Fprintf(os.Stderr, "usage: api-scan <url>\n")
    flag.PrintDefaults()
    os.Exit(2)
}
