package gokubi

import(
  "encoding/json"
  "fmt"
  "os"
)

// @TODO support streaming (io.Reader) encoders and decoders
//       - did not find a method in go-yaml, only marshal/unmarshal

type Data map[string]interface{}


//
// mutators
//

func (d Data) Merge(m Data) error {
  for k, v := range m {
    d[k] = v
  }
  return nil
}


//
// input - input automerges iself
//

func (d Data) Decode(input []byte) error {
  return d.DecodeJSON(input)
}

func (d Data) DecodeJSON(input []byte) error {
  var o Data
  if err := json.Unmarshal(input, &o); err != nil {
    fmt.Fprintf(os.Stderr, "gokubi/Data.DecodeJSON: %v", err)
    return err
  }
  return d.Merge(o)
}



//
// output
//

func (d Data) String() (string, error) {
  return d.EncodeJSON()
}

func (d Data) EncodeJSON() (string, error) {
  bytes, err := json.Marshal(d)
  if err != nil {
    fmt.Fprintf(os.Stderr, "gokubi/Data.EncodeJSON: invalid json: %v", err)
    return "", err
  }
  return string(bytes), nil
}
