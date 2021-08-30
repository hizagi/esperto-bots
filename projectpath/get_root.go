package projectpath

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.ToSlash(filepath.Join(filepath.Dir(b), ".."))
)
