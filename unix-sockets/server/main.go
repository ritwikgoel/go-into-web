package main

import (
    "log"
	"net"
	"os"
)

func echoServer(c net.Conn) {
    for {
        buf := make([]byte, 512)
        nr, err := c.Read(buf)
        if err != nil {
            return
        }

        data := buf[0:nr]
        println(string(data))
        _, err = c.Write(data)
        if err != nil {
            log.Fatal("Write: ", err)
        }
    }
}

func main() {

	if err := os.RemoveAll("/tmp/teamux.sock"); err != nil {
        log.Fatal(err)
    }

    l, err := net.Listen("unix", "/tmp/teamux.sock")
    if err != nil {
        log.Fatal("listen error:", err)
    }

    for {
        fd, err := l.Accept()
        if err != nil {
            log.Fatal("accept error:", err)
        }

        go echoServer(fd)
    }
}