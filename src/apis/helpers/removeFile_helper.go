package helpers

import (
	"os"
	"path"
)

func RemoveFile(dest string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = os.Remove(path.Join(dir, dest))
	if err != nil {
		return err
	}

	return nil
}
