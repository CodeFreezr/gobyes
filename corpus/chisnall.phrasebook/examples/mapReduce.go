package main
import "fmt"
import "os"
import "text/scanner"
import "strings"
//import "sync"

type Partial struct {
	key string
	value string
}
type Result struct {
	token string
	counts map[string] int
}

func Map(fileName string, intermediate chan Partial) {
	file, err := os.Open(fileName)
	if err == nil {
		var s scanner.Scanner
		s.Init(file)
		tok := s.Scan()
		for tok != scanner.EOF {
			intermediate <- Partial{s.TokenText(), fileName}
			tok = s.Scan()
		}
	}
	intermediate <- Partial{"", ""}
}

func Reduce(token string, files []string, final chan Result) {
	counts := make(map[string] int)
	for _, file := range files {
		counts[file]++
	}
	final <- Result{token, counts}
}

func collectPartials(intermediate chan Partial,
                          count int,
                          final chan map[string] map[string] int) {
	intermediates := make(map[string] []string)
	for count > 0 {
		res := <- intermediate
		if res.value == "" && res.key == "" {
			count--
		} else {
			v := intermediates[res.key]
			if v == nil {
				v = make([]string, 0, 10)
			}
			v = append(v, res.value)
			intermediates[res.key] = v
		}
	}
	collect := make(chan Result)
	for token, files := range intermediates {
		go Reduce(token, files, collect)
	}
	results := make(map[string] map[string] int)
	// Collect one result for each goroutine we spawned
	for _, _ = range intermediates {
		r := <- collect
		results[r.token] = r.counts
	}
	final <- results
}

func main() {
	intermediate := make(chan Partial)
	final := make(chan map[string] map[string] int)
	dir, _ := os.Open(".")
	names, _ := dir.Readdirnames(-1)
	go collectPartials(intermediate, len(names), final)
	for _, file := range names {
		if (strings.HasSuffix(file, ".go")) {
			go Map(file, intermediate)
		} else {
			intermediate <- Partial{"", ""}
		}
	}
	result := <- final
	for token, counts := range result {
		fmt.Printf("\n\nToken: %v\n", token)
		total := 0
		for file, count := range counts {
			fmt.Printf("\t%s:%d\n", file, count)
			total += count
		}
		fmt.Printf("Total: %d\n", total)
	}
}
