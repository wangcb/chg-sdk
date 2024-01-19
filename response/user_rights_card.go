package response

type UserRightsCard struct {
	Id              int64        `json:"id"`
	UserId          int64        `json:"user_id"`
	Username        string       `json:"username"`
	Phone           string       `json:"phone"`
	CardId          int          `json:"card_id"`
	CardName        string       `json:"card_name"`
	CardType        string       `json:"card_type"`
	ActivateAt      string       `json:"activate_at"`
	Status          int          `json:"status"`
	ExpirationAt    string       `json:"expiration_at"`
	Version         int64        `json:"version"`
	CreatedAt       string       `json:"created_at"`
	HealthManager   string       `json:"health_manager"`
	FamilyMemberSum int          `json:"family_member_sum"`
	InquiryCount    int          `json:"inquiry_count"`
	Rights          []UserRights `json:"rights"`
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

// UserRightsCardList 用户权益卡列表返回参数结构体
type UserRightsCardList struct {
	Total int64            `json:"total"`
	Data  []UserRightsCard `json:"data"`
}

// RightsChannel 权益渠道列表返回参数结构体
type RightsChannel struct {
	ID          int    `json:"id"`
	ChannelName string `json:"channel_name"`
}
