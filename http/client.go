package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com/", nil)
	if err != nil {
		panic(err)
	}
	//添加请求头
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Mobile Safari/537.36")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}

	resp, err := client.Do(request)
	//resp, err := http.Get("https://www.imooc.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//用 httputil 简化工作
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
