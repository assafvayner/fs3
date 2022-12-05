package fs3processor

import (
	"errors"
	"os"
	"strings"

	"gitlab.cs.washington.edu/assafv/fs3/server/config"
)

func IsPathSafe(path string) bool {
	return len(path) < 1000 &&
		!strings.Contains(path, "..") &&
		!strings.HasPrefix(path, "/") &&
		!strings.ContainsAny(path, "\\\r\t\n'\"()[];:,`|{}?!@#$%^&*+=")
}

func MakeServerSidePath(path string) (fullpath string) {
	fullpath = "/data/"
	if config.IsPrimary() {
		fullpath += "p/"
	} else {
		fullpath += "b/"
	}
	fullpath += path
	return fullpath
}

func FileNotExists(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, os.ErrNotExist)
}

func GetDirFromFilePath(path string) string {
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 {
		return ""
	}
	return path[:lastSlash]
}
