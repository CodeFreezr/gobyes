package main

import "fmt"
import "net/http"
import "os"
import "io"


func main() {
	client := &http.Client{}
	client.CheckRedirect =
		func(req *http.Request, via []*http.Request) error {
		fmt.Fprintf(os.Stderr, "Redirect: %v\n", req.URL);
		return nil
	}
	var url string
	if len(os.Args) < 2 {
		url = "http://golang.org"
	} else {
		url = os.Args[1]
	}
	page, err := client.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return
	}
	io.Copy(os.Stdout, page.Body)
	page.Body.Close()
}
