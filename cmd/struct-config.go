package main

// Config : master configuration file
type Config struct {
	rootFS struct {
		folder     string `yaml:"folder"`
		scanFolder string `yaml:"scan"`
		outFolder  string `yaml:"out"`
		sdFolder   string `yaml:"480"`
		hdFolder   string `yaml:"720"`
		fhdFolder  string `yaml:"1080"`
	}
	logFolder struct {
		LogName string `yaml:"logFolder"`
	}
	watch struct {
		waitInMilliSecs string `yaml:"waitInMilliSecs"`
	}
	SMTP struct {
		Server    string `yaml:"server"`
		From      string `yaml:"from"`
		To        string `yaml:"to"`
		LogSelect string `yaml:"auth"`
	} `yaml:"smtp"`
}
