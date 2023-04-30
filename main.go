package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"testnpipe/testnpipe"
)

var path string

func main() {
	f := "elastic_endpoint"

	switch os := runtime.GOOS; os {
	case "windows":
		path = `\\.\pipe\`
	case "linux":
		path = `/tmp/`
	default:
		fmt.Printf("OS Not supported: %s.\n", os)
		return
	}
	fmt.Printf("OS Detected:  %s \n", runtime.GOOS)

	fmt.Printf("Directory: %s \n", path)
	setupCloseHandler()
	testnpipe.Run(path, f)
}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
