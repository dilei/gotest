package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet() {
	// resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	// req, err := http.NewRequest("AAA", "http://www.01happy.com/demo/accept.php?id=1", nil)
	req, err := http.NewRequest("AAA", "http://10.255.242.83/v2/test.php", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	httpGet()
}
