package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	// GetTotal() (int, error)
	GetList() ([]*Order, error)
}
