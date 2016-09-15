package wcfile

import (
	"os"
	"path"
	"strings"
)

type WordContainsFile struct {
	patterns []*Pattern
	files    map[string]*os.File
}

type Pattern struct {
	FileName string
	Word     string
	Not      bool
}

var (
	FileMode    os.FileMode = 0644
	DirFileMode os.FileMode = 0755
)

func NewContains(patterns []*Pattern) (*WordContainsFile, error) {
	wc := &WordContainsFile{
		patterns: patterns,
		files:    make(map[string]*os.File),
	}

	if err := wc.initFiles(); err != nil {
		return nil, err
	}

	return wc, nil
}

func (wc *WordContainsFile) initFiles() error {
	for _, v := range wc.patterns {
		if dir, _ := path.Split(v.FileName); len(dir) > 0 {
			if err := os.MkdirAll(dir, DirFileMode); err != nil {
				return err
			}
		}

		if _, isExist := wc.files[v.FileName]; !isExist {
			file, err := os.OpenFile(v.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, FileMode)

			if err != nil {
				return err
			}

			wc.files[v.FileName] = file
		}
	}

	return nil
}

func (wc *WordContainsFile) Write(b []byte) (int, error) {
	t := string(b)

	if len(t) > 0 {
		for _, v := range wc.patterns {
			if strings.Contains(t, v.Word) != v.Not {
				return wc.files[v.FileName].Write(b)
			}
		}
	}

	return 0, nil
}

func (wc *WordContainsFile) Close() error {
	for _, v := range wc.files {
		if err := v.Close(); err != nil {
			return err
		}
	}

	return nil
}
