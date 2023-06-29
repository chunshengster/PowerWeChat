package card

import (
	"context"
	"strings"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/card/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/card/response"
)

type Client struct {
	*kernel.BaseClient

	URL               string
	TicketCacheKey    string
	TicketCachePrefix string
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,
	}

	client.TicketCachePrefix = "powerwechat.official_account.card.api_ticket."

	return client, nil

}

func (comp *Client) Colors() {

}

func (comp *Client) Categories() {

}

// 创建卡券
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html
func (comp *Client) Create(ctx context.Context, param *request.RequestCardCreate) (*response.ResponseCardCreate, error) {
	result := &response.ResponseCardCreate{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/create", param, nil, nil, result)

	return result, err

}

// 查看卡券详情
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) Get(ctx context.Context, cardID string) (*response.ResponseCardGet, error) {
	result := &response.ResponseCardGet{}

	param := object.HashMap{
		"card_id": cardID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/get", param, nil, nil, result)

	return result, err

}

// 批量查询卡券列表
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) List(ctx context.Context, offset int, count int, statusList string) (*response.ResponseCardList, error) {
	result := &response.ResponseCardList{}

	param := object.HashMap{
		"offset":      offset,
		"count":       count,
		"status_list": statusList,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/batchget", param, nil, nil, result)

	return result, err

}

// 更改卡券信息接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) Update(ctx context.Context, cardID string, cardType string, card interface{}) (*response.ResponseCardUpdate, error) {
	result := &response.ResponseCardUpdate{}

	param := object.HashMap{
		"card_id":                 cardID,
		strings.ToLower(cardType): card,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/update", param, nil, nil, result)

	return result, err

}

// 删除卡券接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) Delete(ctx context.Context, cardID string) (*response2.ResponseOfficialAccount, error) {
	result := &response2.ResponseOfficialAccount{}

	param := object.HashMap{
		"card_id": cardID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/delete", param, nil, nil, result)

	return result, err

}

// 创建二维码接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html#0
func (comp *Client) CreateQrCode(ctx context.Context, param request.RequestCreateQrCode) (*response.ResponseCreateQrCode, error) {
	result := &response.ResponseCreateQrCode{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "card/qrcode/create", param, nil, nil, result)

	return result, err

}

// 创建子商户
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html#1
func (comp *Client) SumbitMerchant(ctx context.Context, param *request.SubMerchantSubmit) (*response.SubmerchantSubmit, error) {
	result := &response.SubmerchantSubmit{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "card/submerchant/submit", param, nil, nil, result)
	return result, err
}

// 更新子商户接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html#_1-4-%E6%9B%B4%E6%96%B0%E5%AD%90%E5%95%86%E6%88%B7%E6%8E%A5%E5%8F%A3

func (comp *Client) UpdateMerchant(ctx context.Context, param *request.SubMerchantSubmit) (*response.SubmerchantSubmit, error) {
	result := &response.SubmerchantSubmit{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "card/submerchant/update", param, nil, nil, result)
	return result, err
}

//卡券开放类目查询接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html#_1-3%E5%8D%A1%E5%88%B8%E5%BC%80%E6%94%BE%E7%B1%BB%E7%9B%AE%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3

func (comp *Client) OpenCategory(ctx context.Context) (interface{}, error) {
	result := &response.ProtocolCategory{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "card/getapplyprotocol", nil, nil, nil, result)
	return result, err
}

// 查询Code接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#0
func (comp *Client) GetCardCodeInfo(ctx context.Context, param *request.RequestCardCode) (*response.ResponseCardCode, error) {

	result := &response.ResponseCardCode{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "card/code/get", param, nil, nil, result)
	return result, err
}

//获取用户已领取卡券接口
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#1

func (comp *Client) GetCarUserCarddList(ctx context.Context, param *request.RequestCardUserCardList) (*response.ResponseCardUserCardList, error) {
	result := &response.ResponseCardUserCardList{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "card/user/getcardlist", param, nil, nil, result)
	return result, err
}

// 更新会员信息
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html#_7%E6%9B%B4%E6%96%B0%E4%BC%9A%E5%91%98%E4%BF%A1%E6%81%AF

func (comp *Client) UpdateUserMemberCard(ctx context.Context, param *request.RequestCardMemberCardUpdateUser) (*response.ResponseCardMemberCardUpdateUser, error) {
	result := &response.ResponseCardMemberCardUpdateUser{}
	_, err := comp.BaseClient.HttpPostJson(ctx, "card/membercard/updateuser", param, nil, nil, result)
	return result, err
}
