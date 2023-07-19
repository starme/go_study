package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"star/pkg/log"
)

// serverCmd represents the server command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long:  `test aaa`,
	Run:   Handle,
}

func init() {
	RootCmd.AddCommand(testCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type TListener struct{}

func (l TListener) Handler(event interface{}) {
	fmt.Printf("event: %v\n", event)
	log.Info("Handler: ", zap.Any("event", event))
}

func Handle(command *cobra.Command, strings []string) {
	//bus.Register("test", TListener{})
	//bus.Dispatch("test", "this is test event")

	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//log.Info("Shutdown Server ...")
}
