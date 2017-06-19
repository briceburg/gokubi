package hcl

import "github.com/hashicorp/hcl"

func Unmarshal(data []byte, v interface{}) error {
	return hcl.Unmarshal(data, v)
}
