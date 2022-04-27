package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
	"github.com/MrAmperage/ReportBoxAuthenticationService/ORM"
	"github.com/streadway/amqp"
)

func main() {

	AuthenticationService := &ApplicationCore.ApplicationCore{}
	ErrorInitService := AuthenticationService.Init()
	if ErrorInitService != nil {
		fmt.Println(ErrorInitService)
		os.Exit(0)
	}
	ErrorDatabaseConnection := AuthenticationService.WebCore.PostgreSQL.StartDatabaseConnections()
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
		var User ORM.User
		json.Unmarshal(RabbitMQMessage.Body, &User)
		fmt.Println(User.Username)
		PosgreSQL, Error := AuthenticationService.WebCore.PostgreSQL.FindByName("ReportBoxDatabase")
		if Error != nil {

			fmt.Println(Error)
		}
		var Response string
		UserORM := &ORM.UserORM{}
		ResponseUser, Error := UserORM.GetUserByName(PosgreSQL.ConnectionPool, User.Username)
		fmt.Println(ResponseUser.Username)
		Response = ResponseUser.Username
		if Error != nil {

			Response = Error.Error()
		}
		ErrorPublish := Subscribe.ChanelLink.Publish("", RabbitMQMessage.ReplyTo, false, false, amqp.Publishing{
			CorrelationId: RabbitMQMessage.CorrelationId,
			Body:          []byte(Response),
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
