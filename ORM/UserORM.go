package ORM

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"primaryKey;not null;username"`
	Password string `gorm:"not null;password"`
	Enabled  bool   `gorm:"not null;enabled"`
}
type UserORM struct {
	Name           string
	ConnectionLink *gorm.DB
}

func (UserORM UserORM) GetName() string {
	return UserORM.Name
}
func (UserORM *UserORM) SetConnection(ConnectionLink *gorm.DB) {
	UserORM.ConnectionLink = ConnectionLink
}
func (UserORM *UserORM) GetUsers() (Users []User) {
	UserORM.ConnectionLink.Find(&Users)
	return Users

}
func (UserORM *UserORM) GetUserByName(Username string) (UserData User, Error error) {

	Result := UserORM.ConnectionLink.Where(&User{Username: Username}).Take(&UserData)
	if Result.Error != nil {
		return UserData, errors.New("неправильное имя пользователя или пароль")
	}
	return
}
