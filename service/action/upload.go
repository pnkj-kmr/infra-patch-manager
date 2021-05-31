package action

import (
	"bytes"

	"github.com/pnkj-kmr/patch/module/dir"
	"github.com/pnkj-kmr/patch/module/tar"
	"github.com/pnkj-kmr/patch/service/pb"
	"github.com/pnkj-kmr/patch/utility"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CleanRemedyDir cleans the patch folder
func CleanRemedyDir() (err error) {
	d, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "Cannot clean patch directory: %v", err))
	}
	return d.Clean()
}

// ExtractIntoRemedyDir helps to extract tar file into patch folder
func ExtractIntoRemedyDir(t *tar.T) (err error) {
	err = CleanRemedyDir()
	if err != nil {
		return
	}
	return t.Untar(utility.RemedyDirectory)
}

// PostActionAfterUploadFile defines the set of task which need to perform after file upload
func PostActionAfterUploadFile(
	fileName, fileType string, fileData bytes.Buffer,
) (fileList []*pb.FILE, fileSizeWritten int64, err error) {

	// Assets directory - Default patch hold directory
	assetDir, err := dir.New(utility.AssetsDirectory)
	if err != nil {
		return nil, 0, logError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
	}
	// Writeing the file into directory
	fileSizeWritten, err = assetDir.CreateAndWriteFile(fileName+fileType, fileData)
	if err != nil {
		return nil, 0, logError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
	}

	// untaring the uploaded file
	err = ExtractIntoRemedyDir(tar.New(fileName, fileType, utility.AssetsDirectory))
	if err != nil {
		return nil, 0, logError(status.Errorf(codes.Internal, "Unable to extract file into patch directory: %v", err))
	}

	// Patch (remedy) directory
	remedyDir, err := dir.New(utility.RemedyDirectory)
	if err != nil {
		return nil, 0, logError(status.Errorf(codes.Internal, "cannot patch directory.. update config.env file: %v", err))
	}
	// Scan the remedy dir for all files
	files, err := remedyDir.Scan()
	if err != nil {
		return nil, 0, logError(status.Errorf(codes.Internal, "Unable to scan patch directory: %v", err))
	}

	for _, f := range files {
		fileList = append(fileList, convertToFILE(f))
	}
	return
}
