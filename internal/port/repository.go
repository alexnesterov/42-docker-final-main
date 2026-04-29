package port

import "github.com/Yandex-Practicum/42-docker-final/internal/domain"

type ParcelRepository interface {
	Add(domain.Parcel) (int, error)
	Get(number int) (domain.Parcel, error)
	GetByClient(client int) ([]domain.Parcel, error)
	SetStatus(number int, status string) error
	SetAddress(number int, address string) error
	Delete(number int) error
}
