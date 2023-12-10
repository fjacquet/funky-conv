package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/alecthomas/kong"
)

var (
	ConfigFile  string
	Cfg         Config
	programName string
	cli         struct {
		Debug  bool          `help:"Enable debug mode."`
		Config ConfigCommand `cmd help:"Path to configuration file."`
	}
	nbuRoot string
)

func prepareLogs() {
	logFile, err := os.OpenFile(Cfg.Server.LogName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	// log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
}

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

	currentTime := time.Now()
	var version string = currentTime.Format("2006-01-02T15:04:05")

	// program name management
	programName = os.Args[0] + "-" + version

	cmd := exec.Command("ffmpeg", "-i", "input.avi", "output.mp4")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
