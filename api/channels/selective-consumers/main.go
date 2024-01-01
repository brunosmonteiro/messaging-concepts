package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-stomp/stomp"
	"log"
	"messaging-concepts/models"
	"messaging-concepts/utils"
)

const (
	TopicName = "booking-topic"
)

const (
	Hotel         = "hotel"
	Car           = "car"
	Plane         = "plane"
	Selector      = "selector"
	ModelSelector = "model='%s'"
	Model         = "model"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running the application: %s", err.Error())
	}
}

func run() error {
	conn, err := utils.GetConnection(utils.Base)
	if err != nil {
		return err
	}
	defer conn.Disconnect()

	planeSubscription, hotelSubscription, carSubscription, err := getSubscriptions(conn)
	if err != nil {
		return err
	}

	go func() {
		if err = consumePlaneBooking(planeSubscription); err != nil {
			log.Fatalf("error consuming plane booking")
		}
	}()

	go func() {
		if err = consumeHotelBooking(hotelSubscription); err != nil {
			log.Fatalf("error consuming hotel booking")
		}
	}()

	go func() {
		if err = consumeCarBooking(carSubscription); err != nil {
			log.Fatalf("error consuming car booking")
		}
	}()

	if err = sendMockedMessages(conn); err != nil {
		return err
	}

	select {}
}

func consumePlaneBooking(subscription *stomp.Subscription) error {
	for {
		msg, err := subscription.Read()
		if err != nil {
			return fmt.Errorf("error reading plane booking message: %v", err)
		}
		var messageBooking models.BookingMessage[models.PlaneBooking]
		if err := json.Unmarshal(msg.Body, &messageBooking); err != nil {
			return fmt.Errorf("error unmarshaling plane booking: %v", err)
		}
		fmt.Printf("Received a plane booking message! Model: %s, Content: %s\n", messageBooking.Model, messageBooking.Content)
	}
}

func consumeHotelBooking(subscription *stomp.Subscription) error {
	for {
		msg, err := subscription.Read()
		if err != nil {
			return fmt.Errorf("error reading hotel booking message: %v", err)
		}
		var messageBooking models.BookingMessage[models.HotelBooking]
		if err := json.Unmarshal(msg.Body, &messageBooking); err != nil {
			return fmt.Errorf("error unmarshaling hotel booking: %v", err)
		}
		fmt.Printf("Received a hotel booking message! Model: %s, Content: %s\n", messageBooking.Model, messageBooking.Content)
	}
}

func consumeCarBooking(subscription *stomp.Subscription) error {
	for {
		msg, err := subscription.Read()
		if err != nil {
			return fmt.Errorf("error reading car booking message: %v", err)
		}
		var messageBooking models.BookingMessage[models.CarBooking]
		if err := json.Unmarshal(msg.Body, &messageBooking); err != nil {
			return fmt.Errorf("error unmarshaling car booking: %v", err)
		}
		fmt.Printf("Received a car booking message! Model: %s, Content: %s\n", messageBooking.Model, messageBooking.Content)
	}
}

func sendMockedMessages(conn *stomp.Conn) error {
	if err := publishBookingMessage(conn, Plane, models.GenericPlaneBookingInstance()); err != nil {
		return err
	}

	if err := publishBookingMessage(conn, Hotel, models.GenericHotelBookingInstance()); err != nil {
		return err
	}

	if err := publishBookingMessage(conn, Car, models.GenericCarBookingInstance()); err != nil {
		return err
	}

	return nil
}

func getSubscriptions(conn *stomp.Conn) (plane *stomp.Subscription, hotel *stomp.Subscription, car *stomp.Subscription, e error) {
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

func publishBookingMessage[T any](conn *stomp.Conn, bookingModel string, content T) error {
	bookingMsg := models.BookingMessage[T]{
		Model:   bookingModel,
		Content: content,
	}
	return utils.PublishMessage(conn, TopicName, bookingMsg, map[string]string{Model: bookingModel})
}
