package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://guimp.com")
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))
}
