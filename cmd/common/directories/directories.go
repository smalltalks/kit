package directories

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	rootDirectoryEnv = "KIT_DIR"
)

// Directories hold the paths to traverse around KIT structure
type Directories struct {
	Root    string
	Bin     string
	Docs    string
	Scripts string
}

var (
	initialize     sync.Once
	pkgDirectories = Directories{}
)

// Get return directories path to traverse around KIT structure
func Get() Directories {
	initialize.Do(func() {
		root := getRootDirectory()
		pkgDirectories.Root = root
		pkgDirectories.Bin = filepath.Join(root, "bin")
		pkgDirectories.Docs = filepath.Join(root, "docs")
		pkgDirectories.Scripts = filepath.Join(root, "scripts")
	})
	return pkgDirectories
}

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
