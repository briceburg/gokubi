package gokubi

import (
	"fmt"

	"github.com/briceburg/gokubi/formats/bash"
	"github.com/briceburg/gokubi/formats/hcl"
	"github.com/briceburg/gokubi/formats/json"
	"github.com/briceburg/gokubi/formats/xml"
	"github.com/briceburg/gokubi/formats/yaml"
)

// @TODO support streaming (io.Reader) encoders and decoders

// all config is unmarshalled into gokubi.Data
type Data map[string]interface{}

func (d Data) Decode(body []byte, format string) error {
	o := make(map[string]interface{})
	var err error
	switch format {
	case "bash":
		err = bash.Unmarshal(body, &o)
	case "hcl":
		err = hcl.Unmarshal(body, &o)
	case "json":
		err = json.Unmarshal(body, &o)
	case "xml":
		err = xml.Unmarshal(body, &o)
	case "yaml":
		err = yaml.Unmarshal(body, &o)
	default:
		err = fmt.Errorf("gokubi/Decode: unsupported decode format: %v", format)
	}
	if err == nil {
		d.Merge(o)
	}
	return err
}

func (d Data) Encode(format string) (string, error) {
	var bytes []byte
	var err error
	switch format {
	case "bash":
		bytes, err = bash.Marshal(d)
	// case "hcl":
	// 	bytes, err = hcl.Marshal(d)
	case "json":
		bytes, err = json.Marshal(d)
	case "xml":
		bytes, err = xml.Marshal(d)
	case "yaml":
		bytes, err = yaml.Marshal(d)
	default:
		err = fmt.Errorf("gokubi/Encode: unknown encode format: %v", format)
	}
	return string(bytes), err
}

func (d Data) Merge(m Data) error {
	for k, v := range m {
		d[k] = v
	}
	return nil
}

func (d Data) String() (string, error) {
	return d.Encode("json")
}
