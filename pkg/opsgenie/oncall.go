package opsgenie

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"rzwei.net/aak/pkg/config"
	"rzwei.net/aak/pkg/model"
)

func GetOnCalls() (schedule model.Schedule, err string) {
	url := fmt.Sprintf("schedules/%s/on-calls?flat=true", config.GetConfig().ScheduleId)
	resp, httpError := apiClient().Get(url)

	if httpError != nil {
		logrus.Error("Failed to get request: ", httpError)
		return
	}

	if resp.StatusCode() >= 300 {
		err = resp.String()
		return
	}

	json.Unmarshal(resp.Body(), &schedule)
	return
}

func GetRandomOnCall() string {
	rand.Seed(time.Now().Unix())

	onCalls, _ := GetOnCalls()
	users := onCalls.Data.Users
	randomUserEmail := users[rand.Intn(len(users))]
	randomUser := strings.Split(randomUserEmail, "@")[0]

	return randomUser
}
