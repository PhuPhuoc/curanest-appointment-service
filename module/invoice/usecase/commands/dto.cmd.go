package invoicecommands

type PayosWebhookData struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Desc      string `json:"desc"`
	Signature string `json:"signature"`
	Data      struct {
		OrderCode              int64  `json:"orderCode"`
		Amount                 int64  `json:"amount"`
		Description            string `json:"description"`
		AccountNumber          string `json:"accountNumber"`
		Reference              string `json:"reference"`
		TransactionDateTime    string `json:"transactionDateTime"`
		Currency               string `json:"currency"`
		PaymentLinkID          string `json:"paymentLinkId"`
		Code                   string `json:"code"`
		Desc                   string `json:"desc"`
		CounterAccountBankID   string `json:"counterAccountBankId"`
		CounterAccountBankName string `json:"counterAccountBankName"`
		CounterAccountName     string `json:"counterAccountName"`
		CounterAccountNumber   string `json:"counterAccountNumber"`
		VirtualAccountName     string `json:"virtualAccountName"`
		VirtualAccountNumber   string `json:"virtualAccountNumber"`
	} `json:"data"`
}
