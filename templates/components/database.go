package components

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost" // change this to "db" before running compose.yaml
	port     = 5432
	user     = "myUser"
	password = "myPassword"
	dbname   = "myDatabase"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:50"`
	Email string `gorm:"size:50"`
}

var db *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("failed to migrate the schema: %w", err)
	}

	return nil
}

func CreateUser(name, email string) error {
	user := User{Name: name, Email: email}
	result := db.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("failed to insert a new user: %w", result.Error)
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", result.Error)
	}
	return users, nil
}

func GetUserByID(id uint) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", result.Error)
	}
	return &user, nil
}

func UpdateUser(id uint, name, email string) error {
	result := db.Model(&User{}).Where("id = ?", id).Updates(User{Name: name, Email: email})
	if result.Error != nil {
		return fmt.Errorf("failed to update user: %w", result.Error)
	}
	return nil
}

func DeleteUser(id uint) error {
	result := db.Delete(&User{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}
	return nil
}
