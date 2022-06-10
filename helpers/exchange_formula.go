package helpers

// ExchangeFormulaData keeps the data necessary for the conversion
type ExchangeFormulaData struct {
	FromCurrencyRate float64
	ToCurrencyRate   float64
	Amount           float64
}

// ExchangeCalculator holds the basic formula to convert from one currency to another
func ExchangeCalculator(data *ExchangeFormulaData) float64 {
	multipleFromRate := 1 * data.FromCurrencyRate
	divideToRate := 1 / data.ToCurrencyRate
	result := data.Amount * (1 / multipleFromRate) * (1 / divideToRate)
	return result
}
