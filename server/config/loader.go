package config

import (
	"os"
	"fmt"
	"bytes"
	"io/fs"
	"sync"
	"embed"

	"gopkg.in/yaml.v3"
	"dario.cat/mergo"
)

//go:embed *.yaml
var ConfFS embed.FS

var Loader = NewGlobalLoader(ConfFS)

type GlobalLoader struct {
	fs fs.FS
	conf *Config
	mu sync.Mutex
}

func NewGlobalLoader(fs_ fs.FS) *GlobalLoader {
	return &GlobalLoader{
		fs: fs_,
	}
}

func (gl *GlobalLoader) Get() *Config {
	if gl.conf == nil {
		err := gl.load()
		if err != nil {
			panic(err)
		}
	}
	return gl.conf
}

func (gl *GlobalLoader) load() error {
	defer gl.mu.Unlock()
	gl.mu.Lock()

	env := CurrentEnv()
	conf, err := gl.priorityLoad(env)
	if err != nil {
		return err
	}
	gl.conf = conf

	return nil
}

// load priority
// 1. config.yaml
// 2. config.local.yaml
// 3. config.${ENV}.yaml
// 4. config.${ENV}.local.yaml
func (gl *GlobalLoader) priorityLoad(env Env) (*Config, error) {
	conf := map[string]interface{}{}
	c, err := gl.loadYAML("config.yaml")
	if err != nil {
		return nil, err
	}
	mergo.Merge(&conf, &c, mergo.WithOverride)

	paths := []string{
		"config.local.yaml",
		fmt.Sprintf("config.%s.yaml", env),
		fmt.Sprintf("config.%s.local.yaml", env),
	}
	for _, p := range paths {
		c, err := gl.loadYAML(p)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		mergo.Merge(&conf, &c, mergo.WithOverride)
	}

	result := &Config{}
	tmp, err := yaml.Marshal(conf)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(tmp, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (gl *GlobalLoader) loadYAML(path string) (map[string]interface{}, error) {
	f, err := gl.fs.Open(path)
	if err != nil {
		return nil, err
	}
	w := bytes.NewBuffer([]byte{})
	name := "main"
	t := NewYAMLTemplate()
	err = t.Compile(name, f, w)
	if err != nil {
		return nil, err
	}

	c := map[string]interface{}{}
	err = yaml.Unmarshal(w.Bytes(), &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
