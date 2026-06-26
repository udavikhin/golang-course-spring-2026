package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       uint `Gorm:"primaryKey"`
	Email    string
	Password string
}

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}

	seedUsers(db)
	listUsers(db)

	var email, password string
	fmt.Print("email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Print("password: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		log.Println(err)
		return
	}

	if err := login(db, email, password); err != nil {
		fmt.Println("invalid email or password")
		return
	}

	fmt.Printf("logged in as %s", email)
}

func seedUsers(db *gorm.DB) {
	users := []User{
		{
			Email:    "user@mail.com",
			Password: encryptPassword("upass"),
		},
		{
			Email:    "admin@mail.com",
			Password: encryptPassword("apass"),
		},
	}

	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			log.Println(err)
		}
	}
}

func login(db *gorm.DB, email, password string) error {
	var user User
	res := db.First(&user, "email = ?", email)
	if res.Error != nil {
		log.Println(res.Error)
		return res.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func listUsers(db *gorm.DB) {
	var users []User
	db.Find(&users)

	for _, user := range users {
		log.Printf("%+v\n", user)
	}
}

func encryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(bytes)
}
