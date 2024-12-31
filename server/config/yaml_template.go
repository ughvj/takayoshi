package config

import (
	"html/template"
	"io"
	"os"
)

type YAMLTemplate struct{}

func NewYAMLTemplate() *YAMLTemplate {
	return &YAMLTemplate{}
}

func (t *YAMLTemplate) Compile(name string, r io.Reader, w io.Writer) error {
	funcMap := template.FuncMap{
		"env": t.Env, // カスタム関数の登録
	}
	value, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	tmpl, err := template.New("name").Funcs(funcMap).Parse(string(value))
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, nil)
	return err
}

// キーで指定された環境変数を返す。このとき値が空なら、デフォルト引数リストの中から最初のゼロでない値を返す。
//
//	os.Setenv("A", "1")
//
//	a: {{ env "A" }}                 #=> a: 1
//	b: {{ env "B" "val" }}           #=> b: val
//	c: {{ env "C" (env "D") "val" }} #=> c: val
func (t *YAMLTemplate) Env(key string, defaultValues ...string) string {
	val := os.Getenv(key)
	if val == "" {
		values := compact(defaultValues)
		if len(values) == 0 {
			return ""
		}
		return values[0]
	}
	return val
}

func compact[T comparable](list []T) []T {
	result := make([]T, 0)
	var zero T
	for _, v := range list {
		var vv interface{} = v
		if vv != nil && v != zero {
			result = append(result, v)
		}
	}
	return result
}
