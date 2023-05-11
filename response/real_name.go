package response

import "time"

type RealName struct {
	UserId      int64      `json:"user_id"`
	RealName    string     `json:"real_name"`
	Gender      int        `json:"gender"`
	IdType      int        `json:"id_type"`
	IdNumber    string     `json:"id_number"`
	IdCardFront string     `json:"id_card_front"`
	IdCardBack  string     `json:"id_card_back"`
	Birthday    *time.Time `json:"birthday"`
	ExpireTime  *time.Time `json:"expire_time"`
	Address     string     `json:"address"`
	Status      int        `json:"status"`
	Reason      string     `json:"reason"`
	Nation      string     `json:"nation"`
	CreatedAt   *time.Time `json:"created_at"`
}
