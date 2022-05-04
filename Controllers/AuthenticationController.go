package Controllers

import (
	"fmt"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/streadway/amqp"
)

func Authentication(Message amqp.Delivery, ORM []ORMModule.ORMInterface) (Data any, Error error) {
	fmt.Println(len(ORM))

	return WebCore.AuthenticationData{AuthenticationToken: "123"}, nil
}
