package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	u "go_test/utils"
)

type HttpLog struct {
	gorm.Model
	Url          string         `db:"url" json:"url"`
	Method       string         `db:"method" json:"method"`
	ResponseCode uint16         `db:"response_code" json:"response_code"`
	ResponseText string         `db:"response_text" json:"response_text"`
	ResponseTime uint16         `db:"response_time" json:"response_time"`
	QueryParams  postgres.Jsonb `db:"query_params" json:"query_params"`
	Body         postgres.Jsonb `db:"body" json:"body"`
	Headers      postgres.Jsonb `db:"headers" json:"headers"`
}

func (httplog *HttpLog) Validate() (map[string]interface{}, bool) {
	if httplog.Url == "" || httplog.Method == "" || httplog.ResponseCode == 0 {
		return u.Message(false, "url, method, response_code is required"), false
	}
	return u.Message(true, "success"), true
}

func (httplog *HttpLog) CreateLog() bool {
	if _, ok := httplog.Validate(); !ok {
		return ok
	}
	GetDB().Create(httplog)

	return true
}
