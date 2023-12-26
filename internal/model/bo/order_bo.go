package bo

type OrderBusinessObject struct {
	//payType int, payState int8, orderId int64, txId string, currency string, amount float64
	PayType  int     `json:"payType"`
	PayState int8    `json:"payState"`
	OrderId  int64   `json:"orderId"`
	TxId     string  `json:"txId"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type OrderStaticsBo struct {
	SuccessOrderCount        int   `json:"success_order_count"`         //成功订单数
	AllOrderCount            int   `json:"all_order_count"`             //订单数
	TotalFee                 int64 `json:"total_fee"`                   //总充值
	IncrementalFee           int64 `json:"incremental_fee"`             //新用户充值
	TotalPaidUserCount       int   `json:"total_paid_user_count"`       //总支付人数
	IncrementalPaidUserCount int   `json:"incremental_paid_user_count"` //新用户支付人数
	PayType                  int   `json:"payType"`                     //支付大渠道
	ChannelCode              int   `json:"channel_code"`                //支付小渠道
}
