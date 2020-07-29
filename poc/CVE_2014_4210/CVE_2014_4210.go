package CVE_2014_4210

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	//"os"
)

func islive(u string, port string) int {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
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
}

func Run(u string, port string) {
	//args := os.Args
	//if args == nil || len(args) < 3 {
	//	fmt.Println("Usage: WeblogicScanner.exe [IP] [PORT]")
	//	return
	//}
	//url := args[1]
	//port := args[2]

	if islive(u, port) == 200 {
		url := "http://" + u + ":" + port + "/uddiexplorer/"
		fmt.Printf("[+] The target Weblogic UDDI module is exposed!\n[+] The path is: %s \n[+] Please verify the SSRF vulnerability!\n", url)
	} else {
		fmt.Printf("[-] The target Weblogic UDDI module default path does not exist!\n")
	}

}
