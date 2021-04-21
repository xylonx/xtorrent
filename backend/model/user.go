package model

import (
	"backend/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type User struct {
	EmailAddress string `json:"email_address" bson:"email_address" form:"email_address"`
	Password     string `json:"password" bson:"password" form:"password"`
}

func InsertUser(user User) error {
	hashedStr, err := utils.SecurityHashing(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedStr

	_, err = colUser.InsertOne(context.TODO(), user)
	return err
}

func GetUserByEmailAddress(email string) (User, error) {
	var user User
	err := colUser.FindOne(context.TODO(), bson.M{
		"email_address": email,
	}).Decode(&user)

	if err != nil {
		log.Println(err)
		return User{}, err
	}

	return user, nil
}
