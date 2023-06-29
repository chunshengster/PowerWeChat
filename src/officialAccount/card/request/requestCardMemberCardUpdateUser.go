package request

type NotifyOptional struct {
	IsNotifyBonus        bool `json:"is_notify_bonus,omitempty"`
	IsNotifyBalance      bool `json:"is_notify_balance,omitempty"`
	IsNotifyCustomField1 bool `json:"is_notify_custom_field1,omitempty"`
	IsNotifyCustomField2 bool `json:"is_notify_custom_field2,omitempty"`
	IsNotifyCustomField3 bool `json:"is_notify_custom_field3,omitempty"`
}

type RequestCardMemberCardUpdateUser struct {
	Code              string          `json:"code"`
	CardId            string          `json:"card_id"`
	BackgroundPicURL  string          `json:"background_pic_url,omitempty"`
	Bonus             int             `json:"bonus,omitempty"`
	AddBonus          int             `json:"add_bonus,omitempty"`
	Record_Bonus      string          `json:"record_bonus,omitempty"`
	Balance           int             `json:"balance,omitempty"`
	AddBalance        int             `json:"add_balance,omitempty"`
	Record_Balance    string          `json:"record_balance,omitempty"`
	CustomFieldValue1 string          `json:"custom_field_value1,omitempty"`
	CustomFieldValue2 string          `json:"custom_field_value2,omitempty"`
	CustomFieldValue3 string          `json:"custom_field_value3,omitempty"`
	NotifyOptional    *NotifyOptional `json:"notify_optional,omitempty"`
}
