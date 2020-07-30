package db

import (
	"log"

	"github.com/pbillerot/graduel/types"
)

//CreateUser will create a new user, take as input the parameters and
//insert it into database
func CreateUser(username, password, email string) error {
	user := types.User{}
	user.Name = username
	user.Password = password
	user.Email = email
	db := openDb()
	err = db.Insert(&user).Do()
	return err
}

var err error

//ValidUser will check if the user exists in db and if exists if the username password
//combination is valid
func ValidUser(username, password string) bool {
	db := openDb()
	user := types.User{}
	err = db.Select(&user).Where("user_name = ?", username).Do()
	if err != nil {
		log.Fatal(err)
	}

	//If the password matches, return true
	if password == user.Password {
		return true
	}
	//by default return false
	return false
}

//GetUserID will get the user's ID from the database
func GetUserID(username string) (int, error) {
	db := openDb()
	user := types.User{}
	err = db.Select(&user).Where("user_name = ?", username).Do()
	return user.ID, err
}
