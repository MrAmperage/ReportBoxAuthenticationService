package Controllers

import (
	"errors"

	"github.com/streadway/amqp"
)

func Authentication(Message amqp.Delivery) (Data any, Error error) {

	return nil, errors.New("123")
}
