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
	"github.com/mshindle/tidbits/toy"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// whisperCmd represents the whisper command
var whisperCmd = &cobra.Command{
	Use:   "whisper",
	Short: "play whisper adding 1 to every number passed",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Running whisper")
		toy.Whisper()
	},
}

func init() {
	rootCmd.AddCommand(whisperCmd)
}
