package request

import (
	pr "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/refund/request"
)

type RequestPartnerRefund struct {
	SubMchID string `json:"sub_mchid"`
	pr.RequestRefund
}
