package test

import (
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestLogrus(t *testing.T) {
	log.SetReportCaller(true)
	log.WithFields(log.Fields{
		"traceid": "dog",
		"时间":      time.Now().UnixMilli(),
	}).Info("大Tina狗")
}
