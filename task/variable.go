package task

import "path/filepath"

// LocationPatch - default patch location
var LocationPatch string

// LocationRollback - default last rollback location
var LocationRollback string

// LocationTarFiles - default tar files location
var LocationTarFiles string

func init() {
	LocationPatch = filepath.Join("asset", "patch")
	LocationRollback = filepath.Join("asset", "revoke")
	LocationTarFiles = filepath.Join("uploads")
}
