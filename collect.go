package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func (d *dl) collect() error {
	cfg, err := loadCfg()
	if err != nil {
		return err
	}
	d.cfg = cfg

	var dirs []string
	if !parentDir {
		dirs = make([]string, 0, 64)
		err = filepath.Walk(cfg.dir, func(path string, info os.FileInfo, err error) error {
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

	} else {
		dirs = make([]string, 0, 8)
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		for dir != "" {
			dirs = append(dirs, dir)
			if dir == "/" {
				break
			}
			dir = filepath.Dir(dir)
		}
	}

	sort.Strings(dirs)
	d.dirs = dirs
	d.rsl = -1 // refresh immediately
	return nil
}
