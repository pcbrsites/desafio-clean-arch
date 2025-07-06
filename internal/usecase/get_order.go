package usecase

import (
	"github.com/pcbrsites/desafio-clean-arch/internal/entity"
)

type GetOrderListUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderListUseCase(OrderRepository entity.OrderRepositoryInterface) *GetOrderListUseCase {
	return &GetOrderListUseCase{OrderRepository: OrderRepository}
}

func (c *GetOrderListUseCase) Execute() ([]OrderOutputDTO, error) {
	items, err := c.OrderRepository.GetList()
	if err != nil {
		return nil, err
	}
	var lists []OrderOutputDTO

	for _, item := range items {
		dto := OrderOutputDTO{
			ID:         item.ID,
			Price:      item.Price,
			Tax:        item.Tax,
			FinalPrice: item.FinalPrice,
		}
		lists = append(lists, dto)
	}
	return lists, nil
}
