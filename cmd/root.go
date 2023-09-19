package cmd

import (
	"github.com/hdget/hdproject/g"
	"github.com/hdget/hdsdk"
	"github.com/hdget/hdsdk/utils"
	"github.com/spf13/cobra"
)

var env string
var configFile string
var rootCmd = &cobra.Command{
	Long:  "live long description",
	Short: "live short description",
	Use:   "live",
}

func init() {
	cobra.OnInitialize(loadConfig)

	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "running environment, e,g: [prod, sim, pre, test, dev, local]")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file")
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(genCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.LogFatal("root command execute", "err", err)
	}
}

func loadConfig() {
	options := make([]hdsdk.ConfigOption, 0)
	if configFile != "" {
		options = append(options, hdsdk.WithConfigFile(configFile))
	}

	err := hdsdk.NewConfig(g.APP, env, options...).Load(&g.Config)
	if err != nil {
		utils.LogFatal("load config", "err", err)
	}
}
