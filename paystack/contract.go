package paystack

type IBaseContract interface {
	GetAll() error
	FetchById(id string) error
	Create() error
	Update() error
}
