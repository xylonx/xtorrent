package controller

import (
	"backend/env"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func RegisterCtrlFunc(e *echo.Echo) {
	e.PUT("/ctrl/config", updateConfigHandle)
}

var updateLock = sync.Mutex{}

func updateConfigHandle(context echo.Context) error {

	updateLock.Lock()
	defer updateLock.Unlock()

	conf := new(env.Config)

	file, err := os.Open("env/config.yaml")
	if err != nil {
		log.Println(err)
		return err
	}

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return err
	}

	err = yaml.Unmarshal(byteStream, conf)
	if err != nil {
		log.Println(err)
		return err
	}

	conf.Jwt.ExpireDuration = conf.Jwt.ExpireDuration * int64(time.Hour)

	env.Conf = conf

	return context.String(http.StatusOK, "ok")
}
