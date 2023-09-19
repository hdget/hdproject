package cmd

import (
	"github.com/hdget/hdproject/pkg/service"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var genCmd = &cobra.Command{
	Long:    "generate related file",
	PostRun: func(cmd *cobra.Command, args []string) {},
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		baseDir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}

		serviceSourceCodePath := filepath.Join(baseDir, "pkg", "service")
		for _, generator := range service.Generators {
			err = generator.Gen(serviceSourceCodePath)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
	Short: "auto generate related file",
	Use:   "gen",
}
