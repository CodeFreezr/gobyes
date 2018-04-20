package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tidbits",
	Short: "Run sample apps built for self education",
	Long: `Collection of sample applications and code
snippets to enable me to learn a few things.
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal("exiting application")
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	// globally set logging parameters
	logrus.SetOutput(os.Stdout)
}
