package controller

import (
	"backend/model"
	"fmt"
	"net/http"
)

func errorReturnMsg(code int, err error) interface{} {
	return struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    code,
		Message: fmt.Sprintf("%v", err),
	}
}

func logReturnMsg(token string) interface{} {
	return struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Code:    http.StatusOK,
		Message: "success",
		Token:   token,
	}
}

func torrentInsertReturnMsg() interface{} {
	return struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    http.StatusOK,
		Message: "success",
	}
}

func TorrentQueryReturnMsg(count int64, infos []model.TorrentInfo) interface{} {
	return struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Number       int64               `json:"number"`
			TorrentInfos []model.TorrentInfo `json:"torrent_infos"`
		} `json:"data"`
	}{
		Code:    http.StatusOK,
		Message: "success",
		Data: struct {
			Number       int64               `json:"number"`
			TorrentInfos []model.TorrentInfo `json:"torrent_infos"`
		}{
			Number:       count,
			TorrentInfos: infos,
		},
	}
}
