package httpcache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type HttpCache struct {
	sync.Map
}

func NewCacheFromFile(filepath string) HttpCache {
	c := HttpCache{}
	c.FromFile(filepath)
	return c
}

func (c *HttpCache) FromFile(filepath string) {
	dat, err := ioutil.ReadFile(filepath)
	if err == nil {
		err2 := json.Unmarshal(dat, &c)
		if err2 != nil {
			fmt.Fprintln(os.Stderr, "Can't parse json")
		}
	}
}

func (c HttpCache) SaveToFile(filepath string) (n int, err error) {
	b, err := json.Marshal(c)
	if err != nil {
		// fmt.Fprintln(os.Stderr, err)
		return 0, err
	}

	f, err := os.Create(filepath)
	defer f.Close()

	if err != nil {
		// fmt.Fprintln(os.Stderr, err)
		return 0, err
	}

	return f.Write(b)
}

func (c *HttpCache) UnmarshalJSON(data []byte) error {
	var tmpMap map[string]interface{}
	if err := json.Unmarshal(data, &tmpMap); err != nil {
		return err
	}
	for key, value := range tmpMap {
		c.Store(key, value)
	}
	return nil
}

func (c HttpCache) MarshalJSON() ([]byte, error) {
	tmpMap := make(map[string]string)
	c.Range(func(k, v interface{}) bool {
		tmpMap[k.(string)] = v.(string)
		return true
	})
	return json.Marshal(tmpMap)
}
