package controller

import (
	"backend/model"
	"backend/utils"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	ImageDirectory = "assert/image/"
	DirectoryPerm  = 0755
)

func RegisterTorrentFunc(e *echo.Echo) {
	e.POST("/torrent", addTorrentInfoHandle, middleware.JWTWithConfig(utils.JwtConfig))

	e.GET("/torrent", searchTorrentInfoHandle, middleware.JWTWithConfig(utils.JwtConfig))
}

// POST body:
//{
//	"magnet": "",
//	"picture": file1,
//	"name": "",
//	"description": ""
//}
//
// just save filename in db. config Echo to set static resource map.
func addTorrentInfoHandle(context echo.Context) error {

	torrentInfo := new(model.TorrentInfo)
	if err := context.Bind(torrentInfo); err != nil {
		return context.JSON(http.StatusBadRequest, errorReturnMsg(http.StatusBadRequest, err))
	}

	torrentInfo.InsertTime = time.Now().Unix()

	// save picture
	file, err := context.FormFile("picture")
	// it means no pic upload. Set it as default
	if file == nil || err != nil {
		torrentInfo.PicturePath = append(torrentInfo.PicturePath, "default.png")
	} else {
		filename, err := saveImage(file)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, errorReturnMsg(http.StatusInternalServerError, err))
		}
		torrentInfo.PicturePath = append(torrentInfo.PicturePath, filename)
	}

	err = model.InsertTorrentInfo(*torrentInfo)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errorReturnMsg(http.StatusInternalServerError, err))
	}

	return context.JSON(http.StatusOK, torrentInsertReturnMsg())
}

// GET param:
// insertStartTime: int64
// number: int
// name: string
func searchTorrentInfoHandle(context echo.Context) error {
	insertStartTimeStr := context.QueryParam("insertStartTime")
	name := context.QueryParam("name")
	numberStr := context.QueryParam("number")

	filter := model.DefaultTorrentQueryFilter
	options := model.DefaultTorrentQueryOptions

	if insertStartTimeStr != "" {
		insertStartTime, err := strconv.ParseInt(insertStartTimeStr, 10, 64)
		if err != nil {
			return context.JSON(http.StatusBadRequest, errorReturnMsg(http.StatusBadRequest, err))
		}
		filter = bson.D{
			{"insert_time", bson.D{
				{"$gte", insertStartTime},
			}},
		}
	}

	if name != "" {
		name = strings.ToLower(name)
		filter = append(filter, bson.E{Key: "name", Value: name})
	}

	if numberStr != "" {
		number, err := strconv.ParseInt(numberStr, 10, 64)
		if err != nil {
			return context.JSON(http.StatusBadRequest, errorReturnMsg(http.StatusBadRequest, err))
		}
		options = options.SetLimit(number)
	}

	count, err := model.CountTorrentInfo(filter, model.DefaultTorrentCountOptions)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, errorReturnMsg(http.StatusInternalServerError, err))
	}

	infos, err := model.FindTorrentInfo(filter, options)
	if err != nil || infos == nil {
		return context.JSON(http.StatusNotFound, errorReturnMsg(http.StatusNotFound, errors.New("no matched documents")))
	}

	return context.JSON(http.StatusOK, TorrentQueryReturnMsg(count, infos))
}

func saveImage(file *multipart.FileHeader) (string, error) {

	// check the postfix to guarantee that the uploaded file is an image
	matched, err := regexp.MatchString(`.*\.(png|jpg|jpeg|gif|bmp)`, file.Filename)
	if matched == false || err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	file.Filename = strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + file.Filename

	// check the assert directory, if it doesn't exist, make it
	if _, err := os.Open(ImageDirectory); os.IsNotExist(err) {
		_ = os.MkdirAll(ImageDirectory, DirectoryPerm)
	}

	des, err := os.Create(ImageDirectory + file.Filename)
	if err != nil {
		return "", err
	}

	defer des.Close()

	if _, err = io.Copy(des, src); err != nil {
		return "", err
	}

	return file.Filename, nil
}
