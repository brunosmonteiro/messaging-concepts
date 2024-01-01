package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-stomp/stomp"
	"github.com/go-stomp/stomp/frame"
	"log"
	"net"
)

const (
	Base                 = 61613
	PersistentKahaDb     = 61614
	PersistentPostgreSql = 61615
)

const (
	ServerAddr  = "localhost:%d"
	Tcp         = "tcp"
	JsonContent = "application/json"
)

func GetConnection(port int) (*stomp.Conn, error) {
	netConn, err := net.Dial(Tcp, fmt.Sprintf(ServerAddr, port))
	if err != nil {
		return nil, err
	}
	stompConn, err := stomp.Connect(netConn)
	return stompConn, err
}

func PublishMessage[T any](conn *stomp.Conn, topicName string, msg T, headers map[string]string) error {
	byteData, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error marshalling to JSON: %v", err)
		return err
	}

	var headersSlice []func(*frame.Frame) error
	for key, value := range headers {
		headersSlice = append(headersSlice, stomp.SendOpt.Header(key, value))
	}

	err = conn.Send(
		topicName,
		JsonContent,
		byteData,
		headersSlice...,
	)
	return err
}
