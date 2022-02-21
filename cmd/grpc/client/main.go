package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"usersService/internal/proto-files/service"
)

func main() {
	l := hclog.Default()

	con, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Error("connecting to rep microservice", "error", err)
		return
	}
	defer con.Close()
	rc := service.NewUsersServiceClient(con)

	obj := service.AddUserRequest{
		Public: &service.UserPublic{
			Id:        0,
			FirstName: "a",
			LastName:  "b",
			Nickname:  "c",
			Email:     "d",
			Country:   "e",
		},
		Private: &service.UserPrivate{Password: "f"},
	}

	rsp, err := rc.AddUser(context.Background(), &obj)
	if err != nil {
		l.Error("during add request", "error", err)
		return
	}
	l.Info(fmt.Sprintf("%#v", rsp.UserId))
}
