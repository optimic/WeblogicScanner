package main

import (
	"fmt"
	//"io/ioutil"
	//"github.com/asmcos/requests"
	//"net/http"
	"os"

	"poc/CVE_2014_4210"
	"poc/CVE_2016_0638"
	"poc/CVE_2016_3510"
	"poc/CVE_2017_10271"
	"poc/CVE_2017_3248"
	"poc/CVE_2017_3506"
	"poc/CVE_2018_2628"
	"poc/CVE_2018_2893"
	"poc/CVE_2018_2894"
	"poc/CVE_2019_2725"
	"poc/CVE_2019_2729"
)

/*func islive(u string, port string) int {
	client := &http.Client{}
	url := "http://" + u + ":" + port + "/uddiexplorer/"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "ceshi/0.0.1")
	if err != nil {
		return -1
	}
	response, err := client.Do(request)
	if err != nil {
		return -1
	}
	defer response.Body.Close()
	//body, err := ioutil.ReadAll(response.Body)
	//fmt.Println("body:", string(body))

	status := response.StatusCode
	return status
}*/

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s [IP] [PORT]\n", args[0])
		return
	}
	url := args[1]
	port := args[2]
	//
	//if islive(url, port) == 200 {
	//	u := "http://" + url + ":" + port + "/uddiexplorer/"
	//	fmt.Printf("[+] The target Weblogic UDDI module is exposed!\n[+] The path is: %s \n[+] Please verify the SSRF vulnerability!", u)
	//} else {
	//	fmt.Printf("[-] The target Weblogic UDDI module default path does not exist!")
	//}

	CVE_2014_4210.Run(url, port)
	CVE_2016_0638.Run(url, port)
	CVE_2016_3510.Run(url, port)
	CVE_2017_10271.Run(url, port)
	CVE_2017_3248.Run(url, port)
	CVE_2017_3506.Run(url, port)
	CVE_2018_2628.Run(url, port)
	CVE_2018_2893.Run(url, port)
	CVE_2018_2894.Run(url, port)
	CVE_2019_2725.Run(url, port)
	CVE_2019_2729.Run(url, port)

}
