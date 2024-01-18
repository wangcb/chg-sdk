package response

import "time"

type Application struct {
	Id          int64     `json:"ID"`
	Appid       string    `json:"appid"`
	AppName     string    `json:"app_name"`
	AppDescribe string    `json:"app_describe"`
	Status      int       `json:"status"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
