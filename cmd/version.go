package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

type versionInfo struct {
	GitBranch string
	GitCommit string
	GitTag    string
	BuildDate string
	GoVersion string
	Platform  string
}

var (
	gitBranch string
	gitCommit string
	gitTag    string
	buildDate string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version",
	Run: func(cmd *cobra.Command, args []string) {
		newVersion().Print()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func newVersion() *versionInfo {
	return &versionInfo{
		GitBranch: gitBranch,
		GitCommit: gitCommit,
		GitTag:    gitTag,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func (v versionInfo) Print() {
	fmt.Printf("Version: %s\nGitBranch: %s\nGitCommit: %s\nBuild Date: %s\nGo Version: %s\nOS/Arch: %s\n", v.GitTag, v.GitBranch, v.GitCommit, v.BuildDate, v.GoVersion, v.Platform)
}

func (v versionInfo) String() string {
	return v.GitTag
}
