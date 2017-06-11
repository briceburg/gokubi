package xml

import (
	"fmt"
	"reflect"

	"github.com/clbanning/mxj/x2j"
)

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "gokubi/xml: Unmarshal(nil)"
	}
	if e.Type.Kind() != reflect.Ptr {
		return "gokubi/xml: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "gokubi/xml: Unmarshal(nil " + e.Type.String() + ")"
}

func Marshal(in map[string]interface{}) ([]byte, error) {
	return x2j.MapToXml(in)
}

func Unmarshal(data []byte, v *map[string]interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	o, err := x2j.XmlToMap(data)
	if err != nil {
		return err
	}

	for k, vv := range o {
		(*v)[fmt.Sprintf("%v", k)] = vv
	}

	return nil
}
