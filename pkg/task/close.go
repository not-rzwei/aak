package task

import (
	"time"

	"github.com/sirupsen/logrus"
	"rzwei.net/aak/pkg/opsgenie"
)

func DelayedClose(id string, delay time.Duration) {
	logrus.Info("Start delayed CLS: ", id)

	time.Sleep(delay)

	user := opsgenie.GetRandomOnCall()
	logrus.Info("CLS: ", id, " || User: ", user)

	opsgenie.CloseAlert(id, user)
}
