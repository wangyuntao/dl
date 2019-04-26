package main

import (
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func (d *dl) collect() error {
	dir, err := dir()
	if err != nil {
		return err
	}
	dirs := make([]string, 0, 256)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		if strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		dirs = append(dirs, path)
		return nil
	})

	if err != nil {
		return err
	}
	sort.Strings(dirs)

	d.dirs = dirs
	d.rsl = -1 // refresh immediately
	return nil
}

func dir() (string, error) {
	dir := flag.Arg(0)
	if dir != "" {
		return dir, nil
	}

	dir = os.Getenv("DL_DIR")
	if dir != "" {
		return dir, nil
	}

	return os.UserHomeDir()
}
