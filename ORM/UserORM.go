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
type UserORM struct{}

func (UserORM *UserORM) GetUsers(ConnectionPoolLink *gorm.DB) (Users []User) {
	ConnectionPoolLink.Find(&Users)
	return Users

}
func (UserORM *UserORM) GetUserByName(ConnectionPoolLink *gorm.DB, Username string) (UserData User, Error error) {
	Result := ConnectionPoolLink.Where(&User{Username: Username}).Take(&UserData)
	if Result.Error != nil {
		return UserData, errors.New("неправильное имя пользователя или пароль")
	}
	return
}
