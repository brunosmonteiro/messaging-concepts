package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-stomp/stomp"
	"log"
	"messaging-concepts/models"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running the application: %s", err.Error())
	}
}

func run() error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	defer conn.Disconnect()

	planeSubscription, hotelSubscription, carSubscription, err := getPublishers(conn)
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
	planeBooking := models.PlaneBooking{
		FlightNumber: "AB123",
		Departure:    time.Now(),
		Arrival:      time.Now().Add(2 * time.Hour),
	}
	if err := publishBookingMessage(conn, Plane, planeBooking); err != nil {
		return err
	}

	hotelBooking := models.HotelBooking{
		HotelName: "Grand Hotel",
		CheckIn:   time.Now().Add(24 * time.Hour),
		CheckOut:  time.Now().Add(48 * time.Hour),
	}
	if err := publishBookingMessage(conn, Hotel, hotelBooking); err != nil {
		return err
	}

	carBooking := models.CarBooking{
		VehicleModel:   "Tesla Model S",
		PickupLocation: "Downtown Garage",
	}
	if err := publishBookingMessage(conn, Car, carBooking); err != nil {
		return err
	}

	return nil
}
