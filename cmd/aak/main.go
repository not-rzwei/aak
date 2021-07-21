package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rzwei.net/aak/pkg/config"
	"rzwei.net/aak/pkg/model"
	"rzwei.net/aak/pkg/task"
	"rzwei.net/aak/pkg/util"
)

func main() {
	r := gin.New()
	config := &config.Config{}

	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})

	err := env.Parse(config)
	if err != nil {
		logrus.Error("Unable to parse environment variables: ", err)
		return
	}

	r.POST("/", func(c *gin.Context) {
		var alert model.Alert

		if err := c.ShouldBindJSON(&alert); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if c.GetHeader("X-AAK-Key") != config.SecretKey {
			message := "Wrong authentication key"
			c.JSON(http.StatusBadRequest, gin.H{"error": message})
			logrus.Error(message)
			return
		}

		isAutoClose := util.HaveTags(config.CloseTags, alert.Body.Tags)

		go task.StartAck(alert.Body.AlertId)

		if isAutoClose {
			go task.DelayedClose(alert.Body.AlertId, time.Duration(config.CloseDelay)*time.Minute)
		}

		c.JSON(200, gin.H{
			"alert":     alert.Body.Message,
			"tags":      strings.Join(alert.Body.Tags, ","),
			"autoClose": isAutoClose,
			"message":   "Alert processed",
		})
	})

	r.Run()
}
