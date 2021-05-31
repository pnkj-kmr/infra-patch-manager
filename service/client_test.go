package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/pnkj-kmr/patch/service"
	"github.com/pnkj-kmr/patch/service/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestPingClient(t *testing.T) {
	t.Parallel()

	_, addr := startTestPingServer(t)
	pingClient := startTestPingClient(t, addr)

	req := &pb.PingRequest{Msg: "ping"}
	res, err := pingClient.Ping(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, res.GetMsg(), "PONG")
}

func startTestPingServer(t *testing.T) (*service.PatchServer, string) {
	pingServer := service.NewPatchServer()

	grpcServer := grpc.NewServer()
	pb.RegisterPatchServer(grpcServer, pingServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return pingServer, listener.Addr().String()

}

func startTestPingClient(t *testing.T, addr string) pb.PatchClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewPatchClient(conn)

}
