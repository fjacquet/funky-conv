package main

// ConfigCommand type to add to command line parser
type ConfigCommand struct {
	// filePath:  file path for configuration file
	filePath string `arg optional name:"filePath" help:"Paths to list." type:"filePath"`
}
