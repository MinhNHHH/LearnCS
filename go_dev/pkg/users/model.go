package users

import (
	"fmt"
	"time"

	"github.com/MinhNHHH/go_dev/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID           uint      `gorm:"primary_key"`
	Username     string    `gorm:"column:username"`
	LastName     string    `gorm:"column:last_name"`
	PasswordHash string    `gorm:"column:password;not null"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (u *Users) setPassword(password string) error {
	if len(password) == 0 {
		return fmt.Errorf("password should be not empty")
	}

	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *Users) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func GetAllUsers(condition interface{}) (Users, error) {
	database := db.GetDB()
	model := Users{}
	err := database.Where(condition).Find(&model).Error
	return model, err
}

func SaveOne(data interface{}) error {
	database := db.GetDB()

	err := database.Create(data)
	if err != nil {
		return fmt.Errorf("insert error %s\n", err.Error)
	}

	return nil
}

func UpdateRecord(data interface{}, condition interface{}) error {
	database := db.GetDB()
	err := database.Model(&Users{}).Where(condition).Updates(data).Error
	if err != nil {
		return fmt.Errorf("update error: %v", err)
	}
	return nil
}

func DeleteRecord(condition interface{}) error {
	database := db.GetDB()
	model := Users{}
	err := database.Delete(&model, condition).Error
	if err != nil {
		return fmt.Errorf("delete error: %v", err)
	}
	return nil
}

func CreateRecord(data interface{}) error {
	database := db.GetDB()
	err := database.Model(&Users{}).Create(data).Error
	if err != nil {
		return fmt.Errorf("insert error %s\n", err)
	}
	return nil
}
