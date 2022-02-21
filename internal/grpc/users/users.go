package users

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	proto "usersService/internal/proto-files/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Server struct {
	l hclog.Logger
	proto.UnsafeUsersServiceServer
}

func NewUsersServer(l hclog.Logger) *Server {
	return &Server{l, proto.UnimplementedUsersServiceServer{}}
}

func (s *Server) AddUser(ctx context.Context, request *proto.AddUserRequest) (*proto.AddUserResponse, error) {
	s.l.Info("handling add user request", "id", request.Public.Id)
	//TODO Make db call
	uri := "mongodb://127.0.0.1:27017"
	mc := ConnectMongo(uri)

	defer func() {
		if err := mc.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := mc.Database("myDB").Collection("users")
	doc := bson.D{
		{"first_name", request.Public.FirstName},
		{"last_name", request.Public.LastName},
		{"nickname", request.Public.Nickname},
		{"email", request.Public.Email},
		{"country", request.Public.Country},
	}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}

	//
	return &proto.AddUserResponse{
		UserId:    fmt.Sprintf("%v", result),
		CreatedAt: time.Now().String(),
		Err:       nil,
	}, nil
}

func ConnectMongo(uri string) *mongo.Client {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	return client
}
