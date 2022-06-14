package requests

// ConvertRequest holds the body of the request data
type ConvertRequest struct {
	FromCurrency string  `json:"from_currency" binding:"required"`
	ToCurrency   string  `json:"to_currency" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
	Date         string  `json:"date" binding:"required"`
}
