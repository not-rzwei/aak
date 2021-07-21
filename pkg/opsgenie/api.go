package opsgenie

import (
	"github.com/go-resty/resty/v2"
	"rzwei.net/aak/pkg/config"
)

func apiClient() *resty.Request {
	client := resty.New()

	client.SetHostURL("https://api.opsgenie.com/v2")

	authToken := "GenieKey " + config.GetConfig().ApiKey

	req := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", authToken)

	return req
}
