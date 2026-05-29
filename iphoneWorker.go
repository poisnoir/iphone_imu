package main

import (
	"context"
	"encoding/json"
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

	address := "172.20.10.3:" + strconv.Itoa(port)
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
		sig:  make(chan struct{}, 1),
		addr: conn.LocalAddr().String(),
	}
	go worker.run()
	return worker, nil
}

func (iw *IphoneWorker) run() {

	fmt.Println("meow")

	defer iw.conn.Close()
	buffer := make([]byte, 4092)

	for {
		n, _, err := iw.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading packet: %v\n", err)
			continue // Keep running even if one packet fails
		}

		json.Unmarshal(buffer[:n], &iw.data)

		iw.sig <- struct{}{}
		// Convert the payload bytes to
		// a string
		// fmt.Println(iw.data)
	}
}

func (iw *IphoneWorker) GetData(ctx context.Context) IphoneOutput {
	select {
	case <-iw.sig:
		return iw.data
	case <-ctx.Done():
		return iw.data
	}
}

func (iw *IphoneWorker) Address() string {
	return iw.addr
}
