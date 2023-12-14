package response

import "time"

type UserBehaviorStatistic struct {
	ID             int64     `json:"ID"`
	Channel        string    `json:"channel"`
	UserId         int       `json:"user_id"`
	RightsCardId   int       `json:"rights_card_id"`
	EntryAt        time.Time `json:"entry_at"`
	ExternalUserId string    `json:"external_user_id"`
	OrderAmount    int       `json:"order_amount"`
	OrderAt        time.Time `json:"order_at"`
	PatAmount      int       `json:"pat_amount"`
	PayAt          time.Time `json:"pay_at"`
	ScanAt         time.Time `json:"scan_at"`
	LoginAt        time.Time `json:"login_at"`
	RegisterAt     time.Time `json:"register_at"`
}
