package fs3processor

import (
	"errors"
	"os"
	"strings"

	"github.com/assafvayner/fs3/server/app/config"
)

func IsPathSafe(path string) bool {
	return len(path) < 1000 &&
		path != "" &&
		!strings.Contains(path, "..") &&
		!strings.HasPrefix(path, "/") &&
		!strings.ContainsAny(path, "\\\r\t\n'\"()[];:,`|{}?!@#$%^&*+=")
}

func MakeServerSidePath(path, username string) (prefix string) {
	prefix = "/data/"
	if config.IsPrimary() {
		prefix += "p/"
	} else {
		prefix += "b/"
	}
	if username != "" {
		prefix += username + "/"
	}
	return prefix + path
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
