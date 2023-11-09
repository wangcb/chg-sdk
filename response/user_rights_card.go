package response

type UserRightsCard struct {
	Id           int64        `json:"id"`
	UserId       int64        `json:"user_id"`
	CardId       int          `json:"card_id"`
	CardName     string       `json:"card_name"`
	CardType     string       `json:"card_type"`
	ActivateAt   string       `json:"activate_at"`
	ExpirationAt string       `json:"expiration_at"`
	Version      int64        `json:"version"`
	Rights       []UserRights `json:"rights"`
}

type UserRights struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	UserCardId   int64  `json:"user_card_id"`
	RightsNo     string `json:"rights_no"`
	RightsName   string `json:"rights_name"`
	ValueCount   int64  `json:"value_count"`
	ResidueValue int64  `json:"residue_value"`
	ValueUnit    int    `json:"value_unit"`
	IsOnlySelf   int    `json:"is_only_self"`
	WaitDay      int    `json:"wait_day"`
	ActivateAt   string `json:"activate_at"`
	ExpirationAt string `json:"expiration_at"`
}
