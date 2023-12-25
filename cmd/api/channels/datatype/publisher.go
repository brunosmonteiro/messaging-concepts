package datatype

import (
	"encoding/json"
	"fmt"
	"github.com/go-stomp/stomp"
	"log"
	"messaging-concepts/cmd/api/models"
	"net"
)

const (
	ServerAddr  = "localhost:61613"
	TopicName   = "booking-topic"
	JsonContent = "application/json"
)

const (
	Hotel         = "hotel"
	Car           = "car"
	Plane         = "plane"
	Selector      = "selector"
	Tcp           = "tcp"
	ModelSelector = "model='%s'"
	Model         = "model"
)

func GetConnection() (*stomp.Conn, error) {
	netConn, err := net.Dial(Tcp, ServerAddr)
	if err != nil {
		return nil, err
	}
	stompConn, err := stomp.Connect(netConn)
	return stompConn, err
}

func GetPublishers(conn *stomp.Conn) (plane *stomp.Subscription, hotel *stomp.Subscription, car *stomp.Subscription, e error) {
	planeSubscription, err := subscribeToBookingType(conn, Plane)
	if err != nil {
		return nil, nil, nil, err
	}

	hotelSubscription, err := subscribeToBookingType(conn, Hotel)
	if err != nil {
		return nil, nil, nil, err
	}

	carSubscription, err := subscribeToBookingType(conn, Car)
	if err != nil {
		return nil, nil, nil, err
	}

	return planeSubscription, hotelSubscription, carSubscription, nil
}

func subscribeToBookingType(conn *stomp.Conn, bookingType string) (*stomp.Subscription, error) {
	selector := fmt.Sprintf(ModelSelector, bookingType) // Selector based on header
	return conn.Subscribe(TopicName, stomp.AckAuto, stomp.SubscribeOpt.Header(Selector, selector))
}

func PublishBookingMessage[T any](conn *stomp.Conn, bookingModel string, content T) error {
	bookingMsg := models.BookingMessage[T]{
		Model:   bookingModel,
		Content: content,
	}

	header := stomp.SendOpt.Header(Model, bookingModel)

	byteData, err := json.Marshal(bookingMsg)
	if err != nil {
		log.Fatalf("Error marshalling to JSON: %v", err)
	}

	err = conn.Send(
		TopicName,
		JsonContent,
		byteData,
		header,
	)
	if err != nil {
		log.Fatalf("failed to send to server: %v", err)
	}

	return nil
}
