package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "gicli",
	Short: "A simple CLI intended to connect issue trackers and git handling into a single application.",
	Long: "Help and sourcecode can be found at https://github.com/jitesoft/gi-cli",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}