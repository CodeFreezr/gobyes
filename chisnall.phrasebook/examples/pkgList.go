package main
import "fmt"
import "exp/html"
import "net/http"

func main() {
	client := &http.Client{}
	url := "http://golang.org/pkg/"
	page, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	tokenizer := html.NewTokenizer(page.Body)
	foundStart := false
	for {
		ty := tokenizer.Next()
		if ty == html.ErrorToken { break }
		if ty != html.StartTagToken { continue }
		t := tokenizer.Token()
		if t.Data != "a" { continue }
		for _, attr := range t.Attr {
			if "href" == attr.Key {
				if !foundStart ||
				   ((len(attr.Val) > 4) &&
					"http" == attr.Val[0:4] ){
					if ".." == attr.Val {
						foundStart = true
					}
					break
				}
				fmt.Printf("%s\n", attr.Val)
			}
		}
	}
}
