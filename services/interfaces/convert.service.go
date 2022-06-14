package interfaces

import (
	"github.com/navisot/go-exchanger/http/requests"
)

type ConvertServiceInterface interface {
	Convert(request *requests.ConvertRequest) (float64, error)
}
