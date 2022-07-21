package controllers

import (
	"encoding/json"
	"go_test/models"
)

var CreateHttpLog = func(data []byte) bool {
	log := &models.HttpLog{}
	err := json.Unmarshal(data, log)
	if err != nil {
		return false
	}
	return log.CreateLog()
}
