package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	rootDirectoryEnv = "KIT_DIR"
)

var directories = struct {
	root    string
	bin     string
	scripts string
}{}

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
	directories.root = getRootDirectory()
	directories.bin = filepath.Join(directories.root, "bin")
	directories.scripts = filepath.Join(directories.root, "scripts")
}
