package agent

import (
	"log"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
)

func backupExistingRollback() (err error) {
	d, err := entity.NewDir(entity.C.RollbackPath)
	if err != nil {
		log.Println("Unable to load rollback folder", entity.C.RollbackPath, err)
		return err
	}
	assetDir, err := entity.NewDir(entity.C.AssetPath)
	if err != nil {
		log.Println("Unable to load assets folder", entity.C.AssetPath, err)
		return err
	}
	t := entity.NewTar(entity.RandomStringWithTime(0, "ROLLBACK"), ".tar.gz", assetDir.Path())
	return t.Tar([]string{d.Path()})
}

// func postActionAfterUploadFile(
// 	fileName, fileType string, fileData bytes.Buffer,
// ) (fileList []*pb.FILE, fileSizeWritten int64, err error) {
// 	start := time.Now()

// 	// Assets directory - Default patch hold directory
// 	assetDir, err := entity.NewDir(entity.C.AssetPath)
// 	if err != nil {
// 		return nil, 0, grpc.LogError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
// 	}
// 	// Writeing the file into directory
// 	fileSizeWritten, err = assetDir.CreateAndWriteFile(fileName+fileType, fileData)
// 	if err != nil {
// 		return nil, 0, grpc.LogError(status.Errorf(codes.Internal, "cannot save file to assets: %v", err))
// 	}

// 	// // untaring the uploaded file
// 	// err = extractIntoRemedyDir(entity.NewTar(fileName, fileType, entity.C.AssetPath))
// 	// if err != nil {
// 	// 	return nil, 0, grpc.LogError(status.Errorf(codes.Internal, "Unable to extract file into patch directory: %v", err))
// 	// }

// 	// // Patch (remedy) directory
// 	// remedyDir, err := entity.NewDir(entity.C.PatchPath)
// 	// if err != nil {
// 	// 	return nil, 0, grpc.LogError(status.Errorf(codes.Internal, "cannot patch directory.. update config.env file: %v", err))
// 	// }
// 	// // Scan the remedy dir for all files
// 	// files, err := remedyDir.Scan()
// 	// if err != nil {
// 	// 	return nil, 0, grpc.LogError(status.Errorf(codes.Internal, "Unable to scan patch directory: %v", err))
// 	// }

// 	// for _, f := range files {
// 	// 	fileList = append(fileList, grpc.EntityToFILE(f))
// 	// }
// 	log.Println("FILE UPLOAD POST ACTION FOR", fileName+fileType, "T:", time.Since(start))
// 	return
// }
