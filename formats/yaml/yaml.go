package yaml

import (
	"fmt"
	"reflect"

	"github.com/briceburg/gokubi/util"
	"gopkg.in/yaml.v2"
)

func Marshal(in map[string]interface{}) ([]byte, error) {
	return yaml.Marshal(in)
}

// Unmarshal parses the YML-encoded data and stores result in value pointed to by v.
// v must be a pointer to gokubi.Data or map[string]interface{}
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)

	// validate target
	if !util.IsValidUnmarshalPtr(rv) {
		return fmt.Errorf("gokubi/yaml: cannot unmarshal into %s, must be *map[string]interface{}", rv.Type())
	}

	// unmarshal yaml, convert to map[string]iterface{}
	var o map[interface{}]interface{}
	if err := yaml.Unmarshal(data, &o); err != nil {
		return err
	}
	x := util.InterfaceMapToStringMap(o)

	// update target
	t := rv.Elem()
	for key, val := range x {
		t.SetMapIndex(reflect.ValueOf(fmt.Sprintf("%v", key)), reflect.ValueOf(val))
	}

	return nil
}
