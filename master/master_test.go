package master_test

// import (
// 	"context"
// 	"net"
// 	"testing"
// )

// func TestPingClient(t *testing.T) {
// 	t.Parallel()

// 	_, addr := startTestPingServer(t)
// 	pingClient := startTestPingClient(t, addr)

// 	req := &pb.PingRequest{Msg: "ping"}
// 	res, err := pingClient.Ping(context.Background(), req)

// 	require.NoError(t, err)
// 	require.NotNil(t, res)
// 	require.Equal(t, res.GetMsg(), "PONG")
// }

// func startTestPingServer(t *testing.T) (*server.PatchServer, string) {
// 	pingServer := server.NewPatchServer()

// 	grpcServer := grpc.NewServer()
// 	pb.RegisterPatchServer(grpcServer, pingServer)

// 	listener, err := net.Listen("tcp", ":0")
// 	require.NoError(t, err)

// 	go grpcServer.Serve(listener)

// 	return pingServer, listener.Addr().String()

// }

// func startTestPingClient(t *testing.T, addr string) pb.PatchClient {
// 	conn, err := grpc.Dial(addr, grpc.WithInsecure())
// 	require.NoError(t, err)
// 	return pb.NewPatchClient(conn)

// }
