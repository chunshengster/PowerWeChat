package request

type RequestCardCode struct {
	CardID       string `json:"card_id"`
	CardCode     string `json:"code"`
	CheckConsume bool   `json:"check_consume"`
}
