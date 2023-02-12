package domain

type TokopediaWebhookMsg struct {
	MsgId     int64  `json:"msg_id"`
	BuyerId   int64  `json:"buyer_id"`
	Message   string `json:"message"`
	Thumbnail string `json:"thumbnail"`
	FullName  string `json:"full_name"`
	ShopId    int64  `json:"shop_id"`
}
