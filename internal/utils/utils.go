package utils

import (
	"os"
	"path/filepath"
)

func WalkFilesWithExt(root string, extension string, run func(path string, info os.FileInfo) error) (err error) {
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info == nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) == extension {
				return run(path, info)
			}
			return nil
		},
	)
	return err
}
