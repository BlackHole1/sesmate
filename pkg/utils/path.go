package utils

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

func ToAbs(p string, create bool) (string, error) {
	if path.IsAbs(p) {
		return p, nil
	}

	absPath, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(absPath)
	if os.IsNotExist(err) && create {
		if err = os.MkdirAll(absPath, os.ModePerm); err != nil {
			return "", err
		}
		return absPath, nil
	} else if err != nil {
		return "", err
	}

	if !fileInfo.IsDir() {
		return "", errors.New("path is not a directory")
	}

	return absPath, nil
}
