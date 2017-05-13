package gokubi

import(
  "encoding/json"
  "fmt"
  "os"
  "reflect"
)

// all config is unmarshalled into gokubi.Data
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

// @TODO support streaming (io.Reader) encoders and decoders
//       - did not find a method in go-yaml, only marshal/unmarshal
func (d Data) Decode(methodName string, body []byte) error {
  method := reflect.ValueOf(d).MethodByName(methodName)
  result := method.Call([]reflect.Value{reflect.ValueOf(body)})[0]
  if result.IsNil(){
      return nil
    }
  return result.Interface().(error)
}

func (d Data) DecodeJSON(body []byte) error {
  var o Data
  if err := json.Unmarshal(body, &o); err != nil {
    fmt.Fprintf(os.Stderr, "gokubi/Data.DecodeJSON: %v", err)
    return err
  }
  return d.Merge(o)
}

func (d Data) DecodeYML(body []byte) error {
  var o Data
  if err := json.Unmarshal(body, &o); err != nil {
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
