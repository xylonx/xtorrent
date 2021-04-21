package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type RegisterCode struct {
	Code   string `json:"code" bson:"code"`
	Valid  bool   `json:"valid" bson:"valid"`
	UsedBy string `json:"used_by" bson:"used_by"`
}

func InsertRegCode(code string) error {
	registerCode := RegisterCode{
		Code:   code,
		Valid:  true,
		UsedBy: "",
	}

	_, err := colRegisterCode.InsertOne(context.TODO(), registerCode)
	return err
}

func IsRegCodeValid(code string) (bool, error) {
	filter := bson.D{
		{Key: "code", Value: code},
	}

	var result RegisterCode
	err := colRegisterCode.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return result.Valid, nil
}

func ValidateRegCode(code, emailHash string) (bool, error) {
	filter := bson.D{
		{Key: "code", Value: code},
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "used_by", Value: emailHash},
			{Key: "valid", Value: false},
		}},
	}
	result, err := colRegisterCode.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		log.Println(err)
	}

	if result.MatchedCount == 0{
		log.Printf("can't find registerCode %v for user %v\n", code, emailHash)
		return false, nil
	}

	return true, nil
}

