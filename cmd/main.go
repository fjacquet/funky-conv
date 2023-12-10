package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"rsc.io/quote/v4"
)

var (
	// ConfigFile : the special place
	ConfigFile string
	// Cfg : active object
	Cfg         Config
	programName string
	cli         struct {
		Debug bool `help:"Enable debug mode."`
		// Config Sample file
		Config ConfigCommand `cmd help:"Path to configuration file."`
	}
	nbuRoot string
)

func checkParams() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid call, pleases try " + os.Args[0] + " --help ")
		os.Exit(1)
	}

	// command line management
	ctx := kong.Parse(&cli)
	err := ctx.Run(&context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)

	if !fileExists(ConfigFile) {
		fmt.Println("can not find file " + ConfigFile)
		os.Exit(1)
	}

}

func main() {
	fmt.Println(quote.Go())
	currentTime := time.Now()
	var version string = currentTime.Format("2006-01-02T15:04:05")

	// program name management
	programName = os.Args[0] + "-" + version
	prepareFS()

	// cmd := exec.Command("ffmpeg", "-i", "input.avi", "output.mp4")
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
