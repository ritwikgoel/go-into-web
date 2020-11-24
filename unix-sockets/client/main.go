package main

import (
    "io"
    "log"
	"net"
	"bufio"
	"os"
	//"time"
	//"fmt"
)

func reader(r io.Reader) {
    buf := make([]byte, 1024)
    for {
        _, err := r.Read(buf[:])
        if err != nil {
            return
        }
        //println(string(buf[0:n]))
    }
}

func main() {
    c, err := net.Dial("unix", "/tmp/teamux.sock")
    if err != nil {
        panic(err)
    }
    defer c.Close()

    go reader(c)
    for {
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')

        _, err := c.Write([]byte(name))
        if err != nil {
            log.Fatal("write error:", err)
            break
        }
    }
}