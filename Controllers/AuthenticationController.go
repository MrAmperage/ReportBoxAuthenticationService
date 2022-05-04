package Controllers

import (
	"fmt"

	"github.com/MrAmperage/GoWebStruct/WebCore"
	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
	"github.com/MrAmperage/ReportBoxAuthenticationService/ORM"
	"github.com/streadway/amqp"
)

func Authentication(Message amqp.Delivery, ORMs ORMModule.ORMArray) (Data any, Error error) {
	ORMElement, _ := ORMs.FindByName("UserORM")
	UserORM := ORMElement.(*ORM.UserORM)
	User, Error := UserORM.GetUserByName("front")
	if Error != nil {
		return
	}
	fmt.Println(User.Username)

	return WebCore.AuthenticationData{AuthenticationToken: "123"}, nil
}
