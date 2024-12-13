//go:build integration
// +build integration

package sms_test

import (
	"os"

	"github.com/medtrib/qiniu-sdk/v7/auth"

	"github.com/medtrib/qiniu-sdk/v7/sms"
)

var manager *sms.Manager

func init() {
	accessKey := os.Getenv("accessKey")
	secretKey := os.Getenv("secretKey")

	mac := auth.New(accessKey, secretKey)
	manager = sms.NewManager(mac)
}
