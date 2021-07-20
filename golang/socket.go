package main

import (
    "net"
    "fmt"
    "time"
    "bytes"
    "io"
)

func printServerLog(format string, args ...interface{}) {
    fmt.Println("Server: " + fmt.Sprintf(format, args...))
}

func handleConn(conn net.Conn) {
    conn.SetReadDeadline(time.Now().Add(10 * time.Second))
    var buffer bytes.Buffer
    for {
        readBytes := make([]byte, 1)
        _,err := conn.Read(readBytes)
        if err != nil {
            if err == io.EOF {
                printServerLog("read complete")
                break
            }
            panic(err)
        }
        buffer.WriteByte(readBytes[0])
    }
    printServerLog("read content:", buffer.String())
    write(conn, buffer.String())
}

func write(conn net.Conn, str string) {
    var buffer bytes.Buffer
    buffer.WriteString(str)
    conn.Write(buffer.Bytes())
}

func main() {
    var listener, err = net.Listen("tcp", "localhost:1234")
    if err != nil {
        printServerLog("Listen error:%s", err)
    }
    defer listener.Close()
    printServerLog("Start listener for the server: %s", listener.Addr())
    for {
        conn, err := listener.Accept()
        if err != nil {
            printServerLog("Accept Error: %s", err)
        }
        printServerLog("Established a connection %s", conn.RemoteAddr())
        go handleConn(conn)
    }
}
