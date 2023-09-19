package cmd

import (
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/hdget/hdproject/g"
	"github.com/hdget/hdproject/pkg/service"
	"github.com/hdget/hdsdk"
	"github.com/hdget/hdsdk/utils"
	"github.com/spf13/cobra"
)

var runDaprGrpcServerCmd = &cobra.Command{
	Use:   "dapr_grpc",
	Short: "run dapr grpc server",
	Long:  "run dapr grpc server",
	PreRun: func(cmd *cobra.Command, args []string) {
		err := hdsdk.Initialize(g.Config)
		if err != nil {
			utils.LogFatal("sdk initialize", "err", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		runDaprGrpcServer()
	},
}

func runDaprGrpcServer() {
	server, err := daprd.NewService(cliAddress)
	if err != nil {
		hdsdk.Logger.Fatal("new dapr grpc server", "error", err)
	}

	svc, err := service.New()
	if err != nil {
		hdsdk.Logger.Fatal("new service", "error", err)
	}

	err = svc.Initialize(server, service.Generators...)
	if err != nil {
		hdsdk.Logger.Fatal("service initializing", "error", err)
	}

	hdsdk.Logger.Debug("start grpc service", "address", cliAddress)
	if err := server.Start(); err != nil {
		hdsdk.Logger.Fatal("start grpc service", "error", err)
	}

	hdsdk.Logger.Debug("start grpc service", "address", cliAddress)
	if err := server.Start(); err != nil {
		hdsdk.Logger.Fatal("start grpc service", "error", err)
	}
}
