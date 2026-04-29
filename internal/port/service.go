package port

import "github.com/Yandex-Practicum/42-docker-final/internal/domain"

type RegisterRequest struct {
	Client  int    `json:"client"`
	Address string `json:"address"`
}

type ChangeAddressRequest struct {
	Number  int
	Address string `json:"address"`
}

type ParcelService interface {
	Register(RegisterRequest) (domain.Parcel, error)
	GetByClient(client int) ([]domain.Parcel, error)
	PrintClientParcels(client int) error
	NextStatus(number int) error
	ChangeAddress(ChangeAddressRequest) error
	Delete(number int) error
}
