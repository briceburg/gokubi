package util

import (
	"fmt"
	"reflect"
)

// ensure target argument to Unmarshal is a pointer to a map[string]interface{}
func IsValidUnmarshalPtr(rv reflect.Value) bool {
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return false
	}

	// test the pointed to element
	v := rv.Elem()
	if v.Kind() != reflect.Map {
		return false
	}
	mapType := v.Type()
	mapElemType := mapType.Elem().Kind()
	mapKeyType := mapType.Key().Kind()

	return mapKeyType == reflect.String && mapElemType == reflect.Interface
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
