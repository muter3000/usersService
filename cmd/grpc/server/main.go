package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"usersService/internal/grpc/users"
	"usersService/internal/proto-files/service"
)

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}

func main() {
	logger := hclog.Default()

	netListener := getNetListener(9092)

	gs := grpc.NewServer()

	us := users.NewUsersServer(logger)
	service.RegisterUsersServiceServer(gs, us)

	reflection.Register(gs)

	gs.Serve(netListener)
}
