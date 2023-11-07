package response

import "time"

type UserRightsCard struct {
	Id       int64        `json:"id"`
	UserId   int64        `json:"user_id"`
	CardName string       `json:"card_name"`
	CardType string       `json:"card_type"`
	Version  string       `json:"version"`
	Rights   []UserRights `json:"rights"`
}

type UserRights struct {
	Id           int64      `json:"id"`
	UserId       int64      `json:"user_id"`
	UserCardId   int64      `json:"user_card_id"`
	RightsNo     string     `json:"rights_no"`
	RightsName   string     `json:"rights_name"`
	ValueCount   int64      `json:"value_count"`
	ResidueValue int64      `json:"residue_value"`
	ValueUnit    int        `json:"value_unit"`
	IsOnlySelf   int        `json:"is_only_self"`
	WaitDay      int        `json:"wait_day"`
	ActiveTime   *time.Time `json:"active_time"`
	ExpirationAt *time.Time `json:"expiration_at"`
}
