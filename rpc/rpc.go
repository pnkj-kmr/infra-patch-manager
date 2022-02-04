package rpc

import (
	"context"
	"io/fs"
	"log"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
	"github.com/pnkj-kmr/infra-patch-manager/rpc/pb"
	"google.golang.org/grpc"
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

// EntityToFILE converts the desire object
func EntityToFILE(f entity.Entity) *pb.FILE {
	return &pb.FILE{
		Isdir: f.IsDir(),
		Name:  f.Name(),
		Path:  f.Path(),
		Size:  f.Size(),
		Time:  timestamppb.New(f.ModTime()),
	}
}

type _pbfile struct{ f *pb.FILE }

// FILEToEntity - convert FILE into entity
func FILEToEntity(f *pb.FILE) entity.Entity { return &_pbfile{f} }
func (p *_pbfile) Create(string) error      { return nil }
func (p *_pbfile) Path() string             { return p.f.GetPath() }
func (p *_pbfile) Name() string             { return p.f.GetName() }
func (p *_pbfile) Size() int64              { return p.f.GetSize() }
func (p *_pbfile) Mode() fs.FileMode        { return 0 }
func (p *_pbfile) ModTime() time.Time       { return p.f.GetTime().AsTime() }
func (p *_pbfile) IsDir() bool              { return p.f.GetIsdir() }
func (p *_pbfile) Sys() interface{}         { return nil }

// Connection returns the connection with grpc dial
func Connection(address string) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Print(err)
		return
	}
	return
}
