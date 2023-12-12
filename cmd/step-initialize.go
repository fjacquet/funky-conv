package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"gopkg.in/yaml.v2"
)

// ReadFile do read a yaml file
func readFile(Cfg *Config, filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		ProcessError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(Cfg)
	if err != nil {
		ProcessError(err)
	}
}

// initialize warm up
func initialize() {
	var runVersion string

	// fmt.Println(quote.Go())
	currentTime := time.Now()
	runVersion = currentTime.Format("2006-01-02T15:04:05")
	prepareFilesystemStructure()
	prepareLogs()
	// program name management
	programName = os.Args[0] + "-" + runVersion
	InfoLogger("starting at " + runVersion)
	checkParams()
	readFile(&cfg, ConfigFile)

}

// Generate log name
func prepareLogs() {

	logFile, err := os.OpenFile(cfg.logFolder.LogName+"run.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	// log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})
}

// Run in the case of a configuration parameter
func (l *ConfigCommand) Run(ctx *Context) error {
	fmt.Println("config file is ", l.Path)
	ConfigFile = l.Path
	return nil
}

// checkParams verify command line
func checkParams() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid call, pleases try " + os.Args[0] + " --help ")
		os.Exit(1)
	}

	// command line management
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})

	ctx.FatalIfErrorf(err)
	if !fileExists(ConfigFile) {
		fmt.Println("can not find file " + ConfigFile)
		os.Exit(1)
	}
}

// prepareFilesystemStructure : prepare layout
func prepareFilesystemStructure() {
	checkCreateDirectory(cfg.rootFS.folder)
	checkCreateDirectory(cfg.rootFS.scanFolder)
	checkCreateDirectory(cfg.rootFS.outFolder)
	checkCreateDirectory(cfg.rootFS.sdFolder)
	checkCreateDirectory(cfg.rootFS.hdFolder)
	checkCreateDirectory(cfg.rootFS.fhdFolder)
	checkCreateDirectory(cfg.logFolder.LogName)
}
