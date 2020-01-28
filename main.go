package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/fenwickelliott/dirty_socks/model"
)

func serve(c net.Conn) {
	buff := make([]byte, 4)
	_, err := c.Read(buff)
	fatal(err)

	log.Println(string(buff))
	log.Println(c.RemoteAddr())
}

func server() {
	l, err := net.Listen("unix", model.SocketAddress)
	fatal(err)
	defer l.Close()

	for {
		conn, err := l.Accept()
		fatal(err)

		serve(conn)
	}
}

func main() {
	go server()
	time.Sleep(time.Second)
	c, err := net.Dial("unix", model.SocketAddress)
	fatal(err)
	defer c.Close()

	c.Write([]byte("foo"))

	time.Sleep(time.Second)
}

func init() {
	os.RemoveAll(model.SocketAddress)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
