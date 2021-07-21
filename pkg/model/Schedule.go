package model

type Schedule struct {
	Data struct {
		Users []string `json:"onCallRecipients" binding:"required"`
	} `json:"data"`
}
