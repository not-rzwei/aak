package task

import (
	"github.com/sirupsen/logrus"
	"rzwei.net/aak/pkg/opsgenie"
)

func StartAck(id string) {
	user := opsgenie.GetRandomOnCall()

	logrus.Info("ACK: ", id, " || User: ", user)

	opsgenie.PostAcknowledge(id, user)
}
