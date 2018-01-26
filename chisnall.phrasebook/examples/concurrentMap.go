package main
import "fmt"
import "sync"

type request struct {
	key int
	value string
	ret chan string
}
type ConcurrentMap struct {
	ch chan request
	init sync.Once
}
func (cm *ConcurrentMap) Set(key int, value string) string {
	cm.init.Do(func () {
			   cm.ch = make(chan request)
			   go runMap(cm.ch)
			   })
	result := make(chan string)
	cm.ch <- request{key, value, result}
	return <-result
}

func runMap(c chan request) {
	m := make(map[int] string)
	for {
		req := <- c
		old := m[req.key]
		m[req.key] = req.value
		req.ret <- old
	}
}

func main() {
	var m ConcurrentMap
	fmt.Printf("Set %s\n", m.Set(1, "foo"))
	fmt.Printf("Set %s\n", m.Set(1, "bar"))
}
