package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadMockData(filename string, v interface{}) {
	data, _ := ioutil.ReadFile(fmt.Sprintf("/test/mockdata/%s", filename))
	json.Unmarshal(data, v)
}
