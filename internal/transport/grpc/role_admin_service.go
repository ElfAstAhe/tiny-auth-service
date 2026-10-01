package grpc

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// RoleAdminGRPCService implements the pb.AdminRolesServiceServer gRPC interface network delivery boundary.
// It translates incoming administrative binary Protobuf streams into structured Facade DTO criteria targets.
type RoleAdminGRPCService struct {
	pb.UnimplementedAdminRolesServiceServer                        // Embedded fallback forward-compatibility enforcement handle
	roleAdminFacade                         facade.RoleAdminFacade // Application boundary facade router executing administrative role usecases
}

// Compile-time interface compliance verification
var _ pb.AdminRolesServiceServer = (*RoleAdminGRPCService)(nil)

// NewRoleAdminGRPCService acts as a factory constructor mounting required administrative role facade layer dependencies.
func NewRoleAdminGRPCService(roleAdminFacade facade.RoleAdminFacade) *RoleAdminGRPCService {
	return &RoleAdminGRPCService{
		roleAdminFacade: roleAdminFacade,
	}
}

// Find executes a single record retrieval scenario matching the unique identifier parsed from the gRPC wire data.
func (ras *RoleAdminGRPCService) Find(ctx context.Context, req *pb.AdminRolesFindRequest) (*pb.AdminRolesInstanceResponse, error) {
	res, err := ras.roleAdminFacade.Get(ctx, req.GetId())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminRolesInstanceResponse_builder{
		Instance: MapRoleDTOToGRPC(res),
	}.Build(), nil
}

// FindByName executes a single record lookup operation matching the string role name parsed from the incoming Protobuf message payload.
func (ras *RoleAdminGRPCService) FindByName(ctx context.Context, req *pb.AdminRolesFindByNameRequest) (*pb.AdminRolesInstanceResponse, error) {
	res, err := ras.roleAdminFacade.GetByName(ctx, req.GetName())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminRolesInstanceResponse_builder{
		Instance: MapRoleDTOToGRPC(res),
	}.Build(), nil
}

// List handles collection array queries, explicitly casting pagination metric limits and proxying targets to the data transfer mapper layer.
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

// Save coordinates structural role persistence routing, validating ID fields to branch execution between creation and modification steps.
func (ras *RoleAdminGRPCService) Save(ctx context.Context, req *pb.AdminRolesSaveRequest) (*pb.AdminRolesInstanceResponse, error) {
	// PROTO PAYLOAD NOTICE: If the inbound payload carries an unallocated or empty instance node,
	// req.GetInstance() will yield nil, which might cause subsequent nil pointer dereference failures during mapping.
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

// Delete triggers a destructive role eviction operation bound to the administrative criteria sequence received from the network.
func (ras *RoleAdminGRPCService) Delete(ctx context.Context, request *pb.AdminRolesDeleteRequest) (*emptypb.Empty, error) {
	err := ras.roleAdminFacade.Delete(ctx, request.GetId())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}
