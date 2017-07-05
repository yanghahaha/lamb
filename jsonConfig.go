package configs

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//NewConfig new a Config
func NewConfig(file string) Config {
	return Config{file: file}
}

type Config struct {
	file string
	maps map[string]interface{}
}

//Get name pattern key or key.key.key
//
func (c *Config) Get(name string) (bool, interface{}) {
	if c.maps == nil {
		c.read()
	}

	if c.maps == nil {
		return false, nil
	}
	// app.view.path
	keys := strings.Split(name, ".")
	l := len(keys)
	if l == 1 {
		return true, c.maps[name]
	}

	var ret interface{}
	for i := 0; i < l; i++ {
		if i == 0 {
			ret = c.maps[keys[i]]
			if ret == nil {
				return false, nil
			}
		} else {
			m, ok := ret.(map[string]interface{})
			if ok {
				ret = m[keys[i]]
			} else {
				if l == i-1 {
					return true, ret
				}
				return false, nil
			}
		}
	}
	_, res := ret.(interface{})
	if res {
		return true, ret
	}
	return false, nil

}

func (c *Config) read() {
	if !filepath.IsAbs(c.file) {
		file, err := filepath.Abs(c.file)
		if err != nil {
			panic(err)
		}
		c.file = file
	}

	bts, err := ioutil.ReadFile(c.file)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bts, &c.maps)

	if err != nil {
		panic(err)
	}
}
