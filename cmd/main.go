package main

var (
	// ConfigFile : the special place
	ConfigFile string
	// cfg : active object
	cfg         Config
	programName string
	cli         Cli
)

func main() {

	initialize()
	InfoLogger("after init" + ConfigFile)
	watch()

	// cmd := exec.Command("ffmpeg", "-i", "input.avi", "output.mp4")
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
