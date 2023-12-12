package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/radovskyb/watcher"
)

func watch() {
	// Create a new Watcher with the specified options.
	w := watcher.New()
	w.IgnoreHiddenFiles(cfg.watcher.DotFiles)
	

	done := make(chan struct{})
	go func() {
		defer close(done)

		for {
			select {
			case event := <-w.Event:
				// Print the event's info.
				fmt.Println(event)
			case err := <-w.Error:
				if err == watcher.ErrWatchedFileDeleted {
					fmt.Println(err)
					continue
				}
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	fmt.Printf("Watching %d files\n", len(w.WatchedFiles()))

	// Parse the interval string into a time.Duration.
	parsedInterval, err := time.ParseDuration(cfg.watcher.waitInMilliSecs)
	if err != nil {
		log.Fatalln(err)
	}

	closed := make(chan struct{})

	c := make(chan os.Signal)
	signal.Notify(c, os.Kill, os.Interrupt)
	go func() {
		<-c
		w.Close()
		<-done
		fmt.Println("watcher closed")
		close(closed)
	}()

	// Run the command before watcher starts if one was specified.
	// go func() {
	// 	if *cmd != "" && *startcmd {
	// 		c := exec.Command(cmdName, cmdArgs...)
	// 		c.Stdin = os.Stdin
	// 		c.Stdout = os.Stdout
	// 		c.Stderr = os.Stderr
	// 		if err := c.Run(); err != nil {
	// 			if (c.ProcessState == nil || !c.ProcessState.Success()) && *keepalive {
	// 				log.Println(err)
	// 				return
	// 			}
	// 			log.Fatalln(err)
	// 		}
	// 	}
	// }()

	// Start the watching process.
	if err := w.Start(parsedInterval); err != nil {
		log.Fatalln(err)
	}

	<-closed
}
