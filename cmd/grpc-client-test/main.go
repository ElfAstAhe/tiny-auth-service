package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type JWTCredentials struct {
	Token string `json:"token"`
}

func (jt *JWTCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	if jt.Token == "" {
		return map[string]string{}, nil
	}

	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", jt.Token),
	}, nil
}

func (jt *JWTCredentials) RequireTransportSecurity() bool {
	return false
}

func main() {
	creds := &JWTCredentials{}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(creds),
	}
	// conn
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	// auth
	authClient := pb.NewAuthServiceClient(conn)
	// context
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ctx := context.Background()
	// packet
	username := "test"
	password := "password"
	login := pb.AuthLoginRequest_builder{
		Username: &username,
		Password: &password,
	}.Build()
	// make auth
	resp, err := authClient.LoginSimple(ctx, login)
	if err != nil {
		log.Fatalf("fail to login: %v\n", err)
	}
	creds.Token = resp.GetToken()
	fmt.Printf("auth result: [%v]\n", resp)
	// ==== profile service ====
	reqProfile := &emptypb.Empty{}
	profileClient := pb.NewUserServiceClient(conn)
	res, err := profileClient.Profile(ctx, reqProfile)
	if err != nil {
		log.Fatalf("fail to profile: %v\n", err)
	}
	fmt.Printf("profile result: [%v]\n", res)
}
