package gokubi

import (
	"github.com/briceburg/gokubi/formats/bash"
	"github.com/briceburg/gokubi/formats/hcl"
	"github.com/briceburg/gokubi/formats/json"
	"github.com/briceburg/gokubi/formats/xml"
	"github.com/briceburg/gokubi/formats/yaml"
)

// @TODO support streaming (io.Reader) encoders and decoders

// config is unmarshalled into gokubi.Data
type Data map[string]interface{}

//
// Decoding
//

func (d *Data) DecodeBash(body []byte) error {
	return bash.Unmarshal(body, d)
}

func (d *Data) DecodeHCL(body []byte) error {
	return hcl.Unmarshal(body, d)
}

func (d *Data) DecodeJSON(body []byte) error {
	return json.Unmarshal(body, d)
}

func (d *Data) DecodeXML(body []byte) error {
	return xml.Unmarshal(body, d)
}

func (d *Data) DecodeYAML(body []byte) error {
	return yaml.Unmarshal(body, d)
}

//
// Encoding
//

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
