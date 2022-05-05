package Controllers

import (
	"encoding/json"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAuthenticationService/ORM"
	"github.com/streadway/amqp"
)

func Authentication(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {
	var AuthenticationRequest WebCore.AuthenticationRequest
	var AuthenticationResponse WebCore.AuthenticationResponse
	ORMElement, _ := ORMs.FindByName("UserORM")
	UserORM := ORMElement.(*ORM.UserORM)
	json.Unmarshal(Message.Body, &AuthenticationRequest)
	Data, Error = UserORM.GetUserByName(AuthenticationRequest.Username)
	if Error != nil {
		return
	}
	AuthenticationResponse.AuthenticationToken = "Токен авторизации"

	return AuthenticationResponse, Error
}
