//go:build !windows

package testnpipe

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func Run(path, file string) {
	namedPipe := filepath.Join(path, file)
	if _, err := os.Stat(namedPipe); err == nil {
		fmt.Printf("pipe %s already exists, removing \n", namedPipe)
		os.Remove(namedPipe)
	}
	err := syscall.Mkfifo(namedPipe, 0600)
	if err != nil {
		log.Fatal("make named pipe file error:", err)
	}
	fmt.Println("opening pipe")
	f, err := os.OpenFile(namedPipe, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		log.Fatal("open named pipe file error:", err)
	}

	reader := bufio.NewReader(f)
	fmt.Println("new connection found")
	for {
		line, err := reader.ReadBytes('\n')
		if err == nil {
			fmt.Print("load string:" + string(line))
		}
	}
}
