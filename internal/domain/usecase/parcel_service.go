// Package usecase содержит бизнес-логику трекера посылок.
package usecase

import (
	"fmt"
	"time"

	"github.com/Yandex-Practicum/42-docker-final/internal/domain/entity"
	"github.com/Yandex-Practicum/42-docker-final/internal/domain/port"
)

type ParcelService struct {
	repo port.ParcelRepository
}

func NewParcelService(repo port.ParcelRepository) *ParcelService {
	return &ParcelService{repo: repo}
}

func (s *ParcelService) Register(req port.RegisterRequest) (entity.Parcel, error) {
	parcel := entity.Parcel{
		Client:    req.Client,
		Status:    entity.ParcelStatusRegistered,
		Address:   req.Address,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}

	id, err := s.repo.Add(parcel)
	if err != nil {
		return parcel, err
	}

	parcel.Number = id

	fmt.Printf("Новая посылка № %d на адрес %s от клиента с идентификатором %d зарегистрирована %s\n",
		parcel.Number, parcel.Address, parcel.Client, parcel.CreatedAt)

	return parcel, nil
}

func (s *ParcelService) GetByClient(client int) ([]entity.Parcel, error) {
	return s.repo.GetByClient(client)
}

func (s *ParcelService) PrintClientParcels(client int) error {
	parcels, err := s.repo.GetByClient(client)
	if err != nil {
		return err
	}

	fmt.Printf("Посылки клиента %d:\n", client)
	for _, parcel := range parcels {
		fmt.Printf("Посылка № %d на адрес %s от клиента с идентификатором %d зарегистрирована %s, статус %s\n",
			parcel.Number, parcel.Address, parcel.Client, parcel.CreatedAt, parcel.Status)
	}
	fmt.Println()

	return nil
}

func (s *ParcelService) NextStatus(number int) error {
	parcel, err := s.repo.Get(number)
	if err != nil {
		return err
	}

	var nextStatus string
	switch parcel.Status {
	case entity.ParcelStatusRegistered:
		nextStatus = entity.ParcelStatusSent
	case entity.ParcelStatusSent:
		nextStatus = entity.ParcelStatusDelivered
	case entity.ParcelStatusDelivered:
		return nil
	}

	fmt.Printf("У посылки № %d новый статус: %s\n", number, nextStatus)

	return s.repo.SetStatus(number, nextStatus)
}

func (s *ParcelService) ChangeAddress(req port.ChangeAddressRequest) error {
	return s.repo.SetAddress(req.Number, req.Address)
}

func (s *ParcelService) Delete(number int) error {
	return s.repo.Delete(number)
}
