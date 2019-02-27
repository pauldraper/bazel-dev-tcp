package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func HandleListener(ls *net.TCPListener, ch chan *net.TCPConn) {
	for {
		conn, err := ls.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		ch <- conn
	}
}

func HandleConnection(conn *net.TCPConn, targetAddr string) {
	defer conn.Close()

	var err error
	var targetTcpAddr *net.TCPAddr
	var targetConn *net.TCPConn
	for i := 0; i < 20; i++ {
		targetTcpAddr, err = net.ResolveTCPAddr("tcp", targetAddr)
		if err != nil {
			time.Sleep(200 * time.Millisecond)
			continue
		}
		targetConn, err = net.DialTCP("tcp", nil, targetTcpAddr)
		if err != nil {
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer targetConn.Close()
	end := make(chan struct{})
	targetEnd := make(chan struct{})
	go copy(conn, targetConn, end)
	go copy(targetConn, conn, targetEnd)
	<-end
	<-targetEnd
}

func copy(src *net.TCPConn, dst *net.TCPConn, end chan struct{}) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	dst.CloseWrite()
	src.CloseRead()
	end <- struct{}{}
}
