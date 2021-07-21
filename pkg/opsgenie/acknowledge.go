package opsgenie

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func PostAcknowledge(id string, user string) {
	url := fmt.Sprintf("alerts/%s/acknowledge", id)
	body := fmt.Sprintf(`{"user":"%s"}`, user)
	resp, httpError := apiClient().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	if httpError != nil {
		logrus.Error("ACK Failed: ", httpError)
		return
	}

	if resp.StatusCode() >= 300 {
		logrus.Error("ACK Failed: ", resp)
		return
	}
}
