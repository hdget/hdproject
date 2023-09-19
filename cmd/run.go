package cmd

import "github.com/spf13/cobra"

var cliAddress string
var runCmd = &cobra.Command{
	Long:  "run server",
	Short: "run server",
	Use:   "run",
}

func init() {
	runCmd.PersistentFlags().StringVarP(&cliAddress, "address", "d", "", "listen address, e,g: ':8888'")

	runCmd.AddCommand(runDaprGrpcServerCmd)
	runCmd.AddCommand(runDaprHttpServerCmd)
	runCmd.AddCommand(runHttpServerCmd)
}
