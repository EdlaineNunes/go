package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.google.com/robots.txt")
	robots, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", robots)
}
