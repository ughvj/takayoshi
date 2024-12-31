package dml

import (
	"io"
	"io/fs"
	"embed"
)

//go:embed *.sql
var ConfFS embed.FS

var Loader = NewDmlLoader(ConfFS)

type DmlList struct {
	fs fs.FS
}

func NewDmlLoader(fs_ fs.FS) *DmlList {
	return &DmlList{fs_}
}

func (l *DmlList) Get(name string) (string, error) {
	f, err := l.fs.Open(name + ".sql")
	if err != nil {
		return "", nil
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
