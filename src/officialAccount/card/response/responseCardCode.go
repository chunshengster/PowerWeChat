package response

const (
	USER_CARD_CODE_NORMAL   = "NORMAL"
	USER_CARD_CODE_CONSUMED = "CONSUMED"
	USER_CARD_CODE_EXPIRED  = "EXPIRED"
	USER_CARD_CODE_INVALID  = "INVALID"
)

type ResponseCardCodeCard struct {
	CardID    string `json:"card_id"`
	BeginTime uint   `json:"begin_time"`
	EndTime   uint   `json:"end_time"`
}

type ResponseCardCode struct {
	ErrCode    int                   `json:"errcode"`
	ErrMsg     string                `json:"errmsg"`
	Card       *ResponseCardCodeCard `json:"card"`
	OpenId     string                `json:"openid"`
	CanConsume bool                  `json:"can_consume"`
	OuterStr   string                `json:"outer_str"`
}
