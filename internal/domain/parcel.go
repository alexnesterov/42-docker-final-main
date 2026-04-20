package domain

const (
	ParcelStatusRegistered = "registered"
	ParcelStatusSent       = "sent"
	ParcelStatusDelivered  = "delivered"
)

type Parcel struct {
	Number    int
	Client    int
	Status    string
	Address   string
	CreatedAt string
}

type ParcelRepository interface {
	Add(Parcel) (int, error)
	Get(number int) (Parcel, error)
	GetByClient(client int) ([]Parcel, error)
	SetStatus(number int, status string) error
	SetAddress(number int, address string) error
	Delete(number int) error
}

type ParcelService interface {
	Register(client int, address string) (Parcel, error)
	PrintClientParcels(client int) error
	NextStatus(number int) error
	ChangeAddress(number int, address string) error
	Delete(number int) error
}
