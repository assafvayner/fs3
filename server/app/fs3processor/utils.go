package fs3processor

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func IsPathSafe(path string) bool {
	return len(path) < 1000 &&
		path != "" &&
		!strings.Contains(path, "..") &&
		!strings.HasPrefix(path, "/") &&
		!strings.ContainsAny(path, "\\\r\t\n'\"()[];:,`|{}?!@#$%^&*+=")
}

func MakeServerSidePath(path, username string) (prefix string) {
	return fmt.Sprintf("/data/%s/%s", username, path)
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
