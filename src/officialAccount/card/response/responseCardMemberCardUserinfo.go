package response

type CommonField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UserInfo struct {
	CommonFieldList []CommonField `json:"common_field_list"`
	CustomFieldList []string      `json:"custom_field_list"`
}

type ResponseCardMemberCardUserInfo struct {
	ErrCode          int       `json:"errcode"`
	ErrMsg           string    `json:"errmsg"`
	OpenId           string    `json:"openid"`
	NickName         string    `json:"nickname,omitempty"`
	MembershipNumber string    `json:"membership_number"`
	Bonus            int       `json:"bonus"`
	Sex              string    `json:"sex,omitempty"`
	UserInfo         *UserInfo `json:"user_info,omitempty"`
	UserCardStatus   string    `json:"user_card_status"`
}
