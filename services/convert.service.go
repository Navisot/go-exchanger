package services

import "github.com/navisot/go-exchanger/requests"

type ConvertServiceInterface interface {
	Convert(request *requests.ConvertRequest) (float64, error)
}
