package controller

import (
	"context"
	pb "reservation/grpc_gen/reservation"
	"reservation/model"
	"reservation/pkg/service"
)

type reservationServer struct {
	restaurantService service.Restaurant
	pb.UnimplementedReservationServiceServer
}

func NewReservationServer(restaurantService service.Restaurant) *reservationServer {
	return &reservationServer{
		restaurantService: restaurantService,
	}
}

func (h *reservationServer) CreateRestaurant(ctx context.Context, req *pb.CreateRestaurantReq) (*pb.CreateRestaurantRes, error) {
	newRestaurant := model.CreateRestaurantDTO{
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Description: req.Description,
	}

	if err := h.restaurantService.CreateRestaurant(newRestaurant); err != nil {
		return nil, err
	}

	return &pb.CreateRestaurantRes{Success: true}, nil
}
func (h *reservationServer) UpdateRestaurant(ctx context.Context, req *pb.UpdateRestaurantReq) (*pb.UpdateRestaurantRes, error) {
	restaurant := model.UpdateRestaurantDTO{
		Id:          req.Id,
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Description: req.Description,
	}

	if err := h.restaurantService.UpdateRestaurant(restaurant); err != nil {
		return nil, err
	}

	return &pb.UpdateRestaurantRes{Success: true}, nil
}
func (h *reservationServer) RemoveRestaurant(ctx context.Context, req *pb.RemoveRestaurantReq) (*pb.RemoveRestaurantRes, error) {
	if err := h.restaurantService.RemoveRestaurant(req.GetId()); err != nil {
		return nil, err
	}

	return &pb.RemoveRestaurantRes{Success: true}, nil
}
func (h *reservationServer) GetAllRestaurants(ctx context.Context, req *pb.GetRestaurantsReq) (*pb.GetRestaurantsRes, error) {
	var reservationRestaurants []*pb.Restaurant
	restaurants, err := h.restaurantService.GetAllRestaurants()
	if err != nil {
		return nil, err
	}
	for _, r := range restaurants {
		reservationRestaurants = append(reservationRestaurants, &pb.Restaurant{
			Id:          r.Id,
			Name:        r.Name,
			Address:     r.Address,
			PhoneNumber: r.PhoneNumber,
			Description: r.Description,
		})
	}
	return &pb.GetRestaurantsRes{Restaurants: reservationRestaurants}, nil
}
func (h *reservationServer) GetRestaurantById(ctx context.Context, req *pb.GetRestaurantReq) (*pb.GetRestaurantRes, error) {
	restaurant, err := h.restaurantService.GetRestaurantById(req.GetId())
	if err != nil {
		return nil, err
	}
	reservationRestaurant := &pb.Restaurant{
		Id:          restaurant.Id,
		Name:        restaurant.Name,
		Address:     restaurant.Address,
		PhoneNumber: restaurant.PhoneNumber,
		Description: restaurant.Description,
	}
	return &pb.GetRestaurantRes{Restaurant: reservationRestaurant}, nil
}
