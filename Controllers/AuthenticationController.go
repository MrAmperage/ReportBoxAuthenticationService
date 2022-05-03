package Controllers

import (
	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/streadway/amqp"
)

func Authentication(Message amqp.Delivery) (Data any, Error error) {

	return WebCore.AuthenticationData{AuthenticationToken: "123"}, nil
}
