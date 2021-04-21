package model

import (
	"backend/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"time"
)

var (
	colTorrentInfo  *mongo.Collection
	colUser         *mongo.Collection
	colRegisterCode *mongo.Collection
)

func init() {

	client, err := mongo.NewClient(options.Client().ApplyURI(env.Conf.Mongo.Address))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	xtorrentDB := client.Database(env.Conf.Mongo.DBName)
	colTorrentInfo = xtorrentDB.Collection(env.Conf.Mongo.CollectionNameTorrent)
	colUser = xtorrentDB.Collection(env.Conf.Mongo.CollectionNameUser)
	colRegisterCode = xtorrentDB.Collection(env.Conf.Mongo.CollectionNameRegisterCode)
}

func transactionTest() {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	client, err := mongo.NewClient(options.Client().ApplyURI(env.Conf.Mongo.Address))
	if err != nil {
		log.Fatalln(err)
	}

	session, err := client.StartSession()
	if err != nil{
		panic(err)
	}

	defer session.EndSession(context.Background())

	callback := func(sessionContext mongo.SessionContext)(interface{}, error) {
		result, err := colUser.InsertOne(sessionContext, bson.M{})
		if err != nil{
			return nil, err
		}
		log.Println(result)
		return result, nil
	}

	_, err = session.WithTransaction(context.Background(), callback, txnOpts)
	log.Println(err)



}