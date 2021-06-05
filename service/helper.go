package service

import (
	"context"
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/module"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ContextError helps to verify the context error
func ContextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return LogError(status.Error(codes.Canceled, "canceled by sender"))
	case context.DeadlineExceeded:
		return LogError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

// LogError helper method to log the error and returns the error
func LogError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

// // ConvertFILEToF converts the desire object
// func ConvertFILEToF(f *pb.FILE) *file.F {
// 	return nil
// }

// ConvertToFILE converts the desire object
func ConvertToFILE(f module.I) *pb.FILE {
	return &pb.FILE{
		Isdir: f.IsDir(),
		File:  f.Name(),
		Path:  f.Path(),
		Size:  f.Size(),
		Time:  timestamppb.New(f.ModTime()),
	}
}
