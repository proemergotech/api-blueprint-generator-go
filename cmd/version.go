package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/proemergotech/api-blueprint-generator-go/app"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.AppVersion)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.
}
