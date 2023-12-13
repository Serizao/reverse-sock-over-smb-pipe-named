package main

import (
	"flag"
	"net"
	"fmt"
	"log"
	"os"
	"io"
"time"
	socks5 "github.com/armon/go-socks5"
	"github.com/hashicorp/yamux"
		"github.com/Serizao/go-winio"
)

var session *yamux.Session

func main() {

	listen := flag.String("listen", "", "listen port for receiver address:port")
	socks := flag.String("socks", "127.0.0.1:1080", "socks address:port")
	connect := flag.String("connect", "", "connect address:port")
	version := flag.Bool("version", false, "version information")
	flag.Usage = func() {
		fmt.Println("rsocks - reverse socks5 server/client")
		fmt.Println("https://github.com/brimstone/rsocks")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("1) Start rsocks -listen :8080 -socks 127.0.0.1:1080 on the client.")
		fmt.Println("2) Start rsocks -connect client:8080 on the server.")
		fmt.Println("3) Connect to 127.0.0.1:1080 on the client with any socks5 client.")
		fmt.Println("4) Enjoy. :]")
	}

	flag.Parse()

	if *version {
		fmt.Println("rsocks - reverse socks5 server/client")
		fmt.Println("https://github.com/brimstone/rsocks")
		os.Exit(0)
	}

	if *listen != "" {
		log.Println("Starting to listen for clients")
		listenForSocks(*listen)

	}

	if *connect != "" {
		log.Println("Connecting to the far end")

		go connectForSocks(*connect)
		log.Fatal(listenForClients(*socks))

	}

	fmt.Fprintf(os.Stderr, "You must specify a listen port or a connect address")
	os.Exit(1)
}

// Catches yamux connecting to us
func listenForSocks(address string) {
	log.Println("Listening for the far end")
	c := winio.PipeConfig{
  SecurityDescriptor: "S:(ML;;NW;;;LW)D:(A;;0x12019f;;;WD)",
  MessageMode:      false,  // Use message mode so that CloseWrite() is supported
  InputBufferSize:  65536, // Use 64KB buffers to improve performance
  OutputBufferSize: 65536,
}
  ln, err := winio.ListenPipe(`\\.\pipe\testSOCK`,&c)
	if err != nil {
		return
	}
server, err := socks5.New(&socks5.Config{})
	conn, err := ln.Accept()
		if err != nil {
			log.Println("error")
			return
		}
		session, err = yamux.Server(conn, nil)
		if err != nil {
			log.Println("error")
			return
		}
	for {

		stream, err := session.Accept()
		if err != nil {
			log.Println("error")
			return
		}
		go func() {
			err = server.ServeConn(stream)
			if err != nil {
				log.Println(err,"error")
			}
		}()
	}
}

// Catches clients and connects to yamux
// Catches clients and connects to yamux
func listenForClients(address string) error {
	log.Println("Waiting for clients")
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		// TODO dial socks5 through yamux and connect to conn

		if session == nil {
			conn.Close()
			continue
		}
		log.Println("Got a client")

		log.Println("Opening a stream")
		stream, err := session.Open()
		if err != nil {
			return err
		}

		// connect both of conn and stream

		go func() {
			log.Println("Starting to copy conn to stream")
			io.Copy(conn, stream)
			conn.Close()
		}()
		go func() {
			log.Println("Starting to copy stream to conn")
			io.Copy(stream, conn)
			stream.Close()
			log.Println("Done copying stream to conn")
		}()
	}
}
func connectForSocks(address string)  {

	log.Println("Connecting to far end")
	  var d = time.Duration(10 * time.Millisecond)
	conn, err := winio.DialPipe(`\\`+address+`\pipe\testSOCK`, &d)
	if err != nil {
		return
	}


		log.Println("Passing off to socks5")
		session, err = yamux.Client(conn, nil)

return
}
