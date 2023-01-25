package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	versionCommand = &cli.Command{
		Action:      printVersion,
		Name:        "version",
		Usage:       "Print version numbers",
		ArgsUsage:   " ",
		Description: `The output of this command is supposed to be machine-readable.`,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "Print a greeting"
	app.Commands = []*cli.Command{
		versionCommand,
	}
	app.Run(os.Args)
}
func printVersion(ctx *cli.Context) error {
	git, _ := VCS()

	fmt.Println(strings.Title(clientIdentifier))
	fmt.Println("Version:", VersionWithMeta)
	if git.Commit != "" {
		fmt.Println("Git Commit:", git.Commit)
	}
	if git.Date != "" {
		fmt.Println("Git Commit Date:", git.Date)
	}
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Printf("GOPATH=%s\n", os.Getenv("GOPATH"))
	fmt.Printf("GOROOT=%s\n", runtime.GOROOT())
	return nil
}
