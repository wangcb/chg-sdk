package request

type RobotCall struct {
	Phone     string `json:"phone" valid:"phone"`
	VoiceCode string `json:"voice_code" valid:"voice_code"`
	Source    string `json:"source" valid:"source"`
	Remark    string `json:"remark" valid:"remark"`
}
