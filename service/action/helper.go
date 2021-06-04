package action

import (
	"context"
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/module/dir"
	"github.com/pnkj-kmr/infra-patch-manager/module/file"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "canceled by sender"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}

// // ConvertFILEToF converts the desire object
// func ConvertFILEToF(f *pb.FILE) *file.F {
// 	return nil
// }

// ConvertFToFILE converts the desire object
func ConvertFToFILE(f *file.F) *pb.FILE {
	return &pb.FILE{
		Isdir: f.IsDir(),
		File:  f.Name(),
		Path:  f.Path(),
		Size:  f.Size(),
		Time:  timestamppb.New(f.ModTime()),
	}
}

// ConvertDToFILE converts the desire object
func ConvertDToFILE(f *dir.D) *pb.FILE {
	return &pb.FILE{
		Isdir: f.IsDir(),
		File:  f.Name(),
		Path:  f.Path(),
		Size:  f.Size(),
		Time:  timestamppb.New(f.ModTime()),
	}
}
