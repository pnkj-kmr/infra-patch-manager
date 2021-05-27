package record

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	d, err := NewDir(filepath.Join("..", LocationPatch))
	require.NoError(t, err)

	files, err := d.Scan()
	require.NoError(t, err)
	require.NotNil(t, files)

	// files, err := ScanDir(LocationPatch)
	// for _, file := range files {
	// 	fmt.Println(file.path, file.rPath, file.subPath, file.info.Name())
	// }
}
