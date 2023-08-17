package cmd

import (
	"github.com/spf13/cobra"
	"star/pkg/queue"
)

// serverCmd represents the server command
var queueCmd = &cobra.Command{
	Use:   "queue:start",
	Short: "A brief description of your command",
	Long:  `test aaa`,
	Run:   Run,
}

func init() {
	RootCmd.AddCommand(queueCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Run(command *cobra.Command, strings []string) {
	queue.Start()
}
