package model

type Alert struct {
	Body struct {
		AlertId string   `json:"alertId" binding:"required"`
		Message string   `json:"message" binding:"required"`
		Tags    []string `json:"tags" binding:"required"`
	} `json:"alert"`
}
