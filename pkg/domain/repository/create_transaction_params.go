package repository

type NewTransactionCreateParams struct {
	Id     *string `json:"transactionId" query:"id"`
	State  string  `json:"state" query:"state"`
	Amount float64 `json:"amount" query:"amount"`
}

func NewDefaultCreateArticleParams() NewTransactionCreateParams {
	return NewTransactionCreateParams{}
}
