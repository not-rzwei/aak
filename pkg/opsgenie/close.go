package opsgenie

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func CloseAlert(id string, user string) {
	url := fmt.Sprintf("alerts/%s/close", id)
	body := fmt.Sprintf(`{"user":"%s"}`, user)
	resp, httpError := apiClient().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	if httpError != nil {
		logrus.Error("CLS Failed: ", httpError)
		return
	}

	if resp.StatusCode() >= 300 {
		logrus.Error("CLS Failed: ", resp)
		return
	}
}
