package gojsonconf

import (
	"encoding/json"
	"io/ioutil"
)

//load conf
func LoadConf(filename string) *map[string]string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	conf := make(map[string]string)

	json.Unmarshal(bytes, &conf)
	return &conf
}
