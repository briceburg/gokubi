package gokubi

import (
	"fmt"

	"github.com/briceburg/gokubi/formats/bash"
	"github.com/briceburg/gokubi/formats/hcl"
	"github.com/briceburg/gokubi/formats/json"
	"github.com/briceburg/gokubi/formats/xml"
	"github.com/briceburg/gokubi/formats/yaml"
	"github.com/peterbourgon/mergemap"
)

// @TODO support streaming (io.Reader) encoders and decoders

// config is unmarshalled into gokubi.Data
type Data map[string]interface{}

var OutputFormats = []string{"bash", "json", "xml", "yaml"}
var InputFormats = []string{"bash", "hcl", "json", "xml", "yaml"}

//
// Decoding
//

func (d *Data) Decode(body []byte, f string) error {
	switch f {
	case "bash":
		return d.DecodeBash(body)
	case "hcl":
		return d.DecodeHCL(body)
	case "json":
		return d.DecodeJSON(body)
	case "xml", "html":
		return d.DecodeXML(body)
	case "yaml", "yml":
		return d.DecodeYAML(body)
	default:
		return fmt.Errorf("gokubi/Decode: unsupported decode format: %v", f)
	}
}

func (d *Data) DecodeBash(body []byte) error {
	return d.merge(bash.Unmarshal, body)
}

func (d *Data) DecodeHCL(body []byte) error {
	return d.merge(hcl.Unmarshal, body)
}

func (d *Data) DecodeJSON(body []byte) error {
	return d.merge(json.Unmarshal, body)
}

func (d *Data) DecodeXML(body []byte) error {
	return d.merge(xml.Unmarshal, body)
}

func (d *Data) DecodeYAML(body []byte) error {
	return d.merge(yaml.Unmarshal, body)
}

// Recursively merge configuration, descending into maps on key conflicts.
func (d *Data) merge(fn func([]byte, interface{}) error, body []byte) error {
	src := make(map[string]interface{})
	if err := fn(body, &src); err != nil {
		return err
	}
	dst := Data(mergemap.Merge(map[string]interface{}(*d), src))
	d = &dst
	return nil
}

//
// Encoding
//

func (d Data) Encode(f string) ([]byte, error) {
	switch f {
	case "bash":
		return d.EncodeBash()
	//case "hcl":
	//	return d.DecodeHCL()
	case "json":
		return d.EncodeJSON()
	case "xml", "html":
		return d.EncodeXML()
	case "yaml", "yml":
		return d.EncodeYAML()
	default:
		return nil, fmt.Errorf("gokubi/Encode: unsupported encode format: %v", f)
	}
}

func (d Data) EncodeBash() ([]byte, error) {
	return bash.Marshal(d)
}

/*
func (d Data) EncodeHCL() ([]byte, error) {
	return hcl.Marshal(d)
}
*/

func (d Data) EncodeJSON() ([]byte, error) {
	return json.Marshal(d)
}

func (d Data) EncodeXML() ([]byte, error) {
	return xml.Marshal(d)
}

func (d Data) EncodeYAML() ([]byte, error) {
	return yaml.Marshal(d)
}

//
// Rendering
//

func (d Data) String() string {
	bytes, _ := d.EncodeJSON()
	return string(bytes)
}
