package cmd

import (
	"github.com/spf13/cobra"
	"star/internal"
	"star/internal/provider"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long:  `test aaa`,
	Run: func(command *cobra.Command, strings []string) {
		var app internal.Application
		app.Run(provider.Route)
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
