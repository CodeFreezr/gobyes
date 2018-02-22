package limit

import (
	"fmt"

	"github.com/go-redis/redis"
	//"github.com/mshindle/ratelimit"
	"sync"
)

func RunRequest() {
	client := redis.NewClient(&redis.Options{
		Addr:     "firefly.dev:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	var wg sync.WaitGroup
	//rr := &ratelimit.RequestRate{ReplenishRate: 2, Capacity: 2, Client: client}

	for x := 0; x < 4; x++ {
		//initRequest(x, &wg, rr)
	}
	wg.Wait()
}

//func initRequest(id int, wg *sync.WaitGroup, rr *ratelimit.RequestRate) {
//	ok, err := rr.Limit("mike")
//	fmt.Printf("rate request. id=%d ok=%v err=%v\n", id, ok, err)
//	if !ok {
//		return
//	}
//	wg.Add(1)
//	ms := time.Duration(rand.Intn(1000)) * time.Millisecond
//	go func() {
//		fmt.Printf("start. request=%d\n", id)
//		time.Sleep(ms)
//		fmt.Printf("slept. request=%d time=%v\n", id, ms)
//		wg.Done()
//	}()
//}
