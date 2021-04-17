package file

import (
	"github.com/neonima/fftw/pkg/stringster"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

func Walk(path string, recursive bool, extensions ...string) ([]string, error) {
	if recursive {
		return Recursive(path, extensions...)
	}

	return OneLevel(path, extensions...)
}

func OneLevel(path string, extensions ...string) ([]string, error) {
	var files []string
	fc, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, f := range fc {
		if f.IsDir() {
			continue
		}
		if !stringster.IsExtension(f.Name(), extensions...) {
			continue
		}
		files = append(files, f.Name())
	}

	return files, nil
}

func Recursive(path string, extensions ...string) ([]string, error) {
	var files []string
	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if stringster.IsExtension(path, extensions...) {
			files = append(files, d.Name())
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return files, nil
}
