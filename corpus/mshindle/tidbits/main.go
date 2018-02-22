package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"errors"
	"strings"

	"github.com/mshindle/tidbits/dynamic"
	"github.com/mshindle/tidbits/limit"
	"github.com/mshindle/tidbits/retry"
	"github.com/mshindle/tidbits/toy"
	"github.com/urfave/cli"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	app := cli.NewApp()
	app.Name = "tidbits"
	app.Usage = "execute samples of learning code"
	app.Author = "Mike Shindle"
	app.Email = "mshindle@gmail.com"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "search",
			Usage: "run simple google search test",
			Action: func(c *cli.Context) {
				rand.Seed(time.Now().UnixNano())
				fmt.Println("Running Google10 =>")
				search(toy.Google10)
				fmt.Println("Running Google20 =>")
				search(toy.Google20)
			},
		},
		{
			Name:  "whisper",
			Usage: "play whisper adding 1 to every number passed",
			Action: func(c *cli.Context) {
				fmt.Println("Running whisper =>")
				toy.Whisper()
			},
		},
		{
			Name:    "coins",
			Aliases: []string{"c"},
			Usage:   "find min number of coins needed to make a value",
			Flags: []cli.Flag{
				cli.IntFlag{Name: "value, v", Value: 26, Usage: "value in cents to have coins sum to"},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("Running coins =>")
				dynamic.Coins(c.Int("value"))
				return nil
			},
		},
		{
			Name:    "breaker",
			Aliases: []string{"b"},
			Usage:   "run a circuit breaker example",
			Action: func(c *cli.Context) error {
				fmt.Println("Running breaker =>")
				retry.RunBreaker()
				return nil
			},
		},
		{
			Name:    "limit",
			Aliases: []string{"l"},
			Usage:   "run a request rate limiter example",
			Action: func(c *cli.Context) error {
				fmt.Println("Running limiter =>")
				limit.RunRequest()
				return nil
			},
		},
		{
			Name:    "oddword",
			Aliases: []string{"o"},
			Usage:   "parse a string of words - reversing the odd numbered words",
			Action: func(c *cli.Context) error {
				text := strings.Join(c.Args(), " ")
				if !strings.HasSuffix(text, ".") {
					return errors.New("string must terminate with a `.`")
				}
				toy.Oddword(text)
				return nil
			},
		},
		{
			Name:    "gcd",
			Aliases: []string{"g"},
			Usage:   "return the greatest common denominator between two numbers",
			Action: func(c *cli.Context) error {
				fmt.Println("Running gcd => 12 15")
				gcd := toy.GCD(12, 15)
				fmt.Printf("GCD => %d\n", gcd)

				fmt.Println("Running lcm => 12 15")
				lcm := toy.LCM(12, 15)
				fmt.Printf("LCM => %d\n", lcm)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// search excutes the defined search functions
func search(f func(query string) []toy.Result) {
	start := time.Now()
	results := f("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
