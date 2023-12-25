This sample project explores the concept of selective consumers within a canonical data model.

## Context
We work with different types of bookings: planes, hotels and cars. But we don't want to create different
channels for each one of them. Instead of three different datatype channels, we will work with only
one destination, the topic `booking-topic`, and three different subscribers.

## Canonical Data Model
To make things simpler and more manageable, we will work with a canonical data model, that is a specific
data pattern across all booking models. It is simply:
```
type BookingMessage[T any] struct {
	Model   string `json:"model"`
	Content T      `json:"content"`
}
```
The content can either refer to a plane, hotel or car. The type indicates which one. This is informed
by the publisher.

## Selective Consumers
There will be three different subscribers to the topic: plane-subscriber, hotel-subscriber and
car-subscriber.

We are going to use ActiveMQ as our message broker to benefit from its Selectors feature. It basically
subscribes a consumer with a predefined filter, based on an agreed header. In our case, the publisher
will post each message with a header `model: "plane" OR "hotel" OR "car"`: plane, hotel or car.
The subscribers are registered with the selector `model: "plane" OR "hotel" OR "car"`,

## How to Run
% Create a Makefile
% Create a Diagram
