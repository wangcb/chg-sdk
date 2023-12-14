package request

import (
	"github.com/wangcb/chg-sdk/response"
	"time"
)

type UserBehaviorStatisticSearch struct {
	response.UserBehaviorStatistic
	StartCreatedAt  *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt    *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartEntryAt    *time.Time `json:"startEntryAt" form:"startEntryAt"`
	EndEntryAt      *time.Time `json:"endEntryAt" form:"endEntryAt"`
	StartOrderAt    *time.Time `json:"startOrderAt" form:"startOrderAt"`
	EndOrderAt      *time.Time `json:"endOrderAt" form:"endOrderAt"`
	StartPayAt      *time.Time `json:"startPayAt" form:"startPayAt"`
	EndPayAt        *time.Time `json:"endPayAt" form:"endPayAt"`
	StartScanAt     *time.Time `json:"startScanAt" form:"startScanAt"`
	EndScanAt       *time.Time `json:"endScanAt" form:"endScanAt"`
	StartLoginAt    *time.Time `json:"startLoginAt" form:"startLoginAt"`
	EndLoginAt      *time.Time `json:"endLoginAt" form:"endLoginAt"`
	StartRegisterAt *time.Time `json:"startRegisterAt" form:"startRegisterAt"`
	EndRegisterAt   *time.Time `json:"endRegisterAt" form:"endRegisterAt"`
	Page            int        `form:"page" valid:"page"`
	PerPage         int        `form:"per_page" valid:"per_page"`
}
