package grpc

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RoleAdminGRPCService struct {
	pb.UnimplementedAdminRolesServiceServer
	roleAdminFacade facade.RoleAdminFacade
}

var _ pb.AdminRolesServiceServer = (*RoleAdminGRPCService)(nil)

func NewRoleAdminGRPCService(roleAdminFacade facade.RoleAdminFacade) *RoleAdminGRPCService {
	return &RoleAdminGRPCService{
		roleAdminFacade: roleAdminFacade,
	}
}

func (ras *RoleAdminGRPCService) Find(ctx context.Context, req *pb.AdminRolesFindRequest) (*pb.AdminRolesInstanceResponse, error) {
	res, err := ras.roleAdminFacade.Get(ctx, req.GetId())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminRolesInstanceResponse_builder{
		Instance: MapRoleDTOToGRPC(res),
	}.Build(), nil
}

func (ras *RoleAdminGRPCService) FindByName(ctx context.Context, req *pb.AdminRolesFindByNameRequest) (*pb.AdminRolesInstanceResponse, error) {
	res, err := ras.roleAdminFacade.GetByName(ctx, req.GetName())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminRolesInstanceResponse_builder{
		Instance: MapRoleDTOToGRPC(res),
	}.Build(), nil
}

func (ras *RoleAdminGRPCService) List(ctx context.Context, req *pb.AdminRolesListRequest) (*pb.AdminRolesInstancesResponse, error) {
	res, err := ras.roleAdminFacade.List(ctx, int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	offset := req.GetOffset()
	limit := req.GetLimit()

	return pb.AdminRolesInstancesResponse_builder{
		Offset:    &offset,
		Limit:     &limit,
		Instances: MapRoleDTOsToGRPC(res),
	}.Build(), nil
}

func (ras *RoleAdminGRPCService) Save(ctx context.Context, req *pb.AdminRolesSaveRequest) (*pb.AdminRolesInstanceResponse, error) {
	income := MapRoleGRPCToDTO(req.GetInstance())
	var res *dto.RoleDTO
	var err error
	if income.ID == "" {
		res, err = ras.roleAdminFacade.Create(ctx, income)
	} else {
		res, err = ras.roleAdminFacade.Change(ctx, income.ID, income)
	}
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminRolesInstanceResponse_builder{
		Instance: MapRoleDTOToGRPC(res),
	}.Build(), nil
}

func (ras *RoleAdminGRPCService) Delete(ctx context.Context, request *pb.AdminRolesDeleteRequest) (*emptypb.Empty, error) {
	err := ras.roleAdminFacade.Delete(ctx, request.GetId())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}
