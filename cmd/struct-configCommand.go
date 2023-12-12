package main

// ConfigCommand type to add to command line parser
type ConfigCommand struct {
	Path string `arg optional name:"Path" help:"Paths to list." type:"Path"`
}
