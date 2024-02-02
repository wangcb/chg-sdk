package response

type WechatTemplate struct {
	Content              string        `json:"content"`
	TempParams           string        `json:"temp_params"`
	Example              string        `json:"example"`
	KeywordEnumValueList []interface{} `json:"keywordEnumValueList"`
	PriTmplId            string        `json:"priTmplId"`
	Title                string        `json:"title"`
	Type                 int           `json:"type"`
}
