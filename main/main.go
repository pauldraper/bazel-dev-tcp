package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
	status := run()
	os.Exit(status)
}

func run() int {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: ibazel-tcp <listen_address> <target_address>")
		return 1
	}
	listenAddr := os.Args[1]
	targetAddr := os.Args[2]

	profilePath := os.Getenv("IBAZEL_PROFILE_FILE")
	if profilePath == "" {
		fmt.Fprintln(os.Stderr, "IBAZEL_PROFILE_FILE environment variable is empty")
		return 1
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	ls, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer ls.Close()

	eventsFile, err := os.Open(profilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer eventsFile.Close()

	eventsReader := bufio.NewReader(eventsFile)

	buildStatus := BuildStatusBuilding

	msgs := make(chan struct{})
	go StdinMessage(msgs)

	conns := make(chan *net.TCPConn)
	go HandleListener(ls, conns)

	for {
		select {
		case <-msgs:
			for {
				buildStatus, err = UpdateBuildStatus(eventsReader, buildStatus)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return 1
				}

				if buildStatus != BuildStatusBuilding {
					break
				}
				<-msgs
			}
		case conn := <-conns:
			buildStatus, err = UpdateBuildStatus(eventsReader, buildStatus)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 1
			}

			if buildStatus == BuildStatusFailed {
				conn.Close()
			} else {
				go HandleConnection(conn, targetAddr)
			}
		}
	}

	return 0
}
