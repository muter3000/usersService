package test

import (
	"fmt"
	"testing"
	"usersService/internal/grpc/users"
)

func TestMongo(t *testing.T) {
	const uri = "mongodb://127.0.0.1:27017"
	cl := users.ConnectMongo(uri)
	fmt.Printf("%#v", cl)
}
