package env

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Config struct {
	Mongo struct {
		Address                    string `yaml:"address"`
		DBName                     string `yaml:"dbName"`
		CollectionNameTorrent      string `yaml:"collectionNameTorrent"`
		CollectionNameUser         string `yaml:"collectionNameUser"`
		CollectionNameRegisterCode string `yaml:"collectionNameRegisterCode"`
	} `yaml:"mongo"`

	Jwt struct {
		Secret         string `yaml:"secret"`
		ExpireDuration int64  `yaml:"expireDuration"`
	} `yaml:"jwt"`
}

var Conf *Config

func init() {

	Conf = new(Config)

	file, err := os.Open("env/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(byteStream, Conf)
	if err != nil {
		log.Fatal(err)
	}

	Conf.Jwt.ExpireDuration = Conf.Jwt.ExpireDuration * int64(time.Hour)
}
