package httptest

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
)

func FasthttpGet() {
	url := `http://httpbin.org/get`

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println(string(resp))
}

func Get() {
	url := `http://httpbin.org/get`

	resp, err := http.Get(url)
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))
}

func main() {
	Get()
	FasthttpGet()
}
