package invoicecommands

type PayosWebhookData struct {
	Success bool `json:"success"`
	Data    struct {
		PaymentLinkID   string `json:"paymentLinkId"`
		OrderCode       int64  `json:"orderCode"`
		Amount          int64  `json:"amount"`
		AmountPaid      int64  `json:"amountPaid"`
		AmountRemaining int64  `json:"amountRemaining"`
		Status          string `json:"status"`
		CreatedAt       string `json:"createdAt"`
		Transactions    []struct {
			AccountNumber          string `json:"accountNumber"`
			Amount                 int64  `json:"amount"`
			Description            string `json:"description"`
			TransactionDate        string `json:"transactionDate"`
			Reference              string `json:"reference"`
			CounterAccountBankID   string `json:"counterAccountBankId"`
			CounterAccountBankName string `json:"counterAccountBankName"`
			CounterAccountName     string `json:"counterAccountName"`
			CounterAccountNumber   string `json:"counterAccountNumber"`
		} `json:"transactions"`
		CanceledAt   *string `json:"canceledAt"`
		CompletionAt *string `json:"completionAt"`
	} `json:"data"`
	Desc      string `json:"desc"`
	Code      string `json:"code"`
	Signature string `json:"signature"`
}
