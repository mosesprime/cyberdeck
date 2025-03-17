package main

import (
	"flag"
	"log"
	"net"
	"time"
)

func main() {
    listenPort := flag.String("p", "2222", "listen port")

    flag.Parse()

    listener, err := net.Listen("tcp", "0.0.0.0:"+*listenPort)
    if err != nil {
        log.Panicln(err)
    }

    defer listener.Close()
    log.Println("Listening on "+*listenPort)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    log.Println("Hit: "+conn.RemoteAddr().String())
    defer conn.Close()

    start := time.Now().Unix()

    for {
        msgStr := "banana" // TODO: use rand message
        _, err := conn.Write([]byte(msgStr)) // TODO: seperate with /r/n ?
        if err != nil {
            log.Printf("Disconnect: %s Duration: %v\n", conn.RemoteAddr().String(), time.Now().Unix() - start)
            conn.Close()
            break
        }
        time.Sleep(10 * time.Second) // TODO: use rand interval
    }
}
