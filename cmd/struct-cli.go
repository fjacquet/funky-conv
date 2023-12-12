package main

// Cli command line possibilities
type Cli struct {
	Debug  bool          `help:"Enable debug mode."`
	Config ConfigCommand `cmd Config:"Path to configuration file."`
}
