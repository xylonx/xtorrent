package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type TorrentInfo struct {
	Magnet      string   `json:"magnet" bson:"magnet" form:"magnet"`
	PicturePath []string `json:"picture_path" bson:"picture_path" form:"picture_path"`
	Name        string   `json:"name" bson:"name" form:"name"`
	Description string   `json:"description" bson:"description" form:"description"`
	InsertTime  int64    `json:"insert_time" bson:"insert_time" form:"insert_time"`
}

var (
	DefaultTorrentQueryFilter = bson.D{
		{"insert_time", bson.D{
			{"$gte", time.Now().Unix() - int64(time.Hour)*24*30},
		}},
	}

	DefaultTorrentQueryOptions = options.Find().SetSort(bson.D{
		{"insert_time", -1},
	})

	DefaultTorrentCountOptions = options.Count().SetMaxTime(time.Second * 2)
)

func InsertTorrentInfo(info TorrentInfo) error {
	_, err := colTorrentInfo.InsertOne(context.TODO(), info)
	return err
}

func FindTorrentInfo(filter bson.D, opt *options.FindOptions) ([]TorrentInfo, error) {
	cursor, err := colTorrentInfo.Find(context.TODO(), filter, opt)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var torrentInfos []TorrentInfo
	err = cursor.All(context.TODO(), &torrentInfos)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return torrentInfos, nil
}

func CountTorrentInfo(filter bson.D, options *options.CountOptions) (int64, error) {
	return colTorrentInfo.CountDocuments(context.TODO(), filter, options)
}
