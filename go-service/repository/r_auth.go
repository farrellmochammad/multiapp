package repository

import (
	"fmt"
	database "go-service/db"
	model "go-service/model"

	"golang.org/x/crypto/bcrypt"
)

func GetUsers() bool {
	db := database.GetDB()
	selDB, err := db.Query("SELECT Phone,Name,Role FROM Users")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for selDB.Next() {
		var phone string
		var name, password string
		err = selDB.Scan(&phone, &name, &password)
		if err != nil {
			return false
		}
	}

	return true
}

func InsertUser(user model.User) bool {
	db := database.GetDB()
	user.Password = hashPassword([]byte(user.Password))

	sqlStatement := "INSERT INTO Users VALUES(?,?,?,?)"
	_, err := db.Query(sqlStatement, user.Phone, user.Name, user.Role, user.Password)
	defer db.Close()

	if err != nil {
		return false
	} else {
		return true
	}
}

func CheckUserLogin(user model.User) interface{} {
	db := database.GetDB()
	var tag model.User

	fmt.Println(user.Phone)
	fmt.Println(user.Password)
	err := db.QueryRow("SELECT Phone,Password,Name,Role FROM Users where Phone = ?", user.Phone).Scan(&tag.Phone, &tag.Password, &tag.Name, &tag.Role)

	if err != nil {
		return model.User{
			Name: "-1",
		}
	}
	defer db.Close()

	password := user.Password
	hash := tag.Password

	match := checkPasswordHash(password, hash)
	if match {
		return tag
	} else {
		return model.User{
			Name: "-2",
		}
	}

}

func hashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
