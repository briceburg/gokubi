package bash

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Marshal returns a shell parsable representation (var=value encoding) of a map
// arrays are preserved, and non literal values are serialized as JSON strings.
func Marshal(in map[string]interface{}) ([]byte, error) {
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

func Unmarshal(data []byte, v interface{}) error {
	return nil
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
