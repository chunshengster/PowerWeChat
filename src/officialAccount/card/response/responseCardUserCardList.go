package response

type CardCode struct {
	CardID string `json:"card_id"`
	Code   string `json:"code"`
}
type ResponseCardUserCardList struct {
	ErrCode      int        `json:"errcode"`
	ErrMsg       string     `json:"errmsg"`
	CardList     []CardCode `json:"card_list"`
	HasShareCard bool       `json:"has_share_card"`
}
