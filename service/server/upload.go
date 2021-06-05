package server

import (
	"bytes"
	"log"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/module/dir"
	"github.com/pnkj-kmr/infra-patch-manager/module/tar"
	"github.com/pnkj-kmr/infra-patch-manager/service"
	"github.com/pnkj-kmr/infra-patch-manager/service/pb"
	"github.com/pnkj-kmr/infra-patch-manager/utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func cleanRemedyDir() (err error) {
	d, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return service.LogError(status.Errorf(codes.Internal, "Cannot find default patch folder: %v | %v", utility.RemedyDirectory, err))
	}
	return d.Clean()
}

func extractIntoRemedyDir(t *tar.T) (err error) {
	err = cleanRemedyDir()
	if err != nil {
		return service.LogError(status.Errorf(codes.Internal, "Cannot clean default patch folder: %v | %v", utility.RemedyDirectory, err))
	}
	return t.Untar(utility.RemedyDirectory)
}

func postActionAfterUploadFile(
	fileName, fileType string, fileData bytes.Buffer,
) (fileList []*pb.FILE, fileSizeWritten int64, err error) {
	start := time.Now()

	// Assets directory - Default patch hold directory
	assetDir, err := dir.New(utility.AssetsDirectory)
	if err != nil {
		return nil, 0, service.LogError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
	}
	// Writeing the file into directory
	fileSizeWritten, err = assetDir.CreateAndWriteFile(fileName+fileType, fileData)
	if err != nil {
		return nil, 0, service.LogError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
	}

	// untaring the uploaded file
	err = extractIntoRemedyDir(tar.New(fileName, fileType, utility.AssetsDirectory))
	if err != nil {
		return nil, 0, service.LogError(status.Errorf(codes.Internal, "Unable to extract file into patch directory: %v", err))
	}

	// Patch (remedy) directory
	remedyDir, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return nil, 0, service.LogError(status.Errorf(codes.Internal, "cannot patch directory.. update config.env file: %v", err))
	}
	// Scan the remedy dir for all files
	files, err := remedyDir.Scan()
	if err != nil {
		return nil, 0, service.LogError(status.Errorf(codes.Internal, "Unable to scan patch directory: %v", err))
	}

	for _, f := range files {
		fileList = append(fileList, service.ConvertToFILE(f))
	}
	log.Println("FILE UPLOAD POST ACTION FOR", fileName+fileType, "T:", time.Since(start))
	return
}
