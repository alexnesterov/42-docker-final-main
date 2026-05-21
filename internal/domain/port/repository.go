package port

import "github.com/Yandex-Practicum/42-docker-final/internal/domain/entity"

type ParcelRepository interface {
	Add(entity.Parcel) (int, error)
	Get(number int) (entity.Parcel, error)
	GetByClient(client int) ([]entity.Parcel, error)
	SetStatus(number int, status string) error
	SetAddress(number int, address string) error
	Delete(number int) error
}
