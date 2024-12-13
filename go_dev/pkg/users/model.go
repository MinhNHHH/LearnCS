package users

import (
	"fmt"
	"time"

	"github.com/MinhNHHH/go_dev/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID           uint      `gorm:"primary_key"`
	Username     string    `gorm:"column:username"`
	LastName     string    `gorm:"column:last_name"`
	PasswordHash string    `gorm:"column:password;not null"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return fmt.Errorf("password should be not empty")
	}

	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func SaveOne(data interface{}) error {
	database := db.GetDB()

	err := database.Create(data)
	if err != nil {
		return fmt.Errorf("insert error %s\n", err.Error)
	}

	return nil
}

func UpdateRecord(data interface{}) error {
	return nil
}
