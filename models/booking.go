package models

import (
	"fmt"
	"time"
)

type BookingMessage[T any] struct {
	Model   string `json:"model"`
	Content T      `json:"content"`
}

type PlaneBooking struct {
	FlightNumber string    `json:"flightNumber"`
	Departure    time.Time `json:"departure"`
	Arrival      time.Time `json:"arrival"`
}

func GenericPlaneBookingInstance() PlaneBooking {
	return PlaneBooking{
		FlightNumber: "AB123",
		Departure:    time.Now(),
		Arrival:      time.Now().Add(2 * time.Hour),
	}
}

type CarBooking struct {
	VehicleModel   string `json:"vehicleModel"`
	PickupLocation string `json:"pickupLocation"`
}

func GenericCarBookingInstance() CarBooking {
	return CarBooking{
		VehicleModel:   "Tesla Model S",
		PickupLocation: "Downtown Garage",
	}
}

type HotelBooking struct {
	HotelName string    `json:"hotelName"`
	CheckIn   time.Time `json:"checkIn"`
	CheckOut  time.Time `json:"checkOut"`
}

func GenericHotelBookingInstance() HotelBooking {
	return HotelBooking{
		HotelName: "Grand Hotel",
		CheckIn:   time.Now().Add(24 * time.Hour),
		CheckOut:  time.Now().Add(48 * time.Hour),
	}
}

func (pb PlaneBooking) String() string {
	return fmt.Sprintf("FlightNumber: %s, Departure: %s, Arrival: %s",
		pb.FlightNumber,
		pb.Departure.Format("2006-01-02 15:04:05"),
		pb.Arrival.Format("2006-01-02 15:04:05"))
}

func (cb CarBooking) String() string {
	return fmt.Sprintf("VehicleModel: %s, PickupLocation: %s",
		cb.VehicleModel,
		cb.PickupLocation)
}

func (hb HotelBooking) String() string {
	return fmt.Sprintf("HotelName: %s, CheckIn: %s, CheckOut: %s",
		hb.HotelName,
		hb.CheckIn.Format("2006-01-02 15:04:05"),
		hb.CheckOut.Format("2006-01-02"))
}
