package services

import (
	"encoding/xml"
	"errors"
	"github.com/navisot/go-exchanger/helpers"
	"github.com/navisot/go-exchanger/requests"
	"github.com/navisot/go-exchanger/xml_config"
	"net/http"
	"strconv"
)

var (
	envelope xml_config.Envelope
)

type ECBConvert struct{}

// NewECBConvertService returns a new ConvertServiceInterface
func NewECBConvertService() ConvertServiceInterface {
	return &ECBConvert{}
}

// Convert does the conversion to get the result
func (ecb *ECBConvert) Convert(data *requests.ConvertRequest) (float64, error) {

	if err := parseXML(); err != nil {
		return 0, err
	}

	formulaData, err := findExchangeRates(data)

	if err != nil {
		return 0, err
	}

	result := helpers.ExchangeCalculator(formulaData)

	return result, nil
}

// parseXML gets the XML file and decodes it
func parseXML() error {

	res, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")

	if err != nil {
		return err
	}

	if err := xml.NewDecoder(res.Body).Decode(&envelope); err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

// findExchangeRates finds the exchange rates based on a given date
func findExchangeRates(data *requests.ConvertRequest) (*helpers.ExchangeFormulaData, error) {

	var exchangeRate helpers.ExchangeFormulaData

	// Adds the amount
	exchangeRate.Amount = data.Amount

	// Loop through the cubes to find exchange rates
	for _, timeCube := range envelope.Cube.TimeCube {
		if data.Date == timeCube.Time {
			for _, currencyRateCube := range timeCube.CurrencyRateCube {
				if currencyRateCube.Currency == data.FromCurrency {
					exchangeRate.FromCurrencyRate, _ = strconv.ParseFloat(currencyRateCube.Rate, 64)
				}
				if currencyRateCube.Currency == data.ToCurrency {
					exchangeRate.ToCurrencyRate, _ = strconv.ParseFloat(currencyRateCube.Rate, 64)
				}
			}
		}
	}

	// error
	if exchangeRate.FromCurrencyRate == 0 || exchangeRate.ToCurrencyRate == 0 {
		errors.New("could not find exchange rate for the selected date")
	}

	return &exchangeRate, nil
}
