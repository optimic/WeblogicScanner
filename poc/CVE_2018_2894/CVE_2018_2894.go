package CVE_2018_2894

import (
	"fmt"
	"net/http"
)

var VUL = "CVE-2018-2894"

func islive(u string, port string) int {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	url := "http://" + u + ":" + port + "/ws_utc/resources/setting/options/general"
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

	if islive(u, port) != 404 {
		fmt.Printf("[+] The target weblogic has a JAVA deserialization vulnerability: %s\n", VUL)
	} else {
		fmt.Printf("[-] Target weblogic not detected %s\n", VUL)
	}

}
