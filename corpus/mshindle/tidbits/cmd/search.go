// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"time"

	"github.com/mshindle/tidbits/toy"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "run a simple google search test",
	Long:  `run a simple google search test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running Google10 =>")
		search(toy.Google10)
		fmt.Println("Running Google20 =>")
		search(toy.Google20)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}

// search excutes the defined search functions
func search(f func(query string) []toy.Result) {
	start := time.Now()
	results := f("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
