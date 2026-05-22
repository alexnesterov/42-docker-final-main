// Package port содержит интерфейсы и типы запросов для взаимодействия между слоями.
package port

import "github.com/Yandex-Practicum/42-docker-final/internal/domain/entity"

type RegisterRequest struct {
	Client  int    `json:"client"`
	Address string `json:"address"`
}

type ChangeAddressRequest struct {
	Number  int
	Address string `json:"address"`
}

type ParcelService interface {
	Register(RegisterRequest) (entity.Parcel, error)
	GetByClient(client int) ([]entity.Parcel, error)
	PrintClientParcels(client int) error
	NextStatus(number int) error
	ChangeAddress(ChangeAddressRequest) error
	Delete(number int) error
}
