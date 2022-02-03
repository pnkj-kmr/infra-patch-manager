package agent

import (
	"bytes"

	"github.com/pnkj-kmr/infra-patch-manager/entity"
)

// PatchAgent declare all agent actions
type PatchAgent interface {
	RightsCheck() (bool, error)
	// File write into system
	WriteUploaded(entity.Entity, bytes.Buffer) (entity.Entity, error)
	// patch related
	PatchNow() error
	VerifyPatched() ([]entity.File, bool, error)
	// rollback related
	PatchRollback() error
	VerifyRollbacked() (bool, error)
	// patch file related
	PatchExtract(string, string) error
	VerifyExtracted() ([]entity.File, bool, error)
	// List available patches
	ListAssets() ([]entity.Entity, error)
}
