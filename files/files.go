package files

import (
	"os/user"
	"path/filepath"
	"strings"
)

func ParsePath(rawPath string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	if rawPath == "~" {
		rawPath = dir
	} else if strings.HasPrefix(rawPath, "~/") {
		rawPath = filepath.Join(dir, rawPath[2:])
	}

	return rawPath
}
