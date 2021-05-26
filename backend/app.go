package main

import (
	"backend/controller"
	"backend/env"
	"backend/model"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func main() {
	fmt.Println(env.Conf.Mongo)
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
		},
	}))

	controller.RegisterLogFunc(e)
	controller.RegisterTorrentFunc(e)
	controller.RegisterTestFunc(e)
	controller.RegisterCtrlFunc(e)

	e.Server.Addr = ":5678"

	e.Static("/", "assert/image")

	e.Logger.Fatal(gracehttp.Serve(e.Server))
}

func testTorrenthandle(context echo.Context) error {
	var torrents []model.TorrentInfo

	number := context.QueryParam("number")

	if number == "2" {
		for i := 0; i < 8; i++ {
			torrent1 := model.TorrentInfo{
				Magnet:      "xx",
				PicturePath: []string{"1615100849687681679_xx.png"},
				Name:        "xx",
				Description: "xx",
				InsertTime:  time.Now().Unix(),
			}
			torrents = append(torrents, torrent1)
		}
	} else {
		for i := 0; i < 8; i++ {
			torrent1 := model.TorrentInfo{
				Magnet:      "lock",
				PicturePath: []string{"1615389652670235842_lock.png"},
				Name:        "lock",
				Description: "lock",
				InsertTime:  time.Now().Unix(),
			}
			torrents = append(torrents, torrent1)
		}
	}

	return context.JSON(200, controller.TorrentQueryReturnMsg(16, torrents))
}

func testRegisterHandle(context echo.Context) error {
	postBody := new(model.User)
	if err := context.Bind(postBody); err != nil {
		return err
	}

	err := model.InsertUser(*postBody)
	if err != nil {
		return err
	}

	return context.String(200, "")
}
