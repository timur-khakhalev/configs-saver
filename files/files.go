package files

import (
	"os"
	"strings"
	//"path/filepath"
)

func EnsurePathExists(rawPath string, isDir bool) bool {
	/*if rawPath == nil {
		return false
	}*/
	var path string
	if strings.Contains(rawPath, "${HOME}") {
		path = strings.Replace(rawPath, "${HOME}", os.Getenv("HOME"), 1)
	} else {
		path = rawPath
	}
	f, err := os.Stat(path)
	if isDir {
		return err == nil && f.IsDir()
	}
		return err == nil && !f.IsDir()
}

func GetStats (path string) {

}
