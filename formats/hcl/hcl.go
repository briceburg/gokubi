package hcl

import "github.com/hashicorp/hcl"

func Unmarshal(data []byte, v *map[string]interface{}) error {
	return hcl.Unmarshal(data, v)
}
