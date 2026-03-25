package grpc

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserAdminGRPCService struct {
	pb.UnimplementedAdminUsersServiceServer
	userAdminFacade facade.UserAdminFacade
}

func NewUserAdminGRPCService(adminFacade facade.UserAdminFacade) *UserAdminGRPCService {
	return &UserAdminGRPCService{
		userAdminFacade: adminFacade,
	}
}

func (uas *UserAdminGRPCService) Find(ctx context.Context, req *pb.AdminUsersFindRequest) (*pb.AdminUsersInstanceResponse, error) {
	res, err := uas.userAdminFacade.Get(ctx, req.GetId())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminUsersInstanceResponse_builder{
		Instance: MapUserDTOToGRPC(res),
	}.Build(), nil
}

func (uas *UserAdminGRPCService) FindByName(ctx context.Context, req *pb.AdminUsersFindByNameRequest) (*pb.AdminUsersInstanceResponse, error) {
	res, err := uas.userAdminFacade.GetByName(ctx, req.GetName())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminUsersInstanceResponse_builder{
		Instance: MapUserDTOToGRPC(res),
	}.Build(), nil
}

func (uas *UserAdminGRPCService) List(ctx context.Context, req *pb.AdminUsersListRequest) (*pb.AdminUsersInstancesResponse, error) {
	res, err := uas.userAdminFacade.List(ctx, int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	offset := req.GetOffset()
	limit := req.GetLimit()
	return pb.AdminUsersInstancesResponse_builder{
		Offset:    &offset,
		Limit:     &limit,
		Instances: MapUserDTOsToGRPC(res),
	}.Build(), nil
}

func (uas *UserAdminGRPCService) Save(ctx context.Context, req *pb.AdminUsersSaveRequest) (*pb.AdminUsersInstanceResponse, error) {
	income := MapUserGRPCToDTO(req.GetInstance())
	var res *dto.UserDTO
	var err error
	if income.ID == "" {
		res, err = uas.userAdminFacade.Create(ctx, income)
	} else {
		res, err = uas.userAdminFacade.Change(ctx, income.ID, income)
	}
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminUsersInstanceResponse_builder{
		Instance: MapUserDTOToGRPC(res),
	}.Build(), nil
}

func (uas *UserAdminGRPCService) Delete(ctx context.Context, req *pb.AdminUsersDeleteRequest) (*emptypb.Empty, error) {
	err := uas.userAdminFacade.Delete(ctx, req.GetId())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}
