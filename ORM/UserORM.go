package ORM

import (
	"errors"

	"github.com/MrAmperage/GoWebStruct/WebCore/Modules/ORMModule"
)

type User struct {
	Username string `gorm:"primaryKey;not null;username"`
	Password string `gorm:"not null;password"`
	Enabled  bool   `gorm:"not null;enabled"`
}
type UserORM struct {
	ORMModule.ORM
}

func (UserORM *UserORM) GetUsers() (Users []User) {
	UserORM.ConnectionLink.Find(&Users)
	return Users

}
func (UserORM *UserORM) GetUserByName(Username string) (UserData User, Error error) {

	Result := UserORM.ConnectionLink.Where(&User{Username: Username}).Take(&UserData)
	if Result.Error != nil {
		return UserData, errors.New("Неправильное имя пользователя или пароль")
	}
	return
}
