package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"usersService/internal/proto-files/service"
)

func TestUsers(t *testing.T) {
	con, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	uc := service.NewUsersServiceClient(con)

	ur := service.AddUserRequest{
		Public: &service.UserPublic{
			Id:        0,
			FirstName: "A",
			LastName:  "B",
			Nickname:  "C",
			Email:     "D",
			Country:   "E",
		},
		Private: &service.UserPrivate{Password: "F"},
	}

	resp, err := uc.AddUser(context.Background(), &ur)
	if err != nil {
		log.Println(resp.Err.Code, resp.Err.Message)
		log.Fatal(err)
	}

	fmt.Printf("id=%s createdAt=%s\n", resp.UserId, resp.CreatedAt)
}
