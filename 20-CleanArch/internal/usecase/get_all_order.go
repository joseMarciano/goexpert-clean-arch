package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type GetAllOrderOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type GetAllOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetAllOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetAllOrderUseCase {
	return &GetAllOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetAllOrderUseCase) Execute() []GetAllOrderOutput {
	var orders = c.OrderRepository.FindAll()
	var output []GetAllOrderOutput

	for _, o := range orders {
		output = append(output, GetAllOrderOutput{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}
	//dto := OrderOutputDTO{
	//	ID:         order.ID,
	//	Price:      order.Price,
	//	Tax:        order.Tax,
	//	FinalPrice: order.Price + order.Tax,
	//}
	//
	//c.OrderCreated.SetPayload(dto)
	//c.EventDispatcher.Dispatch(c.OrderCreated)

	return output
}
