package main

import (
	"context"
	"fmt"
	"net"
	"strconv"
)

type IphoneWorker struct {
	conn net.UDPConn
	addr string
	data IphoneOutput
	sig  chan struct{}
}

func NewIphoneWorker(port int) (*IphoneWorker, error) {

	address := "127.0.0.1:" + strconv.Itoa(port)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, fmt.Errorf("Error resolving address: %v\n", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, fmt.Errorf("Error listening on %s: %v\n", address, err)
	}

	worker := &IphoneWorker{
		conn: *conn,
		addr: conn.LocalAddr().String(),
	}
	go worker.run()
	return nil, nil
}

func (iw *IphoneWorker) run() {

	defer iw.conn.Close()
	buffer := make([]byte, 4092)

	for {
		n, remoteAddr, err := iw.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading packet: %v\n", err)
			continue // Keep running even if one packet fails
		}

		// Convert the payload bytes to a string
		message := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", remoteAddr.String(), message)
	}

}

func (iw *IphoneWorker) GetData(ctx context.Context) IphoneOutput {
	select {
	case <-iw.sig:
	case <-ctx.Done():
	}

	return iw.data
}

func (iw *IphoneWorker) Address() string {
	return iw.addr
}
