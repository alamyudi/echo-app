package controllers

import (
	"github.com/alamyudi/echo-app/echokit"
	"github.com/alamyudi/echo-app/mobilerestkit/helpers"
	"github.com/sirupsen/logrus"
)

// Request type
type (
	// LoginRequest for login request
	LoginRequest struct {
		Email      string `validate:"required,email" json:"email"`
		Password   string `validate:"required" json:"password"`
		MacAddress string `validate:"required" json:"mac_address"`
	}
)

type (
	// DownloadLog to post metadata after download
	DownloadLog struct {
		RevNum     string `validate:"required" json:"rev_num"`
		LastDate   string `validate:"required" json:"last_date"`
		LastEdit   string `validate:"required" json:"last_edit"`
		LastTime   string `validate:"required" json:"last_time"`
		MacAddress string `validate:"required" json:"mac_address"`
	}
)

// response type
type (
	// MessageResponse for response message
	MessageResponse struct {
		Title   string `json:"title"`
		Message string `json:"message"`
	}

	// MessageWithPayloadResponse for response message with payload
	MessageWithPayloadResponse struct {
		Title   string      `json:"title"`
		Message string      `json:"message"`
		Payload interface{} `json:"payload"`
	}
)

func init() {
	logrus.Info("Init Controller")
}

var conf helpers.Config
var iKit *echokit.EchoKit

// SetIAPManager to pass echokit and configs
func SetIAPManager(nk *echokit.EchoKit, cnf helpers.Config) {
	conf = cnf
	iKit = nk
}
