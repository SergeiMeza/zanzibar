package zanzibar

import (
	"encoding/json"
	"io"

	"go.uber.org/config"
	"fmt"
)

type StaticConfig struct {
	name string
	data map[string]interface{}
}

func newJSONProvider(name string, readers []io.Reader) (*StaticConfig, error) {
	data := make(map[string]interface{})

	for _, reader := range readers {
		decoder := json.NewDecoder(reader)
		err := decoder.Decode(&data)
		if err != nil {
			return nil, err
		}
	}

	return &StaticConfig{
		name: name,
		data: data,
	}, nil
}

func NewJSONProvider(name string, readers []io.Reader) (*StaticConfig, error) {
	return newJSONProvider(name, readers)
}

// the Name of the provider (YAML, Env, etc)
func (j *StaticConfig) Name() string {
	return j.name
}

// Get pulls a config value
func (j *StaticConfig) Get(key string) config.Value {
	val, ok := j.data[key]
	return config.NewValue(j, key, val, ok)
}

func (j *StaticConfig) MustGetString(key string) string {
	var s string
	val := j.Get(key)
	if !val.HasValue() {
		panic(fmt.Errorf("%q not found", key))
	}
	err := val.Populate(&s)
	if err != nil {
		panic(fmt.Errorf("%q unexpected error %s", key, err.Error()))
	}
	return s
}

func (j *StaticConfig) MustGetBoolean(key string) bool {
	var s bool
	val := j.Get(key)
	if !val.HasValue() {
		panic(fmt.Errorf("%q not found", key))
	}
	err := val.Populate(&s)
	if err != nil {
		panic(fmt.Errorf("%q unexpected error %s", key, err.Error()))
	}
	return s
}



func (j *StaticConfig) MustGetInt(key string) int64 {
	var s int64
	val := j.Get(key)
	if !val.HasValue() {
		panic(fmt.Errorf("%q not found", key))
	}
	err := val.Populate(&s)
	if err != nil {
		panic(fmt.Errorf("%q unexpected error %s", key, err.Error()))
	}
	return s
}

