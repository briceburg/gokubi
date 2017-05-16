package gokubi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// EnvMarshal returns the shell environment var=value encoding of in map
// flattens map values into JSON strings
// values cotaining a basic type or an array of basic types are preserved.
func EnvMarshal(in map[string]interface{}) ([]byte, error) {
	var b bytes.Buffer
	for k, v := range in {
		v, err := jsonVal(v)
		b.WriteString(fmt.Sprintf("%v=%v\n", k, v))
		if err != nil {
			return b.Bytes(), err
		}
	}
	return b.Bytes(), nil
}

func jsonVal(v interface{}) (string, error) {
	if t := reflect.TypeOf(v).Kind(); isNumeric(t) {
		// value is numeric
		return fmt.Sprintf("%v", v), nil
	} else if t == reflect.String {
		// value is a string
		return fmt.Sprintf("%q", v), nil
	} else if t == reflect.Array || t == reflect.Slice {
		// value is an array
		a, ok := v.([]interface{})
		if !ok {
			return "", fmt.Errorf("unable to iterate type")
		}
		var vals = make([]string, len(a))
		for i, v := range a {
			v, err := jsonVal(v)
			if err != nil {
				return "", err
			}
			vals[i] = v
		}
		return fmt.Sprintf("( %s )", strings.Join(vals, " ")), nil
	} else {
		// value is a map
		bytes, err := json.Marshal(v)
		return fmt.Sprintf("%q", string(bytes)), err
	}
}

func spaceSep(a []interface{}) string {
	var b bytes.Buffer
	for _, v := range a {
		b.WriteString(fmt.Sprintf("ARRVAL%T", v))
	}
	return b.String()
}

func isNumeric(t reflect.Kind) bool {
	return (t >= reflect.Bool && t <= reflect.Float64)
}

func isBasic(t reflect.Kind) bool {
	return isNumeric(t) || t == reflect.String
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
