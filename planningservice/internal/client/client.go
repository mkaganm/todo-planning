package client

import (
	"context"
	"github.com/mkaganm/providergrpc"
	"google.golang.org/grpc"
	"log"
)

// GetTasks is a function that returns tasks
func GetTasks(port string) ([]*providergrpc.Task, error) {

	// connection to the gRPC server
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Default().Println("Error while connecting to the server: ", err)
		return nil, err
	}

	// close the connection
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Default().Println("Error while closing the connection: ", err)
		}
	}(conn)

	// create a new Provider client
	client := providergrpc.NewProviderClient(conn)

	// GetTasks RPC call
	response, err := client.GetTasks(context.Background(), &providergrpc.Request{})
	if err != nil {
		log.Default().Println("Error when calling the GetTasks ", err)
		return nil, err
	}

	tasks := response.Tasks

	return tasks, nil

}
