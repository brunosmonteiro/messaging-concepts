package main

import (
	"github.com/go-stomp/stomp"
	"log"
	"messaging-concepts/models"
	"messaging-concepts/utils"
)

const (
	TopicName  = "booking-topic"
	Persistent = "persistent"
	True       = "true"
)

const (
	Hotel = "hotel"
	Plane = "plane"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running the application: %s", err.Error())
	}
}

func run() error {
	connKahaDb, err := utils.GetConnection(utils.PersistentKahaDb)
	if err != nil {
		return err
	}
	defer connKahaDb.Disconnect()

	connPostgreSql, err := utils.GetConnection(utils.PersistentPostgreSql)
	if err != nil {
		return err
	}
	defer connPostgreSql.Disconnect()

	if err = sendMockedMessages(connKahaDb, connPostgreSql); err != nil {
		return err
	}

	return nil
}

func sendMockedMessages(connKahaDb *stomp.Conn, connPostgreSql *stomp.Conn) error {
	if err := publishBookingMessage(connKahaDb, Plane, models.GenericPlaneBookingInstance()); err != nil {
		return err
	}

	if err := publishBookingMessage(connPostgreSql, Hotel, models.GenericHotelBookingInstance()); err != nil {
		return err
	}

	return nil
}

func publishBookingMessage[T any](conn *stomp.Conn, bookingModel string, content T) error {
	bookingMsg := models.BookingMessage[T]{
		Model:   bookingModel,
		Content: content,
	}
	return utils.PublishMessage(conn, TopicName, bookingMsg, map[string]string{Persistent: True})
}
