package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const cfgFileFmt = `
dir     = %s
history = %d
`

type cfg struct {
	dir     string
	history int
}

func loadCfg() (*cfg, error) {
	dir, err := cfgDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(dir, "dl.cfg")
	bs, err := ioutil.ReadFile(path)

	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
		hdir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		s := fmt.Sprintf(cfgFileFmt, hdir, 5)
		err = ioutil.WriteFile(path, []byte(s), 0644)
		if err != nil {
			return nil, err
		}
		return parseCfg(s)
	}

	if err != nil {
		return nil, err
	}

	return parseCfg(string(bs))
}

func parseCfg(s string) (*cfg, error) {
	cfg := &cfg{}

	ss := strings.Split(s, "\n")
	for _, s := range ss {
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "//") {
			continue
		}

		kv := strings.Split(s, "=")
		if len(kv) != 2 {
			continue
		}

		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])

		if k == "" || v == "" {
			continue
		}

		switch k {
		case "dir":
			cfg.dir = v
		case "history":
			i, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			cfg.history = i
		}
	}

	if cfg.dir == "" {
		d, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		cfg.dir = d
	}

	if cfg.history <= 0 {
		cfg.history = 5
	}

	return cfg, nil
}

func cfgDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, ".dl"), nil
}
