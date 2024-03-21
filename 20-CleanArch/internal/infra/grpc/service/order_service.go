package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetAllOrderUseCase usecase.GetAllOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getAllOrderUseCase usecase.GetAllOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetAllOrderUseCase: getAllOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(_ context.Context, _ *pb.Void) (*pb.ListOrderResponse, error) {
	output := s.GetAllOrderUseCase.Execute()
	var response []*pb.OrderResponse

	for _, orderOutput := range output {
		response = append(response, &pb.OrderResponse{
			Id:         orderOutput.ID,
			Price:      float32(orderOutput.Price),
			Tax:        float32(orderOutput.Tax),
			FinalPrice: float32(orderOutput.FinalPrice),
		})
	}

	return &pb.ListOrderResponse{
		Orders: response,
	}, nil
}
