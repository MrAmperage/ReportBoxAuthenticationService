package main

import (
	"fmt"
	"os"

	"github.com/MrAmperage/GoWebStruct/ApplicationCore"
	"github.com/MrAmperage/ReportBoxAuthenticationService/Controllers"
	"github.com/MrAmperage/ReportBoxAuthenticationService/ORM"
)

func main() {

	AuthenticationService := &ApplicationCore.ApplicationCore{}
	ErrorInitService := AuthenticationService.Init()
	if ErrorInitService != nil {
		fmt.Println(ErrorInitService)
		os.Exit(0)
	}
	UserORM := &ORM.UserORM{}
	UserORM.SetName("UserORM")
	ReportBoxDatabase, Error := AuthenticationService.WebCore.PostgreSQL.FindByName("ReportBoxDatabase")
	if Error != nil {
		fmt.Println(Error)
	}

	ReportBoxDatabase.ORMs.Add(UserORM)
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
	Subscribe.MessageEmmiter.Handler("Authentication", Controllers.Authentication).Method("POST")

	Subscribe.MessageProcessing(&ReportBoxDatabase.ORMs)

	ErrorWebServer := AuthenticationService.StartWebServer()
	if ErrorInitService != nil {

		fmt.Println(ErrorWebServer)
		os.Exit(0)
	}
}
