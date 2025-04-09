package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

func TestWalkFilesWithExt(t *testing.T) {
	path := runpath.PARENT.Path()
	err := WalkFilesWithExt(path, ".go", func(path string, info os.FileInfo) error {
		t.Log(path)
		return nil
	})
	require.NoError(t, err)
}
