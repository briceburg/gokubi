package yaml

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "gokubi/yaml: Unmarshal(nil)"
	}
	if e.Type.Kind() != reflect.Ptr {
		return "gokubi/yaml: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "gokubi/yaml: Unmarshal(nil " + e.Type.String() + ")"
}

func Marshal(in map[string]interface{}) ([]byte, error) {
	return yaml.Marshal(in)
}

func Unmarshal(data []byte, v *map[string]interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	var o map[interface{}]interface{}
	if err := yaml.Unmarshal(data, &o); err != nil {
		return err
	}

	for k, vv := range o {
		(*v)[fmt.Sprintf("%v", k)] = safeValue(vv)
	}

	return nil
}

/**
safe functions inspired by https://github.com/elastic/beats/blob/6435194af9f42cbf778ca0a1a92276caf41a0da8/libbeat/common/mapstr.go#L117
  (c)Elasticsearch, license: APL2
**/

func InterfaceMapToStringMap(in map[interface{}]interface{}) map[string]interface{} {
	d := make(map[string]interface{})
	for k, v := range in {
		d[fmt.Sprintf("%v", k)] = safeValue(v)
	}
	return d
}

func safeArray(in []interface{}) []interface{} {
	a := make([]interface{}, len(in))
	for i, v := range in {
		a[i] = safeValue(v)
	}
	return a
}

func safeValue(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return safeArray(v)
	case map[interface{}]interface{}:
		return InterfaceMapToStringMap(v)
	default:
		return v
	}
}
