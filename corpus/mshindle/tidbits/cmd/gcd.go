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

// gcdCmd represents the gcd command
var gcdCmd = &cobra.Command{
	Use:     "gcd",
	Aliases: []string{"g"},
	Short:   "return the greatest common denominator between two numbers",
	Long:    `return the greatest common denominator between two numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Running gcd => 12 15")
		gcd := toy.GCD(12, 15)
		logrus.WithField("gcd", gcd).Info("calculated")

		logrus.Info("Running lcm => 12 15")
		lcm := toy.LCM(12, 15)
		logrus.WithField("lcm", lcm).Info("calculated")
	},
}

func init() {
	rootCmd.AddCommand(gcdCmd)
}
