package yandex

import (
	"encoding/json"
	"github.com/SkurkovPavel/Homie/internal/storage"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)


func SendTvCommand(w http.ResponseWriter, logger *logrus.Logger) {

	url := "https://api.iot.yandex.net/v1.0/scenarios/%s/actions"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + "token"

	// Create a new request using http
	req, err := http.NewRequest("POST", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Println("Error on response.\n[ERROR] -", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Println("Error while reading the response bytes:", err)
	}

	var status storage.RequestStatus
	if err = json.Unmarshal(body, &status); err != nil {
		logger.Println("Error while unmarshal the response body:", err)
	}
	_, err = io.WriteString(w, "ok")
	if err != nil {
		logger.Error(err)
	}

	logger.Infof("yandex command result: id: %s status: %s", status.RequestID, status.Status)
}
