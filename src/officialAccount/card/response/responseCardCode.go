package response

const (
	USER_CARD_CODE_NORMAL   = "NORMAL"
	USER_CARD_CODE_CONSUMED = "CONSUMED"
	USER_CARD_CODE_EXPIRED  = "EXPIRED"
	USER_CARD_CODE_INVALID  = "INVALID"
)

type ResponseCardCodeCard struct {
	CardID           string `json:"card_id"`
	BeginTime        uint   `json:"begin_time"`
	EndTime          uint   `json:"end_time"`
	Bonus            int    `json:"bonus,omitempty"`
	Code             string `json:"code,omitempty"`
	MembershipNumber string `json:"membership_number,omitempty"`
}

type ResponseCardCode struct {
	ErrCode          int                   `json:"errcode"`
	ErrMsg           string                `json:"errmsg"`
	Card             *ResponseCardCodeCard `json:"card"`
	OpenId           string                `json:"openid"`
	CanConsume       bool                  `json:"can_consume"`
	OuterStr         string                `json:"outer_str"`
	UserCardStatus   string                `json:"user_card_status,omitempty"`
	QRCode           string                `json:"QRCode,omitempty"`
	BackgroundPicURL string                `json:"background_pic_url,omitempty"`
	UnionID          string                `json:"unionid,omitempty"`
}
