package main

import(
	"fmt"
	"net"
	"os"
	"log"
	"bufio"
)

const(
	connHost="localhost"
	connPort="8080"
	connType="tcp"
)


func handleConnection(conn net.Conn) {
    buffer, err := bufio.NewReader(conn).ReadBytes('\n')

    if err != nil {
        fmt.Println("Client left.")
        conn.Close()
        return
    }

    log.Println("Client message:", string(buffer[:len(buffer)-1]))

    conn.Write(buffer)

    handleConnection(conn)
}

func main(){
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	l,err:=net.Listen(connType,connHost+":"+connPort)
	if(err !=nil){
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer l.Close();
	for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println("Error connecting:", err.Error())
            return
        }
        fmt.Println("Client connected.")

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")
		go handleConnection(c)
	}
}