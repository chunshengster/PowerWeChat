package response

type ResponseCardMemberCardUpdateUser struct {
	ErrCode       int    `json:"errcode"`
	ErrMsg        string `json:"errmsg"`
	ResultBonus   int    `json:"result_bonus,omitempty"`
	ResultBalance int    `json:"result_balance,omitempty"`
	Openid        string `json:"openid"`
}
