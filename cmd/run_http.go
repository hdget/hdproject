package cmd

import (
	"github.com/hdget/hdproject/g"
	"github.com/hdget/hdsdk"
	"github.com/hdget/hdsdk/lib/ws"
	"github.com/hdget/hdsdk/utils"
	"github.com/spf13/cobra"
)

var runHttpServerCmd = &cobra.Command{
	Long:    "run normal http server",
	PostRun: func(cmd *cobra.Command, args []string) {},
	PreRun: func(cmd *cobra.Command, args []string) {
		err := hdsdk.Initialize(g.Config)
		if err != nil {
			utils.LogFatal("sdk initialize", "err", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		runHttpServer()
	},
	Short: "run normal http server",
	Use:   "http",
}

func runHttpServer() {
	ws.SetReleaseMode()
	svc := ws.NewHttpServer(hdsdk.Logger, cliAddress)
	svc.AddRoutes(nil)
	svc.Run()
}
