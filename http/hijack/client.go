package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
)

func main()  {
	stop := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Kill, os.Interrupt)

	stdin := bufio.NewReader(os.Stdin)
	stdout := os.Stdout

	logger := log.New(stdout, "", log.Lshortfile)

	rawurl := "http://127.0.0.1:8080/hijack"

	u , err := url.Parse(rawurl)

	if nil != err {
		logger.Fatal(err)
	}

	conn, err := net.Dial("tcp", u.Host)
	if nil != err {
		logger.Fatal(err)
	}

	defer func() {
		conn.Close()
		logger.Println("connection closed")
	}()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		logger.Fatal(err)
	}

	if err := req.Write(conn) ; nil != err {
		logger.Fatal(err)
	}

	go func() {
		<- sigCh
		stop <- struct {}{}
	}()

	go func(ogger *log.Logger) {
		_, err := io.Copy(logger.Writer(), conn)
		if nil != err {
			logger.Printf("remote -> stdout: err=%s\n", err.Error())
			stop <- struct{}{}
		}
	}(logger)

	go func(reader io.Reader, logger *log.Logger) {
		_, err := io.Copy(conn, reader)
		if nil != err {
			logger.Printf("stdin -> remote: err=%s\n", err.Error())
			stop <- struct{}{}
		}
	}(stdin, logger)

	logger.Println("wait for stop channel")

	<- stop
}
