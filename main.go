package main

import (
	"fmt"
	"os"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
	"github.com/streadway/amqp"
)

func main() {

	AuthenticationService := &ApplicationCore.ApplicationCore{}
	ErrorInitService := AuthenticationService.Init()
	if ErrorInitService != nil {
		fmt.Println(ErrorInitService)
		os.Exit(0)
	}
	ErrorDatabaseConnection := AuthenticationService.StartDatabaseConnections()
	if ErrorDatabaseConnection != nil {

		fmt.Println(ErrorDatabaseConnection)
		os.Exit(0)
	}
	ErrorRabbitMQ := AuthenticationService.StartRabbitMQ()
	if ErrorRabbitMQ != nil {

		fmt.Println(ErrorRabbitMQ)
		os.Exit(0)
	}
	Subscribe, Error := AuthenticationService.WebCore.RabbitMQ.RabbitMQChanel.GetSubscribeByQueueName("AuthenticationQueue")
	if Error != nil {
		fmt.Println(Error)
	}
	Subscribe.MessageProcessing(func(RabbitMQMessage amqp.Delivery) {
		fmt.Println(string(RabbitMQMessage.Body))
		ErrorPublish := Subscribe.ChanelLink.Publish("", RabbitMQMessage.ReplyTo, false, false, amqp.Publishing{
			CorrelationId: RabbitMQMessage.CorrelationId,
			Body:          []byte("Ответное сообщение"),
		})
		if ErrorPublish != nil {
			fmt.Println(ErrorPublish)
		}

	})

	ErrorWebServer := AuthenticationService.StartWebServer()
	if ErrorInitService != nil {

		fmt.Println(ErrorWebServer)
		os.Exit(0)
	}
}
