package client

import (
	"fmt"
	"github.com/begizi/vch-server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"time"
)

type Client interface {
	Tunnel() error
}

type VCHClient struct {
	apiClient pb.VCHClient
}

func (c *VCHClient) Tunnel() error {
	stream, err := c.apiClient.Tunnel(context.Background(), &pb.TunnelRequest{})
	if err != nil {
		return nil
	}

	// Block and run forever
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("Error", err)
			return err
		}
		body := response.GetResponse()
		fmt.Println(body)
	}
	return nil
}

func NewVCHClient(host string) (*VCHClient, error) {
	conn, err := grpc.Dial(host,
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(time.Second*10),
	)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	return &VCHClient{
		apiClient: pb.NewVCHClient(conn),
	}, nil
}
