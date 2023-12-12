package main

import "time"

// Pause : wait in millisec
func Pause() {
	var duration time.Duration
	var errMsg error
	duration, errMsg = time.ParseDuration(cfg.watcher.waitInMilliSecs)
	if errMsg != nil {
		PanicLoggerErr(errMsg)
	}

	time.Sleep(duration)
}
