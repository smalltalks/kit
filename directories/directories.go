package directories

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	rootDirectoryEnv = "KIT_DIR"
)

var (
	// Root is Kit root directory
	Root string
	// Bin is Kit bin directory
	Bin string
	// Scripts is Kit scripts directory
	Scripts string
)

func getRootDirectory() string {
	dir, ok := os.LookupEnv(rootDirectoryEnv)
	if !ok || strings.TrimSpace(dir) == "" {
		panic(fmt.Sprintf("missing environment variable %s", rootDirectoryEnv))
	}

	stat, err := os.Stat(dir)
	if err != nil {
		panic(err)
	}

	if !stat.IsDir() {
		panic(fmt.Sprintf("%s=%s is not a directory", rootDirectoryEnv, dir))
	}

	return dir
}

func init() {
	Root = getRootDirectory()
	Bin = filepath.Join(Root, "bin")
	Scripts = filepath.Join(Root, "scripts")
}
