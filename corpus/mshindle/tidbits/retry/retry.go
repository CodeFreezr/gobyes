package retry

import (
	"encoding/json"
	"net/http"
	"os"
)

type Book struct {
	Id     int
	Title  string
	Author string
}

// listenAndServe sets up a mini web server that serves a predetermined response.
func listenAndServe(addr string, statusCode int, book *Book) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		if statusCode == http.StatusOK {
			output, _ := json.MarshalIndent(book, "", "  ")
			w.Write(output)
		}
		return
	})
	http.ListenAndServe(addr, mux)
}

func RunBreaker(hosts ...string) {
	mobyDick := &Book{Id: 1, Title: "Moby Dick", Author: "Herman Melville"}

	// set up the bad server
	go listenAndServe(":8080", http.StatusInternalServerError, mobyDick)
	// set up the good server
	go listenAndServe(":8081", http.StatusOK, mobyDick)

	// try and get a resource
	client := NewClient(hosts...)

	resp, err := client.Get("/1")
	if err != nil {
		panic(err)
	}
	resp.Write(os.Stdout)
}
