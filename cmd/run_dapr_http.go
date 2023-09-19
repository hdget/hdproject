package cmd

import (
	"github.com/pkg/errors"
	"hdproject/g"
	"hdproject/pkg/service"
	"net/http"

	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/hdget/hdsdk"
	"github.com/hdget/hdsdk/utils"
	"github.com/spf13/cobra"
)

var runDaprHttpServerCmd = &cobra.Command{
	Use:   "dapr_http",
	Short: "run dapr http server",
	Long:  "run dapr http server",
	PreRun: func(cmd *cobra.Command, args []string) {
		err := hdsdk.Initialize(g.Config)
		if err != nil {
			utils.LogFatal("sdk initialize", "err", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		runDaprHttpServer()
	},
}

func runDaprHttpServer() {
	server := daprd.NewService(cliAddress)
	if server == nil {
		hdsdk.Logger.Fatal("new dapr http server", "error", "empty http server")
	}

	svc, err := service.New()
	if err != nil {
		hdsdk.Logger.Fatal("new service", "error", err)
	}

	err = svc.Initialize(server, service.Generators...)
	if err != nil {
		hdsdk.Logger.Fatal("service initializing", "error", err)
	}

	hdsdk.Logger.Debug("start http service", "address", cliAddress)
	if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		hdsdk.Logger.Fatal("start http service", "error", err)
	}
}
