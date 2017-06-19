package xml

import (
	"fmt"
	"reflect"

	"github.com/briceburg/gokubi/util"
	"github.com/clbanning/mxj/x2j"
)

func Marshal(in map[string]interface{}) ([]byte, error) {
	return x2j.MapToXml(in)
}

// Unmarshal parses the XML-encoded data and stores result in value pointed to by v.
// v must be a pointer to gokubi.Data or map[string]interface{}
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)

	// validate target
	if !util.IsValidUnmarshalPtr(rv) {
		return fmt.Errorf("gokubi/xml: cannot unmarshal into %s, must be *map[string]interface{}", rv.Type())
	}

	// unmarshal xml
	o, err := x2j.XmlToMap(data)
	if err != nil {
		return err
	}

	// update target
	t := rv.Elem()
	for key, val := range o {
		t.SetMapIndex(reflect.ValueOf(fmt.Sprintf("%v", key)), reflect.ValueOf(val))
	}

	return nil
}
