//go:build windows

package testnpipe

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/Microsoft/go-winio"
)

func Run(path, file string) {
	namedPipe := fmt.Sprintf("%s%s", path, file)
	l, err := winio.ListenPipe(namedPipe, nil)
	if err != nil {
		log.Fatal("error creating listener: ", err)
	}
	defer l.Close()
	log.Printf("listener started on: %v\n", namedPipe)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("error accepting new connection to listener:", err)
		}
		go handleClient(conn)
	}
}

func handleClient(c net.Conn) {
	defer c.Close()
	log.Printf("client connected [%s]", c.RemoteAddr().Network())

	buf := make([]byte, 8*1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("read error: %v\n", err)
			}
			break
		}
		str := string(buf[:n])
		log.Printf("read %d bytes: %q\n", n, str)
	}
	log.Println("client disconnected")
}
