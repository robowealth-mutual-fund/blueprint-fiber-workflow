package blueprint

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) connectToService(host string) grpc.ClientConnInterface {
	msgMaxSize := c.Config.Connection.MaximumMessageSize * 1024 * 1024
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(c.Config.Connection.GRPCTimeout)*time.Second,
	)

	defer cancel()
	target := fmt.Sprintf("dns:///%s", c.Config.Connection.BluePrintService)

	conn, err := grpc.DialContext(ctx,
		target,
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(msgMaxSize),
			grpc.MaxCallSendMsgSize(msgMaxSize),
			grpc.MaxRetryRPCBufferSize(msgMaxSize),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [ { "round_robin": {} } ]}`),
		grpc.WithBlock(),
	)

	if err != nil {
		slog.Error("did not connect: %v", err)
	}

	slog.Info("Connect to service on", host)

	return conn
}
