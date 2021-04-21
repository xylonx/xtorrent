package controller

import (
	"backend/model"
	"backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func RegisterTestFunc(e *echo.Echo) {
	e.POST("/test/login", testLoginHandle)
	e.GET("/test/torrents", testTorrenthandle, middleware.JWTWithConfig(utils.JwtConfig))
	e.POST("/test/upload", testHandle)
}

func testLoginHandle(context echo.Context) error {
	token, err := utils.GenerateJwtToken("xylon@xlonx.com")
	if err != nil{
		return err
	}

	return context.JSON(200, logReturnMsg(token))
}

func testTorrenthandle(context echo.Context) error {
	var torrents []model.TorrentInfo

	number := context.QueryParam("number")

	if number ==  "2"{
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
	}else {
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


	return context.JSON(200, TorrentQueryReturnMsg(16, torrents))
}


func testHandle(context echo.Context) error {
	torrentInfo := new(model.TorrentInfo)
	if err := context.Bind(torrentInfo); err != nil {
		return err
	}

	torrentInfo.InsertTime = time.Now().Unix()

	file, err := context.FormFile("picture")

	// it means no pic upload. Set it as default
	if err != nil || file == nil {
		torrentInfo.PicturePath = append(torrentInfo.PicturePath, "assert/default.png")
	} else {
		// check the postfix to guarantee that the uploaded file is an image
		matched, err := regexp.MatchString(`.*\.(png|jpg|jpeg|gif|bmp)`, file.Filename)
		if matched == false || err != nil {
			return context.JSON(http.StatusBadRequest, "not matched")
		}

		src, err := file.Open()
		if err != nil {
			return context.JSON(http.StatusInternalServerError, "")
		}
		defer src.Close()

		file.Filename = strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + file.Filename

		// check the assert directory, if it doesn't exist, make it
		if _, err := os.Open("assert/image/"); os.IsNotExist(err) {
			err = os.MkdirAll("assert/image/", 0755)
		}

		des, err := os.Create("assert/image/" + file.Filename)
		if err != nil {
			return err
		}

		defer des.Close()

		if _, err = io.Copy(des, src); err != nil {
			return err
		}
	}

	return context.JSON(200, torrentInsertReturnMsg())

}
