package main
import "fmt"
import "net/http"
import "text/template"

type webCounter struct {
	count chan int
	template *template.Template
}
func NewCounter() *webCounter {
	counter := new(webCounter)
	counter.count = make(chan int, 1)
	go func() {
		for i:=1 ;; i++ { counter.count <- i }
	}()
	counter.template, _ = template.ParseFiles("counter.html")
	return counter
}
func (w *webCounter) ServeHTTP(r http.ResponseWriter, rq *http.Request) {
	if rq.URL.Path != "/" {
		r.WriteHeader(http.StatusNotFound)
		return
	}
	w.template.Execute(r, struct{Counter int}{<-w.count})
}
func main() {
	err := http.ListenAndServe(":8000", NewCounter())
	if err != nil {
		fmt.Printf("Server failed: ", err.Error())
	}
}
