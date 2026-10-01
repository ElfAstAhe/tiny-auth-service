package grpc

import (
	"context"

	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserAdminGRPCService implements the pb.AdminUsersServiceServer gRPC interface network delivery boundary.
// It translates incoming administrative binary Protobuf streams into structured Facade DTO criteria targets.
type UserAdminGRPCService struct {
	pb.UnimplementedAdminUsersServiceServer                        // Embedded fallback forward-compatibility enforcement handle
	userAdminFacade                         facade.UserAdminFacade // Application boundary facade router executing administrative identity usecases
}

// Compile-time interface compliance verification
var _ pb.AdminUsersServiceServer = (*UserAdminGRPCService)(nil)

// NewUserAdminGRPCService acts as a factory constructor mounting required administrative user facade layer dependencies.
func NewUserAdminGRPCService(adminFacade facade.UserAdminFacade) *UserAdminGRPCService {
	return &UserAdminGRPCService{
		userAdminFacade: adminFacade,
	}
}

// Find executes a single record retrieval scenario matching the unique identifier parsed from the gRPC wire data.
func (uas *UserAdminGRPCService) Find(ctx context.Context, req *pb.AdminUsersFindRequest) (*pb.AdminUsersInstanceResponse, error) {
	res, err := uas.userAdminFacade.Get(ctx, req.GetId())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminUsersInstanceResponse_builder{
		Instance: MapUserDTOToGRPC(res),
	}.Build(), nil
}

// FindByName executes a single record lookup operation matching the string username parsed from the incoming Protobuf message payload.
func (uas *UserAdminGRPCService) FindByName(ctx context.Context, req *pb.AdminUsersFindByNameRequest) (*pb.AdminUsersInstanceResponse, error) {
	res, err := uas.userAdminFacade.GetByName(ctx, req.GetName())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return pb.AdminUsersInstanceResponse_builder{
		Instance: MapUserDTOToGRPC(res),
	}.Build(), nil
}

// List handles collection array queries, explicitly casting pagination metric limits and proxying targets to the data transfer mapper layer.
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

// Save coordinates structural identity persistence routing, validating ID fields to branch execution between creation and modification steps.
func (uas *UserAdminGRPCService) Save(ctx context.Context, req *pb.AdminUsersSaveRequest) (*pb.AdminUsersInstanceResponse, error) {
	// PROTO PAYLOAD NOTICE: If the inbound payload carries an unallocated or empty instance node,
	// req.GetInstance() will yield nil, which might cause subsequent nil pointer dereference failures during mapping.
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

// Delete triggers a destructive identity eviction operation bound to the administrative criteria sequence received from the network.
func (uas *UserAdminGRPCService) Delete(ctx context.Context, req *pb.AdminUsersDeleteRequest) (*emptypb.Empty, error) {
	err := uas.userAdminFacade.Delete(ctx, req.GetId())
	if err != nil {
		return nil, MapToGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}
