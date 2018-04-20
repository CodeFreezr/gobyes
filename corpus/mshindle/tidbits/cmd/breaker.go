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
	"github.com/mshindle/tidbits/retry"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// breakerCmd represents the breaker command
var breakerCmd = &cobra.Command{
	Use:   "breaker",
	Short: "run a circuit breaker example",
	Long: `Example of how a circuit breaker client looks in golang.
To use the internal services, specify the hosts as

  localhost:8080 localhost:8081

The 8080 host will always return unavailable. The 80801 should always succeed.
`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("running breaker")
		retry.RunBreaker(args...)
	},
}

func init() {
	rootCmd.AddCommand(breakerCmd)
}
