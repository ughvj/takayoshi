package dml

import (
	"io"
	"io/fs"
	"embed"
	"bytes"

	"text/template"
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

func (l *DmlList) EmbedAndGet(name string, embedded string) (string, error) {
	f, err := l.fs.Open(name + ".sql")
	if err != nil {
		return "", nil
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", nil
	}

	funcMap := template.FuncMap{
		"EMBEDDED": func() string {
			return embedded
		},
	}
	tmpl, err := template.New("name").Funcs(funcMap).Parse(string(b))
	if err != nil {
		return "", err
	}
	w := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(w, nil)
	if err != nil {
		return "", err
	}

	return w.String(), nil
}
