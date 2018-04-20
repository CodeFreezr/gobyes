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
	"strings"

	"github.com/mshindle/tidbits/toy"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// oddwordCmd represents the oddword command
var oddwordCmd = &cobra.Command{
	Use:     "oddword",
	Aliases: []string{"o"},
	Short:   "parse a string of words reversing the odd numbered words",
	Long: `oddword takes all the arguments on the command line and
reverses the odd numbered words.

Consider a character set consisting of letters, a space, and a point. Words consist of one or more,
but at most 20 letters. An input text consists of one or more words separated from each other by one or more
spaces and terminated by 0 or more spaces followed by a point. Input should be read from, and including, the
first letter of the first word, up to and including the terminating point. The output text is to be produced
such that successive words are separated by a single space with the last word being terminated by a single point.
Odd words are copied in reverse order while even words are merely echoed. For example, the input string
 : whats the matter with kansas.
becomes
 : whats eht matter htiw kansas.`,
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		if !strings.HasSuffix(text, ".") {
			logrus.Fatal("string must terminate with a `.`")
		}
		toy.Oddword(text)
	},
}

func init() {
	rootCmd.AddCommand(oddwordCmd)
}
