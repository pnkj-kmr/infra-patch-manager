package agent_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/pnkj-kmr/infra-patch-manager/rpc/pb"
// 	"github.com/stretchr/testify/require"
// )

// func TestPingServer(t *testing.T) {
// 	t.Parallel()

// 	testCases := []struct {
// 		name string
// 		msg  string
// 		ok   string
// 	}{
// 		{
// 			name: "OK",
// 			msg:  "ping",
// 			ok:   "PONG",
// 		},
// 		{
// 			name: "FAIL",
// 			msg:  "",
// 			ok:   "",
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()

// 			req := &pb.PingReq{Msg: tc.msg}
// 			server := NewPatchServer()
// 			res, err := server.Ping(context.Background(), req)

// 			require.NoError(t, err)
// 			require.NotNil(t, res)
// 			require.Equal(t, res.GetMsg(), tc.ok)
// 		})
// 	}
// }
